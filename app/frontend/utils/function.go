package utils

import (
	"context"
)

func GetUserIdFromCtx(ctx context.Context) uint64 {
	userId := ctx.Value(IdentityKey)
	if userId != nil {
		return userId.(uint64)
	}
	return 0
}
