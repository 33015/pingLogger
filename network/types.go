package network

import probing "github.com/prometheus-community/pro-bing"

type Pinger interface {
	Run() error
	Statistics() *probing.Statistics
}
