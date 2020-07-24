package cmd

import (
	"net/http"
	
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/version"
	
	"github.com/cosmos/cosmos-sdk/client"
	rest2 "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	
	internal "github.com/saiSunkari19/aicumen/client/config"
	"github.com/saiSunkari19/aicumen/client/rest"
)

func handlerRestServer(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rest-server",
		Short: "Start rest-server for the given node",
		RunE: func(cmd *cobra.Command, args []string) error {
			
			version.Name = config.KeyBaseServiceName
			
			log.Info().Str("address", config.RestPort).Msg("starting rest-server")
			keyring, _ := cmd.PersistentFlags().GetString(FlagKeyringBackend)
			viper.Set(FlagKeyringBackend, keyring)
			log.Info().Str("keyring-backend", keyring).Msg("keyring backend type")
			
			internal.SetInitConfig(config)
			cliCtx, err := internal.CreateCLIContextFromConfig(config, cdc)
			if err != nil {
				return err
			}
			
			rest.RegisterRoutes(cliCtx, &router)
			client.RegisterRoutes(cliCtx.CliCtx, &router)
			rest2.RegisterTxRoutes(cliCtx.CliCtx, &router)
			
			log.Info().Str("address", config.RestPort).Msg("server started")
			return http.ListenAndServe(config.RestPort, &router)
		},
	}
	
	cmd.PersistentFlags().String(FlagKeyringBackend, DefaultKeyringBackend, "Select keyring's backend (os|file|test)")
	viper.BindPFlag(FlagKeyringBackend, cmd.Flags().Lookup(FlagKeyringBackend))
	
	return cmd
}
