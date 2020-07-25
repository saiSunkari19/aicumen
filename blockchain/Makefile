PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=app \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=appd \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=appcli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

all: install

create-wallet:
	appcli keys add relayer --keyring-backend test

build:
	go build -o bin/appd ./cmd/appd

unit-test:
	go test -mod=readonly ./... -v

# TODO : Think a diffrent way to install
install: go.sum
	go build -mod=mod $(BUILD_FLAGS) -o $${GOBIN-$${GOPATH-$$HOME/go}/bin}/appd ./cmd/appd
	go build -mod=mod $(BUILD_FLAGS) -o $${GOBIN-$${GOPATH-$$HOME/go}/bin}/appcli ./cmd/appcli

go.sum: go.mod
	echo "--> Ensure dependencies have not been modified"
	GO111MODULE=on go mod verify

build-linux:
	GOARCH=amd64 GOOS=linux go build -o bin/appd ./cmd/appd
	GOARCH=amd64 GOOS=linux  go build -o bin/appcli ./cmd/appcli

docker-test: build-linux
	docker build -f Dockerfile.test -t saiSunkari19/blockchain:latest .

lint:
	echo "--> Running linter"
	golangci-lint run
	go mod verify

.PHONY: install unit-test