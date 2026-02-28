package peers

import (
	"context"
	"fmt"
	"net/netip"

	"github.com/thebeyond-net/node-agent/pkg/keypair"
)

func (uc *Interactor) CreatePeer(ctx context.Context) (string, string, error) {
	serverConfig, err := uc.vpn.GetServerConfig()
	if err != nil {
		return "", "", fmt.Errorf("get server config: %w", err)
	}

	ip, releaseIP, err := uc.allocateIP(ctx)
	if err != nil {
		return "", "", err
	}

	defer func() {
		if err != nil {
			releaseIP()
		}
	}()

	keypair, err := keypair.Generate()
	if err != nil {
		return "", "", fmt.Errorf("generate keys: %w", err)
	}

	if err = uc.vpn.RegisterPeer(keypair.Public, ip.String()); err != nil {
		return "", "", fmt.Errorf("register peer: %w", err)
	}

	clientConfig, err := uc.vpn.BuildClientConfig(keypair.Private, ip.String(), serverConfig)
	if err != nil {
		_ = uc.vpn.RemovePeer(keypair.Public)
		return "", "", fmt.Errorf("build client config: %w", err)
	}

	return clientConfig, keypair.Public, nil
}

func (uc *Interactor) allocateIP(ctx context.Context) (netip.Addr, func(), error) {
	ip, err := uc.ipRepo.Allocate(ctx, uc.network)
	if err != nil {
		return netip.Addr{}, nil, fmt.Errorf("allocate ip: %w", err)
	}

	release := func() {
		_ = uc.ipRepo.Release(ctx, ip)
	}

	if err := uc.ipRepo.Reserve(ctx, ip, "", uc.network); err != nil {
		release()
		return netip.Addr{}, nil, fmt.Errorf("reserve ip: %w", err)
	}

	return ip, release, nil
}
