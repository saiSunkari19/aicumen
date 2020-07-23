package config

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rs/zerolog/log"

	"github.com/saiSunkari19/aicumen/blockchain/types"
)

type Context struct {
	CliCtx context.CLIContext
}

type Config struct {
	KeyName            string `toml:"key-name"`
	RPCAddress         string `toml:"rpc-addr"`
	ChainID            string `toml:"chain-id"`
	KeyBaseServiceName string `toml:"keybase-service-name"`
	HomeDir            string `toml:"home"`
	RestPort           string `toml:"port"`
	BroadcastMode      string `toml:"broadcast-mode"`
}

func SetInitConfig(cfg *Config) {

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(types.Bech32PrefixAccAddr, types.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(types.Bech32PrefixValAddr, types.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(types.Bech32PrefixConsAddr, types.Bech32PrefixConsPub)
	config.Seal()

	log.Info().Str("keyring-service-name", sdk.KeyringServiceName()).Msg("keyring service name")

}
