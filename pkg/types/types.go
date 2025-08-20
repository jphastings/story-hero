package types

import (
	"encoding/hex"
	"fmt"

	"github.com/dop251/goja"
)

// Ensure tygo is installed: go install github.com/gzuidhof/tygo/cmd@latest
//go:generate tygo generate

type Story struct {
	Title  string  `mapstructure:"title"`
	Groups []Group `mapstructure:"groups"`
}

type Group struct {
	Title           string                             `mapstructure:"title"`
	Songs           []MD5Hash                          `mapstructure:"songs"`
	IsUnlocked      func(goja.FunctionCall) goja.Value `mapstructure:"isUnlocked,omitempty" tstype:"UnlockFunc"`
	LockedMessage   string                             `mapstructure:"lockedMessage,omitempty"`
	ShowLockedSongs bool                               `mapstructure:"showLockedSongs,omitempty"`
	UnlockAction    func(goja.FunctionCall) goja.Value `mapstructure:"unlockAction,omitempty" tstype:"ActionFunc"`
}

type SongPlay struct {
	ID        MD5Hash        `mapstructure:"id"`
	PlayCount uint           `mapstructure:"playCount"`
	Scores    map[uint]Score `mapstructure:"scores" tstype:"PerInstrumentScore"`
}

type Score struct {
	Difficulty uint `mapstructure:"difficulty" tstype:"Difficulty"`
	Percentage uint `mapstructure:"percentage" tstype:"number"`
	Speed      uint `mapstructure:"speed" tstype:"number"`
	Stars      uint `mapstructure:"stars"`
	Score      uint `mapstructure:"score"`
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
