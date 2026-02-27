package amneziawg

import (
	"context"

	"github.com/thebeyond-net/node-agent/internal/core/application/amneziawg"
	v1 "github.com/thebeyond-net/node-agent/pkg/amneziawg/v1"
	"github.com/thebeyond-net/node-agent/pkg/amneziawg/v1/amneziawgv1connect"
)

type AmneziaWGServiceServer struct {
	amneziawgv1connect.UnimplementedAmneziaWGServiceHandler
	usecase amneziawg.UseCase
}

func NewAmneziaWGServiceServer(usecase amneziawg.UseCase) *AmneziaWGServiceServer {
	return &AmneziaWGServiceServer{usecase: usecase}
}

func (s *AmneziaWGServiceServer) CreatePeer(ctx context.Context, req *v1.CreatePeerRequest) (*v1.CreatePeerResponse, error) {
	clientConf, pubKey, err := s.usecase.CreatePeer(ctx)
	return &v1.CreatePeerResponse{
		Success: true,
		Config:  clientConf,
		Pubkey:  pubKey,
	}, err
}

func (s *AmneziaWGServiceServer) DeletePeer(ctx context.Context, req *v1.DeletePeerRequest) (*v1.DeletePeerResponse, error) {
	err := s.usecase.DeletePeer(ctx, req.GetPubkey())
	if err != nil {
		return &v1.DeletePeerResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &v1.DeletePeerResponse{
		Success: true,
	}, nil
}
