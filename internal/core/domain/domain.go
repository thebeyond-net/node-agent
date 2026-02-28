package domain

import (
	"net/netip"
	"time"
)

type ServerConfig struct {
	PublicKey  string `json:"publicKey"`
	ListenPort string `json:"listenPort"`
}

type Peer struct {
	PublicKey string
	AllowedIP netip.Addr
	CreatedAt time.Time
}
