package conf

import (
	"github.com/spf13/viper"
	"strings"
	"template.com/restapi/internal/logger"
)

type AppConfig struct {
	AppName        string `mapstructure:"app_name"`
	Db             DB     `mapstructure:"db"`
	DefaultTimeOut int    `mapstructure:"default_time_out"`
	Port           int    `mapstructure:"port"`
	BasePath       string `mapstructure:"base_path"`
	Version        string `mapstructure:"version"`
}

// DB holds the database connection configuration.
type DB struct {
	URL  string `mapstructure:"url"`
	Port string `mapstructure:"port"`
	Name string `mapstructure:"name"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
}

func LoadConfig(path string) (*AppConfig, error) {
	return getViperConfig(path)
}

// readingConfigFromExternalVault better management for secrets
// for this simple example I will work only with ENVs
func readingConfigFromExternalVault() (*AppConfig, error) {
	// like AWS manager... https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieving-secrets-go-sdk.html
	return nil, nil
}

// reading config file with Viper. Setting sensitive values as env variables
func getViperConfig(path string) (*AppConfig, error) {
	l := logger.L.WithGroup("viper_config")
	// Set up viper to read the config.yaml file
	if path == "" {
		path = "./conf/"
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			l.Warn("Config file not found")
		} else {
			l.Warn("Config file was found but another error was produced")
		}
		return nil, err
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	var c AppConfig
	err := viper.Unmarshal(&c)
	if err != nil {
		l.Error("unable to decode into struct", "error:", err)
		return nil, err
	}

	return &c, nil
}
