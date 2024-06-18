#!/bin/bash

# Ethereum Smart Contract Go Bindings Generation Script

# Define variables
REPO_URL="https://github.com/Taraxa-project/bridge.git"
REPO_DIR="bridge"
CONTRACTS=("TaraClient" "EthClient" "BeaconLightClient" "BridgeBase")  # Add more contract names here as needed
BINDINGS_DIR="./bindings"

if [ -z ${BRANCH+x} ]; then 
    echo "BRANCH is unset. Using the default branch 'master', you can set it using 'export BRANCH=your_branch'";
    BRANCH="master"
fi

echo "Generating bindings from the branch '$BRANCH'"; 

# Clone the repository
echo "Cloning the repository..."
git clone --recursive $REPO_URL $REPO_DIR
cd $REPO_DIR || exit

git fetch
git checkout $BRANCH
git reset --hard origin/$BRANCH
git submodule update --init --recursive

# Compile the smart contracts
echo "Compiling the smart contracts..."
forge build

# Check if jq is installed
if ! command -v jq &> /dev/null; then
    echo "jq is not installed. Please install jq before running this script."
    exit 1
fi

# Install abigen if not already installed
if ! command -v abigen &> /dev/null; then
    echo "Installing abigen..."
    go install github.com/ethereum/go-ethereum/cmd/abigen@latest

    # Add GOPATH/bin to PATH
    export PATH=$PATH:$(go env GOPATH)/bin
fi

# Ensure the bindings directory exists
echo "Creating bindings directory: $BINDINGS_DIR"
mkdir -p $BINDINGS_DIR

# Function to generate Go bindings for a contract
generate_bindings() {
    local contract_name=$1
    local contract_dir="out/${contract_name}.sol"
    local contract_json="$contract_dir/${contract_name}.json"
    local abi_file="${contract_name}.abi"
    local go_package="${contract_name}"
    local go_output_file="${go_package}.go"
    local binding_output_file="../${BINDINGS_DIR}/${contract_name}/${contract_name}.go"

    # Check if contract JSON exists
    if [ ! -f "$contract_json" ]; then
        echo "Contract JSON file not found: $contract_json"
        return
    fi

    # Generate the contract ABI
    echo "Generating the ABI for $contract_name..."
    jq .abi $contract_json > $abi_file

    # Generate Go bindings
    echo "Generating Go bindings for $contract_name..."
    abigen --abi=$abi_file --pkg=$go_package --out=$go_output_file

    # Copy the generated Go file to the bindings directory
    echo "Moving $go_output_file to $binding_output_file"
    mv $go_output_file $binding_output_file

    echo "Go bindings have been generated successfully: $binding_output_file"
}

# Generate bindings for each contract
for contract in "${CONTRACTS[@]}"; do
    generate_bindings "$contract"
done

# Remove the cloned repository
cd ..
# echo "Removing the repository directory: $REPO_DIR"
# rm -rf $REPO_DIR

echo "All Go bindings have been generated and copied to $BINDINGS_DIR."
