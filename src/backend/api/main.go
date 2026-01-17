package main

import (
	"log"
	"os"
)

func main() {
	config := config{
		address: ":8080",
		db:      dbConfig{},
	}

	api := application{
		config: config,
	}
	
	if error := api.run(api.mount()) ; error != nil {
		log.Printf("server has failed to start, error %s", error)
		os.Exit(1)
	}
}
