package clonehero

import (
	"fmt"
	"io/fs"
)

type SongCache struct{}

func OpenSongCache(f *fs.File) (*SongCache, error) {
	return nil, fmt.Errorf("not yet implemented")
}
