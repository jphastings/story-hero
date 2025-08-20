package clonehero

import (
	"path/filepath"
)

func configDir(filename string) string {
	return filepath.Join(homedir, "Clone Hero", filename)
}

func SupportDirPath(filename string) string {
	return filepath.Join(
		homedir,
		"Library",
		"Application Support",
		"com.srylain.CloneHero",
		filename,
	)
}
