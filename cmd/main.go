package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jangidRkt08/go-Ecom_Prod-API/internal/env"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	 "database/sql"
    _ "github.com/lib/pq"
)

// func main(){
// 	godotenv.Load() 
//     err := goose.Up(db, "./internal/adapters/postgresql/migrations")
//     if err != nil {
//         slog.Error("goose migration failed", "error", err)
//         os.Exit(1)
//     }
// 	ctx := context.Background()
// 	cfg := config{
// 		addr : ":" + os.Getenv("PORT"),
// 		db: dbConfig{
// 			dsn: env.GetString("GOOSE_DBSTRING"),
// 	},
// }

// 	// Structure logging
// 	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))	
// 	slog.SetDefault(logger)


// 	// Database
// 	conn, err := pgx.Connect(ctx, cfg.db.dsn)
// 	if err != nil {
// 		slog.Error("Unable to connect to database", "error", err)
// 		os.Exit(1)
// 	}
// 	defer conn.Close(ctx)
// 	logger.Info("Connected to database", "dsn", cfg.db.dsn)

// 	api:= application{
// 		config: cfg,
// 		db:     conn,
// 	}
	
// 	if err :=api.run(api.mount()); err != nil{
// 		slog.Error("Server failed to start", "error", err)
// 		os.Exit(1)
// 	}

	
// }

func main() {
    godotenv.Load()

    ctx := context.Background()

    cfg := config{
        addr: ":" + env.GetString("PORT"),
        db: dbConfig{
            dsn: env.GetString("GOOSE_DBSTRING"),
        },
    }

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
    slog.SetDefault(logger)

    // Connect using database/sql for Goose
    sqlDB, err := sql.Open("postgres", cfg.db.dsn)
    if err != nil {
        slog.Error("sql.Open failed", "error", err)
        os.Exit(1)
    }

    // Run Goose migrations

    if err := goose.Up(sqlDB, "./internal/adapters/postgresql/migrations"); err != nil {
        slog.Error("goose migration failed", "error", err)
        os.Exit(1)
    }

    // 3️⃣ Connect using pgx for the application
    conn, err := pgx.Connect(ctx, cfg.db.dsn)
    if err != nil {
        slog.Error("Unable to connect to database", "error", err)
        os.Exit(1)
    }
    defer conn.Close(ctx)

    logger.Info("Connected to database", "dsn", cfg.db.dsn)

    api := application{
        config: cfg,
        db:     conn,
    }

    if err := api.run(api.mount()); err != nil {
        slog.Error("Server failed", "error", err)
        os.Exit(1)
    }
}
