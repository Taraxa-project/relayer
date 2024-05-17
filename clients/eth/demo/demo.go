package main

import (
	"log"

	"github.com/Taraxa-project/relayer/clients/client_base"
	eth_client "github.com/Taraxa-project/relayer/clients/eth"
	"github.com/Taraxa-project/relayer/clients/eth/eth_net_config"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	log.Print("Eth client demo")

	netConfig, err := eth_net_config.GenNetConfig(client_base.Testnet)
	if err != nil {
		log.Fatal(err)
	}

	var config eth_client.EthClientConfig
	config.NetConfig = *netConfig
	config.BridgeContractAddress = common.HexToAddress("0x0")
	config.TaraClientContractAddress = common.HexToAddress("0x0")

	ethClient, err := eth_client.NewEthClient(config, client_base.Http, nil)
	if err != nil {
		log.Fatal(err)
	}

	pendingPillarBlock, err := ethClient.TaraClientContractClient.GetFinalizedPillarBlock()
	if err != nil {
		log.Print("GetFinalizedPillarBlock err: ", err)
	} else {
		log.Printf("GetFinalizedPillarBlock: %d\n\n", pendingPillarBlock)
	}

	stateWithProof, err := ethClient.BridgeContractClient.GetStateWithProof()
	if err != nil {
		log.Print("GetStateWithProof err: ", err)
	} else {
		log.Printf("GetStateWithProof: %d\n\n", stateWithProof)
	}
}
