package storymode

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/jphastings/story-hero/pkg/types"
)

const (
	songIniFile = "song.ini"
	lockSuffix  = ".storylocked"
)

func (ss *Stories) ToggleVisibility(songID types.MD5Hash, doUnlock bool) error {
	unlockedSongPath, isUnlocked := ss.unlockedSongs[songID]
	lockedSongPath, isLocked := ss.lockedSongs[songID]

	switch {
	case isUnlocked && doUnlock:
		// Ignore duplicate locked song
		return nil

	case !isUnlocked && !doUnlock:
		// Ignore if the song is also absent in the locked songs, it should have been caught at launch
		return nil

	case isUnlocked && !doUnlock:
		songPath := songFile(unlockedSongPath)
		// Ignore duplicate locked song

		if _, err := os.Stat(songPath); os.IsNotExist(err) {
			return fmt.Errorf("a file has been renamed elsewhere since this program started, please restart: %w", err)
		} else if err != nil {
			return err
		} else {
			if err := os.Rename(songPath, songPath+lockSuffix); err != nil {
				return fmt.Errorf("unable to rename to lock %s: %w", songPath, err)
			}
		}

		// Note: the player will need to "Scan songs" to have this reflected in their game state; which is annoying
		ss.lockedSongs[songID] = unlockedSongPath
		delete(ss.unlockedSongs, songID)

		return nil

	case isLocked && doUnlock:
		songPath := songFile(lockedSongPath)

		// This would never be called if there was a duplicate song in the visible songs set (because this case is last)
		if err := os.Rename(songPath+lockSuffix, songPath); err != nil {
			return fmt.Errorf("unable to rename to unlock %s: %w", songPath+lockSuffix, err)
		}

		// Note: the player will need to "Scan songs" to have this reflected in their game state; which is annoying
		ss.unlockedSongs[songID] = lockedSongPath
		delete(ss.lockedSongs, songID)

		return nil
	}

	return fmt.Errorf("the song (%s) couldn't be found, cannot make it visible", songID)
}

func songFile(songPath string) string {
	if filepath.Ext(songPath) == ".sng" {
		return songPath
	}
	return filepath.Join(songPath, songIniFile)
}

func (ss *Stories) findSongs(songsPaths []string, filename string) (map[types.MD5Hash]string, error) {
	list := make(map[types.MD5Hash]string)

	for _, songsPath := range songsPaths {
		err := filepath.Walk(songsPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() || filepath.Base(path) != filename {
				return nil
			}

			// TODO: Handle .sng files
			songDir := filepath.Dir(path)

			nf, err := os.Open(filepath.Join(songDir, "notes.mid"))
			if err != nil {
				return fmt.Errorf("unable to open the motes.mid file to calculate the songID of %s: %w", songDir, err)
			}
			defer nf.Close()

			hash := md5.New()
			if _, err := io.Copy(hash, nf); err != nil {
				return fmt.Errorf("unable to calculate the songID of %s: %w", songDir, err)
			}
			songID := types.MD5HashFromBytes([16]byte(hash.Sum(nil)))

			list[songID] = songDir

			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("error scanning songs path %s: %w", songsPath, err)
		}
	}

	return list, nil
}
