package storymode_test

import (
	"os"
	"testing"

	"github.com/jphastings/story-hero/pkg/storymode"
	"github.com/stretchr/testify/assert"
)

func TestParseGH1Fixture(t *testing.T) {
	file, err := os.Open("fixtures/gh1.ts")
	assert.NoError(t, err, "Failed to open fixture file")
	defer file.Close()

	s, err := storymode.LoadStory(file, nil, nil)
	assert.NoError(t, err, "Failed to parse story")

	assert.Equal(t, "Guitar Hero", s.Story.Title)
}
