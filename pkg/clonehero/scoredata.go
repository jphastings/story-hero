package clonehero

import (
	"fmt"
	"io/fs"
)

type ScoreData struct {
}

func OpenScoreData(f *fs.File) (*ScoreData, error) {
	return nil, fmt.Errorf("not yet implemented")
}
