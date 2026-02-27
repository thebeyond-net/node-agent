package ports

import "context"

type AmneziaWG interface {
	CreatePeer(ctx context.Context) (clientConf, clientPubKey string, err error)
	DeletePeer(ctx context.Context, pubKey string) error
}
