package main

import (
	"crypto/rand"
	"encoding/hex"
)

func randomHexBytes(size int) ([]byte, error) {
	b, err := randomBytes(size)

	if err != nil {
		return b, err
	}
	return []byte(hex.EncodeToString(b)), nil
}

func randomBytes(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return b, err
	}
	return b, nil
}
