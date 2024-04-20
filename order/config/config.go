package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Configration struct {
	DbUrl string
	Port  int
	Env   string
}

func New() *Configration {
	godotenv.Load("../")
	port := os.Getenv("port")
	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}
	DbUrl := os.Getenv("dns")
	env := os.Getenv("env")
	return &Configration{
		DbUrl: DbUrl,
		Port:  portInt,
		Env:   env,
	}
}
