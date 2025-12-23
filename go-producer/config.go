package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	RabbitMQ RabbitMQConfig `yaml:"rabbitmq"`
}

type RabbitMQConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Queue    string `yaml:"queue"`
}

func loadConfig() (*Config, error) {
	// 优先使用本地配置文件（会被gitignore），如果不存在则使用线上配置文件
	configFile := "config.local.yaml"
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		configFile = "config.prod.yaml"
		log.Printf("使用线上配置文件: %s", configFile)
	} else {
		log.Printf("使用本地配置文件: %s", configFile)
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	return &config, nil
}

func (c *RabbitMQConfig) ConnectionString() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", c.Username, c.Password, c.Host, c.Port)
}

