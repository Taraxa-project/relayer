# Relayer

Relayer is a Go service that connects the Ethereum and Taraxa networks by monitoring bridge and light-client contracts on both chains and forwarding the required data between them.

It is intended to run as a long‑lived process (or container) using a configured Ethereum RPC endpoint, Taraxa node, and bridge contract addresses.

## Requirements

- Go 1.24 or newer
- Access to:
  - An Ethereum execution RPC endpoint (HTTP/HTTPS)
  - An Ethereum beacon node endpoint (HTTP/HTTPS)
  - A Taraxa node WebSocket endpoint
- Contract addresses for:
  - Beacon light client on Taraxa
  - EthClient on Taraxa
  - Tara bridge on Taraxa
  - TaraClient on Ethereum
  - Eth bridge on Ethereum

 Docker (optional) is supported via the provided `Dockerfile`.

## Configuration

The relayer is configured via environment variables and/or command‑line flags. Flags, when specified, override the corresponding environment variables.

### Core environment variables

- `ETHEREUM_API_ENDPOINT`  
  Ethereum JSON‑RPC endpoint (e.g. `https://eth-mainnet.example.com/v2/<api-key>`).

- `BEACON_LIGHT_CLIENT_ADDRESS`  
  Address of the BeaconLightClient contract on the Taraxa chain.

- `ETH_CLIENT_ON_TARA_ADDRESS`  
  Address of the EthClient contract on the Taraxa chain.

- `TARA_BRIDGE_ADDRESS`  
  Address of the Tara bridge contract on the Taraxa chain.

- `TARA_CLIENT_ON_ETH_ADDRESS`  
  Address of the TaraClient contract on the Ethereum chain.

- `ETH_BRIDGE_ADDRESS`  
  Address of the Eth bridge contract on the Ethereum chain.

- `TARAXA_NODE_URL`  
  Taraxa node WebSocket URL (e.g. `wss://ws.mainnet.taraxa.io`).

- `BEACON_NODE_ENDPOINT`  
  Ethereum beacon node endpoint (e.g. `https://beacon.mainnet.taraxa.io`).

- `PRIVATE_KEY`  
  Hex‑encoded private key used by the relayer to sign transactions. **Keep this secret and never commit it to version control.**

- `ETH_GAS_PRICE_LIMIT`  
  Maximum gas price, in wei, that the relayer will accept for Ethereum transactions.  
  If unset, a default (15 gwei) is used.

- `PILLAR_BLOCKS_IN_BATCH`  
  Number of pillar blocks included in a batch when relaying to Ethereum.  
  If unset, a default of `20` is used.

- `LOG_LEVEL`  
  Log level; one of `trace`, `debug`, `info`, `warn`, `error`, `fatal`.  
  Defaults to `info` if unset.

### Command‑line flags

All of the above configuration values can also be provided as flags:

- `--ethereum_api_endpoint`
- `--beacon_light_client_address`
- `--eth_client_on_tara_address`
- `--tara_bridge_address`
- `--tara_client_on_eth_address`
- `--eth_bridge_address`
- `--taraxa_node_url`
- `--private_key`
- `--beacon_node_endpoint`
- `--eth_gas_price_limit`
- `--pillar_blocks_in_batch`
- `--log_level`

Flags override the corresponding environment variables when both are set.

### Using a `.env` file

This project uses `github.com/joho/godotenv`, so a `.env` file in the project root will be loaded automatically on startup.  
You can copy the existing `.env.example` as a template and replace the values with your own endpoints and contract addresses:

```bash
cp .env.example .env
# edit .env.local with your own values
```

Ensure that any `.env` files containing secrets are **not** committed to version control.

## Building and running locally

### Build with Makefile (recommended)

The Makefile provides a simple wrapper around common tasks:

```bash
make build
```

This produces the `relayer` binary at `build/relayer`.

Run the binary (after setting up your environment variables or `.env` file):

```bash
./build/relayer
```

### Run directly with `go run`

You can also run the service directly without producing a binary:

```bash
go run ./...
```

or:

```bash
go run main.go
```

In all cases, the process will continue running until it receives `SIGINT` or `SIGTERM` (e.g. Ctrl‑C).

## Running with Docker

A multi‑stage `Dockerfile` is provided to build and run the relayer inside a minimal Alpine image.

### Build the image

```bash
docker build -t relayer .
```

### Run the container

Using an env file:

```bash
docker run --rm \
  --env-file .env \
  --name relayer \
  relayer
```

Logs are written both to stdout and to log files under the `logs/` directory in the container’s working directory.

## Development tasks

- **Lint**:

  ```bash
  make lint
  ```

  This will install `golangci-lint` into `./bin` if needed and run it over the codebase.

- **Run tests**:

  ```bash
  make check
  ```

- **Tidy Go modules**:

  ```bash
  make tidy
  ```
