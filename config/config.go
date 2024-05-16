package config

import (
	"encoding/json"
	"os"

	"github.com/Taraxa-project/relayer/internal/to_eth"
	"github.com/Taraxa-project/relayer/internal/to_tara"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

type RelayerConfig struct {
	TaraRelayerConfig to_tara.RelayerConfig `json:"tara_relayer"`
	EthRelayerConfig  to_eth.RelayerConfig  `json:"eth_relayer"`
}

type WalletConfig struct {
	Address    common.Address `json:"address"`
	PrivateKey string         `json:"private_key"`
}

func LoadConfigs(configPath string, walletPath string) (relayerConfig RelayerConfig, walletConfig WalletConfig) {
	configData, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal("Cannot read config file ", configPath, ", err: ", err)
	}

	walletData, err := os.ReadFile(walletPath)
	if err != nil {
		log.Fatal("Cannot read wallet file ", walletPath, ", err: ", err)
	}

	err = json.Unmarshal([]byte(configData), &relayerConfig)
	if err != nil {
		log.Fatal("Cannot parse config file data. Err: ", err, ", data: ", configData)
	}

	err = json.Unmarshal([]byte(walletData), &walletConfig)
	if err != nil {
		log.Fatal("Cannot parse wallet file data. Err: ", err, ", data: ", walletData)
	}

	return
}
