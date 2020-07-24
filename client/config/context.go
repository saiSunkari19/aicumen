package config

import (
	"errors"
	"net/http"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/rs/zerolog/log"
)

var (
	errInvalidGasAdjustment = errors.New("invalid gas adjustment")
)

func CreateCLIContextFromConfig(config *Config, cdc *codec.Codec) (*CLI, error) {

	kb, err := keyring.New("app", viper.GetString(flags.FlagKeyringBackend), config.HomeDir, nil)
	if err != nil {
		panic(err)
	}
	log.Info().Str("nodeURI", config.RPCAddress).Str("chainID", config.ChainID).Msg("set config params")

	cli := NewCLI(cdc, kb, config)
	_client, verifier, err := NewVerifier(config.HomeDir, config.ChainID, config.RPCAddress)
	if err != nil {
		panic(err)
	}
	cli.CliCtx.Client = _client
	cli.CliCtx.Verifier = verifier

	cli.CliCtx = cli.CliCtx.WithCodec(cdc)
	cli.CliCtx = cli.CliCtx.WithFromName(config.KeyName)
	cli.CliCtx = cli.CliCtx.WithNodeURI(config.RPCAddress)
	cli.CliCtx = cli.CliCtx.WithChainID(config.ChainID)
	cli.CliCtx = cli.CliCtx.WithBroadcastMode(config.BroadcastMode)
	cli.CliCtx = cli.CliCtx.WithTrustNode(true)

	info, err := cli.CliCtx.Keyring.Key(config.KeyName)
	if err != nil {
		panic(err)
	}
	cli.CliCtx = cli.CliCtx.WithFromAddress(info.GetAddress())

	log.Info().Msg("created cliCtx with verifier")
	return cli, nil
}

func BuildSignAndBroadCast(w http.ResponseWriter, cli *CLI, br rest.BaseReq, msgs []sdk.Msg, password string) {
	gasAdj, ok := rest.ParseFloat64OrReturnBadRequest(w, br.GasAdjustment, flags.DefaultGasAdjustment)
	if !ok {
		return
	}

	simAndExec, gas, err := flags.ParseGas(br.Gas)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	account, err := cli.GetAccount(cli.CliCtx.GetFromAddress())
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	txBldr := types.NewTxBuilder(
		authclient.GetTxEncoder(cli.CliCtx.Codec), account.GetAccountNumber(), account.GetSequence(), gas, gasAdj,
		br.Simulate, br.ChainID, br.Memo, br.Fees, br.GasPrices,
	)

	if br.Simulate || simAndExec {
		if gasAdj < 0 {
			rest.WriteErrorResponse(w, http.StatusBadRequest, errInvalidGasAdjustment.Error())
			return
		}

		txBldr, err = authclient.EnrichWithGas(txBldr, cli.CliCtx, msgs)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if br.Simulate {
			rest.WriteSimulationResponse(w, cli.CliCtx.Codec, txBldr.Gas())
			return
		}
	}

	txBytes, err := txBldr.BuildAndSign(cli.CliCtx.GetFromName(), password, msgs)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := cli.CliCtx.BroadcastTx(txBytes)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	rest.PostProcessResponseBare(w, cli.CliCtx, res)

}
