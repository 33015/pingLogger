package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"mkgh.eu/pingLogger/config"
	"mkgh.eu/pingLogger/db"
	"mkgh.eu/pingLogger/network"
)

func main() {
	checkInterval := config.PingInterval * time.Second // Intervall zwischen Pings
	daba, targets := initialize()
	log.Println("Ping-Logger mit folgenden Zielen gestartet")
	for _, target := range targets {
		log.Println(target.Name, target.Address)
	}
	defer daba.Close()

	// Status-Tracking für alle Ziele
	lastStatuses := make(map[string]string)

	for {
		for _, target := range targets {
			currentStatus := network.CheckPingWithPinger(target.GetAddress(), network.PingerCreator)
			// currentStatus := network.CheckPing(target.Address)
			log.Println(currentStatus)
			// Prüfen, ob der Status sich geändert hat
			if lastStatus, exists := lastStatuses[target.GetAddress()]; !exists || currentStatus != lastStatus {
				lastStatuses[target.GetAddress()] = currentStatus
				logMessage := fmt.Sprintf("%s - [%s] (%s) Status geändert: %s", time.Now().Format(time.RFC3339), target.GetName(), target.GetAddress(), currentStatus)
				fmt.Println(logMessage)

				// Änderung in die Datenbank schreiben
				err := db.LogChange(daba, target, currentStatus)
				if err != nil {
					log.Printf("Fehler beim Schreiben in die Datenbank: %v\n", err)
				}
			}
		}

		time.Sleep(checkInterval)
	}
}

func initialize() (*sql.DB, []config.Target) {
	// Ziele aus JSON-Datei einlesen
	targets, err := config.LoadTargets(config.JsonFile)
	if err != nil {
		log.Fatalf("Fehler beim Laden der Ziele: %v\n", err)
	}
	// SQLite-Datenbank initialisieren
	db, err := db.InitializeDatabase(config.DBFileName)
	if err != nil {
		log.Fatalf("Fehler beim Initialisieren der Datenbank: %v\n", err)
	}
	return db, targets
}
