package hash

import "golang.org/x/crypto/bcrypt"

type Bcrypt struct {
	//
}

type IBcrypt interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) error
}

func New() IBcrypt {
	return &Bcrypt{}
}

func (b *Bcrypt) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (b *Bcrypt) ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
