package main

import (
	"log"
	"math/big"

	"github.com/Taraxa-project/relayer/clients/bridge_contract_client"
	"github.com/Taraxa-project/relayer/clients/client_base"
	"github.com/Taraxa-project/relayer/clients/eth/eth_net_config"
	"github.com/ethereum/go-ethereum/common"
)

//	"tara_client": {
//		"http_url": "https://rpc-pr-2618.prnet.taraxa.io",
//		"ws_url": "wss://ws-pr-2618.prnet.taraxa.io/",
//		"chain_id": 200,
//		"bridge_contract_address": "0xFBC597EEf68722E05bbC1e52264103b416551dFB"
//	},
//
//	"eth_client": {
//		"http_url": "https://holesky.drpc.org",
//		"chain_id": 17000,
//		"tara_client_contract_address": "0x515d5e39a9FfF8dBBD84C8064ea3Bc4ad2610442",
//		"bridge_contract_address": "0x762dA247D9F269b1689d4baaD587243eccF7910c"
//	}
func main() {
	log.Print("Bridge contract client demo")

	config, err := eth_net_config.GenNetConfig(client_base.Testnet)
	if err != nil {
		log.Fatal(err)
	}
	config.ChainID = big.NewInt(200)
	config.HttpUrl = "https://rpc-pr-2618.prnet.taraxa.io"
	config.WsUrl = "wss://ws-pr-2618.prnet.taraxa.io/"

	config.ContractAddress = common.HexToAddress("0xFBC597EEf68722E05bbC1e52264103b416551dFB")

	bridgeContractClient, err := bridge_contract_client.NewBridgeContractClient(*config, client_base.WebSocket)
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
