package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/criptycpizza7/habit-tracker/common"
	"go.yaml.in/yaml/v3"
)

type IConfig interface {
	ServiceName() string
}

func path() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	path := filepath.Join(wd, common.CONFIG_PATH)
	return path, nil
}

func readFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func loadServiceConfig(raw_full_config []byte, service_name string) ([]byte, error) {
	config_map := make(map[string]any)

	err := yaml.Unmarshal(raw_full_config, &config_map)
	if err != nil {
		return []byte(""), err
	}

	raw_config_map, ok := config_map[service_name].(map[string]any)
	if !ok {
		return []byte(""), fmt.Errorf("service config %s not found", service_name)
	}

	raw_config, err := yaml.Marshal(raw_config_map)
	if err != nil {
		return []byte(""), err
	}
	return raw_config, nil
}

func MustLoadConfig(config IConfig) {
	config_path, err := path()
	if err != nil {
		panic(err)
	}

	raw_full_config, err := readFile(config_path)
	if err != nil {
		panic(err)
	}

	raw_config, err := loadServiceConfig(raw_full_config, config.ServiceName())
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(raw_config, config)
	if err != nil {
		panic(err)
	}
}
