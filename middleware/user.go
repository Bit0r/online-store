package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthUser(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID, ok := session.Get("userID").(uint64)
	if ok {
		ctx.Set("userID", userID)
	} else {
		ctx.Status(http.StatusUnauthorized)
		ctx.Abort()
	}
}

func AuthUserRedirect(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID, ok := session.Get("userID").(uint64)
	if ok {
		ctx.Set("userID", userID)
	} else {
		ctx.Redirect(http.StatusFound, "/user/log-in")
		ctx.Abort()
	}
}
