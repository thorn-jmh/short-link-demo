package controller

import (
	"github.com/gin-gonic/gin"
	"go-svc-tpl/api/dto"
	"go-svc-tpl/utils/jwt"
	"go-svc-tpl/utils/stacktrace"
)

var (
	COOKIE_NAME = "access_token"
	USER_ID_KEY = "uid"
)

var BasicAuthMidware = func() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := parseToken(ctx); err != nil {
			dto.ResponseFail(ctx, err)
			return
		}
		ctx.Next()
	}
}

var ParseTokenMidware = func() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_ = parseToken(ctx)
		ctx.Next()
	}
}

func parseToken(ctx *gin.Context) error {
	token, err := ctx.Cookie(COOKIE_NAME)
	if err != nil {
		return stacktrace.PropagateWithCode(err, dto.ErrPrivilege, "Unauthorized")

	}

	claims, err := jwt.Parse(token)
	if err != nil {
		return stacktrace.PropagateWithCode(err, dto.ErrPrivilege, "Unauthorized")
	}

	ctx.Set(USER_ID_KEY, claims.UserID)
	return nil
}
