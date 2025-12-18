package main

import (
	"context"

	"github.com/johnnybgoode/breadzilla/internal/api"
	"github.com/johnnybgoode/breadzilla/pkg/common"
	"github.com/johnnybgoode/breadzilla/pkg/database"
	"github.com/johnnybgoode/breadzilla/pkg/server"
)

func main() {
	ctx := context.Background()
	var config database.Config
	config.ProcessFromEnv(ctx)
	db := common.Must(database.Connect(&config))

	server := server.NewServer(":3000", db)
	server.ApplyRoutes(api.Routes)

	server.Start()
}
