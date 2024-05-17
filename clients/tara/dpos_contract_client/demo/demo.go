package main

import (
	"log"

	"github.com/Taraxa-project/relayer/clients/client_base"
	dpos_contract_client "github.com/Taraxa-project/relayer/clients/tara/dpos_contract_client"
	"github.com/Taraxa-project/relayer/clients/tara/tara_net_config"
)

func main() {
	log.Print("Dpos client demo")

	config, err := tara_net_config.GenNetConfig(client_base.Mainnet)
	if err != nil {
		log.Fatal(err)
	}

	dposContractClient, err := dpos_contract_client.NewDposContractClient(*config, nil)
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
