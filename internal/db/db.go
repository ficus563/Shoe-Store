package db
package db

import (
    "database/sql"
    "os"
    "path/filepath"

    _ "modernc.org/sqlite"
)

func Open(path string) (*sql.DB, error) {
    if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
        return nil, err
    }
    return sql.Open("sqlite", path)
}

func Migrate(db *sql.DB) error {
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS products (
            id          TEXT PRIMARY KEY,
            name        TEXT NOT NULL,
            type        TEXT,
            gender      TEXT,
            size        TEXT,
            color       TEXT,
            description TEXT,
            brand       TEXT,
            season      TEXT,
            price       REAL NOT NULL DEFAULT 0
        );
    `)
    return err
}
