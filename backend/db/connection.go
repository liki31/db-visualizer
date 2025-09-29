package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DBConnection struct {
	DB *sql.DB
}

func ConnectPostgres(host string, port int, user, password, dbname string) (*DBConnection, error) {
	fmt.Printf("Connecting to database %s at Host-%s:password-%d as user %s\n", dbname, host, port, user)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return &DBConnection{DB: db}, nil
}