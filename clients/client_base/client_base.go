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

type CommunicationProtocol uint8

const (
	Http = iota
	WebSocket
)

type NetConfig struct {
	HttpUrl         string         `json:"http_url"`
	WsUrl           string         `json:"ws_url"`
	ChainID         *big.Int       `json:"chain_id"`
	ContractAddress common.Address `json:"contract_address"`
}

type Transactor struct {
	TransactOpts *bind.TransactOpts
	Address      common.Address
	Nonce        uint64
}

type ClientBase struct {
	EthClient *ethclient.Client
	Config    NetConfig
}

func NewClientBase(config NetConfig, communicationProtocol CommunicationProtocol) (*ClientBase, error) {
	var err error
	var networkUrl string

	switch communicationProtocol {
	case Http:
		if config.HttpUrl == "" {
			return nil, errors.New("config.HttpUrl not configured")
		}
		networkUrl = config.HttpUrl
		break
	case WebSocket:
		if config.WsUrl == "" {
			return nil, errors.New("config.WsUrl not configured")
		}
		networkUrl = config.WsUrl
		break
	default:
		return nil, errors.New("invalid communicationProtocol argument")
	}

	clientBase := new(ClientBase)
	clientBase.EthClient, err = ethclient.Dial(networkUrl)
	if err != nil {
		return nil, err
	}
	clientBase.Config = config

	return clientBase, nil
}

func (ClientBase *ClientBase) NewTransactor(privateKeyStr string) (*Transactor, error) {
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

	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, ClientBase.Config.ChainID)
	if err != nil {
		return nil, err
	}

	nonce, err := ClientBase.EthClient.PendingNonceAt(context.Background(), address)
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

func (ClientBase *ClientBase) CreateNewTransactOpts(transactor *Transactor) (*bind.TransactOpts, error) {
	nonce, err := ClientBase.EthClient.PendingNonceAt(context.Background(), transactor.Address)
	if err != nil {
		return nil, err
	}

	gasPrice, err := ClientBase.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	maxNonce := transactor.Nonce
	if nonce > maxNonce {
		maxNonce = nonce
	}

	transactOpts := new(bind.TransactOpts)
	*transactOpts = *transactor.TransactOpts

	transactOpts.Nonce = big.NewInt(int64(maxNonce))
	transactOpts.GasLimit = uint64(300000) // in units
	transactOpts.GasPrice = gasPrice

	// Increment transactos nonce for the next tx
	transactor.Nonce++

	return transactOpts, nil
}
