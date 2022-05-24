package grpc

import (
	"google.golang.org/grpc/metadata"
)

func getLoginSessionID(md metadata.MD) string {
	registered := md.Get("session")
	return registered[0]
}
