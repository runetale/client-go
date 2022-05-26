package grpc

import (
	"fmt"

	"google.golang.org/grpc/metadata"
)

func getLoginSessionID(md metadata.MD) string {
	registered := md.Get("session_id")
	fmt.Println(registered)
	return registered[0]
}
