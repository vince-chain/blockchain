#!/bin/bash

KEY="mykey"
CHAINID="vince_5000-1"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t vince-datadir.XXXXX)

echo "create and add new keys"
./vinced keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
vince./vinced init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./vinced add-genesis-account \
"$(./vinced keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000avince \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./vinced gentx $KEY 1000000000000000000avince --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./vinced collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./vinced validate-genesis --home $DATA_DIR

echo "starting vince node $i in background ..."
./vinced start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started vince node"
tail -f /dev/null