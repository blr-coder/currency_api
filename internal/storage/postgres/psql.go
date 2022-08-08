package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
	// DB driver
	_ "github.com/lib/pq"
)

const (
	Timeout = 3 * time.Second
)

func NewDB(dbHost, dbPort, dbName, dbUser, dbUserPassword string) (*sqlx.DB, error) {
	dbURI := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		dbHost,
		dbPort,
		dbName,
		dbUser,
		dbUserPassword,
	)
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(60)
	db.SetMaxOpenConns(60)
	db.SetConnMaxLifetime(5 * time.Minute)

	newSQLX := sqlx.NewDb(db, "postgres")
	return newSQLX, nil
}
