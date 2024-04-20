package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Configration struct {
	Port  int
	Env   string
	OrderUrl string
	ProductUrl string
}

func New() *Configration {
	godotenv.Load("../")
	port := os.Getenv("port")
	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}
	env := os.Getenv("env")
	OrderUrl := os.Getenv("order_url")
	ProductUrl := os.Getenv("product_url")
	return &Configration{
		Port:  portInt,
		Env:   env,
		OrderUrl: OrderUrl,
		ProductUrl: ProductUrl,
	}
}
