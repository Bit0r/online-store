package controller

import (
	"net/http"

	"github.com/Bit0r/online-store/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func setupUser() {
	userGroup := router.Group("/user")
	userGroup.Any("/log-in", handleLogin)
	userGroup.Any("/sign-up", handleSignUp)
}

func handleLogin(ctx *gin.Context) {
	switch ctx.Request.Method {
	case http.MethodGet:
		files := []string{"layout-no-nav.html", "log-in.html"}
		ctx.Set("tpl_files", files)
	case http.MethodPost:
		id := model.VerifyUser(ctx.PostForm("name"), ctx.PostForm("passwd"))
		if id != 0 {
			session := sessions.Default(ctx)
			session.Set("userID", id)
			session.Save()
		} else {
			ctx.Status(http.StatusUnauthorized)
		}
	}
}

func handleSignUp(ctx *gin.Context) {
	switch ctx.Request.Method {
	case http.MethodGet:
		files := []string{"layout-no-nav.html", "sign-up.html"}
		ctx.Set("tpl_files", files)
	case http.MethodPost:
		id := model.AddUser(ctx.PostForm("name"), ctx.PostForm("passwd"))
		if id != 0 {
			session := sessions.Default(ctx)
			session.Set("userID", id)
			session.Save()
		} else {
			ctx.Status(http.StatusUnauthorized)
		}
	}
}
