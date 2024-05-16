package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/Taraxa-project/relayer/clients/client_base"
	eth_client "github.com/Taraxa-project/relayer/clients/eth"
	tara_client "github.com/Taraxa-project/relayer/clients/tara"
	"github.com/Taraxa-project/relayer/config"
	"github.com/Taraxa-project/relayer/internal/logging"
	"github.com/Taraxa-project/relayer/internal/to_eth"
	"github.com/Taraxa-project/relayer/internal/to_tara"

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

	TaraxaNodeURL     string
	PrivateKey        string
	LightNodeEndpoint string
}

func main() {
	var cfg Config

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	var log_level string

	// TODO: unite config processing
	configPath := pflag.String("config", "configs/customnet.json", "Path to the config file")
	walletPath := pflag.String("wallet", "configs/wallet.json", "Path to the wallet file with private key")

	pflag.StringVar(&cfg.EthereumAPIEndpoint, "ethereum_api_endpoint", os.Getenv("ETHEREUM_API_ENDPOINT"), "Ethereum API endpoint")
	pflag.StringVar(&cfg.BeaconLightClientAddress, "beacon_light_client_address", os.Getenv("BEACON_LIGHT_CLIENT_ADDRESS"), "Address of the BeaconLightClient contract on Taraxa chain")
	pflag.StringVar(&cfg.EthClientOnTaraAddress, "eth_client_on_tara_address", os.Getenv("ETH_CLIENT_ON_TARA_ADDRESS"), "Address of the EthClient contract on Taraxa chain")
	pflag.StringVar(&cfg.TaraBridgeAddress, "tara_bridge_address", os.Getenv("TARA_BRIDGE_ADDRESS"), "Address of the Tara bridge contract on Taraxa chain")
	pflag.StringVar(&cfg.TaraClientOnEthAddress, "tara_client_on_eth_address", os.Getenv("TARA_CLIENT_ON_ETH_ADDRESS"), "Address of the TaraClient contract on Ethereum chain")
	pflag.StringVar(&cfg.EthBridgeAddress, "eth_bridge_address", os.Getenv("ETH_BRIDGE_ADDRESS"), "Address of the Eth bridge contract on Ethereum chain")
	pflag.StringVar(&cfg.TaraxaNodeURL, "taraxa_node_url", os.Getenv("TARAXA_NODE_URL"), "Taraxa node URL")
	pflag.StringVar(&cfg.PrivateKey, "private_key", os.Getenv("PRIVATE_KEY"), "Private key")
	pflag.StringVar(&cfg.LightNodeEndpoint, "light_node_endpoint", os.Getenv("LIGHT_NODE_ENDPOINT"), "Light node endpoint")
	pflag.StringVar(&log_level, "log_level", os.Getenv("LOG_LEVEL"), "log level. could be only [trace, debug, info, warn, error, fatal]")
	pflag.Parse()

	relayerConfig, walletConfig := config.LoadConfigs(*configPath, *walletPath)

	// Create shared clients
	taraClient, ethClient := createSharedClients(relayerConfig)

	// Create eth transactor
	ethTransactor, err := ethClient.NewTransactor(walletConfig.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	// // Create tara transactor
	// taraTransactor, err := taraClient.NewTransactor(walletConfig.PrivateKey)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	data_dir := "./"
	logging.Config(filepath.Join(data_dir, "logs"), log_level)

	log.Printf("Starting relayer with config: %+v", cfg)

	ctx, cancel := context.WithCancel(context.Background())

	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		log.Fatalf("Failed to convert private key: %v", err)
	}

	taraRelayer, err := to_tara.NewRelayer(&to_tara.Config{
		BeaconNodeEndpoint:    cfg.EthereumAPIEndpoint,
		EthRPCURL:             cfg.EthereumAPIEndpoint,
		TaraxaRPCURL:          cfg.TaraxaNodeURL,
		BeaconLightClientAddr: common.HexToAddress(cfg.BeaconLightClientAddress),
		EthBridgeAddr:         common.HexToAddress(cfg.EthBridgeAddress),
		TaraxaBridgeAddr:      common.HexToAddress(cfg.TaraBridgeAddress),
		EthClientOnTaraAddr:   common.HexToAddress(cfg.EthClientOnTaraAddress),
		Key:                   privateKey,
		LightNodeEndpoint:     cfg.LightNodeEndpoint,
	})
	if err != nil {
		panic(err)
	}

	ethRelayer, err := to_eth.NewRelayer(taraClient, ethClient, ethTransactor)

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

func createSharedClients(relayerConfig config.RelayerConfig) (taraClient *tara_client.TaraClient, ethClient *eth_client.EthClient) {
	var err error

	taraClient, err = tara_client.NewTaraClient(relayerConfig.TaraClientConfig, client_base.WebSocket)
	if err != nil {
		log.Fatal("NewTaraClient err: ", err)
	}

	// Tara client contract client on ethereum
	ethClient, err = eth_client.NewEthClient(relayerConfig.EthClientConfig, client_base.Http)
	if err != nil {
		log.Fatal("NewEthClient err: ", err)
	}

	return
}
