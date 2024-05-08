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
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	EthereumAPIEndpoint   string
	TaraxaContractAddress string
	TaraxaNodeURL         string
	Key                   string
	LightNodeEndpoint     string
}

func main() {
	// Bind flags to viper
	pflag.String("ethereum_api_endpoint", "", "Ethereum API endpoint")
	pflag.String("taraxa_contract_address", "", "Taraxa contract address")
	pflag.String("taraxa_node_url", "", "Taraxa node URL")
	pflag.String("key", "", "Private key")
	pflag.String("light_node_endpoint", "", "Light node endpoint")

	// Read config from environment variables
	viper.AutomaticEnv()

	// Bind environment variables to viper
	viper.BindEnv("ethereum_api_endpoint", "ETHEREUM_API_ENDPOINT")
	viper.BindEnv("taraxa_contract_address", "TARAXA_CONTRACT_ADDRESS")
	viper.BindEnv("taraxa_node_url", "TARAXA_NODE_URL")
	viper.BindEnv("key", "KEY")
	viper.BindEnv("light_node_endpoint", "LIGHT_NODE_ENDPOINT")

	// Parse flags
	pflag.Parse()

	// Read config from flags
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	privateKey, err := crypto.HexToECDSA(config.Key)
	if err != nil {
		log.Fatalf("Failed to convert private key: %v", err)
	}

	relayer, err := relayer.NewRelayer(&relayer.RelayerConfig{
		BeaconNodeEndpoint: config.EthereumAPIEndpoint,
		TaraxaNodeURL:      config.TaraxaNodeURL,
		TaraxaContractAddr: common.HexToAddress(config.TaraxaContractAddress),
		Key:                privateKey,
		LightNodeEndpoint:  config.LightNodeEndpoint,
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

	relayer.Start(ctx)

	// Keep the main goroutine running until an interrupt is received
	<-ctx.Done()
}
