package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string `yaml:"address" env:"HTTP_SERVER_ADDRESS" env-default:"localhost:8082"`
}

type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-default:"production"`
	StoragePath string     `yaml:"storage_path" env:"STORAGE_PATH" env-default:"/tmp"`
	HTTPServer  HTTPServer `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	// If config path is not set, check if it is passed as a flag
	if configPath == "" {
		flags := flag.String("config", "", "path to config file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("config path is required and not set")
		}
	}

	// Load config from file
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file not found at %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("failed to load config: %v", err.Error())
	}

	return &cfg
}
