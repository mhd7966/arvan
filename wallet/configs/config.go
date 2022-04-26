package configs

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"os"
)

var Cfg Config

type Config struct {
	Debug bool `env:"DEBUG" env-default:"False"`
	Redis struct {
		Addr string `env:"REDIS_ADD" env-default:"redis:6379"`
	}
	Postgres struct {
		Port string `env:"PG_PORT" env-default:"5432"`
		Host string `env:"PG_HOST" env-default:"postgres"`
		Name string `env:"PG_NAME" env-default:"wallet"`
		User string `env:"PG_USER" env-default:"admin"`
		Pass string `env:"PG_PASS" env-default:"postgres_password"`
	}
	Log struct {
		LogLevel   string `env:"LOG_LEVEL" env-default:"debug"`
		OutputType string `env:"LOG_OUTPUT_TYPE" env-default:"stdout"`
		OutputAdd  string `env:"LOG_FILE_Add" env-default:"/log.txt"`
	}
	CodeService struct {
		Host string `env:"CODESERVICE_HOST" env-default:"localhost:3001"`
	}
}

func SetConfig() {

	if _, err := os.Stat(".env"); err == nil {
		cleanenv.ReadConfig(".env", &Cfg)
		logrus.Info("Set config from .env file")
	} else {
		cleanenv.ReadEnv(&Cfg)
		logrus.Info("Set config from Config struct values")
	}

}
