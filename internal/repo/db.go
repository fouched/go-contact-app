package repo

import (
	"database/sql"
	"time"
)

// db is the database connection pool
var db *sql.DB

func CreateDbPool(dsn string) (*sql.DB, error) {
	// no error thrown even host or db does not exist
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(5)
	conn.SetConnMaxLifetime(5 * time.Minute)

	// do a real test to see if we have a db conn
	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	db = conn

	return conn, nil
}
