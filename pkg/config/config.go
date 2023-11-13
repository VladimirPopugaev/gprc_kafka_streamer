package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Env    string `yaml:"env"`
	Server Server `yaml:"server"`
}

type Server struct {
	Address string `yaml:"address"`
}

func New() (*Config, error) {
	const op = "config.New"

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return nil, fmt.Errorf("CONFIG_PATH is not set")
	}

	err := validConfigPath(configPath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("%s: file is not open. %w", op, err)
	}
	defer func() { _ = file.Close() }()

	var cfg Config
	configDecoder := yaml.NewDecoder(file)

	if err := configDecoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("%s: decode fault. %w", op, err)
	}

	return &cfg, nil
}

func validConfigPath(path string) error {
	const op = "config.ValidConfigPath"

	fileInfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if fileInfo.IsDir() {
		return fmt.Errorf("%s: it is directory. You need file", op)
	}

	return nil
}
