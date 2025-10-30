package config

import "github.com/spf13/viper"

type Config struct {
	DbUsers    string `mapstructure:""`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
}

func Load() (Config, error) {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	cnf := Config{}
	err := viper.Unmarshal(&cnf)

	return cnf, err
}
