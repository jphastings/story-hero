package clonehero

import (
	"os"
)

const (
	scoreDataFile = "scoredata.bin"
	settingsFile  = "settings.ini"
)

func Load() (*Settings, *SongCache, *ScoreData, error) {
	sif, err := os.Open(configDir(settingsFile))
	if err != nil {
		return nil, nil, nil, err
	}
	si, err := OpenSettings(sif)
	if err != nil {
		return nil, nil, nil, err
	}

	sc, err := OpenSongCache(SupportDirPath(si.Game.CacheFile))
	if err != nil {
		return nil, nil, nil, err
	}

	sd, err := OpenScoreData(SupportDirPath(scoreDataFile))
	if err != nil {
		return nil, nil, nil, err
	}

	return si, sc, sd, nil
}
