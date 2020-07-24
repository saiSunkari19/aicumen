package cmd

import (
	"fmt"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/saiSunkari19/aicumen/blockchain/app"

	internal "github.com/saiSunkari19/aicumen/client/config"
)

const (
	logLevelJSON       = "json"
	logLevelText       = "text"
	cfgFile            = "config.toml"
	FlagcfgFile        = "config"
	FlagKeyringBackend = "keyring-backend"

	DefaultKeyringBackend = KeyringBackendOS
	KeyringBackendOS      = "os"
)

var (
	logLevel  string
	logFormat string
	config    *internal.Config
	cfgPath   string
)
var router = *mux.NewRouter()

var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "Client implementation for app",
}

func init() {

	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", zerolog.InfoLevel.String(), "logging-level")
	rootCmd.PersistentFlags().StringVar(&logFormat, "log-format", logLevelJSON, "logging  format must be json or text")
	rootCmd.PersistentFlags().StringVar(&cfgPath, FlagcfgFile, cfgFile, "set config file")

}

func Execute() {
	_, cdc := app.MakeCodecs()

	rootCmd.AddCommand(handlerRestServer(cdc))

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		return initConfig(rootCmd)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
