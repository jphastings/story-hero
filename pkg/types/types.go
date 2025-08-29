package types

import (
	"encoding/hex"
	"fmt"

	"github.com/dop251/goja"
)

// Ensure tygo is installed: go install github.com/gzuidhof/tygo/cmd@latest
//go:generate tygo generate

type Story struct {
	Title  string  `json:"title" mapstructure:"title"`
	Groups []Group `json:"groups" mapstructure:"groups"`
}

type Group struct {
	Title           string                             `json:"title" mapstructure:"title"`
	Songs           []MD5Hash                          `json:"songs" mapstructure:"songs"`
	IsUnlocked      func(goja.FunctionCall) goja.Value `json:"isUnlocked,omitempty" mapstructure:"isUnlocked,omitempty" tstype:"UnlockFunc"`
	LockedMessage   string                             `json:"lockedMessage,omitempty" mapstructure:"lockedMessage,omitempty"`
	ShowLockedSongs bool                               `json:"showLockedSongs,omitempty" mapstructure:"showLockedSongs,omitempty"`
	UnlockAction    func(goja.FunctionCall) goja.Value `json:"unlockAction,omitempty" mapstructure:"unlockAction,omitempty" tstype:"ActionFunc"`
}

type SongPlay struct {
	ID        MD5Hash        `json:"id" mapstructure:"id"`
	PlayCount uint           `json:"playCount" mapstructure:"playCount"`
	Scores    map[uint]Score `json:"scores" mapstructure:"scores" tstype:"PerInstrumentScore"`
}

type Score struct {
	Difficulty uint `json:"difficulty" mapstructure:"difficulty" tstype:"Difficulty"`
	Percentage uint `json:"percentage" mapstructure:"percentage" tstype:"number"`
	Speed      uint `json:"speed" mapstructure:"speed" tstype:"number"`
	Stars      uint `json:"stars" mapstructure:"stars"`
	Score      uint `json:"score" mapstructure:"score"`
}

type MD5Hash string

func (m MD5Hash) ToBytes() ([]byte, error) {
	b, err := hex.DecodeString(string(m))
	if err != nil {
		return nil, err
	}
	if len(b) != 16 {
		return nil, fmt.Errorf("songID isn't the correct length")
	}
	return b, nil
}

func MD5HashFromBytes(id [16]byte) MD5Hash {
	return MD5Hash(hex.EncodeToString(id[:]))
}

type Song struct {
	ID     MD5Hash
	Path   string
	Title  string
	Artist string
}
