package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitPostgres(ctx context.Context) (*pgxpool.Pool, error){
	// Replace with your DSN or load from config
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		dsn = "postgres://nervix:pharmacist@localhost:5432/nervix_db?sslmode=disable"
	}
	// postgresURL := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)
	connPool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to PostgreSQL ✅")

	return connPool, nil
}
