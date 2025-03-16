package config

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func LoadViperConfig() error {
	// replace nested values
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigType("yaml")

	err := viper.ReadConfig(bytes.NewBuffer(PropertiesFile))
	if err != nil {
		return fmt.Errorf("error while loading config: %w", err)
	}

	viper.AutomaticEnv() // load from env

	return nil
}
