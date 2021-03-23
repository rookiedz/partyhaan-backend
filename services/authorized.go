package services

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(password, salt string) (string, error) {
	b := make([]byte, len(salt+password))
	_, err := rand.Read(b[:])
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(sha256.New().Sum(b)), nil
}
func GenerateSalt(size uint64) (string, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(sha256.New().Sum(b)), nil
}
func ComparePassword() (bool, error) {
	return false, nil
}
