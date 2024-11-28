package config

import (
	"os"
	"testing"
)

func TestLoadTargets(t *testing.T) {
	// Testdaten vorbereiten
	validJSON := `[{"name": "Google", "address": "google.com"}, {"name": "GitHub", "address": "github.com"}]`
	invalidJSON := `[{name: "Google", "address": "google.com"}]` // Falsches JSON-Format
	nonexistentFile := "nonexistent.json"

	// Temporäre Datei für validen JSON-Inhalt
	validFile := "valid_targets.json"
	err := os.WriteFile(validFile, []byte(validJSON), 0644)
	if err != nil {
		t.Fatalf("Fehler beim Erstellen der Testdatei: %v", err)
	}
	defer os.Remove(validFile)

	// Test: Valides JSON
	t.Run("Valid JSON", func(t *testing.T) {
		targets, err := LoadTargets(validFile)
		if err != nil {
			t.Errorf("Erwartete kein Fehler, erhielt: %v", err)
		}
		if len(targets) != 2 {
			t.Errorf("Erwartete 2 Targets, erhielt: %d", len(targets))
		}
		if targets[0].Name != "Google" || targets[0].Address != "google.com" {
			t.Errorf("Erwartete Target 'Google', erhielt: %+v", targets[0])
		}
	})

	// Test: Ungültiges JSON
	t.Run("Invalid JSON", func(t *testing.T) {
		invalidFile := "invalid_targets.json"
		err := os.WriteFile(invalidFile, []byte(invalidJSON), 0644)
		if err != nil {
			t.Fatalf("Fehler beim Erstellen der Testdatei: %v", err)
		}
		defer os.Remove(invalidFile)

		_, err = LoadTargets(invalidFile)
		if err == nil {
			t.Error("Erwartete einen Fehler, erhielt keinen")
		}
	})

	// Test: Datei existiert nicht
	t.Run("Nonexistent File", func(t *testing.T) {
		_, err := LoadTargets(nonexistentFile)
		if err == nil {
			t.Error("Erwartete einen Fehler, erhielt keinen")
		}
	})
}
