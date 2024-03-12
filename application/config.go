package application

import (
	"log"

	"github.com/spf13/viper"
)

// Config has environment stored
type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	Env        string `mapstructure:"ENV"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPass     string `mapstructure:"DB_PASS"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
}

// NewConfig creates a new environment
func NewConfig() Config {

	env := Config{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("☠️ cannot read configuration")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
	}

	return env
}
