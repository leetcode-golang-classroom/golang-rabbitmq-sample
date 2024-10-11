package internal

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	RABBITMQ_URL string `mapstructure:"RABBITMQ_URL"`
	PORT         string `mapstructure:"PORT"`
	QUEUE_NAME   string `mapstructure:"QUEUE_NAME"`
}

var AppConfig *Config

func loadConfig() {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")
	failOnError(v.BindEnv("PORT"), "Failed on Bind PORT")
	failOnError(v.BindEnv("RABBITMQ_URL"), "Failed on Bind RABBITMQ_URL")
	failOnError(v.BindEnv("QUEUE_NAME"), "Failed on QUEUE_NAME")
	err := v.ReadInConfig()
	if err != nil {
		failOnError(err, "Load from environment variable")
	}
	v.AutomaticEnv()

	err = v.Unmarshal(&AppConfig)
	if err != nil {
		failOnError(err, "Failed to read enivroment")
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func init() {
	loadConfig()
}
