package main

import (
	"context"
	"log"
	"os"

	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5"
	"github.com/jeevaprakashdr/image-gallery/services/env"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config := config{
		address: ":8080",
		db: dbConfig{
			connectionString: env.GetString(
				"GOOSE_DBSTRING",
				"host=localhost user=postgres password=postgres dbname=gallery sslmode=disable"),
		},
	}

	conn, err := pgx.Connect(ctx, config.db.connectionString)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	log.Printf("connected to database %s", config.db.connectionString)

	wsConn, _, err := websocket.DefaultDialer.Dial("ws://localhost:5000/ws", nil)
	if err != nil {
		log.Fatal("WebSocket dial error:", err)
	}
	defer wsConn.Close()

	api := application{
		config:             config,
		db:                 conn,
		wsClientConnection: wsConn,
	}

	if error := api.run(api.mount()); error != nil {
		log.Printf("server has failed to start, error %s", error)
		os.Exit(1)
	}
}
