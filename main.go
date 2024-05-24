package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"relayer/internal/logging"
	"relayer/internal/to_eth"
	"relayer/internal/to_tara"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"github.com/spf13/pflag"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	EthereumAPIEndpoint string

	BeaconLightClientAddress string
	EthClientOnTaraAddress   string
	TaraBridgeAddress        string

	TaraClientOnEthAddress string
	EthBridgeAddress       string

	TaraxaNodeURL      string
	PrivateKey         string
	BeaconNodeEndpoint string
}

func main() {
	var config Config

	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Warn("Error loading .env file")
	}
	var log_level string
	pflag.StringVar(&config.EthereumAPIEndpoint, "ethereum_api_endpoint", os.Getenv("ETHEREUM_API_ENDPOINT"), "Ethereum API endpoint")
	pflag.StringVar(&config.BeaconLightClientAddress, "beacon_light_client_address", os.Getenv("BEACON_LIGHT_CLIENT_ADDRESS"), "Address of the BeaconLightClient contract on Taraxa chain")
	pflag.StringVar(&config.EthClientOnTaraAddress, "eth_client_on_tara_address", os.Getenv("ETH_CLIENT_ON_TARA_ADDRESS"), "Address of the EthClient contract on Taraxa chain")
	pflag.StringVar(&config.TaraBridgeAddress, "tara_bridge_address", os.Getenv("TARA_BRIDGE_ADDRESS"), "Address of the Tara bridge contract on Taraxa chain")
	pflag.StringVar(&config.TaraClientOnEthAddress, "tara_client_on_eth_address", os.Getenv("TARA_CLIENT_ON_ETH_ADDRESS"), "Address of the TaraClient contract on Ethereum chain")
	pflag.StringVar(&config.EthBridgeAddress, "eth_bridge_address", os.Getenv("ETH_BRIDGE_ADDRESS"), "Address of the Eth bridge contract on Ethereum chain")
	pflag.StringVar(&config.TaraxaNodeURL, "taraxa_node_url", os.Getenv("TARAXA_NODE_URL"), "Taraxa node URL")
	pflag.StringVar(&config.PrivateKey, "private_key", os.Getenv("PRIVATE_KEY"), "Private key")
	pflag.StringVar(&config.BeaconNodeEndpoint, "beacon_node_endpoint", os.Getenv("BEACON_NODE_ENDPOINT"), "Beacon node endpoint")
	log_level_env := os.Getenv("LOG_LEVEL")
	if log_level_env == "" {
		log_level_env = "debug"
	}
	pflag.StringVar(&log_level, "log_level", log_level_env, "log level. could be only [trace, debug, info, warn, error, fatal]")
	pflag.Parse()

	data_dir := "./"
	logging.Config(filepath.Join(data_dir, "logs"), log_level)

	log.Printf("Starting relayer with config: %+v", config)

	ctx, cancel := context.WithCancel(context.Background())

	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.Fatalf("Failed to convert private key: %v", err)
	}

	taraRelayer, err := to_tara.NewRelayer(&to_tara.Config{
		BeaconNodeEndpoint:    config.BeaconNodeEndpoint,
		EthRPCURL:             config.EthereumAPIEndpoint,
		TaraxaRPCURL:          config.TaraxaNodeURL,
		BeaconLightClientAddr: common.HexToAddress(config.BeaconLightClientAddress),
		EthBridgeAddr:         common.HexToAddress(config.EthBridgeAddress),
		TaraxaBridgeAddr:      common.HexToAddress(config.TaraBridgeAddress),
		EthClientOnTaraAddr:   common.HexToAddress(config.EthClientOnTaraAddress),
		Key:                   privateKey,
	})
	if err != nil {
		panic(err)
	}

	ethRelayer, err := to_eth.NewRelayer(&to_eth.Config{
		TaraxaRPCURL:          config.TaraxaNodeURL,
		EthRPCURL:             config.EthereumAPIEndpoint,
		TaraxaClientOnEthAddr: common.HexToAddress(config.TaraClientOnEthAddress),
		TaraxaBridgeAddr:      common.HexToAddress(config.TaraBridgeAddress),
		EthBridgeAddr:         common.HexToAddress(config.EthBridgeAddress),
		Key:                   privateKey,
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
		taraRelayer.Close()
		ethRelayer.Close()

		// Additional cleanup can be done here
		cancel() // Cancel the context to stop any ongoing operations

		fmt.Println("Shutdown complete.")
		os.Exit(0)
	}()

	taraRelayer.Start(ctx)
	ethRelayer.Start(ctx)
	// Keep the main goroutine running until an interrupt is received
	<-ctx.Done()
}
