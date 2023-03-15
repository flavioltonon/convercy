package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	// Database config
	viper.SetDefault("database.kind", "mongodb")
	viper.SetDefault("database.uri", "mongodb://localhost:27017")

	// Server config
	viper.SetDefault("server.address", ":8000")
	viper.SetDefault("server.development_environment", true)

	// OpenExchangeRates config
	viper.SetDefault("openexchangerates.base_url", "https://openexchangerates.org/api")
}

type Config struct {
	Database struct {
		Kind string
		Name string
		URI  string
	}

	Server struct {
		Address                string
		DevelopmentEnvironment bool `mapstructure:"development_environment"`
	}

	OpenExchangeRates struct {
		AppID   string `mapstructure:"app_id"`
		BaseURL string `mapstructure:"base_url"`
	} `mapstructure:"openexchangerates"`
}

// FromFile creates a new config from a given file
func FromFile(path string) (*Config, error) {
	var config Config

	// If the path to the file is not absolute, we should look for the file from the current working directory
	if !filepath.IsAbs(path) {
		workdir, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		path = filepath.Join(workdir, filepath.Clean(path))
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file at %s: %v", path, err)
	}

	defer f.Close()

	// Set viper config file extension. Since filepath.Ext returns an extension preceded by a dot (e.g. ".yaml"),
	// we have to trim it manually.
	viper.SetConfigType(strings.TrimLeft(filepath.Ext(path), "."))

	// Read environment variables to get default values
	viper.AutomaticEnv()

	if err := viper.ReadConfig(f); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
