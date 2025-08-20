package clonehero

import (
	"fmt"
	"io"
	"os"

	"github.com/jphastings/clone-hero-storymode/pkg/types"
)

type SongCache struct {
	f            *os.File
	lookupTables map[lookupAttr][]string
	Songs        map[types.MD5Hash]*types.Song
}

type lookupAttr int

const (
	lookupAttrTitle  lookupAttr = 0
	lookupAttrArtist lookupAttr = 1
	// Others not needed at this time
)

func OpenSongCache(f *os.File) (*SongCache, error) {
	sc := &SongCache{f: f}
	if err := sc.Reload(); err != nil {
		return nil, err
	}
	return sc, nil
}

func (sc *SongCache) Reload() error {
	_, err := sc.f.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("unable to interact with the SongCache file: %w", err)
	}

	// Skip header (20 bytes)
	if _, err := sc.f.Seek(20, io.SeekCurrent); err != nil {
		return fmt.Errorf("failed to skip header: %w", err)
	}

	if err := sc.readLookupTables(); err != nil {
		return err
	}

	songCount, err := readUint32(sc.f)
	if err != nil {
		return err
	}

	sc.Songs = make(map[types.MD5Hash]*types.Song)

	for i := uint(0); i < songCount; i++ {
		song, err := sc.readSong()
		if err != nil {
			return err
		}

		sc.Songs[song.ID] = song
	}

	return nil
}

func (sc *SongCache) readLookupTable() (lookupAttr, []string, error) {
	tableType, err := readUint8(sc.f)
	if err != nil {
		return 0, nil, err
	}

	entryCount, err := readUint32(sc.f)
	if err != nil {
		return 0, nil, err
	}

	table := make([]string, entryCount)
	for i := uint(0); i < entryCount; i++ {
		entry, err := readPrefixLengthString(sc.f)
		if err != nil {
			return 0, nil, err
		}

		table[i] = entry
	}

	return lookupAttr(tableType), table, nil
}

const countLookupTables = 7

func (sc *SongCache) readLookupTables() error {
	sc.lookupTables = make(map[lookupAttr][]string)

	for i := 0; i < countLookupTables; i++ {
		tType, table, err := sc.readLookupTable()
		if err != nil {
			return err
		}
		sc.lookupTables[tType] = table
	}

	return nil
}

func (sc *SongCache) readSong() (*types.Song, error) {
	path, err := readPrefixLengthString(sc.f)
	if err != nil {
		return nil, err
	}

	// Skip unknown (checksum?) (16 bytes)
	if err := skipBytes(sc.f, 16); err != nil {
		return nil, err
	}

	// Skip chart
	if err := skipPrefixLengthString(sc.f); err != nil {
		return nil, err
	}

	// Skip Unknown-a
	if err := skipBytes(sc.f, 1); err != nil {
		return nil, err
	}

	title, err := sc.readLookupValue(lookupAttrTitle)
	if err != nil {
		return nil, err
	}

	artist, err := sc.readLookupValue(lookupAttrArtist)
	if err != nil {
		return nil, err
	}

	// Skip album, genre, year, charter, playlist offsets (4 bytes each = 20 bytes)
	// Skip unknown-b (8 bytes)
	// Skip instrument difficulty (13 bytes)
	// Skip preview offset (4 bytes)
	if err := skipBytes(sc.f, 45); err != nil {
		return nil, err
	}

	// Skip icon (varint string)
	if err := skipPrefixLengthString(sc.f); err != nil {
		return nil, err
	}

	// Skip unknown-c (8 bytes)
	// Skip song length (4 bytes)
	// Skip unknown-d (8 bytes)
	if err := skipBytes(sc.f, 20); err != nil {
		return nil, err
	}

	// Skip game name (varint string)
	if err := skipPrefixLengthString(sc.f); err != nil {
		return nil, err
	}

	// Skip delimiter (1 byte)
	if err := skipBytes(sc.f, 1); err != nil {
		return nil, err
	}

	// Read chart MD5 (16 bytes)
	songID, err := readSongID(sc.f)
	if err != nil {
		return nil, err
	}

	return &types.Song{
		ID:     songID,
		Path:   path,
		Title:  title,
		Artist: artist,
	}, nil
}

func (sc *SongCache) readLookupValue(la lookupAttr) (string, error) {
	offset, err := readUint32(sc.f)
	if err != nil {
		return "", err
	}
	table := sc.lookupTables[la]
	if len(table) < int(offset) {
		return "", fmt.Errorf("offset (%d) is larger than table (%d: %d)", offset, la, len(table))
	}
	return table[offset], nil
}
