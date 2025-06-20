package main

import (
	"context"
	"log"
	"time"

	"medicare/utility/config"

	"medicare/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"medicare/internal/cmd"
)

func main() {
	// Initialize PostgreSQL connection
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	pool, err := config.InitPostgres(ctx)
	if err != nil{
		log.Printf("Unable to connect to database: %v\n", err)
	}

	defer pool.Close()
	// Initialize SQLC with the PostgreSQL DB
	logic.InitSQLC(pool)

	cmd.Main.Run(gctx.GetInitCtx())
}
