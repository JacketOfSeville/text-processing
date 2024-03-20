package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

/*
Parse the configuration file (config.yml) and return the configuration object so that it can be used in the application.
*/
type Config struct {
	Database struct {
		Uri string `yaml:"uri"`
	} `yaml:"database"`

	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
}

func ConfigFromFile(path string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
