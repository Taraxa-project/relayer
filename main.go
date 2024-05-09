package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"relayer/to_eth"
	"relayer/to_tara"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	EthereumAPIEndpoint string

	BeaconLightClientAddress string
	EthClientOnTaraAddress   string
	TaraBridgeAddress        string

	TaraClientOnEthAddress string
	EthBridgeAddress       string

	TaraxaNodeURL     string
	PrivateKey        string
	LightNodeEndpoint string
}

func main() {
	var config Config

	// Bind flags to viper
	pflag.StringVar(&config.EthereumAPIEndpoint, "ethereum_api_endpoint", "", "Ethereum API endpoint")
	pflag.StringVar(&config.BeaconLightClientAddress, "beacon_light_client_address", "", "Address of the BeaconLightClient contract on Taraxa chain")
	pflag.StringVar(&config.EthClientOnTaraAddress, "eth_client_on_tara_address", "", "Address of the EthClient contract on Taraxa chain")
	pflag.StringVar(&config.TaraBridgeAddress, "tara_bridge_address", "", "Address of the Tara bridge contract on Taraxa chain")
	pflag.StringVar(&config.TaraClientOnEthAddress, "tara_client_on_eth_address", "", "Address of the TaraClient contract on Ethereum chain")
	pflag.StringVar(&config.EthBridgeAddress, "eth_bridge_address", "", "Address of the Eth bridge contract on Ethereum chain")
	pflag.StringVar(&config.TaraxaNodeURL, "taraxa_node_url", "", "Taraxa node URL")
	pflag.StringVar(&config.PrivateKey, "private_key", "", "Private key")
	pflag.StringVar(&config.LightNodeEndpoint, "light_node_endpoint", "", "Light node endpoint")
	// Parse flags
	pflag.Parse()

	// Read config from environment variables
	viper.AutomaticEnv()

	// Bind environment variables to viper
	viper.BindEnv("ethereum_api_endpoint", "ETHEREUM_API_ENDPOINT")
	viper.BindEnv("beacon_light_client_address", "BEACON_LIGHT_CLIENT_ADDRESS")
	viper.BindEnv("eth_client_on_tara_address", "ETH_CLIENT_ON_TARA_ADDRESS")
	viper.BindEnv("tara_bridge_address", "TARA_BRIDGE_ADDRESS")
	viper.BindEnv("tara_client_on_eth_address", "TARA_CLIENT_ON_ETH_ADDRESS")
	viper.BindEnv("eth_bridge_address", "ETH_BRIDGE_ADDRESS")
	viper.BindEnv("taraxa_node_url", "TARAXA_NODE_URL")
	viper.BindEnv("private_key", "PRIVATE_KEY")
	viper.BindEnv("light_node_endpoint", "LIGHT_NODE_ENDPOINT")

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	log.Printf("Starting relayer with config: %+v", config)

	ctx, cancel := context.WithCancel(context.Background())

	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.Fatalf("Failed to convert private key: %v", err)
	}

	taraRelayer, err := to_tara.NewRelayer(&to_tara.Config{
		BeaconNodeEndpoint:  config.EthereumAPIEndpoint,
		TaraxaRPCURL:        config.TaraxaNodeURL,
		EthClientOnTaraAddr: common.HexToAddress(config.EthClientOnTaraAddress),
		Key:                 privateKey,
		LightNodeEndpoint:   config.LightNodeEndpoint,
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
		LightNodeEndpoint:     config.LightNodeEndpoint,
	})
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
