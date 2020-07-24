Blockchain Server
===


**Required [Go1.14+](https://golang.org/dl/)**

## Installation

```bash
   go get -u github.com/saiSunkari19/aicumen
```
- Go to the directory 
```bash
   cd $GOPATH/github.com/saiSunkari19/aicumen
   cd blockchain
```
- Install the Project
```bash
    make install
```

## Setup
- Creating Account

```bash
    appcli keys add relayer --keyring-backend test
```
- Init chain
```bash
    bash entry-point.sh chain-1 $(appcli keys show relayer --keyring-backed test) stake
```

- Start chain
> Note: starts at 26657 by default
```bash
    appd start
```