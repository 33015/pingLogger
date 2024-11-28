package db

import (
	"database/sql"
	"time"

	"mkgh.eu/pingLogger/config"
)

// logChange speichert eine Änderung in der SQLite-Datenbank
func LogChange(db *sql.DB, target config.TargetInterface, status string) error {
	query := `INSERT INTO connectivity_changes (timestamp, target_name, target_address, status) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, time.Now().Format(time.RFC3339), target.GetName(), target.GetAddress(), status)
	return err
}
