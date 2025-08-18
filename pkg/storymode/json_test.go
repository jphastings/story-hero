package storymode_test

import (
	"os"
	"testing"

	"github.com/jphastings/clone-hero-storymode/pkg/storymode"
	"github.com/stretchr/testify/assert"
)

func TestParseGH1Fixture(t *testing.T) {
	file, err := os.Open("fixtures/gh1.ini")
	assert.NoError(t, err, "Failed to open fixture file")
	defer file.Close()

	story, err := storymode.Parse(file)
	assert.NoError(t, err, "Failed to parse story")

	assert.Equal(t, "Guitar Hero", story.Title)
}
