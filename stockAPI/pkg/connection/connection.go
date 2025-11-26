package connection

import (
    "database/sql"
    "fmt"
    "os"

    "github.com/joho/godotenv"
    _ "github.com/lib/pq" 
)

var db *sql.DB

func connectDB() error {
    if err := godotenv.Load(); err != nil {
        return fmt.Errorf("failed to load .env file: %w", err)
    }

    dbUrl := os.Getenv("POSTGRESQL_URL")
    if dbUrl == "" {
        return fmt.Errorf("POSTGRESQL_URL not set in environment")
    }

    d, err := sql.Open("postgres", dbUrl)
    if err != nil {
        return fmt.Errorf("failed to open connection: %w", err)
    }

    if err := d.Ping(); err != nil {
        return fmt.Errorf("failed to ping database: %w", err)
    }

    db = d
    fmt.Println("Connected to Postgres successfully")
    return nil
}

func Connection() (*sql.DB, error) {
    if db == nil {
        if err := connectDB(); err != nil {
            return nil, err
        }
    }
    return db, nil
}
