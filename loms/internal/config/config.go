package config

import (
	"os"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v3"
)

type ConfigStruct struct {
	Token string `yaml:"token"`
	Ports struct {
		Http string `yaml:"http"`
		Grpc string `yaml:"grpc"`
	} `yaml:"ports"`
	DBConnectURL string `yaml:"db_connect_url"`
	Kafka        struct {
		Brokers []string `yaml:"brokers"`
		Topic   string   `yaml:"topic"`
	} `yaml:"kafka"`
}

var ConfigData ConfigStruct

func Init() error {
	rawYAML, err := os.ReadFile("config.yml")
	if err != nil {
		return errors.WithMessage(err, "reading config file")
	}

	err = yaml.Unmarshal(rawYAML, &ConfigData)
	if err != nil {
		return errors.WithMessage(err, "parsing yaml")
	}

	return nil
}
