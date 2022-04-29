package controller

import (
	"log"
	"net/http"

	"github.com/Bit0r/online-store/middleware"
	"github.com/Bit0r/online-store/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func setupUser() {
	userGroup := router.Group("/user")
	userGroup.Any("/log-in", handleLogin)
	userGroup.Any("/sign-up", handleSignUp)
	userGroup.GET("/orders", middleware.AuthUserRedirect, getOrders(false))
	userGroup.GET("/log-out", middleware.AuthUser, handleLogout)
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

func handleLogout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Options(sessions.Options{MaxAge: -1})
	err := session.Save()
	if err == nil {
		ctx.Redirect(http.StatusFound, "/")
	} else {
		log.Println(err)
		ctx.Status(http.StatusInternalServerError)
	}
}
