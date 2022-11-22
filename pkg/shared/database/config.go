package database

import "time"

type ConfigDatabase struct {
	Username           string        `json:"username"`
	Password           string        `json:"password"`
	Name               string        `json:"name"`
	Schema             string        `json:"schema"`
	Host               string        `json:"host"`
	Port               int           `json:"port"`
	MinIdleConnections int           `json:"minIdleConnections"`
	MaxOpenConnections int           `json:"maxOpenConnections"`
	ConnMaxLifetime    time.Duration `json:"connMaxLifetime"`
	DebugMode          bool          `json:"debugMode"`
}
