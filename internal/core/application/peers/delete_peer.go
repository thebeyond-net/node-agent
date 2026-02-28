package peers

import (
	"context"
	"fmt"
)

func (uc *Interactor) DeletePeer(ctx context.Context, pubKey string) error {
	if err := uc.vpn.RemovePeer(pubKey); err != nil {
		return fmt.Errorf("remove peer: %w", err)
	}
	_ = uc.ipRepo.ReleaseByPublicKey(ctx, pubKey)
	return nil
}
