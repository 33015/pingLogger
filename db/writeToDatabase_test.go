package db

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

// Simulierte Target-Struktur aus dem config-Paket
type Target struct {
	Name    string
	Address string
}

func (t Target) GetName() string {
	return t.Name
}

func (t Target) GetAddress() string {
	return t.Address
}

func TestLogChange(t *testing.T) {
	// Temporäre SQLite-Datenbank
	tempDBFile := "test_logchange.db"
	defer os.Remove(tempDBFile) // Datenbankdatei nach dem Test entfernen

	// Datenbank initialisieren
	db, err := sql.Open("sqlite3", tempDBFile)
	if err != nil {
		t.Fatalf("Fehler beim Öffnen der Datenbank: %v", err)
	}
	defer db.Close()

	// Tabelle erstellen
	query := `
	CREATE TABLE IF NOT EXISTS connectivity_changes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		timestamp TEXT NOT NULL,
		target_name TEXT NOT NULL,
		target_address TEXT NOT NULL,
		status TEXT NOT NULL
	);`
	_, err = db.Exec(query)
	if err != nil {
		t.Fatalf("Fehler beim Erstellen der Tabelle: %v", err)
	}

	// Testziel definieren
	target := Target{Name: "Google", Address: "google.com"}
	status := "reachable"

	// Funktion aufrufen
	err = LogChange(db, target, status)
	if err != nil {
		t.Errorf("Fehler bei LogChange: %v", err)
	}

	// Überprüfen, ob der Eintrag in der Datenbank ist
	row := db.QueryRow(`SELECT target_name, target_address, status FROM connectivity_changes WHERE target_name = ?`, target.Name)
	var name, address, loggedStatus string
	err = row.Scan(&name, &address, &loggedStatus)
	if err != nil {
		t.Errorf("Fehler beim Abrufen des Eintrags: %v", err)
	}

	// Einträge vergleichen
	if name != target.GetName() || address != target.GetAddress() || loggedStatus != status {
		t.Errorf("Erwartete Einträge: %v, %v, %v; Erhalten: %v, %v, %v",
			target.Name, target.Address, status, name, address, loggedStatus)
	}
}

func TestLogChange_NoTable(t *testing.T) {
	// Temporäre SQLite-Datenbank
	tempDBFile := "test_no_table.db"
	defer os.Remove(tempDBFile) // Datenbankdatei nach dem Test entfernen

	// Datenbank initialisieren
	db, err := sql.Open("sqlite3", tempDBFile)
	if err != nil {
		t.Fatalf("Fehler beim Öffnen der Datenbank: %v", err)
	}
	defer db.Close()

	// Testziel definieren
	target := Target{Name: "Google", Address: "google.com"}
	status := "reachable"

	// Funktion aufrufen (ohne Tabelle)
	err = LogChange(db, target, status)
	if err == nil {
		t.Error("Erwarteter Fehler, da die Tabelle fehlt, erhielt jedoch keinen")
	}
}
