.PHONY: build run test clean
.ALL: run

APP_NAME := emugo-8
APP_PATH := cmd/emulator

run:
	@echo "Running emulator..."
	go run $(APP_PATH)/main.go

test:
	@echo "Running tests..."
	go test -v ./...

build:
	@echo "Building emulator..."
	go build -o $(APP_NAME) $(APP_PATH)/main.go

clean:
	@echo "Cleaning build..."
	rm -f $(APP_NAME)
