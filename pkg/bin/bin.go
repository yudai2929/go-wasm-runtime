package bin

import (
	"encoding/binary"
	"errors"
)

// LeU32 parses a little-endian uint32 from the byte slice and returns the
func LeU32(b []byte) (remaining []byte, value uint32, err error) {
	if len(b) < 4 {
		return nil, 0, errors.New("byte slice too short to parse uint32")
	}
	value = binary.LittleEndian.Uint32(b[:4])
	remaining = b[4:]
	return
}

// Tag checks if the input byte slice starts with the expected tag and returns the remaining byte slice.
func Tag(input []byte, expectedTag []byte) (remaining []byte, err error) {
	if len(input) < len(expectedTag) {
		return nil, errors.New("input byte slice is too short")
	}
	if !equal(input[:len(expectedTag)], expectedTag) {
		return nil, errors.New("input byte slice does not match the expected tag")
	}
	return input[len(expectedTag):], nil
}

// equal checks if two byte slices are equal.
func equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
