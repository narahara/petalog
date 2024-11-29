package config

import (
	"os"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
)

type ConfigDb struct {
	Host     string `yaml:"host" validate:"required"`
	Port     int    `yaml:"port" validate:"required"`
	User     string `yaml:"user" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	Database string `yaml:"database" validate:"required"`
}

type Config struct {
	Stage string   `yaml:"stage" validate:"required"`
	Db    ConfigDb `yaml:"db" validate:"required"`
	Port  int      `yaml:"port" validate:"required"`
}

func LoadConfig(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(b, config)
	if err != nil {
		return nil, err
	}

	// バリデーションを実行
	validate := validator.New()
	err = validate.Struct(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
