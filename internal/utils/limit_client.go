package utils

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
)

type GasLimitClient struct {
	*ethclient.Client
	limit *big.Int
}

func NewLimitClient(client *ethclient.Client, limit *big.Int) *GasLimitClient {
	return &GasLimitClient{Client: client, limit: limit}
}

// SendTransaction checks if the transaction gas price or gas tip cap + gas fee cap exceeds the set limit
func (c *GasLimitClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	if tx.GasPrice() != nil && tx.GasPrice().Cmp(c.limit) > 0 {
		return fmt.Errorf("gas price(%s) exceeds limit(%s)", tx.GasPrice().String(), c.limit.String())
	}
	if tx.GasTipCap() != nil && tx.GasFeeCap() != nil {
		totalGasPrice := new(big.Int).Add(tx.GasTipCap(), tx.GasFeeCap())
		if totalGasPrice.Cmp(c.limit) > 0 {
			return fmt.Errorf("gas tip cap + gas fee cap(%s) exceeds limit(%s)", totalGasPrice.String(), c.limit.String())
		}
	}
	return c.Client.SendTransaction(ctx, tx)
}

func (c *GasLimitClient) GetGethClient() *gethclient.Client {
	return gethclient.New(c.Client.Client())
}
