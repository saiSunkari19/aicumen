package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	
	"github.com/BurntSushi/toml"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	
	conf "github.com/saiSunkari19/aicumen/client/config"
)

func initConfig(cmd *cobra.Command, ) error {
	logLvl, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		fmt.Println(err)
	}
	
	zerolog.SetGlobalLevel(logLvl)
	
	switch logFormat {
	case logLevelJSON:
	case logLevelText:
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	default:
		fmt.Errorf("invalid format :%s", logFormat)
	}
	
	config = &conf.Config{}
	viper.SetConfigFile(cfgFile) // TODO: accept any type of config file
	if err := viper.ReadInConfig(); err != nil {
		log.Info().Err(err).Msg("unable to read the config existed config")
		return err
	}
	file, err := ioutil.ReadFile(viper.ConfigFileUsed())
	if err != nil {
		log.Err(err).Msg("i/o parse error")
		return err
	}
	
	_, err = toml.Decode(string(file), config)
	if err != nil {
		log.Err(err).Msg("yaml unmarshal error")
		return err
	}
	
	log.Info().Msg("config.toml successfully parsed")
	return nil
}
