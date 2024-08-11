# Makefile

# Variables
APP_NAME := tinyrp
MAIN_FILE := cmd/tinyrp/main.go
SAMPLE_A := sample_A
SAMPLE_B := sample_B
SAMPLE_APP_FILE := cmd/sample/main.go
PORT := 8080

# Default target
all: build

# Build the application
build:
	go build -o $(APP_NAME) $(MAIN_FILE);
	go build -o '$(SAMPLE_A)' $(SAMPLE_APP_FILE);
	go build -o '$(SAMPLE_B)' $(SAMPLE_APP_FILE);

run: build
	./tinyrp

# Clean the build artifacts
clean:
	rm -f $(APP_NAME) $(SAMPLE_APP_NAME)

# Set environment variables and run the application
run_sample_a:
	NAME=A PORT=8080 ./$(SAMPLE_A)

run_sample_b:
	NAME=B PORT=8081 ./$(SAMPLE_B)

.PHONY: all build run clean run_with_env