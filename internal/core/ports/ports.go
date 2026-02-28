package ports

import (
	"context"
	"net/netip"

	"github.com/thebeyond-net/node-agent/internal/core/domain"
)

type VPNProvider interface {
	GetServerConfig() (domain.ServerConfig, error)
	RegisterPeer(pubKey, allowedIP string) error
	RemovePeer(pubKey string) error
	BuildClientConfig(privKey, address string, srv domain.ServerConfig) (string, error)
}

type IPAllocationRepository interface {
	Allocate(ctx context.Context, prefix netip.Prefix) (netip.Addr, error)
	Reserve(ctx context.Context, ip netip.Addr, publicKey string, prefix netip.Prefix) error
	Release(ctx context.Context, ip netip.Addr) error
	ReleaseByPublicKey(ctx context.Context, publicKey string) error
}

type PeerUseCase interface {
	CreatePeer(ctx context.Context) (clientConf string, pubKey string, err error)
	DeletePeer(ctx context.Context, pubKey string) error
}
