package util

import (
	"gopkg.in/yaml.v2"
	"os"
)

func GenerateDefaultServerConfig() {
	config := ServerConfig{
		Port: 8080,
		Database: DatabaseConfig{
			Host:     "localhost",
			User:     "choi",
			Password: "123456",
			DBName:   "postgres",
			Port:     5432,
		},
	}

	yamlData, err := yaml.Marshal(config)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("config.yaml", yamlData, 0755)
}

func NewServerConifg() ServerConfig {
	return ServerConfig{}
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Port     int    `yaml:"port"`
}

type ServerConfig struct {
	Port     int            `yaml:"port"`
	Database DatabaseConfig `yaml:"database"`
}

func (sc *ServerConfig) LoadConfig() error {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, sc)
	if err != nil {
		return err
	}

	return nil
}
