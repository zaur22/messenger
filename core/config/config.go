package config

import (
	"github.com/spf13/viper"
)

var (
	DBUser     string
	DBName     string
	DBSSLMode  string
	DBPassword string
	DBHost     string
	DBPort     string
)

func init() {
	viper.AutomaticEnv()

	viper.SetDefault("DB_USER", "user")
	DBUser = viper.GetString("DB_USER")

	viper.SetDefault("DB_NAME", "testdb")
	DBName = viper.GetString("DB_NAME")

	viper.SetDefault("DB_PASSWORD", "password")
	DBPassword = viper.GetString("DB_PASSWORD")

	viper.SetDefault("SSL_MODE", "disable")
	DBSSLMode = viper.GetString("SSL_MODE")

	viper.SetDefault("DB_HOST", "localhost")
	DBHost = viper.GetString("DB_HOST")

	viper.SetDefault("DB_PORT", "5432")
	DBPort = viper.GetString("DB_PORT")

}
