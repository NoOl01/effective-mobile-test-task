package config

import (
	"log"
	"os"
)

type Config struct {
	DbUser     string
	DbHost     string
	DbPass     string
	DbPort     string
	DbName     string
	DbSslMode  string
	ServerPort string
}

var Env = &Config{}

func LoadEnv() {
	Env.DbUser = os.Getenv("DB_USER")
	Env.DbHost = os.Getenv("DB_HOST")
	Env.DbPass = os.Getenv("DB_PASS")
	Env.DbPort = os.Getenv("DB_PORT")
	Env.DbName = os.Getenv("DB_NAME")
	Env.DbSslMode = os.Getenv("DB_SSL_MODE")
	Env.ServerPort = os.Getenv("SERVER_PORT")

	log.Println(".env config loaded")
}
