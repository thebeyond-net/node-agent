package ports

import (
	"context"
	"net/netip"

	"github.com/thebeyond-net/node-agent/internal/core/domain"
)

type VPNProvider interface {
	GetServerConfig() (domain.ServerConfig, error)
	SetPeerBandwidth(ip string, bandwidth int) error
	RegisterPeer(publicKey, allowedIP string) error
	RemovePeer(publicKey, allowedIP string) error
	BuildClientConfig(privKey, address string, srv domain.ServerConfig) (string, error)
}

type IPAllocationRepository interface {
	Allocate(ctx context.Context, prefix netip.Prefix) (netip.Addr, error)
	Reserve(ctx context.Context, ip netip.Addr, publicKey string, prefix netip.Prefix) error
	Release(ctx context.Context, ip netip.Addr) error
	ReleaseByPublicKey(ctx context.Context, publicKey string) (netip.Addr, error)
}

type PeerUseCase interface {
	CreatePeer(ctx context.Context, bandwidth int) (clientConf string, publicKey string, err error)
	DeletePeer(ctx context.Context, publicKey string) error
}
