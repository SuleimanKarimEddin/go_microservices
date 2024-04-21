package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	
	db, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("connected")

}
