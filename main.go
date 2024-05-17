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
	"github.com/joho/godotenv"

	"github.com/spf13/pflag"

	log "github.com/sirupsen/logrus"
)

func main() {
	//var cfg Config

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	//var log_level string

	// TODO: unite config processing - do we want config file, cmd args or both ?
	configPath := pflag.String("config", "configs/customnet.json", "Path to the config file")
	walletPath := pflag.String("wallet", "configs/wallet.json", "Path to the wallet file with private key")
	logLevel := pflag.String("log_level", os.Getenv("LOG_LEVEL"), "log level. could be only [trace, debug, info, warn, error, fatal]")

	// pflag.StringVar(&cfg.EthereumAPIEndpoint, "ethereum_api_endpoint", os.Getenv("ETHEREUM_API_ENDPOINT"), "Ethereum API endpoint")
	// pflag.StringVar(&cfg.BeaconLightClientAddress, "beacon_light_client_address", os.Getenv("BEACON_LIGHT_CLIENT_ADDRESS"), "Address of the BeaconLightClient contract on Taraxa chain")
	// pflag.StringVar(&cfg.EthClientOnTaraAddress, "eth_client_on_tara_address", os.Getenv("ETH_CLIENT_ON_TARA_ADDRESS"), "Address of the EthClient contract on Taraxa chain")
	// pflag.StringVar(&cfg.TaraBridgeAddress, "tara_bridge_address", os.Getenv("TARA_BRIDGE_ADDRESS"), "Address of the Tara bridge contract on Taraxa chain")
	// pflag.StringVar(&cfg.TaraClientOnEthAddress, "tara_client_on_eth_address", os.Getenv("TARA_CLIENT_ON_ETH_ADDRESS"), "Address of the TaraClient contract on Ethereum chain")
	// pflag.StringVar(&cfg.EthBridgeAddress, "eth_bridge_address", os.Getenv("ETH_BRIDGE_ADDRESS"), "Address of the Eth bridge contract on Ethereum chain")
	// pflag.StringVar(&cfg.TaraxaNodeURL, "taraxa_node_url", os.Getenv("TARAXA_NODE_URL"), "Taraxa node URL")
	// pflag.StringVar(&cfg.PrivateKey, "private_key", os.Getenv("PRIVATE_KEY"), "Private key")
	// pflag.StringVar(&cfg.LightNodeEndpoint, "light_node_endpoint", os.Getenv("LIGHT_NODE_ENDPOINT"), "Light node endpoint")
	// pflag.StringVar(&log_level, "log_level", os.Getenv("LOG_LEVEL"), "log level. could be only [trace, debug, info, warn, error, fatal]")
	pflag.Parse()

	relayerConfig, walletConfig := config.LoadConfigs(*configPath, *walletPath)

	// Create shared clients
	taraClient, ethClient := createSharedClients(relayerConfig, &walletConfig.PrivateKey, &walletConfig.PrivateKey)

	data_dir := "./"
	logging.Config(filepath.Join(data_dir, "logs"), *logLevel)

	log.Printf("Starting relayer with config: %+v", relayerConfig)

	ctx, cancel := context.WithCancel(context.Background())

	taraRelayer, err := to_tara.NewRelayer(taraClient, ethClient, relayerConfig.TaraRelayerConfig)
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

func createSharedClients(relayerConfig config.RelayerConfig, taraNetworkPrivateKey *string, ethNetworkPrivateKey *string) (taraClient *tara_client.TaraClient, ethClient *eth_client.EthClient) {
	var err error

	taraClient, err = tara_client.NewTaraClient(*relayerConfig.TaraRelayerConfig.TaraClientConfig, client_base.WebSocket, taraNetworkPrivateKey)
	if err != nil {
		log.Fatal("NewTaraClient err: ", err)
	}

	// Tara client contract client on ethereum
	ethClient, err = eth_client.NewEthClient(*relayerConfig.EthRelayerConfig.EthClientConfig, client_base.Http, ethNetworkPrivateKey)
	if err != nil {
		log.Fatal("NewEthClient err: ", err)
	}

	return
}
