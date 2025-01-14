#!/bin/sh

# Enable debugging and error handling
set +x
set -e

mkdir cosmos
cd cosmos
git clone https://github.com/cosmos/cosmos-sdk
cd cosmos-sdk
#git pull
#local
#make build

#docker build . -t simd:v2.0.0-beta.1
#docker run --rm -it simd:v2.0.0-beta.1 simd version

docker build . -t 0.52.0-rc.1
docker run --rm -it simd:0.52.0-rc.1 simd version
