package clonehero

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/jphastings/story-hero/pkg/types"
)

type ScoreData struct {
	path      string
	SongPlays map[types.MD5Hash]types.SongPlay
}

func OpenScoreData(path string) (*ScoreData, error) {
	sd := &ScoreData{path: path}
	if err := sd.Reload(); err != nil {
		return nil, err
	}
	return sd, nil
}

// A blocking function that watches for the specific file to change, and calls the given function when it does
// It also calls the callback immediately upon being run, as a first pass.
func (sd *ScoreData) Watch(callback func() error) error {
	log.Println("INFO: performing callback")
	if err := callback(); err != nil {
		return err
	}
	log.Println("INFO: callback complete")

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	err = watcher.Add(sd.path)
	if err != nil {
		return err
	}

	var mu sync.Mutex
	var debounceTimer *time.Timer
	const debounceDelay = 100 * time.Millisecond

	for {
		select {
		case event := <-watcher.Events:
			if !event.Has(fsnotify.Write) {
				continue
			}

			// Reset the timer if it's already going
			if debounceTimer != nil {
				debounceTimer.Stop()
			}
			debounceTimer = time.AfterFunc(debounceDelay, func() {
				mu.Lock()
				defer mu.Unlock()

				if err := sd.Reload(); err != nil {
					log.Printf("ERROR: (while reloading Score Data in watcher callback) %s\n", err.Error())
					return
				}

				log.Println("INFO: performing callback")

				if err := callback(); err != nil {
					log.Printf("ERROR: (while performing watcher callback) %s\n", err.Error())
					return
				}

				log.Println("INFO: callback complete")
			})

		case err := <-watcher.Errors:
			return err
		}
	}
}

func (sd *ScoreData) Reload() error {
	f, err := os.Open(sd.path)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("unable to interact with the ScoreData file: %w", err)
	}

	header := make([]byte, 8)
	if _, err := f.Read(header); err != nil {
		return fmt.Errorf("failed to read header: %w", err)
	}

	if !bytes.Equal(header[:4], []byte{0x41, 0x65, 0x34, 0x01}) {
		return fmt.Errorf("the Score Data magic number isn't as expected, stopping parsing")
	}

	// Song count (4 bytes, little endian)
	songCount := binary.LittleEndian.Uint32(header[4:8])

	sd.SongPlays = make(map[types.MD5Hash]types.SongPlay)

	for i := uint32(0); i < songCount; i++ {
		songID, err := readSongID(f)
		if err != nil {
			return err
		}

		instrumentCount, err := readUint8(f)
		if err != nil {
			return err
		}

		playCount, err := readUint24(f)
		if err != nil {
			return err
		}

		songPlay := types.SongPlay{
			ID:        songID,
			PlayCount: playCount,
			Scores:    make(map[uint]types.Score),
		}

		// Read instruments
		for j := uint(0); j < instrumentCount; j++ {
			instrumentType, err := readUint16(f)
			if err != nil {
				return err
			}

			difficulty, err := readUint8(f)
			if err != nil {
				return err
			}

			percentage, err := readUint16(f)
			if err != nil {
				return err
			}

			speed, err := readUint16(f)
			if err != nil {
				return err
			}

			stars, err := readUint8(f)
			if err != nil {
				return err
			}

			// Skip unneeded values
			if err := skipBytes(f, 4); err != nil {
				return err
			}

			score, err := readUint32(f)
			if err != nil {
				return err
			}

			songPlay.Scores[instrumentType] = types.Score{
				Difficulty: difficulty,
				Percentage: percentage,
				Speed:      speed,
				Stars:      stars,
				Score:      score,
			}
		}

		sd.SongPlays[songID] = songPlay
	}

	return nil
}
