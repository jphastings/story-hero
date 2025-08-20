package storymode

import (
	"reflect"

	"github.com/dop251/goja"
	"github.com/jphastings/story-hero/pkg/types"
	"github.com/mitchellh/mapstructure"
)

func (s *StoryHeroState) defineStory(storyMap goja.Value) error {
	config := &mapstructure.DecoderConfig{
		DecodeHook: decodeHookStringToMD5Hash(),
		Result:     &s.Story,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(storyMap.Export())
}

func decodeHookStringToMD5Hash() mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}
		if t != reflect.TypeOf(types.MD5Hash("")) {
			return data, nil
		}

		songID, ok := data.(types.MD5Hash)
		if !ok {
			return data, nil
		}

		// Make sure it can be decoded into an MD5Hash
		if _, err := songID.ToBytes(); err != nil {
			return nil, err
		}

		return songID, nil
	}
}
