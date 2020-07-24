module github.com/saiSunkari19/aicumen/client

go 1.14

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/cosmos/cosmos-sdk v0.34.4-0.20200530180557-ba70f4d4dc2e
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/golangci/golangci-lint v1.27.0 // indirect
	github.com/gorilla/mux v1.7.4
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rs/zerolog v1.19.0
	github.com/saiSunkari19/aicumen v0.0.0-20200722114847-d53b92a7e6ae
	github.com/saiSunkari19/aicumen/blockchain v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.0
	github.com/tendermint/go-amino v0.15.1 // indirect
	github.com/tendermint/tendermint v0.33.4
	github.com/tendermint/tm-db v0.5.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
)

replace github.com/saiSunkari19/aicumen/blockchain => ./../blockchain
