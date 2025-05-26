package config

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog/log"
)

// Koanf instance
var k = koanf.New(".")

var Cfg Config

// InitConfig initializes the configuration by loading from the YAML file and environment variables
func InitConfig() {
	// Get the config file path from the environment variable or use the default
	configFile := os.Getenv("AUTH5_CONFIG")
	if configFile == "" {
		configFile = "auth5.yml"
	}
	log.Info().Str("config_file", configFile).Msg("Loading config file")

	// Load configuration from YAML file
	if err := k.Load(file.Provider(configFile), yaml.Parser()); err != nil {
		log.Fatal().Err(err).Msg("Error loading config file")
	}

	// Load environment variables and merge into the loaded config
	if err := k.Load(env.Provider("AUTH5_", ".", func(s string) string {
		return s
	}), nil); err != nil {
		log.Error().Err(err).Msg("Error loading environment variables")
	}

	// Unmarshal the configuration into the Config struct
	if err := k.Unmarshal("", &Cfg); err != nil {
		log.Fatal().Err(err).Msg("Error unmarshalling config")
	}

	// Validate the configuration
	validate := validator.New()
	if err := validate.Struct(Cfg); err != nil {
		log.Fatal().Err(err).Msg("Configuration validation failed")
	}
}
