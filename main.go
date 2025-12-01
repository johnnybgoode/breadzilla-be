package main

import (
	"context"
	"log"

	"github.com/johnnybgoode/breadzilla/internal/api"
	"github.com/johnnybgoode/breadzilla/pkg/database"
	"github.com/johnnybgoode/breadzilla/pkg/server"
)

func main() {
	ctx := context.Background()
	var config database.Config
	config.ProcessFromEnv(ctx)
	db, err := database.Connect(&config)
	if err != nil {
		log.Fatal(err)
	}

	server := server.NewServer(":8080", db)
	api.AddRoutes(server)

	server.Start()
}