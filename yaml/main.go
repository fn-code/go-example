package main

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	HTTPS    bool   `mapstructure:"https"`
	CertFile string `mapstructure:"certFile"`
	KeyFile  string `mapstructure:"keyFile"`
}

type DatabaseConfig struct {
	MaxOpenConn     int           `mapstructure:"maxOpenConn"`
	MaxIdleConn     int           `mapstructure:"maxIdleConn"`
	ConnMaxLifetime time.Duration `mapstructure:"connMaxLifetime"`
	SslMode         string        `mapstructure:"sslMode"`
}

type LogConfig struct {
	LogFolder string `mapstructure:"logFolder"`
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Logger   LogConfig
}

func main() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var cfg Config
	viper.Unmarshal(&cfg)

	fmt.Println(cfg)
}
