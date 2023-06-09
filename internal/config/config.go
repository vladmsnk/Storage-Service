package config

import (
	"flag"
	"git.miem.hse.ru/1206/app"
	"git.miem.hse.ru/1206/app/logger"
	"git.miem.hse.ru/1206/app/storage/stpg"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	App      App            `yaml:"mode"`
	Logger   logger.Config  `yaml:"logger"`
	Storage  Storage        `yaml:"storage"`
	GRPC     app.GRPCConfig `yaml:"grpc"`
	Postgres stpg.Config    `yaml:"psql"`
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

type Storage struct {
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	AccessKeyId     string `yaml:"access_key_id"`
	SecretAccessKey string `yaml:"secret_access_key"`
	UseSsl          bool   `yaml:"use_ssl"`
	BucketName      string `yaml:"bucket_name"`
}
