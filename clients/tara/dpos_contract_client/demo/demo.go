package main

import (
	"log"

	"github.com/Taraxa-project/relayer/clients/client_base"
	dpos_contract_client "github.com/Taraxa-project/relayer/clients/tara/dpos_contract_client"
	"github.com/Taraxa-project/relayer/clients/tara/tara_net_config"
)

func main() {
	log.Print("Dpos client demo")

	// var config client_base.NetConfig
	// config.HttpUrl = "http://localhost:7017"
	// config.ChainID = big.NewInt(649)
	// config.ContractAddress = common.HexToAddress("0x00000000000000000000000000000000000000FE")

	config, err := tara_net_config.GenNetConfig(client_base.Mainnet)
	if err != nil {
		log.Fatal(err)
	}

	dposContractClient, err := dpos_contract_client.NewDposContractClient(*config, client_base.Http)
	if err != nil {
		log.Fatal(err)
	}

	totalEligibleVotesCount, err := dposContractClient.GetTotalEligibleVotesCount()
	if err != nil {
		log.Print("GetTotalEligibleVotesCount err: ", err)
	} else {
		log.Printf("GetTotalEligibleVotesCount: %d\n\n", totalEligibleVotesCount)
	}
}
