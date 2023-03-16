package config

import "github.com/spf13/viper"

type Config struct {
	OpenAIToken string  `mapstructure:"OPENAI_TOKEN"`
	BotToken    string  `mapstructure:"BOT_TOKEN"`
	Admins      []int64 `mapstructure:"ADMINS"`
}

func NewConfig() (*Config, error) {
	err := parseEnv()
	if err != nil {
		return nil, err
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func parseEnv() error {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err == nil {
		return nil
	}
	if err := viper.BindEnv("OPENAI_TOKEN"); err != nil {
		return err
	}
	if err := viper.BindEnv("BOT_TOKEN"); err != nil {
		return err
	}
	if err := viper.BindEnv("ADMINS"); err != nil {
		return err
	}
	return nil

}