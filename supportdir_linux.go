package main

import (
	"os"
	"path/filepath"
)

var homedir string

func init() {
	var err error
	homedir, err = os.UserHomeDir()
	check(err)
}

func supportDirPath(filename string) string {
	return filepath.Join(
		homedir,
		".config",
		"unity3d",
		"srylain Inc_",
		"Clone Hero",
		filename,
	)
}
