package storymode

import (
	"fmt"
	"io"

	"github.com/clarkmcc/go-typescript"
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

	if err := s.parseStory(r); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *StoryHeroState) Usable() bool {
	for _, g := range s.Story.Groups {
		for _, songID := range g.Songs {
			if !s.songPresent(songID) {
				return false
			}
		}
	}
	return true
}

func (s *StoryHeroState) parseStory(r io.Reader) error {
	jsCode, err := typescript.Transpile(r, typescript.WithVersion("v4.9.3"))
	if err != nil {
		return fmt.Errorf("failed to transpile TypeScript: %w", err)
	}

	if _, err := s.vm.RunString(jsCode); err != nil {
		return fmt.Errorf("failed to execute transpiled Story definition: %w", err)
	}

	return nil
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
