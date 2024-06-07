package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var configInterface *ConfigSchema

const configFile = "reblog.yml"

type ServerConfig struct {
	Port    int  `yaml:"port"`
	Prefork bool `yaml:"prefork"`
}

type DBConfig struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"pass"`
}

type DashboardConfig struct {
	Enable bool `yaml:"enable"`
}

type ConfigSchema struct {
	Dev       bool            `yaml:"dev"`
	Server    ServerConfig    `yaml:"server"`
	DB        DBConfig        `yaml:"db"`
	Dashboard DashboardConfig `yaml:"dashboard"`
}

func Config() *ConfigSchema {
	if configInterface == nil {
		InitConfig()
	}

	return configInterface
}

func (c *ConfigSchema) SaveConfig() {
	configFile, err := os.OpenFile(configFile, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer configFile.Close()

	configBytes, err := yaml.Marshal(c)

	if err != nil {
		panic(err)
	}

	_, err = configFile.Write(configBytes)

	if err != nil {
		panic(err)
	}
}

func (c *ConfigSchema) LoadConfig() error {
	configFile, err := os.ReadFile(configFile)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configFile, c)

	return err
}

func DefuaulConfig() *ConfigSchema {
	return &ConfigSchema{
		Dev: false,
		Server: ServerConfig{
			Port:    3000,
			Prefork: true,
		},
		DB: DBConfig{
			Type:     "mysql",
			Host:     "localhost",
			Port:     3306,
			Name:     "reblog",
			User:     "reblog",
			Password: "reblog",
		},
		Dashboard: DashboardConfig{
			Enable: true,
		},
	}
}

func InitConfig() {
	_, err := os.Stat(configFile)

	if os.IsNotExist(err) {
		DefuaulConfig().SaveConfig()
	}

	configInterface = &ConfigSchema{}

	configInterface.LoadConfig()
}
