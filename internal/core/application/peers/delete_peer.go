package peers

import (
	"context"
	"fmt"
)

func (uc *Interactor) DeletePeer(ctx context.Context, pubKey string) error {
	ip, err := uc.ipRepo.ReleaseByPublicKey(ctx, pubKey)
	if err != nil {
		return fmt.Errorf("release ip: %w", err)
	}

	if err := uc.vpn.RemovePeer(pubKey, ip.String()); err != nil {
		return fmt.Errorf("remove peer: %w", err)
	}

	return nil
}
