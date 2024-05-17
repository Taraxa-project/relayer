package tara_net_config

import (
	"errors"

	"github.com/Taraxa-project/relayer/clients/client_base"
	"github.com/ethereum/go-ethereum/common"
)

func GenNetConfig(network client_base.Network) (*client_base.NetConfig, error) {
	config := new(client_base.NetConfig)

	switch network {
	case client_base.Mainnet:
		config.Url = "https://rpc.mainnet.taraxa.io"
		config.ContractAddress = common.HexToAddress("0x00000000000000000000000000000000000000FE")
		break
	case client_base.Testnet:
		config.Url = "https://rpc.testnet.taraxa.io"
		config.ContractAddress = common.HexToAddress("0x00000000000000000000000000000000000000FE")
		break
	case client_base.Devnet:
		config.Url = "https://rpc.devnet.taraxa.io"
		config.ContractAddress = common.HexToAddress("0x00000000000000000000000000000000000000FE")
		break
	default:
		return nil, errors.New("Invalid network argument")
	}

	return config, nil
}
