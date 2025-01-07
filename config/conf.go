package config

import (
	"gin-sample-framework/pkg/utils"

	"github.com/spf13/viper"
)

var (
	Configuration Config
)

type Config struct {
	Name string `json:"name" yaml:"name"`
	Env  string `json:"env" yaml:"env" `

	Jwt struct {
		Secret string `json:"secret" yaml:"secret" mapstructure:"secret"`
		Expire int    `json:"expire" yaml:"expire" mapstructure:"expire"`
	} `json:"jwt" yaml:"jwt"`

	Jaeger struct {
		Enabled bool   `json:"enabled" yaml:"enabled" mapstructure:"enabled"`
		Host    string `json:"host" yaml:"host" mapstructure:"host"`
		Port    string `json:"port" yaml:"port" mapstructure:"port"`
	} `json:"jaeger" yaml:"jaeger"`

	DB struct {
		MySQL struct {
			Host     string `json:"host" yaml:"host" mapstructure:"host"`
			Port     int    `json:"port" yaml:"port" mapstructure:"port"`
			Username string `json:"username" yaml:"username" mapstructure:"username"`
			Password string `json:"password" yaml:"password" mapstructure:"password"`
			DBName   string `json:"dbname" yaml:"dbname" mapstructure:"dbname"`
		} `json:"mysql" yaml:"mysql"`
	} `json:"db" yaml:"db"`
}

func Init() error {
	var (
		confAddrYaml = utils.PathJoin("./config", "config.yaml")
		err          error
	)
	viper.SetConfigFile(confAddrYaml)
	if err = viper.ReadInConfig(); err != nil {
		return err
	}
	if err = viper.Unmarshal(&Configuration); err != nil {
		return err
	}

	return nil
}
