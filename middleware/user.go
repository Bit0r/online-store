package middleware

import (
	"net/http"

	"github.com/Bit0r/online-store/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthUser(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID, ok := session.Get("userID").(uint64)
	if ok {
		ctx.Set("userID", userID)
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
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

func Permission(privilege string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetUint64("userID")
		if !model.HasPrivilege(userID, privilege) {
			ctx.AbortWithStatus(http.StatusForbidden)
		}
	}
}
