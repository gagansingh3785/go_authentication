package config

import (
	"github.com/gagansingh3785/go_authentication/constants"
	"os"
)

var GlobalConfig Config

type Config struct {
	SendGrid SendGrid
}

type SendGrid struct {
	APIHost     string
	APIKey      string
	APIEndpoint string
}

func InitConfig() {
	GlobalConfig.SetupMailGridConfig()
}

func (c *Config) SetupMailGridConfig() {
	c.SendGrid.APIHost = os.Getenv(constants.SENDGRID_API_HOST)
	c.SendGrid.APIKey = os.Getenv(constants.SENDGRID_API_KEY)
	c.SendGrid.APIEndpoint = os.Getenv(constants.SENDGRID_API_ENDPOINT)
}
