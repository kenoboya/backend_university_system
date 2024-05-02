package hash

import (
	"fmt"

	"golang.org/x/crypto/sha3"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type SHA256Hasher struct {
	salt string
}

func NewSHA256Hasher(salt string) *SHA256Hasher {
	return &SHA256Hasher{salt}
}

func (h *SHA256Hasher) Hash(password string) (string, error) {
	hash := sha3.New256()
	if _, err := hash.Write([]byte(password)); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum([]byte(h.salt))), nil
}
