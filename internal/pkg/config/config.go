package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Load(configDir, env string) (*viper.Viper, error) {
	conf := viper.New()
	conf.AddConfigPath(configDir)
	conf.SetConfigName("default")
	if err := conf.ReadInConfig(); err != nil {
		if _, ok := err.(*viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("configs directory %s not found", configDir))
		}
		panic(err)
	}

	conf.SetConfigName(env)
	if err := conf.MergeInConfig(); err != nil {
		if _, ok := err.(*viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}
	return conf, nil
}
