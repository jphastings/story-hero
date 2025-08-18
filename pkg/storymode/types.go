package storymode

import (
	"encoding/hex"
	"encoding/json"
	"math/big"
)

// Ensure tygo is installed: go install github.com/gzuidhof/tygo/cmd@latest
//go:generate tygo generate

type Story struct {
	Title  string  `json:"title"`
	Groups []Group `json:"groups"`
}

type Group struct {
	Title           string    `json:"title"`
	Songs           []MD5Hash `json:"songs"`
	IsUnlocked      string    `json:"isUnlocked,omitempty" tstype:"UnlockFunc"`
	LockedMessage   string    `json:"lockedMessage,omitempty"`
	ShowLockedSongs bool      `json:"showLockedSongs,omitempty"`
	UnlockAction    string    `json:"unlockAction,omitempty" tstype:"ActionFunc"`
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

type MD5Hash [16]byte

func (m MD5Hash) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(m[:]))
}
