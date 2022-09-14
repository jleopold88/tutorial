package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Cfg struct {
	DBHost     string
	DBPort     int
	DBName     string
	DBPassword string
	DBUser     string
	GRPCSERVER string
}

func NewConfig() *Cfg {
	if err := godotenv.Load("config/.env"); err != nil {
		panic(err)
	}
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	return &Cfg{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     port,
		DBName:     os.Getenv("DB_NAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBUser:     os.Getenv("DB_USER"),
		GRPCSERVER: os.Getenv("GRPC_SERVER"),
	}
}
