#!/bin/bash
clear

set -x
set -e

go build
air go run .