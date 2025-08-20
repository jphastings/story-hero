package clonehero

import (
	"path/filepath"
)

// TODO: Support portable mode: https://wiki.clonehero.net/books/clone-hero-manual/page/adding-custom-songs
func configDir(filename string) string {
	return filepath.Join(homedir, "Documents", "Clone Hero", filename)
}

func SupportDirPath(filename string) string {
	return filepath.Join(
		homedir,
		"AppData",
		"LocalLow",
		"srylain Inc_",
		"Clone Hero",
		filename,
	)
}
