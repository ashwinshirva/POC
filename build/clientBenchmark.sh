#!/usr/bin/env bash

cd ./../src/client/services/
echo "=================================================================================================================="
echo "Starting benchmark for stream rpc..."
cd stream && go test -bench=./serverStreaming_test.go && cd ..
sleep 1
echo "=================================================================================================================="
echo "Starting benchmark for generic rpc..."
cd generic && go test -bench=./generic_test.go
cd ..
sleep 1
echo "=================================================================================================================="
echo "Starting benchmark for compression rpc..."
pwd
cd compression && go test -bench=./compression_test.go && cd ..
sleep 1
echo "=================================================================================================================="