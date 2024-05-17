package eth_net_config

import (
	"errors"

	"github.com/Taraxa-project/relayer/clients/client_base"
	"github.com/ethereum/go-ethereum/common"
)

func GenNetConfig(network client_base.Network) (*client_base.NetConfig, error) {
	config := new(client_base.NetConfig)

	switch network {
	case client_base.Mainnet:
		// TODO:
		return nil, errors.New("Mainnet not supported")
		break
	case client_base.Testnet:
		config.Url = "https://holesky.drpc.org"
		config.ContractAddress = common.HexToAddress("0x0")
		break
	case client_base.Devnet:
		// TODO
		return nil, errors.New("Devnet not supported")
		break
	default:
		return nil, errors.New("Invalid network argument")
	}

	return config, nil
}
