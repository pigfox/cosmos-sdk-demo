#!/bin/sh

# Enable debugging and error handling
set -x
set -e
echo "---------Current script name: $0---------"
source_path="/home/peter/.simapp/config/genesis.json"
target_path="/home/peter/Documents/Mercor/Project1/cosmos-sdk-demo/go-generated-genesis/genesis.json"
ls -l "$source_path"
cp -rf "$source_path" "$target_path"
ls -l "$target_path"