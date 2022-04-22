package controller

import (
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
	userGroup.GET("/orders", middleware.AuthUserRedirect, handleUserOrders)
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

func handleUserOrders(ctx *gin.Context) {
	session := sessions.Default(ctx)
	id, _ := session.Get("userID").(uint64)

	orders, _ := model.GetOrders(id, 0, 0)

	for idx := range orders {
		var status string
		switch orders[idx].Status {
		case "unpaid":
			status = "未付款"
		case "failed":
			status = "失败"
		case "transit":
			status = "正在运输"
		case "success":
			status = "成功"
		}
		orders[idx].Status = status
	}

	ctx.Set("tpl_files", []string{"layout.html", "navbar.html", "orders.html", "pagination.html"})
	ctx.Set("tpl_data", orders)

	paging, _ := ctx.MustGet("paging").(middleware.Paging)
	paging.Total = 1
	ctx.Set("paging", paging)
}
