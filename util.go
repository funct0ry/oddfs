package main

import (
	"crypto/rand"
	"encoding/hex"
)

func randomHexBytes(size uint64) ([]byte, error) {
	b, err := randomBytes(size)

	if err != nil {
		return nil, err
	}
	return []byte(hex.EncodeToString(b))[0:size], nil
}

func randomBytes(size uint64) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
