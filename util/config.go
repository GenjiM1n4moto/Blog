package util

import (
	"github.com/spf13/viper"
)

// Config struct represents all configuration of the application
// The values are read by Viper from a config file or environment variables
type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()       // Check if env variable matches
	err = viper.ReadInConfig() // Read config
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config) // Unmarshal config, before unmarshal, it's better to do AutomaticEnv
	return config, err
}
