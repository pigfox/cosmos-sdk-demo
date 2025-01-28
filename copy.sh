#!/bin/sh
clear

set -x
set -e

file="genesis.json"
ls -l $file
cp ~/.simd/config/$file  ./$file
ls -l $file