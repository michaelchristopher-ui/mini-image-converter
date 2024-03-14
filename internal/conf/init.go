package conf

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var cfg = new(Config)

func GetConfig() *Config {
	return cfg
}

func Init(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading config file, %s", err)
	}

	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return fmt.Errorf("unable to decode into struct, %v", err)
	}
	return nil
}
