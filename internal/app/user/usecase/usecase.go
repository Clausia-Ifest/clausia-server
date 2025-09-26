package usecase

import (
	"github.com/Clausia-Ifest/clausia-server/internal/domain/contract"
	"github.com/Clausia-Ifest/clausia-server/pkg/hash"
	jwt "github.com/Clausia-Ifest/clausia-server/pkg/token"
	"github.com/Clausia-Ifest/clausia-server/pkg/transactor"
)

type UUser struct {
	tx     transactor.ITransactor
	bcrypt hash.IBcrypt
	token  jwt.IJWT
	ru     contract.IRUser
}

func New(tx transactor.ITransactor, bcrypt hash.IBcrypt, token jwt.IJWT, ru contract.IRUser) contract.IUUser {
	return &UUser{
		tx:     tx,
		bcrypt: bcrypt,
		token:  token,
		ru:     ru,
	}
}
