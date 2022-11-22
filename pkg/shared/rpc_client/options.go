package rpc_client

import "time"

type Options struct {
	Address      string        `json:"address"`
	Timeout      time.Duration `json:"timeout"`
	DebugMode    bool          `json:"debugMode"`
	WithProxy    bool          `json:"withProxy"`
	ProxyAddress string        `json:"proxyAddress"`
	SkipTLS      bool          `json:"skipTLS"`
}
