package storymode

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jphastings/story-hero/pkg/clonehero"
	"github.com/jphastings/story-hero/pkg/types"
)

const (
	storySuffix = ".story.ts"
)

type Stories struct {
	songCache     *clonehero.SongCache
	scoreData     *clonehero.ScoreData
	unlockedSongs map[types.MD5Hash]string
	lockedSongs   map[types.MD5Hash]string

	ActiveStories   []*StoryHeroState
	InactiveStories []string
}

func LoadStories(songsPaths []string, sc *clonehero.SongCache, sd *clonehero.ScoreData) (Stories, error) {
	ss := Stories{
		songCache: sc,
		scoreData: sd,
	}

	if list, err := ss.findSongs(songsPaths, songIniFile); err == nil {
		ss.unlockedSongs = list
	} else {
		return ss, err
	}

	if list, err := ss.findSongs(songsPaths, songIniFile+lockSuffix); err == nil {
		ss.lockedSongs = list
	} else {
		return ss, err
	}

	for _, songsPath := range songsPaths {
		err := filepath.WalkDir(songsPath, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() || !strings.HasSuffix(path, storySuffix) {
				return nil
			}

			sf, err := os.Open(path)
			if err != nil {
				return err
			}
			defer sf.Close()

			s, err := LoadStory(sf, sc, sd, ss.songPresent)
			if err != nil {
				log.Printf("WARN: Story file at '%s' ignored: %v", path, err)
				return nil
			}

			if s.Usable() {
				ss.ActiveStories = append(ss.ActiveStories, s)
			} else {
				ss.InactiveStories = append(ss.InactiveStories, s.Story.Title)
			}

			return nil
		})
		if err != nil {
			return ss, err
		}
	}

	return ss, nil
}

func (ss Stories) songPresent(songID types.MD5Hash) bool {
	if _, isUnlocked := ss.unlockedSongs[songID]; isUnlocked {
		return true
	}

	if _, isLocked := ss.lockedSongs[songID]; isLocked {
		return true
	}

	return false
}

func (ss Stories) SongVisibility() (map[types.MD5Hash]bool, error) {
	totalAvail := make(map[types.MD5Hash]bool)

	// Make sure scores are up-to-date
	if err := ss.scoreData.Reload(); err != nil {
		return nil, err
	}

	for _, s := range ss.ActiveStories {
		sAvail, err := s.SongVisibility()
		if err != nil {
			return nil, err
		}

		for k, v := range sAvail {
			// There will be no overlaps, as stories with overlapping songs are rejected
			totalAvail[k] = v
		}
	}

	return totalAvail, nil
}
