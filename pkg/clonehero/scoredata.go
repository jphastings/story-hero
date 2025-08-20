package clonehero

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math/big"
	"os"

	"github.com/jphastings/story-hero/pkg/types"
)

type ScoreData struct {
	f         *os.File
	SongPlays map[types.MD5Hash]types.SongPlay
}

func OpenScoreData(f *os.File) (*ScoreData, error) {
	sd := &ScoreData{f: f}
	if err := sd.Reload(); err != nil {
		return nil, err
	}
	return sd, nil
}

func (sd *ScoreData) Reload() error {
	_, err := sd.f.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("unable to interact with the ScoreData file: %w", err)
	}

	header := make([]byte, 8)
	if _, err := sd.f.Read(header); err != nil {
		return fmt.Errorf("failed to read header: %w", err)
	}

	if !bytes.Equal(header[:4], []byte{0x41, 0x65, 0x34, 0x01}) {
		return fmt.Errorf("the Score Data magic number isn't as expected, stopping parsing")
	}

	// Song count (4 bytes, little endian)
	songCount := binary.LittleEndian.Uint32(header[4:8])

	sd.SongPlays = make(map[types.MD5Hash]types.SongPlay)

	for i := uint32(0); i < songCount; i++ {
		songID, err := readSongID(sd.f)
		if err != nil {
			return err
		}

		instrumentCount, err := readUint8(sd.f)
		if err != nil {
			return err
		}

		playCount, err := readUint24(sd.f)
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
			instrumentType, err := readUint16(sd.f)
			if err != nil {
				return err
			}

			difficulty, err := readUint8(sd.f)
			if err != nil {
				return err
			}

			percentage, err := sd.readPercentage()
			if err != nil {
				return err
			}

			stars, err := readUint8(sd.f)
			if err != nil {
				return err
			}

			// Skip unneeded values
			if err := skipBytes(sd.f, 4); err != nil {
				return err
			}

			score, err := readUint32(sd.f)
			if err != nil {
				return err
			}

			songPlay.Scores[instrumentType] = types.Score{
				Difficulty: difficulty,
				Percentage: percentage,
				Stars:      stars,
				Score:      score,
			}
		}

		sd.SongPlays[songID] = songPlay
	}

	return nil
}

func (sd *ScoreData) readPercentage() (*big.Rat, error) {
	numeratorBuf := make([]byte, 2)
	if _, err := sd.f.Read(numeratorBuf); err != nil {
		return nil, fmt.Errorf("failed to read score numerator: %w", err)
	}
	numerator := binary.LittleEndian.Uint16(numeratorBuf)

	denominatorBuf := make([]byte, 2)
	if _, err := sd.f.Read(denominatorBuf); err != nil {
		return nil, fmt.Errorf("failed to read score denominator: %w", err)
	}
	denominator := binary.LittleEndian.Uint16(denominatorBuf)

	return big.NewRat(int64(numerator), int64(denominator)), nil
}
