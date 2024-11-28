package network

import (
	"fmt"
	"testing"

	probing "github.com/prometheus-community/pro-bing"
)

type MockPinger struct {
	runError        error
	packetsReceived int
}

func (m *MockPinger) Run() error {
	return m.runError
}

func (m *MockPinger) Statistics() *probing.Statistics {
	return &probing.Statistics{PacketsRecv: m.packetsReceived}
}

func TestCheckPing(t *testing.T) {
	mockPinger := func(target string) (Pinger, error) {
		if target == "unreachable.com" {
			return &MockPinger{runError: nil, packetsReceived: 0}, nil
		}
		if target == "error.com" {
			return nil, fmt.Errorf("failed to create pinger")
		}
		return &MockPinger{runError: nil, packetsReceived: 1}, nil
	}

	tests := []struct {
		target   string
		expected string
	}{
		{"reachable.com", "Erreichbar"},
		{"unreachable.com", "Nicht erreichbar"},
		{"error.com", "Nicht erreichbar"},
	}

	for _, test := range tests {
		result := CheckPingWithPinger(test.target, mockPinger)
		if result != test.expected {
			t.Errorf("checkPing(%q) = %q; want %q", test.target, result, test.expected)
		}
	}
}
