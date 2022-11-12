package hasher

import "golang.org/x/crypto/bcrypt"

type Hasher struct {
	hashCost int
}

func NewHasher(hashCost int) *Hasher {
	return &Hasher{hashCost}
}

func (h *Hasher) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.hashCost)
	return string(bytes), err
}

func (h *Hasher) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
