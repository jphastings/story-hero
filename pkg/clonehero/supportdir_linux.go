package clonehero

import (
	"path/filepath"
)

func configDir(filename string) string {
	return filepath.Join(homedir, ".clonehero", filename)
}

func SupportDirPath(filename string) string {
	return filepath.Join(
		homedir,
		".config",
		"unity3d",
		"srylain Inc_",
		"Clone Hero",
		filename,
	)
}
