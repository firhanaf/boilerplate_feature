package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
)

var JWTSecret string

type AppConfig struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string
	JWTSecret  string
}

func InitConfig() *AppConfig {
	return ReadENV()
}

func ReadENV() *AppConfig {
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("DB_USERNAME"); found && val != "" {
		app.DBUsername = val
		isRead = false
	}
	if val, found := os.LookupEnv("DB_PASSWORD"); found && val != "" {
		app.DBPassword = val
		isRead = false
	}
	if val, found := os.LookupEnv("DB_HOST"); found && val != "" {
		app.DBHost = val
		isRead = false
	}
	if val, found := os.LookupEnv("DB_PORT"); found && val != "" {
		app.DBPort, _ = strconv.Atoi(val)
		isRead = false
	}
	if val, found := os.LookupEnv("DB_NAME"); found && val != "" {
		app.DBName = val
		isRead = false
	}
	if val, found := os.LookupEnv("JWT_SECRET"); found && val != "" {
		app.JWTSecret = val
	}
	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")
		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error to read config: ", err.Error())
			return nil
		}
		app.DBUsername = viper.GetString("DB_USERNAME")
		app.DBPassword = viper.GetString("DB_PASSWORD")
		app.DBHost = viper.GetString("DB_HOST")
		app.DBPort = viper.GetInt("DB_PORT")
		app.DBName = viper.GetString("DB_NAME")
		app.JWTSecret = viper.GetString("JWT_SECRET")
	}
	JWTSecret = app.JWTSecret
	return &app
}
