package storymode

import (
	"fmt"
	"io"
	"log"

	"github.com/dop251/goja"

	"github.com/jphastings/story-hero/pkg/clonehero"
	"github.com/jphastings/story-hero/pkg/types"
)

type StoryHeroState struct {
	SongCache *clonehero.SongCache
	ScoreData *clonehero.ScoreData
	Story     *types.Story

	vm          *goja.Runtime
	state       goja.Value
	songPresent func(types.MD5Hash) bool
}

func LoadStory(r io.Reader, sc *clonehero.SongCache, sd *clonehero.ScoreData, songPresent func(types.MD5Hash) bool) (*StoryHeroState, error) {
	s := &StoryHeroState{
		SongCache:   sc,
		ScoreData:   sd,
		songPresent: songPresent,
	}

	s.prepareVM()

	if err := s.loadHelpers(); err != nil {
		return nil, err
	}

	if err := s.executeTypescript(r); err != nil {
		return nil, err
	}

	if s.Story == nil {
		return nil, fmt.Errorf("no story was defined")
	}

	return s, nil
}

func (s *StoryHeroState) Usable() bool {
	if s.Story == nil {
		return false
	}

	for _, g := range s.Story.Groups {
		for _, songID := range g.Songs {
			if !s.songPresent(songID) {
				log.Printf("WARN: song '%s' unavailable, story '%s' is unplayable\n", songID, s.Story.Title)
				return false
			}
		}
	}
	return true
}

func (s *StoryHeroState) SongVisibility() (map[types.MD5Hash]bool, error) {
	availability := make(map[types.MD5Hash]bool)

	for _, g := range s.Story.Groups {
		for _, songID := range g.Songs {
			av, err := s.isUnlocked(g, songID)
			if err != nil {
				return nil, err
			}
			availability[songID] = av
		}
	}

	return availability, nil
}
