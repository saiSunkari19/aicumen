package config

import (
	"fmt"
	"io/ioutil"
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	keys "github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/lite"
	"github.com/tendermint/tendermint/lite/proxy"
	"github.com/tendermint/tendermint/rpc/client/http"
)

type CLI struct {
	CliCtx context.CLIContext
}

func NewCLI(cdc *codec.Codec, kb keys.Keyring, config *Config) *CLI {
	return &CLI{
		CliCtx: context.CLIContext{
			Codec:         cdc,
			Output:        os.Stdout,
			Keyring:       kb,
			OutputFormat:  "json",
			BroadcastMode: config.BroadcastMode,
			SkipConfirm:   true,
		},
	}
}
func NewVerifier(dir, id, address string) (*http.HTTP, *lite.DynamicVerifier, error) {
	root, err := ioutil.TempDir(dir, "lite_")
	if err != nil {
		return nil, nil, err
	}

	c, err := http.New(address, "/websocket")
	if err != nil {
		return nil, nil, err
	}

	verifier, err := proxy.NewVerifier(id, root, c, log.NewNopLogger(), 10)
	if err != nil {
		return nil, nil, err
	}

	return c, verifier, nil
}

func (c *CLI) GetAccount(address sdk.AccAddress) (auth.BaseAccount, error) {
	bytes, err := c.CliCtx.Codec.MarshalJSON(auth.NewQueryAccountParams(address))
	if err != nil {
		return auth.BaseAccount{}, err
	}

	res, _, err := c.CliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", auth.QuerierRoute, auth.QueryAccount), bytes)
	if err != nil {
		return auth.BaseAccount{}, nil
	}

	var account auth.BaseAccount
	if err := c.CliCtx.Codec.UnmarshalJSON(res, &account); err != nil {
		return auth.BaseAccount{}, err
	}

	return account, nil
}
