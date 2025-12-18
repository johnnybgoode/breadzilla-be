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

	config := new(database.Config)
	common.Must[any](nil, config.ProcessFromEnv(&ctx))
	db := common.Must(database.Connect(config))

	server := server.NewServer(&ctx, db)
	server.ApplyRoutes(api.Routes)

	server.Start()
}
