package hash

import (
	"golang.org/x/crypto/bcrypt"
)

// Hasher defines password hashing interface
type Hasher interface {
	Hash(password string) (string, error)
	Compare(password, hash string) error
}

type bcryptHasher struct {
	cost int
}

// NewBcryptHasher creates a new bcrypt hasher
func NewBcryptHasher(cost int) Hasher {
	if cost < bcrypt.MinCost || cost > bcrypt.MaxCost {
		cost = bcrypt.DefaultCost
	}
	return &bcryptHasher{cost: cost}
}

// Hash hashes a password using bcrypt
func (h *bcryptHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Compare compares a password with a hash
func (h *bcryptHasher) Compare(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
