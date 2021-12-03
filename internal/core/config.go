package core

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database DatabaseConfig `json:"database"`
	Cache    CacheConfig    `json:"cache"`
	Auth     AuthConfig     `json:"auth"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"dbname"`
}

type CacheConfig struct {
	Host              string `json:"host"`
	Password          string `json:"password"`
	DatabaseSelection int    `json:"database-selection"`
}

type AuthConfig struct {
	Realm          string `json:"realm"`
	SecretKey      string `json:"secret-key"`
	TokenTimeout   int32  `json:"token-timeout"`
	RefreshTimeout int32  `json:"refresh-timeout"`
}

func LoadConfig(yamlFilePath string, cfg *Config) error {
	err := readYamlFile(yamlFilePath, cfg)
	if err != nil {
		return err
	}

	return nil
}

func readYamlFile(filepath string, cfg *Config) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return err
	}

	return nil
}
