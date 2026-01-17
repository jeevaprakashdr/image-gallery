package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jeevaprakashdr/image-gallery/internal/env"
)

func main() {
	ctx := context.Background()
	config := config{
		address: ":8080",
		db: dbConfig{
			connectionString: env.GetString(
				"GOOSE_DBSTRING",
				"host=localhost user=postgres password=gallery dbname=gallery sslmode=disable"),
		},
	}
	
	conn, err := pgx.Connect(ctx, config.db.connectionString)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	log.Printf("connected to database %s", config.db.connectionString)

	api := application{
		config: config,
	}

	if error := api.run(api.mount()); error != nil {
		log.Printf("server has failed to start, error %s", error)
		os.Exit(1)
	}
}
