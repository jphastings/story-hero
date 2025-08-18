package storymode

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/jphastings/clone-hero-storymode/pkg/clonehero"
)

func Parse(r io.Reader) (*Story, error) {
	return nil, fmt.Errorf("not yet implemented")
}

func (s *Story) Write(w io.Writer) error {
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.SetEscapeHTML(false)
	return e.Encode(s)
}

func (*Story) Usable() (bool, []string, error) {
	return false, nil, fmt.Errorf("not yet implemented")
}

func (s *Story) SongAvailability(sc *clonehero.SongCache, sd *clonehero.ScoreData) (map[MD5Hash]bool, error) {
	return nil, fmt.Errorf("not yet implemented")
}
