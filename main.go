package main

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/jphastings/clone-hero-storymode/pkg/clonehero"
	"github.com/jphastings/clone-hero-storymode/pkg/storymode"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	songCacheFile   = "songcache.bin"
	scoreDataFile   = "scoredata.bin"
	hiddenSongsFile = "hiddensongs.bin"
)

func main() {
	sf, err := os.Open("./pkg/storymode/fixtures/gh1.ts")
	check(err)
	scf, err := os.Open(supportDirPath(songCacheFile))
	check(err)
	sdf, err := os.Open(supportDirPath(scoreDataFile))
	check(err)

	sc, err := clonehero.OpenSongCache(scf)
	check(err)
	sd, err := clonehero.OpenScoreData(sdf)
	check(err)

	s, err := storymode.LoadStory(sf, sc, sd)
	check(err)

	watchFile(supportDirPath(scoreDataFile), func() error {
		if err := s.ScoreData.Reload(); err != nil {
			return err
		}

		songs, err := s.SongAvailability()
		if err != nil {
			return err
		}

		hidden, err := os.OpenFile(supportDirPath(hiddenSongsFile), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer hidden.Close()

		for songID, isUnlocked := range songs {
			song, ok := sc.Songs[songID]
			if !ok || song == nil {
				log.Printf("ERROR: missing song: %s\n", songID)
				continue
			}
			if !isUnlocked {
				songIDBytes, err := songID.ToBytes()
				if err != nil {
					return err
				}

				if _, err := hidden.Write(songIDBytes); err != nil {
					return err
				}
			}
		}

		return nil
	})
}

// A blocking function that watches for the specific file to change, and calls the given function when it does
// It also calls the callback immediately upon being run, as a first pass.
func watchFile(path string, callback func() error) error {
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

	err = watcher.Add(path)
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

				log.Println("INFO: performing callback")
				if err := callback(); err != nil {
					log.Printf("WARN: (while performing watcher callback) %s\n", err.Error())
				} else {
					log.Println("INFO: callback complete")
				}
			})

		case err := <-watcher.Errors:
			return err
		}
	}
}
