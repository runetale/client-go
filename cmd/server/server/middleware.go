package grpcserver

import (
	"context"
	"errors"

	client "github.com/Notch-Technologies/wizy/auth0"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
	auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

type MiddlewareManager interface {
	Authenticate(ctx context.Context) (newCtx context.Context, err error)
}

type Middleware struct {
	client *client.Auth0Client
}

func NewMiddlware(client *client.Auth0Client) *Middleware {
	return &Middleware{
		client: client,
	}
}

func (m *Middleware) Authenticate(ctx context.Context) (newCtx context.Context, err error) {
	sub, err := auth.AuthFromMD(ctx, "sub")
	if err != nil {
		return nil, err
	}

	accessToken, err := m.client.GetAuth0ManagementAccessToken()
	if err != nil {
		return nil, errors.New(domain.ErrCanNotGetAccessToken.Error())
	}

	isAdmin, err := m.client.IsAdmin(sub, accessToken)
	if err != nil || !isAdmin {
		return nil, errors.New(domain.ErrNotEnoughPermission.Error())
	}

	newCtx = context.WithValue(ctx, "sub", sub)
	return newCtx, nil
}
