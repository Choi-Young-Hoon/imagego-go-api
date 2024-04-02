package util

import (
	"gopkg.in/yaml.v2"
	"os"
)

var defaultConfig = NewServerConifg()

func LoadConfig() error {
	return defaultConfig.LoadConfig()
}

func GetServerConfig() ServerConfig {
	return defaultConfig
}

func GenerateDefaultServerConfig() {
	config := ServerConfig{
		Port:           8080,
		ImageServerUrl: "http://192.168.219.103:8080",
		ImageDir:       "img",
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
	ImageServerUrl string         `yaml:"image_server_url"`
	ImageDir       string         `yaml:"image_dir"`
	Port           int            `yaml:"port"`
	Database       DatabaseConfig `yaml:"database"`
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
