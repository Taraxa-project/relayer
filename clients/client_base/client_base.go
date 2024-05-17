package client_base

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Network uint8

const (
	Mainnet = iota
	Testnet
	Devnet
)

type NetConfig struct {
	Url             string         `json:"url"`
	ContractAddress common.Address `json:"contract_address"`
}

type Transactor struct {
	TransactOpts *bind.TransactOpts
	Address      common.Address
	Nonce        uint64
}

type ClientBase struct {
	EthClient  *ethclient.Client
	transactor *Transactor
	Config     NetConfig
}

func NewClientBase(config NetConfig, privateKeyStr *string) (*ClientBase, error) {
	var err error

	clientBase := new(ClientBase)
	clientBase.EthClient, err = ethclient.Dial(config.Url)
	if err != nil {
		return nil, err
	}

	if privateKeyStr != nil {
		transactor, err := clientBase.NewTransactor(*privateKeyStr)
		if err != nil {
			return nil, errors.New("Unable to create transactor")
		}
		clientBase.transactor = transactor
	}

	clientBase.Config = config

	return clientBase, nil
}

func (cb *ClientBase) NewTransactor(privateKeyStr string) (*Transactor, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	chainID, err := cb.EthClient.ChainID(context.Background())
	if err != nil {
		return nil, errors.New("failed to retrieve chain ID")
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	nonce, err := cb.EthClient.PendingNonceAt(context.Background(), address)
	if err != nil {
		return nil, err
	}

	transactor := new(Transactor)
	transactor.Address = address
	transactor.Nonce = nonce
	transactor.TransactOpts = new(bind.TransactOpts)
	*transactor.TransactOpts = *transactOpts

	return transactor, nil
}

func (cb *ClientBase) GenTransactOpts() (*bind.TransactOpts, error) {
	nonce, err := cb.EthClient.PendingNonceAt(context.Background(), cb.transactor.Address)
	if err != nil {
		return nil, err
	}

	gasPrice, err := cb.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	maxNonce := cb.transactor.Nonce
	if nonce > maxNonce {
		maxNonce = nonce
	}

	transactOpts := new(bind.TransactOpts)
	*transactOpts = *cb.transactor.TransactOpts

	transactOpts.Nonce = big.NewInt(int64(maxNonce))
	transactOpts.GasLimit = uint64(300000) // in units
	transactOpts.GasPrice = gasPrice

	// Increment transactos nonce for the next tx
	cb.transactor.Nonce++

	return transactOpts, nil
}
