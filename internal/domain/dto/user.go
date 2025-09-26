package dto

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
}

type GetUserParams struct {
	ID    uuid.UUID `db:"id"`
	Email string    `db:"email"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"email,min=5,max=60"`
	Password string `json:"password" validate:"min=8,max=255"`
}

type SignInResponse struct {
	AccessToken string `json:"access_token"`
	User        *User  `json:"user"`
}
