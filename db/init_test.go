package db

import (
	"os"
	"testing"
)

func TestInitializeDatabase(t *testing.T) {
	// Temporäre Testdatei erstellen
	tempDBFile := "test_db.sqlite"
	defer os.Remove(tempDBFile) // Datei nach dem Test entfernen

	t.Run("Database Creation", func(t *testing.T) {
		db, err := InitializeDatabase(tempDBFile)
		if err != nil {
			t.Fatalf("Fehler bei der Initialisierung der Datenbank: %v", err)
		}
		defer db.Close()

		// Überprüfen, ob die Datei erstellt wurde
		if _, err := os.Stat(tempDBFile); os.IsNotExist(err) {
			t.Errorf("Die Datenbankdatei wurde nicht erstellt: %v", err)
		}

		// Überprüfen, ob die Tabelle existiert
		query := `SELECT name FROM sqlite_master WHERE type='table' AND name='connectivity_changes';`
		row := db.QueryRow(query)
		var tableName string
		err = row.Scan(&tableName)
		if err != nil {
			t.Errorf("Die Tabelle 'connectivity_changes' wurde nicht erstellt: %v", err)
		}
	})

	t.Run("Existing Database", func(t *testing.T) {
		// Datenbank erneut initialisieren
		db, err := InitializeDatabase(tempDBFile)
		if err != nil {
			t.Fatalf("Fehler bei der Initialisierung einer bestehenden Datenbank: %v", err)
		}
		defer db.Close()

		// Überprüfen, dass die Tabelle nicht doppelt erstellt wurde (kein Fehler geworfen)
		query := `SELECT name FROM sqlite_master WHERE type='table' AND name='connectivity_changes';`
		row := db.QueryRow(query)
		var tableName string
		err = row.Scan(&tableName)
		if err != nil || tableName != "connectivity_changes" {
			t.Errorf("Die Tabelle 'connectivity_changes' existiert nicht wie erwartet: %v", err)
		}
	})
}

func TestInitializeDatabase_ErrorHandling(t *testing.T) {
	// Einen ungültigen Pfad testen
	invalidPath := "/invalid_path/test_db.sqlite"

	_, err := InitializeDatabase(invalidPath)
	if err == nil {
		t.Errorf("Erwarteter Fehler bei ungültigem Pfad, erhielt jedoch keinen")
	}
}
