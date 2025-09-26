package entity

import (
	"time"

	"github.com/Clausia-Ifest/clausia-server/internal/domain/dto"
	"github.com/Clausia-Ifest/clausia-server/internal/domain/enum"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id"`
	FullName  string    `db:"full_name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      enum.Role `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	TotalData int64 `db:"total_data"`
}

func (e *User) ParseDTO() *dto.User {
	return &dto.User{
		ID:       e.ID,
		FullName: e.FullName,
		Email:    e.Email,
		Role:     e.Role.String(),
	}
}
