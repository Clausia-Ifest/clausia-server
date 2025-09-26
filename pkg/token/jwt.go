package token

import (
	"errors"
	"fmt"
	"time"

	wib "github.com/Clausia-Ifest/clausia-server/pkg/time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
	Role     string    `json:"role"`
	Email    string    `json:"email"`
}

type Claims struct {
	User User `json:"user"`
	jwt.RegisteredClaims
}

type JWT struct {
	sk []byte
	d  int
}

type IJWT interface {
	Encode(id uuid.UUID, fullName, role, email string) (string, error)
	Decode(signedToken string) (*Claims, error)
}

func New(secretKey string, duration int) *JWT {
	return &JWT{
		sk: []byte(secretKey),
		d:  duration,
	}
}

func (j *JWT) Encode(id uuid.UUID, fullName, role, email string) (string, error) {
	claims := &Claims{
		User: User{
			ID:       id,
			FullName: fullName,
			Email:    email,
			Role:     role,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   fmt.Sprintf("%s:%s", role, id),
			Issuer:    "clausia-ifest",
			IssuedAt:  jwt.NewNumericDate(wib.Now()),
			ExpiresAt: jwt.NewNumericDate(wib.Now().Add(time.Duration(j.d) * time.Minute)),
		},
	}

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := unsignedToken.SignedString(j.sk)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *JWT) Decode(signedToken string) (*Claims, error) {
	decoded, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(_ *jwt.Token) (any, error) {
		return j.sk, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return &Claims{}, errors.New("token is expired")
		}

		return &Claims{}, errors.New("invalid signature")
	}

	if !decoded.Valid {
		return &Claims{}, errors.New("invalid token")
	}

	claims, ok := decoded.Claims.(*Claims)
	if !ok {
		return &Claims{}, errors.New("error parse claims")
	}

	return claims, nil
}
