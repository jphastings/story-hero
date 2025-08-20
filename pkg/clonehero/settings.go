package clonehero

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

type Settings struct {
	Game struct {
		CacheFile string `ini:"cache_file"`
	} `ini:"game"`

	Directories struct {
		Paths []string
	}
}

func OpenSettings(f *os.File) (*Settings, error) {
	cfg, err := ini.Load(f)
	if err != nil {
		return nil, err
	}

	var si Settings
	if err := cfg.MapTo(&si); err != nil {
		return nil, err
	}

	if si.Game.CacheFile == "" {
		return nil, fmt.Errorf("the settings.ini file doesn't contain a cache_file declaration")
	}

	// Manually collect paths
	dirSection := cfg.Section("directories")

	si.Directories.Paths = []string{configDir("Songs")}
	for _, key := range dirSection.Keys() {
		if !strings.HasPrefix(key.Name(), "path") {
			continue
		}
		if key.String() == "" {
			continue
		}

		si.Directories.Paths = append(si.Directories.Paths, key.String())
	}

	return &si, nil
}
