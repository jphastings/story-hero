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

func (ss *Stories) ToggleVisibility(songID types.MD5Hash, makeVisible bool) error {
	song, isVisible := ss.songCache.Songs[songID]
	hiddenSongPath, isHidden := ss.hiddenSongs[songID]

	switch {
	case isVisible && makeVisible:
		// Ignore duplicate hidden song
		return nil

	case !isVisible && !makeVisible:
		// Ignore if the song is also absent in the hidden songs, it should have been caught at launch
		return nil

	case isVisible && !makeVisible:
		// Ignore duplicate hidden song
		songFile := songFile(song.Path)
		if _, err := os.Stat(songFile); !os.IsNotExist(err) {
			if err := os.Rename(songFile, songFile+lockSuffix); err != nil {
				return fmt.Errorf("unable to rename to lock %s: %w", songFile, err)
			}
		}

		// Note: the player will need to "Scan songs" to have this reflected in their game state; which is annoying
		ss.hiddenSongs[songID] = song.Path
		delete(ss.songCache.Songs, songID)

		return nil

	case isHidden && makeVisible:
		// This would never be called if there was a duplicate song in the visible songs set (because this case is last)
		songFile := songFile(hiddenSongPath)
		if err := os.Rename(songFile+lockSuffix, songFile); err != nil {
			return fmt.Errorf("unable to rename to unlock %s: %w", songFile+lockSuffix, err)
		}

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

func (ss *Stories) findHiddenSongs(songsPaths []string) error {
	ss.hiddenSongs = make(map[types.MD5Hash]string)

	for _, songsPath := range songsPaths {
		err := filepath.Walk(songsPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() || filepath.Ext(path) != lockSuffix {
				return nil
			}

			// TODO: Handle .sng files
			// Remove the song.ini.storylocked filename
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

			ss.hiddenSongs[songID] = songDir

			return nil
		})
		if err != nil {
			return fmt.Errorf("error scanning songs path %s: %w", songsPath, err)
		}
	}

	return nil
}
