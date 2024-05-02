package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"relayer/relayer"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	ethereumAPIEndpoint   = "https://responsive-weathered-morning.ethereum-holesky.quiknode.pro/4e7ca7b018c76a5ee1041bd2eb9125293d8ec3a1"
	taraxaContractAddress = "0xC77919c1c830FB8176246e547C36546866ae0f92"
	taraxaNodeURL         = "https://rpc-pr-2618.prnet.taraxa.io"
	key                   = "fc6c309495809b69ce77b3250cacfef94d28698d8fb425501a59836fe30fab1d"
	lightNodeEndpoint     = "https://beacon-pr-2618.prnet.taraxa.io"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatalf("Failed to convert private key: %v", err)
	}

	relayer, err := relayer.NewRelayer(&relayer.RelayerConfig{
		BeaconNodeEndpoint: ethereumAPIEndpoint,
		TaraxaNodeURL:      taraxaNodeURL,
		TaraxaContractAddr: common.HexToAddress(taraxaContractAddress),
		Key:                privateKey,
		LightNodeEndpoint:  lightNodeEndpoint,
	})
	if err != nil {
		panic(err)
	}

	// Handle interrupt signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signals
		fmt.Println("\nReceived an interrupt, closing connection...")

		// Perform cleanup
		relayer.Close()

		// Additional cleanup can be done here
		cancel() // Cancel the context to stop any ongoing operations

		fmt.Println("Shutdown complete.")
		os.Exit(0)
	}()

	// relayer.Start(ctx)
	relayer.UpdateLightClient(48806, false)

	// Keep the main goroutine running until an interrupt is received
	<-ctx.Done()
}
