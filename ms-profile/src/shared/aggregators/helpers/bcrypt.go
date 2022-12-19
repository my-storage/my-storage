package helpers

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/my-storage/ms-profile/src/shared/protocols/cryptography"
)

type BcryptAdapter struct {
	cryptography.Hashing
}

func (bc *BcryptAdapter) CreateHash(payload string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(payload), 14)
	return string(bytes), err
}

func (bc *BcryptAdapter) CompareHash(payload string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(payload))
	return err == nil
}
