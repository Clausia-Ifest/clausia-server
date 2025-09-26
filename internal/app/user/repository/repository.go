package repository

import (
	"github.com/Clausia-Ifest/clausia-server/internal/domain/contract"
	querybuilder "github.com/Clausia-Ifest/clausia-server/pkg/query_builder"
)

type RUser struct {
	qb querybuilder.IQB
}

func New(qb querybuilder.IQB) contract.IRUser {
	return &RUser{
		qb: qb,
	}
}
