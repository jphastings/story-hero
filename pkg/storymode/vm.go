package storymode

import (
	"embed"
	"fmt"
	"io"

	"github.com/clarkmcc/go-typescript"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"

	"github.com/go-viper/mapstructure/v2"
	"github.com/jphastings/story-hero/pkg/types"
)

//go:embed implementations.ts
var helpersFile embed.FS

func (s *StoryHeroState) prepareVM() {
	s.vm = goja.New()

	req := require.NewRegistry()
	req.Enable(s.vm)
	console.Enable(s.vm)

	s.vm.Set("exports", s.vm.NewObject())

	s.vm.Set("defineStory", s.defineStory)

	s.vm.Set("story", func() map[string]interface{} { return jsIfy(s.vm, *s.Story) })

	s.vm.Set("useState", s.useState)
	s.vm.Set("plays", s.plays)
}

func (s *StoryHeroState) loadHelpers() error {
	f, err := helpersFile.Open("implementations.ts")
	if err != nil {
		return fmt.Errorf("unable to load the TypeSCript helpers from memory: %w", err)
	}

	return s.executeTypescript(f)
}

func (s *StoryHeroState) executeTypescript(r io.Reader) error {
	jsCode, err := typescript.Transpile(r, typescript.WithVersion("v4.9.3"))
	if err != nil {
		return fmt.Errorf("failed to transpile TypeScript: %w", err)
	}

	if _, err := s.vm.RunString(jsCode); err != nil {
		if ex, ok := err.(*goja.Exception); ok {
			return fmt.Errorf("failed to execute transpiled TypeScript: %w", ex)
		}
		return fmt.Errorf("failed to execute transpiled TypeScript: %w", err)
	}

	return nil
}

func jsIfy(vm *goja.Runtime, v interface{}) map[string]interface{} {
	var jsReady map[string]interface{}
	// TODO: This needs to be deep mapped. See https://github.com/go-viper/mapstructure/pull/53
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
		This:      s.vm.ToValue(jsIfy(s.vm, g)),
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
	plays, isPlayed := s.ScoreData.SongPlays[songID]
	if isPlayed {
		return jsIfy(s.vm, plays)
	}
	return nil
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
