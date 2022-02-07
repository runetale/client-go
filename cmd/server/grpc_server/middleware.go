package grpcserver

import (
	"context"
	"errors"
	"strings"

	"github.com/Notch-Technologies/wizy/client"
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

	if !strings.HasPrefix(sub, "auth0") {
		return nil, errors.New(domain.ErrInvalidValue.Error())
	}

	accessToken, err := m.client.GetAuth0ManagementAccessToken()
	if err != nil {
		return nil, errors.New(domain.ErrInvalidValue.Error())
	}

	isAdmin, err := m.client.IsAdmin(sub, accessToken)
	if err != nil || !isAdmin {
		return nil, errors.New(domain.ErrInvalidValue.Error())
	}

	newCtx = context.WithValue(ctx, "sub", sub)
	return newCtx, nil
}
