package main

import (
	"log"

	"github.com/Taraxa-project/relayer/clients/client_base"
	"github.com/Taraxa-project/relayer/clients/eth/eth_net_config"
	tara_client_contract_client "github.com/Taraxa-project/relayer/clients/eth/tara_client_contract_client"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	log.Print("Tara client demo")

	config, err := eth_net_config.GenNetConfig(client_base.Testnet)
	if err != nil {
		log.Fatal(err)
	}
	config.ContractAddress = common.HexToAddress("0x402908C007aAC2fAf83D57945ff95cF2de49b359")

	taraClientContractClient, err := tara_client_contract_client.NewTaraClientContractClient(*config, client_base.Http)
	if err != nil {
		log.Fatal(err)
	}

	finalizedPillarBlock, err := taraClientContractClient.GetFinalizedPillarBlock()
	if err != nil {
		log.Print("GetFinalizedPillarBlock err: ", err)
	} else {
		log.Printf("GetFinalizedPillarBlock: %d\n\n", finalizedPillarBlock)
	}

	// // TODO: add here valid private key
	// privateKey := ""

	// transactor, err := taraClientContractClient.NewTransactor(privateKey)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// pillarBlockData := []byte("0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000b00000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000021000000000000000000000000000000000000000000000000000000000000002c00000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000001ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000003fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000005fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb000000000000000000000000290decd9548b62a8d60345a988386fc84ba6bc9500000000000000000000000000000000000000000000000000000000486d7a74000000000000000000000000b10e2d527612073b26eecdfd717e6a320cf44b4afffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe493f000000000000000000000000405787fa12a823e0f2b7631cc41b3ba8828b3321ffffffffffffffffffffffffffffffffffffffffffffffffffffffffaf53b57e000000000000000000000000c2575a0e9e593c00f959f8c92f12db2869c3395a000000000000000000000000000000000000000000000000000000003b77acd10000000000000000000000008a35acfbc15ff81a39ae7d344fd709f28e8600b4000000000000000000000000000000000000000000000000000000001bc4b73e")
	// tx, err := taraClientContractClient.FinalizeBlocks(transactor, pillarBlockData)
	// if err != nil {
	// 	log.Print("AddPendingBlock err: ", err)
	// } else {
	// 	log.Println("AddPendingBlock: ", tx)
	// }
}
