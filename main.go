package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	eth_client "github.com/Taraxa-project/relayer/clients/eth"
	tara_client "github.com/Taraxa-project/relayer/clients/tara"
	"github.com/Taraxa-project/relayer/internal/logging"
	"github.com/Taraxa-project/relayer/internal/to_eth"
	"github.com/Taraxa-project/relayer/internal/to_tara"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"

	"github.com/spf13/pflag"

	log "github.com/sirupsen/logrus"
)

func main() {
	var taraCfg to_tara.RelayerConfig
	var ethCfg to_eth.RelayerConfig

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Tara to Eth relayer config params
	pflag.StringVar(&ethCfg.Url, "ethereum_api_endpoint", os.Getenv("ETHEREUM_API_ENDPOINT"), "Ethereum API endpoint")
	taraClientContractAddress := pflag.String("tara_client_on_eth_address", os.Getenv("TARA_CLIENT_ON_ETH_ADDRESS"), "Address of the TaraClient contract on Ethereum chain")
	ethBridgeContractAddress := pflag.String("eth_bridge_address", os.Getenv("ETH_BRIDGE_ADDRESS"), "Address of the Eth bridge contract on Ethereum chain")

	// Eth to Tara relayer config params
	pflag.StringVar(&taraCfg.BeaconNodeEndpoint, "beacon_light_client_address", os.Getenv("BEACON_LIGHT_CLIENT_ADDRESS"), "Address of the BeaconLightClient contract on Taraxa chain")
	ethClientContractAddress := pflag.String("eth_client_on_tara_address", os.Getenv("ETH_CLIENT_ON_TARA_ADDRESS"), "Address of the EthClient contract on Taraxa chain")
	taraBridgeContractAddress := pflag.String("tara_bridge_address", os.Getenv("TARA_BRIDGE_ADDRESS"), "Address of the Tara bridge contract on Taraxa chain")
	pflag.StringVar(&taraCfg.Url, "taraxa_node_url", os.Getenv("TARAXA_NODE_URL"), "Taraxa node URL")
	pflag.StringVar(&taraCfg.LightNodeEndpoint, "light_node_endpoint", os.Getenv("LIGHT_NODE_ENDPOINT"), "Light node endpoint")
	privateKey := pflag.String("private_key", os.Getenv("PRIVATE_KEY"), "Private key")
	logLevel := pflag.String("log_level", os.Getenv("LOG_LEVEL"), "log level. could be only [trace, debug, info, warn, error, fatal]")

	pflag.Parse()

	ethCfg.TaraClientContractAddress = common.HexToAddress(*taraClientContractAddress)
	ethCfg.BridgeContractAddress = common.HexToAddress(*ethBridgeContractAddress)
	taraCfg.EthClientContractAddress = common.HexToAddress(*ethClientContractAddress)
	taraCfg.BridgeContractAddress = common.HexToAddress(*taraBridgeContractAddress)

	// Create shared clients
	taraClient, ethClient := createSharedClients(taraCfg, ethCfg, privateKey, privateKey)

	data_dir := "./"
	logging.Config(filepath.Join(data_dir, "logs"), *logLevel)

	log.Printf("Starting tara relayer with config: %+v, eth relayer with config: %+v", taraCfg, ethCfg)

	ctx, cancel := context.WithCancel(context.Background())

	taraRelayer, err := to_tara.NewRelayer(taraClient, ethClient, taraCfg)
	if err != nil {
		panic(err)
	}

	ethRelayer, err := to_eth.NewRelayer(taraClient, ethClient)
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

	// taraRelayer.Start(ctx)
	ethRelayer.Start(ctx)
	// Keep the main goroutine running until an interrupt is received
	<-ctx.Done()
}

func createSharedClients(taraCfg to_tara.RelayerConfig, ethCfg to_eth.RelayerConfig, taraNetworkPrivateKey *string, ethNetworkPrivateKey *string) (taraClient *tara_client.TaraClient, ethClient *eth_client.EthClient) {
	var err error

	taraClient, err = tara_client.NewTaraClient(*taraCfg.TaraClientConfig, taraNetworkPrivateKey)
	if err != nil {
		log.Fatal("NewTaraClient err: ", err)
	}

	// Tara client contract client on ethereum
	ethClient, err = eth_client.NewEthClient(*ethCfg.EthClientConfig, ethNetworkPrivateKey)
	if err != nil {
		log.Fatal("NewEthClient err: ", err)
	}

	return
}
