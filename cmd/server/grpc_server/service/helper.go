package service

import "context"

func getSub(ctx context.Context) string {
	userId := ctx.Value("sub").(string)
	return userId
}
