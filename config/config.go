package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Token struct {
			Expiration int `yaml:"expiration"`
		} `yaml:"token"`
	} `yaml:"app"`
}

type Token struct {
	Expiration int
}

func Load(path, filename string) (*Config, error) {
	route := fmt.Sprintf("%s%s", path, filename)
	file, err := os.Open(route)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	var config Config
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
