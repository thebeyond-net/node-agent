package amneziawg

import (
	"context"

	"github.com/thebeyond-net/node-agent/internal/core/ports"
)

type UseCase interface {
	CreatePeer(ctx context.Context) (clientConf, clientPubKey string, err error)
	DeletePeer(ctx context.Context, pubKey string) error
}

type Interactor struct {
	adapter ports.AmneziaWG
}

func NewInteractor(adapter ports.AmneziaWG) UseCase {
	return &Interactor{adapter}
}

func (i *Interactor) CreatePeer(ctx context.Context) (clientConf, clientPubKey string, err error) {
	return i.adapter.CreatePeer(ctx)
}

func (i *Interactor) DeletePeer(ctx context.Context, pubKey string) error {
	return i.adapter.DeletePeer(ctx, pubKey)
}
