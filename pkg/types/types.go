package types

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/dop251/goja"
)

// Ensure tygo is installed: go install github.com/gzuidhof/tygo/cmd@latest
//go:generate tygo generate

type Story struct {
	Title  string  `json:"title"`
	Groups []Group `json:"groups"`
}

type Group struct {
	Title           string                             `json:"title"`
	Songs           []MD5Hash                          `json:"songs"`
	IsUnlocked      func(goja.FunctionCall) goja.Value `json:"isUnlocked,omitempty" tstype:"UnlockFunc"`
	LockedMessage   string                             `json:"lockedMessage,omitempty"`
	ShowLockedSongs bool                               `json:"showLockedSongs,omitempty"`
	UnlockAction    func(goja.FunctionCall) goja.Value `json:"unlockAction,omitempty" tstype:"ActionFunc"`
}

type SongPlay struct {
	ID        MD5Hash        `json:"id"`
	PlayCount uint           `json:"playCount"`
	Scores    map[uint]Score `json:"scores" tstype:"PerInstrumentScore"`
}

type Score struct {
	Difficulty uint     `json:"difficulty" tstype:"Difficulty"`
	Percentage *big.Rat `json:"percentage" tstype:"number"`
	Stars      uint     `json:"stars"`
	Score      uint     `json:"score"`
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
