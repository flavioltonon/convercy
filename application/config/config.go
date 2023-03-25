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
	viper.SetDefault("database.uri", "mongodb://database:27017")

	// Server config
	viper.SetDefault("server.address", ":8080")
	viper.SetDefault("server.development_environment", true)

	// OpenExchangeRates config
	viper.SetDefault("openexchangerates.app_id", "")
	viper.SetDefault("openexchangerates.base_url", "https://openexchangerates.org/api")

	// Set all configuration keys to be settable with environment variables. All environment variables have the same name as
	// the configuration keys in uppercase and with dots (".") replaced by underscores ("_"). Example:
	//
	// openexchangerates.base_url => OPENEXCHANGERATES_BASE_URL
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	for _, key := range viper.AllKeys() {
		viper.MustBindEnv(key)
	}
}

type Config struct {
	Database struct {
		Kind string
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

	if err := viper.ReadConfig(f); err != nil {
		return nil, err
	}

	return New()
}

// New creates a new Config from configuration keys default values and their respective environment variables
func New() (*Config, error) {
	var config Config

	// Read environment variables to get default values
	viper.AutomaticEnv()

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
