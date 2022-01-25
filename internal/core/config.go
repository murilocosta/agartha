package core

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Database DatabaseConfig `yaml:"database"`
	Cache    CacheConfig    `yaml:"cache"`
	Auth     AuthConfig     `yaml:"auth"`
}

type DatabaseConfig struct {
	Driver   string `yaml:"driver"`
	Username string `yaml:"username" envconfig:"DATABASE_USERNAME"`
	Password string `yaml:"password" envconfig:"DATABASE_PASSWORD"`
	Host     string `yaml:"host" envconfig:"DATABASE_HOST"`
	Port     string `yaml:"port" envconfig:"DATABASE_PORT"`
	DbName   string `yaml:"dbname" envconfig:"DATABASE_NAME"`
	SslMode  string `yaml:"ssl-mode"`
}

type CacheConfig struct {
	Enabled           bool   `yaml:"enabled"`
	Host              string `yaml:"host"`
	Password          string `yaml:"password"`
	DatabaseSelection int    `yaml:"database-selection"`
}

type AuthConfig struct {
	Realm          string `yaml:"realm"`
	SecretKey      string `yaml:"secret-key"`
	TokenTimeout   int32  `yaml:"token-timeout"`
	RefreshTimeout int32  `yaml:"refresh-timeout"`
}

func LoadConfig(yamlFilePath string, cfg *Config) error {
	err := readYamlFile(yamlFilePath, cfg)
	if err != nil {
		return err
	}

	// Should override the values with environment variables
	err = envconfig.Process("", cfg)
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
