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
	ethereumAPIEndpoint   = "https://restless-spring-uranium.quiknode.pro/31384766eed9e4977424797ae8b020ee6dd6b9eb"
	taraxaContractAddress = "0x0DC0d841F962759DA25547c686fa440cF6C28C61"
	taraxaNodeURL         = "https://rpc.devnet.taraxa.io"
	key                   = "fc6c309495809b69ce77b3250cacfef94d28698d8fb425501a59836fe30fab1d"
	// lightNodeEndpoint = "http://unstable.holesky.beacon-api.nimbus.team/" // Holesky
	lightNodeEndpoint = "https://www.lightclientdata.org"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatalf("Failed to convert private key: %v", err)
	}

	relayer, err := relayer.NewRelayer(&relayer.RelayerConfig{
		ethereumAPIEndpoint,
		taraxaNodeURL,
		common.HexToAddress(taraxaContractAddress),
		privateKey,
		lightNodeEndpoint})
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

	relayer.Start(ctx)
	// relayer.UpdateNewHeader(8826112)

	// Keep the main goroutine running until an interrupt is received
	<-ctx.Done()
}
