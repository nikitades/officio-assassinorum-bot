package conf

import (
	"context"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	BotToken           string `env:"BOT_TOKEN,required"`
	BotName            string `env:"BOT_NAME,required"`
	Mode               string `env:"MODE,default=polling"`
	Port               int    `env:"PORT,default=8080"`
	DatabaseConnection string `env:"DB_CONNECTION,required"`
}

var CurrentConfig = new(Config)

//App global initialization, no matter which entrypoint is the app started from
func init() {
	rand.Seed(time.Now().UnixNano())

	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := envconfig.Process(ctx, CurrentConfig); err != nil {
		log.Fatal(err)
	}

	CurrentConfig.BotName = strings.ReplaceAll(CurrentConfig.BotName, "@", "")
}
