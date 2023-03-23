package config

import (
	"fmt"
	"github.com/joomcode/errorx"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

// ConfigProvider Provider interface for retrieving service configuration
type ConfigProvider interface {
	GetStringOrDefault(path string, defaultValue string) string
	GetString(path string) string
	GetInt(path string) int
	GetBool(path string) bool
	GetEnv(path string) string
}

type LocalConfigProvider struct {
	config *viper.Viper
}

func NewLocalConfigProvider() ConfigProvider {
	provider := &LocalConfigProvider{
		config: viper.New(),
	}

	provider.initConfig()

	return provider
}

func (c *LocalConfigProvider) GetStringOrDefault(path string, defaultValue string) string {
	if v := c.config.GetString(path); v == "" {
		return defaultValue
	} else {
		return v
	}
}

func (c *LocalConfigProvider) GetString(path string) string {
	return c.config.GetString(path)
}

func (c *LocalConfigProvider) GetInt(path string) int {
	return c.config.GetInt(path)
}

func (c *LocalConfigProvider) GetBool(path string) bool {
	return c.config.GetBool(path)
}

func (c *LocalConfigProvider) GetEnv(path string) string {
	return fmt.Sprint(c.config.Get(path))
}

func (c *LocalConfigProvider) initConfig() {
	log.Info("Initializing service configuration")

	c.config.SetConfigFile("yml")
	c.config.SetConfigName("config")
	c.config.AddConfigPath(".")
	c.config.AddConfigPath("./resources")

	if err := c.config.ReadInConfig(); err != nil {
		log.Warn(errorx.Decorate(err, "Failed to load configuration file."))
	}

	c.config.AutomaticEnv()
}
