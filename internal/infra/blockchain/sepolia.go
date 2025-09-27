package blockchain

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type SourceOfTruth struct {
	auth     *bind.TransactOpts
	instance *Sourceoftruth
}

type ISourceOfTruth interface {
	Insert(hashFile string) (*types.Transaction, error)
	Check(hashFile string) error
}

func New(rpcURL, privateKeyHex, contractAddress string) ISourceOfTruth {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	chainID := big.NewInt(11155111)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	instance, err := NewSourceoftruth(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatal(err)
	}

	return &SourceOfTruth{
		auth:     auth,
		instance: instance,
	}
}

func (s *SourceOfTruth) Insert(hashFile string) (*types.Transaction, error) {
	t, err := s.instance.StoreHash(s.auth, hashFile)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (s *SourceOfTruth) Check(hashFile string) error {
	_, err := s.instance.StoredHashes(&bind.CallOpts{
		Context: context.Background(),
	}, hashFile)

	return err
}
