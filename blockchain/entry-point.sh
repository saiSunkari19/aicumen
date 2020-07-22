#!/bin/sh

CHAINID=$1
GENACCT=$2
DENOM=$3

if [ -z "$1" ]; then
  echo "Need to input chain id..."
  exit 1
fi

if [ -z "$2" ]; then
  echo "Need to input genesis account address..."
  exit 1
fi

if [ -z "$3" ]; then
  echo "Need to input staking/gov denom..."
  exit 1
fi

# Build genesis file incl account for passed address
coins="100000000000$DENOM,100000000000testtoken"
appd init --chain-id $CHAINID $CHAINID --stake-denom $DENOM
appcli keys add validator --keyring-backend="test"
appd add-genesis-account validator $coins --keyring-backend="test"
appd add-genesis-account $GENACCT $coins --keyring-backend="test"
appd gentx --name validator --amount 100000000$DENOM --keyring-backend="test"
appd collect-gentxs

# Set proper defaults and change ports
sed -i 's/"leveldb"/"goleveldb"/g' ~/.appd/config/config.toml
sed -i 's#"tcp://127.0.0.1:26657"#"tcp://0.0.0.0:26657"#g' ~/.appd/config/config.toml
sed -i 's/timeout_commit = "5s"/timeout_commit = "1s"/g' ~/.appd/config/config.toml
sed -i 's/timeout_propose = "3s"/timeout_propose = "1s"/g' ~/.appd/config/config.toml
sed -i 's/index_all_keys = false/index_all_keys = true/g' ~/.appd/config/config.toml

# Start the chain
# appd start --pruning=nothing
