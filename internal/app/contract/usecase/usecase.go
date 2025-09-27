package usecase

import (
	"github.com/Clausia-Ifest/clausia-server/internal/domain/contract"
	"github.com/Clausia-Ifest/clausia-server/internal/infra/blockchain"
	"github.com/Clausia-Ifest/clausia-server/internal/infra/storage"
	"github.com/Clausia-Ifest/clausia-server/pkg/transactor"
	clausiapb "github.com/Clausia-Ifest/clausia-server/proto"
)

type UContract struct {
	tx         transactor.ITransactor
	s3         storage.IS3
	grpc       clausiapb.ClausIAClient
	blockchain blockchain.ISourceOfTruth
	rc         contract.IRContract
	rd         contract.IRDocument
}

func New(tx transactor.ITransactor, s3 storage.IS3, grpc clausiapb.ClausIAClient, blockchain blockchain.ISourceOfTruth, rc contract.IRContract, rd contract.IRDocument) contract.IUContract {
	return &UContract{
		tx:         tx,
		s3:         s3,
		grpc:       grpc,
		blockchain: blockchain,
		rc:         rc,
		rd:         rd,
	}
}
