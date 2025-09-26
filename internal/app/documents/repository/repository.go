package repository

import (
	"github.com/Clausia-Ifest/clausia-server/internal/domain/contract"
	querybuilder "github.com/Clausia-Ifest/clausia-server/pkg/query_builder"
)

type RDocument struct {
	qb querybuilder.IQB
}

func New(qb querybuilder.IQB) contract.IRDocument {
	return &RDocument{
		qb: qb,
	}
}
