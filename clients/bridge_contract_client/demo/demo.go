package main

import (
	"log"

	"github.com/Taraxa-project/relayer/clients/bridge_contract_client"
	"github.com/Taraxa-project/relayer/clients/client_base"
	"github.com/Taraxa-project/relayer/clients/eth/eth_net_config"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	log.Print("Bridge contract client demo")

	config, err := eth_net_config.GenNetConfig(client_base.Testnet)
	if err != nil {
		log.Fatal(err)
	}
	//config.Url = "https://rpc-pr-2618.prnet.taraxa.io"
	config.Url = "wss://ws-pr-2618.prnet.taraxa.io/"

	config.ContractAddress = common.HexToAddress("0xFBC597EEf68722E05bbC1e52264103b416551dFB")

	bridgeContractClient, err := bridge_contract_client.NewBridgeContractClient(*config, nil)
	if err != nil {
		log.Fatal(err)
	}

	stateWithProof, err := bridgeContractClient.GetStateWithProof()
	if err != nil {
		log.Print("GetStateWithProof err: ", err)
	} else {
		log.Printf("GetStateWithProof: %d\n\n", stateWithProof)
	}
}
