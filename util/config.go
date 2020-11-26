package util

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Website  string   `yaml:"website"`
	Check    string   `yaml:"check"`
	Keywords []string `yaml:"keywords"`
	Interval string   `yaml:"interval"`
	Trigger  Trigger  `yaml:"trigger"`
}

type Trigger struct {
	Event string `yaml:"event"`
	Key   string `yaml:"key"`
}

func ReadConfig(path string) (*Config, error) {
	config := &Config{}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
