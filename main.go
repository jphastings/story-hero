package main

import (
	"log"
	"strings"

	"github.com/jphastings/story-hero/pkg/clonehero"
	"github.com/jphastings/story-hero/pkg/storymode"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	si, sc, sd, err := clonehero.Load()
	check(err)

	ss, err := storymode.LoadStories(si.Directories.Paths, sc, sd)
	check(err)

	if len(ss.ActiveStories) == 0 {
		log.Fatalf("FATAL: No stories found to load in %s\n", strings.Join(si.Directories.Paths, ", "))
	}

	for _, s := range ss.ActiveStories {
		log.Printf("INFO: Story running: %s\n", s.Story.Title)
	}

	for _, s := range ss.InactiveStories {
		log.Printf("WARN: Story unable to run (missing songs): %s\n", s)
	}

	sd.Watch(func() error {
		songs, err := ss.SongVisibility()
		if err != nil {
			return err
		}

		for songID, isNowUnlocked := range songs {
			if err := ss.ToggleVisibility(songID, isNowUnlocked); err != nil {
				action := "lock"
				if isNowUnlocked {
					action = "unlock"
				}
				log.Printf("WARN: Was unable to %s song with ID %s: %s\n", action, songID, err)
			}
		}

		return nil
	})
}
