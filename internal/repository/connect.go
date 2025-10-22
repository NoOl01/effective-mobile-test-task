package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"testovoe/internal/config"
)

func Connect() *pgxpool.Pool {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.Env.DbUser, config.Env.DbPass, config.Env.DbHost, config.Env.DbPort, config.Env.DbName, config.Env.DbSslMode)

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("connect failed: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("ping failed: %v", err)
	}

	log.Println("Database connected")
	return pool
}
