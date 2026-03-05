package interceptors

import (
	"context"
	"errors"

	"connectrpc.com/connect"
)

func NewAuthInterceptor(secret string) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if req.Spec().IsClient {
				return next(ctx, req)
			}

			if req.Header().Get("Authorization") != secret {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("invalid or missing auth secret"),
				)
			}

			return next(ctx, req)
		}
	}
}
