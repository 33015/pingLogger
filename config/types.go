package config

const (
	DBFileName   = "connectivity_changes.db"
	JsonFile     = "targets.json"
	PingInterval = 30
)

type TargetInterface interface {
	GetName() string
	GetAddress() string
}

type Target struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (t Target) GetName() string {
	return t.Name
}

func (t Target) GetAddress() string {
	return t.Address
}
