package storymode

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/jphastings/story-hero/pkg/clonehero"
	"github.com/jphastings/story-hero/pkg/types"
)

type Stories struct {
	songCache   *clonehero.SongCache
	scoreData   *clonehero.ScoreData
	hiddenSongs map[types.MD5Hash]string

	ActiveStories   []*StoryHeroState
	InactiveStories []string
}

func LoadStories(songsPaths []string, sc *clonehero.SongCache, sd *clonehero.ScoreData) (Stories, error) {
	ss := Stories{
		songCache: sc,
		scoreData: sd,
	}

	if err := ss.findHiddenSongs(songsPaths); err != nil {
		return ss, err
	}

	for _, songsPath := range songsPaths {
		err := filepath.WalkDir(songsPath, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() && strings.HasSuffix(path, ".story.ts") {
				sf, err := os.Open(path)
				if err != nil {
					return err
				}
				defer sf.Close()

				s, err := LoadStory(sf, sc, sd, ss.songPresent)
				if err != nil {
					return err
				}

				if s.Usable() {
					ss.ActiveStories = append(ss.ActiveStories, s)
				} else {
					ss.InactiveStories = append(ss.InactiveStories, s.Story.Title)
				}

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
	if _, isVisible := ss.songCache.Songs[songID]; isVisible {
		return true
	}

	if _, isHidden := ss.hiddenSongs[songID]; isHidden {
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
