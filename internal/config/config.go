package config

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Goticker : General configuration node
type Goticker struct {
	Refresh string `yaml:"refresh"`
}

// HTTPConfig : HTTP exposed apps configuration node
type HTTPConfig struct {
	Address                 string `yaml:"address"`
	CookieAuthenticationKey string `yaml:"cookieAuthenticationKey"`
	CookieEncryptionKey     string `yaml:"cookieEncryptionKey"`
	CookieMaxAge            int    `yaml:"cookieMaxAge"`
	TemplateDir             string `yaml:"templateDir"`
	PublicDir               string `yaml:"publicDir"`
}

// Config : Configuration object
type Config struct {
	Goticker Goticker   `yaml:"goticker"`
	HTTP     HTTPConfig `yaml:"http"`
}

// CreateDefault : Create default configuration
func CreateDefault() *Config {
	return &Config{
		Goticker{
			Refresh: "10",
		},
		HTTPConfig{
			Address:                 ":3000",
			CookieAuthenticationKey: "",
			CookieEncryptionKey:     "",
			CookieMaxAge:            int((time.Hour * 1).Seconds()), // 1 hour
			TemplateDir:             "template",
			PublicDir:               "public/dist",
		},
	}
}

// NewFromFile retrieves the configuration from the given file
func NewFromFile(filepath string) (*Config, error) {
	config := CreateDefault()

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		config, err := Dump(config, filepath)
		check(err)
		return config, nil
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, errors.Wrap(err, "[Config] Could not unmarshal configuration")
	}

	return config, nil
}

// Dump : write configuration to file
func Dump(config *Config, filepath string) (*Config, error) {
	data, err := yaml.Marshal(config)
	if err != nil {
		return nil, errors.Wrap(err, "[Config] Could not dump config")
	}

	f, err := os.Create(filepath)
	check(err)

	_, err = f.Write(data)
	check(err)

	return config, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
