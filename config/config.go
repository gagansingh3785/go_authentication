package config

import (
	"fmt"
	"os"
)

var GlobalConfig Config

type DBConfig struct {
	Host       string
	Port       string
	DBName     string
	DBPassword string
	DBUser     string
	SSLMode    string
}

type SendGrid struct {
	APIHost     string
	APIKey      string
	APIEndpoint string
}

type Config struct {
	DBConfig DBConfig
}

func SetupDBConfig() DBConfig {
	dbConfig := DBConfig{
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBUser:     os.Getenv("DB_USER"),
		SSLMode:    os.Getenv("SSL_MODE"),
	}
	return dbConfig
}

func InitConfig() {
	GlobalConfig.DBConfig = SetupDBConfig()
	fmt.Printf("%+v \n", GlobalConfig.DBConfig)
}
