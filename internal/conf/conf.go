package conf

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	BotToken string `env:"BOT_TOKEN,required"` //todo envconfig
}

//New creates the instance of the config filled with the values of .env file
func New() (*Config, error) {
	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	cfg := new(Config)

	if err := envconfig.Process(ctx, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
