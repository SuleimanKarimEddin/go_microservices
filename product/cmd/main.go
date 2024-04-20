package main

import (
	"context"
	"fmt"
	"product/config"
	"product/internal/adapters/db/generated"
	"product/internal/adapters/grpc"
	"product/internal/application/core/api"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config := config.New()
	db, err := pgxpool.New(context.Background(), config.DbUrl)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		err = db.Ping(context.Background())
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		panic(err)
	}
	dbAdatpter := generated.New(db)
	apiAdapter := api.New(dbAdatpter)
	grpcServer := grpc.New(config.Port, apiAdapter)
	fmt.Println("server is running on port: ", config.Port)
	err = grpcServer.Run()
	if err != nil {
		panic(err)
	}
}
