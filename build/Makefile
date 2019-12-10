SHELL := /bin/bash

all: run_client_benchmark clean

run_client_benchmark:
	@echo "------------------------------------------------------------------------------------------------"
	@echo "BUILD: Setting GO Environment Variables"
	@echo ""
	@sh ./startServer.sh
	@sh ./clientBenchmark.sh
	@echo ""
	@echo "------------------------------------------------------------------------------------------------"

clean:
	@echo "------------------------------------------------------------------------------------------------"
	@echo "Clean up initiated..."
	@echo "Stopping server..."
	@fuser -k 50051/tcp
	@echo "Clean up complete..."
	@echo "------------------------------------------------------------------------------------------------"