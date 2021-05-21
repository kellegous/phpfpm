package config

import "time"

type ProcessManagerType string

const (
	TypeStatic ProcessManagerType = "static"

	TypeOnDemand ProcessManagerType = "ondemand"

	TypeDynamic ProcessManagerType = "dynamic"
)

type ProcessManager struct {
	Type               ProcessManagerType
	MaxChildren        int
	StartServers       int
	MinSpareServers    int
	MaxSpareServers    int
	ProcessIdleTimeout time.Duration
	MaxRequests        int
	StatusPath         string
}
