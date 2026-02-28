package peers

import (
	"fmt"
	"net/netip"

	"github.com/thebeyond-net/node-agent/internal/core/ports"
)

type Interactor struct {
	vpn     ports.VPNProvider
	ipRepo  ports.IPAllocationRepository
	network netip.Prefix
}

func NewInteractor(
	vpn ports.VPNProvider,
	ipRepo ports.IPAllocationRepository,
	networkCIDR string,
) (ports.PeerUseCase, error) {
	prefix, err := netip.ParsePrefix(networkCIDR)
	if err != nil {
		return nil, fmt.Errorf("invalid network cidr %q: %w", networkCIDR, err)
	}

	return &Interactor{
		vpn:     vpn,
		ipRepo:  ipRepo,
		network: prefix,
	}, nil
}
