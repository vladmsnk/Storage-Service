package config

import (
	"flag"
	"git.miem.hse.ru/1206/app"
	"git.miem.hse.ru/1206/app/logger"
	"git.miem.hse.ru/1206/app/storage/s3"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// Ну эта директория особо в пояснениях не нуждается.
// Здесь находится структура конфига и его инициализация.

type Config struct {
	// Структура конфига
	App     App            `yaml:"mode"`
	Logger  logger.Config  `yaml:"logger"`
	Storage s3.Config      `yaml:"storage"`
	GRPC    app.GRPCConfig `yaml:"grpc"`
}

func Init() *Config {
	configPath := flag.String("c", "config.yml", "Path to configuration file")
	flag.Parse()

	data, err := os.ReadFile("etc/" + *configPath)
	if err != nil {
		log.Fatalf("error reading config file: %s", err.Error())
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("error unmarshalling config file: %s", err.Error())
	}

	return &cfg
}

type App struct {
	Mode string `yaml:"mode"`
}
