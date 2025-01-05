package config

import (
    "database/sql"

    "log"
    "os"

    _ "github.com/lib/pq" 
)

var Database *sql.DB

func ConnectDB() *sql.DB {
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        log.Fatal("DATABASE_URL is not set")
    }

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    if err := db.Ping(); err != nil {
        log.Fatalf("Cannot ping database: %v", err)
    }

    log.Println("Connected to the database successfully!")
    Database = db
    return db
}