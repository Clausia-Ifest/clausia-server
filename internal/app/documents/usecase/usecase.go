package usecase

import (
	"github.com/Clausia-Ifest/clausia-server/internal/domain/contract"
	"github.com/Clausia-Ifest/clausia-server/internal/infra/storage"
	"github.com/Clausia-Ifest/clausia-server/pkg/transactor"
	clausiapb "github.com/Clausia-Ifest/clausia-server/proto"
)

type UDocument struct {
	tx   transactor.ITransactor
	s3   storage.IS3
	grpc clausiapb.ClausIAClient
	rd   contract.IRDocument
}

func New(tx transactor.ITransactor, s3 storage.IS3, grpc clausiapb.ClausIAClient, rd contract.IRDocument) contract.IUDocument {
	return &UDocument{
		tx:   tx,
		s3:   s3,
		grpc: grpc,
		rd:   rd,
	}
}
