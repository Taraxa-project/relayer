# Beacon Light Client Go Bindings

This README provides instructions for generating Go bindings for the Beacon Light Client Ethereum Smart Contract. Follow these steps to clone the repository, prepare your environment, compile the smart contract, generate the ABI, and create Go bindings.

## Prerequisites

Before you start, ensure you have the following tools installed on your system:

- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- [Foundry](https://book.getfoundry.sh/getting-started/installation.html)
- [Go Ethereum (`abigen`)](https://geth.ethereum.org/docs/install-and-build/installing-geth)
- [`jq`](https://stedolan.github.io/jq/download/)

## Setup Instructions

### Step 1: Clone the Repository

First, clone the Beacon Light Client repository to your local machine:

```bash
git clone https://github.com/Taraxa-project/beacon-light-client.git
cd beacon-light-client
```

### Step 2: Compile the Smart Contract

Compile the smart contract using Foundry's `forge` to generate the contract artifacts:

```bash
forge build
```

### Step 3: Generate the Contract ABI

Extract the ABI from the generated JSON artifact using `jq`:

```bash
jq .abi out/BeaconLightClient.sol/BeaconLightClient.json > BeaconLightClient.abi
```

This command creates a `BeaconLightClient.abi` file containing the ABI for the Beacon Light Client contract.

### Step 4: Install `abigen`

If you haven't already installed `abigen`, you can do so by running:

```bash
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

Ensure that your `$GOPATH/bin` or `$GOBIN` is in your `PATH`.

```
export PATH=$PATH:$(go env GOPATH)/bin
```

### Step 5: Generate Go Bindings

With the ABI file ready, use `abigen` to generate Go bindings:

```bash
abigen --abi=BeaconLightClient.abi --pkg=BeaconLightClient --out=BeaconLightClient.go
```

This command generates a Go file named `BeaconLightClient.go` containing Go bindings for the Beacon Light Client contract.
