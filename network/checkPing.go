package network

import (
	"log"

	probing "github.com/prometheus-community/pro-bing"
)

type RealPinger struct {
	*probing.Pinger
}

func (r *RealPinger) Statistics() *probing.Statistics {
	return r.Pinger.Statistics()
}

func PingerCreator(targetAddress string) (Pinger, error) {
	pinger, err := probing.NewPinger(targetAddress)
	if err != nil {
		return nil, err
	}
	pinger.SetPrivileged(true) // Aktivieren von ICMP (falls notwendig)
	return &RealPinger{Pinger: pinger}, nil
}

// checkPing sendet einen Ping und gibt den Status ("Erreichbar" oder "Nicht erreichbar") zurück
func CheckPing(targetAddress string) string {
	pinger, err := probing.NewPinger(targetAddress)
	if err != nil {
		log.Println("NewPinger:", err)
		return "Nicht erreichbar"
	}
	pinger.Count = 1
	pinger.SetPrivileged(true)

	err = pinger.Run()
	if err != nil || pinger.Statistics().PacketsRecv == 0 {
		log.Println("Run:", err)
		return "Nicht erreichbar"
	}
	return "Erreichbar"
}

func CheckPingWithPinger(targetAddress string, createPinger func(string) (Pinger, error)) string {
	pinger, err := createPinger(targetAddress)
	if err != nil {
		log.Println("NewPinger:", err)
		return "Nicht erreichbar"
	}
	err = pinger.Run()
	if err != nil || pinger.Statistics().PacketsRecv == 0 {
		log.Println("Run:", err)
		return "Nicht erreichbar"
	}
	return "Erreichbar"
}
