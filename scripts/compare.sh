#!/bin/sh

# Enable debugging and error handling
set -x  # Show commands as they execute
set -e  # Exit script on any error
clear
echo "---------Current script name: $0---------"
validFile="./config/genesis.json"
goGeneratedFile="/home/peter/.simapp/config/genesis.json"

diff "$validFile" "$goGeneratedFile"