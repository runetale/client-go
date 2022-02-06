package grpcserver

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Notch-Technologies/wizy/client"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
	auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

type MiddlewareManager interface {
	Authenticate(ctx context.Context) (newCtx context.Context, err error)
}

type Middleware struct {
}

func NewMiddlware() *Middleware {
	return &Middleware{
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

	accessToken, err := client.GetAuth0ManagementAccessToken()
	if err != nil {
		return nil, errors.New(domain.ErrInvalidValue.Error())
	}

	fmt.Println(accessToken)

	isAdmin, err := client.IsAdmin(sub, accessToken)
	if err != nil || !isAdmin {
		return nil, errors.New(domain.ErrInvalidValue.Error())
	}

	newCtx = context.WithValue(ctx, "sub", sub)
	return newCtx, nil
}
