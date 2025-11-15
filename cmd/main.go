package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jangidRkt08/go-Ecom_Prod-API/internal/env"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load() 
	ctx := context.Background()
	cfg := config{
		addr : ":" + os.Getenv("PORT"),
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING"),
	},
}

	// Structure logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))	
	slog.SetDefault(logger)


	// Database
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		slog.Error("Unable to connect to database", "error", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)
	logger.Info("Connected to database", "dsn", cfg.db.dsn)

	api:= application{
		config: cfg,
		db:     conn,
	}
	
	if err :=api.run(api.mount()); err != nil{
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}

	
}
