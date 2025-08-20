package storymode

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/go-viper/mapstructure/v2"
	"github.com/jphastings/clone-hero-storymode/pkg/types"
)

func (s *StoryHeroState) prepareVM() {
	s.vm = goja.New()

	s.vm.Set("exports", s.vm.NewObject())

	s.vm.Set("defineStory", s.defineStory)

	s.vm.Set("useState", s.useState)
	s.vm.Set("plays", s.plays)

	s.vm.Set("totalStars", s.songPlayIterator(maxStars))
	s.vm.Set("totalScore", s.songPlayIterator(maxScore))

	s.vm.Set("countMeetingScore", s.countMeeting(maxScore))
	s.vm.Set("countMeetingPercentage", s.countMeeting(maxPercentage))
	s.vm.Set("countMeetingStars", s.countMeeting(maxStars))
}

type SongPlayIterator func(types.SongPlay) uint

func (s *StoryHeroState) songPlayIterator(fn SongPlayIterator) func() uint {
	return func() uint {
		n := uint(0)
		for _, g := range s.Story.Groups {
			for _, songID := range g.Songs {
				play, ok := s.ScoreData.SongPlays[songID]
				if ok {
					n += fn(play)
				}
			}
		}
		return n
	}
}

func (s *StoryHeroState) countMeeting(fn SongPlayIterator) func(uint, bool) uint {
	return func(meeting uint, only bool) uint {
		return s.songPlayIterator(func(sp types.SongPlay) uint {
			if only && fn(sp) >= meeting {
				return 1
			} else if !only && fn(sp) > meeting {
				return 1
			}
			return 0
		})()
	}
}

func maxScore(sp types.SongPlay) uint {
	maxScore := uint(0)
	for _, score := range sp.Scores {
		if score.Score > maxScore {
			maxScore = score.Score
		}
	}
	return maxScore
}

func maxStars(sp types.SongPlay) uint {
	maxStars := uint(0)
	for _, score := range sp.Scores {
		if score.Stars > maxStars {
			maxStars = score.Stars
		}
	}
	return maxStars
}

// Kludge for now, percentages are treated as uint; which they mostly are…
func maxPercentage(sp types.SongPlay) uint {
	maxPercentage := float64(0)
	for _, score := range sp.Scores {
		percentage, _ := score.Percentage.Float64()
		if percentage > maxPercentage {
			maxPercentage = percentage
		}
	}
	return uint(maxPercentage)
}

func jsIfy(vm *goja.Runtime, v interface{}) map[string]interface{} {
	var jsReady map[string]interface{}
	if err := mapstructure.Decode(v, &jsReady); err != nil {
		jsError := vm.NewGoError(fmt.Errorf("unable to convert to JS VM error: %w", err))
		panic(jsError)
	}

	return jsReady
}

// A helper that allows Go to call the isUnlocked function of a group without having to use/interpret goja values
func (s *StoryHeroState) isUnlocked(g types.Group, songID types.MD5Hash) (bool, error) {
	// Always unlocked if not specified
	if g.IsUnlocked == nil {
		return true, nil
	}

	// TODO: Can I custom unmarshal with mapstructure so that these are 'normal' functions?
	// Feels like I should be able to define them as ActionFunc and UnlockFunc in Go, and have mapstructure do the goja.FunctionCall stuff behind the scenes
	arg := goja.FunctionCall{
		This:      goja.Undefined(),
		Arguments: []goja.Value{s.vm.ToValue(songID)},
	}
	val := g.IsUnlocked(arg).Export()

	isUnlocked, isBool := val.(bool)
	if !isBool {
		return false, fmt.Errorf("the isUnlocked function did not return a boolean")
	}

	return isUnlocked, nil
}

func (s *StoryHeroState) plays(songID types.MD5Hash) map[string]interface{} {
	// TODO: Real song lookup
	sp := &types.SongPlay{
		PlayCount: 2,
	}

	return jsIfy(s.vm, sp)
}

func (s *StoryHeroState) useState(init goja.Callable) []goja.Callable {
	newState, err := init(goja.Undefined(), s.state)
	if err != nil {
		jsError := s.vm.NewGoError(fmt.Errorf("unable initialize state: %w", err))
		panic(jsError)
	}
	s.state = newState

	getState := func(goja.Value, ...goja.Value) (goja.Value, error) {
		return s.vm.ToValue(s.state), nil
	}

	updateState := func(this goja.Value, args ...goja.Value) (goja.Value, error) {
		if len(args) == 0 {
			return nil, fmt.Errorf("updateState expects a single function argument")
		}
		updater, isFunc := goja.AssertFunction(args[0])
		if !isFunc {
			return nil, fmt.Errorf("updateState expects a single function argument")
		}

		newState, err := updater(goja.Undefined(), s.state)
		if err != nil {
			return nil, err
		}
		s.state = newState

		return nil, nil
	}

	return []goja.Callable{getState, updateState}
}
