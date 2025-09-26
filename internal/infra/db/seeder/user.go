package seeder

import (
	"time"

	"github.com/Clausia-Ifest/clausia-server/internal/domain/entity"
	"github.com/Clausia-Ifest/clausia-server/internal/domain/enum"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var hashedPassword, _ = bcrypt.GenerateFromPassword([]byte("123123123"), bcrypt.DefaultCost)

var Users = []entity.User{
	{
		ID:        uuid.MustParse("01997fb9-c4d5-754b-a0d6-e622347f66f6"),
		FullName:  "Admin ILCS",
		Email:     "admin@ilcs.co.id",
		Password:  string(hashedPassword),
		Role:      enum.RoleAdmin,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        uuid.MustParse("01997fb9-c4d5-7e41-bce7-923260bfa945"),
		FullName:  "Legal ILCS",
		Email:     "legal@ilcs.co.id",
		Password:  string(hashedPassword),
		Role:      enum.RoleLegal,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        uuid.MustParse("01997fb9-c4d5-7a31-9853-29f91ed2810a"),
		FullName:  "Manager ILCS",
		Email:     "manager@ilcs.co.id",
		Password:  string(hashedPassword),
		Role:      enum.RoleManager,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
