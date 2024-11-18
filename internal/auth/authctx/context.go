package authctx

import (
	"context"
	"github.com/aridae/gophermart-diploma/internal/model"
)

const (
	_userCtxKey = "USER_CONTEXT_KEY"
)

func ContextWithUser(ctx context.Context, user model.User) context.Context {
	return context.WithValue(ctx, _userCtxKey, user)
}

func GetUserFromContext(ctx context.Context) (model.User, bool) {
	user, ok := ctx.Value(_userCtxKey).(model.User)
	if !ok {
		return model.User{}, false
	}

	return user, true
}
