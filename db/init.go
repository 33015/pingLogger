package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// initializeDatabase erstellt oder öffnet die SQLite-Datenbank und legt die Tabelle an, falls diese nicht existiert
func InitializeDatabase(dbFileName string) (*sql.DB, error) {
	if _, err := os.Stat(dbFileName); os.IsNotExist(err) {
		file, err := os.Create(dbFileName)
		if err != nil {
			return nil, err
		}
		file.Close()
	}
	db, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		return nil, err
	}

	query := `
		CREATE TABLE IF NOT EXISTS connectivity_changes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			timestamp TEXT NOT NULL,
			target_name TEXT NOT NULL,
			target_address TEXT NOT NULL,
			status TEXT NOT NULL
		);
	`
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}
	return db, nil
}
