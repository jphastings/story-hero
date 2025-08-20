package clonehero

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/jphastings/story-hero/pkg/types"
)

func readSongID(r io.Reader) (types.MD5Hash, error) {
	var songIDBytes [16]byte
	if _, err := r.Read(songIDBytes[:]); err != nil {
		return "", fmt.Errorf("failed to read song ID: %w", err)
	}
	return types.MD5HashFromBytes(songIDBytes), nil
}

func readUint8(r io.Reader) (uint, error) {
	buf := make([]byte, 1)
	if _, err := r.Read(buf); err != nil {
		return 0, fmt.Errorf("failed to read uint8: %w", err)
	}
	return uint(buf[0]), nil
}

func readUint16(r io.Reader) (uint, error) {
	buf := make([]byte, 2)
	if _, err := r.Read(buf); err != nil {
		return 0, fmt.Errorf("failed to read uint16: %w", err)
	}
	return uint(binary.LittleEndian.Uint16(buf)), nil
}

func readUint24(r io.Reader) (uint, error) {
	buf := make([]byte, 3)
	if _, err := r.Read(buf); err != nil {
		return 0, fmt.Errorf("failed to read uint24: %w", err)
	}
	return uint(buf[0]) | uint(buf[1])<<8 | uint(buf[2])<<16, nil
}

func readUint32(r io.Reader) (uint, error) {
	buf := make([]byte, 4)
	if _, err := r.Read(buf); err != nil {
		return 0, fmt.Errorf("failed to read uint32: %w", err)
	}
	return uint(binary.LittleEndian.Uint32(buf)), nil
}

func readVarUint(r io.Reader) (uint, error) {
	byteValue := uint(128) // Start with max to enter loop, will be overwritten immediately
	u := uint(0)
	shift := 0

	var err error
	for byteValue > 127 {
		byteValue, err = readUint8(r)
		if err != nil {
			return 0, err
		}

		u += (byteValue & 0x7F) << shift
		shift += 7
	}

	return u, nil
}

func readPrefixLengthString(r io.Reader) (string, error) {
	length, err := readVarUint(r)
	if err != nil || length == 0 {
		return "", err
	}

	strBytes := make([]byte, length)
	if _, err := io.ReadFull(r, strBytes); err != nil {
		return "", fmt.Errorf("failed to read string data: %w", err)
	}

	return string(strBytes), nil
}

func skipPrefixLengthString(rs io.ReadSeeker) error {
	length, err := readVarUint(rs)
	if err != nil {
		return err
	}

	_, err = rs.Seek(int64(length), io.SeekCurrent)
	return err
}

func skipBytes(rs io.ReadSeeker, byteCount int64) error {
	if _, err := rs.Seek(byteCount, io.SeekCurrent); err != nil {
		return fmt.Errorf("failed to interact with file: %w", err)
	}
	return nil
}
