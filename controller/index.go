package controller

import (
	"github.com/Bit0r/online-store/middleware"
	"github.com/Bit0r/online-store/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func setupBooks() {
	router.GET("/index/*category", func(ctx *gin.Context) {
		var step uint64 = 12
		data := struct {
			Category   string
			Categories []string
			model.Books
		}{}

		// 填充类别信息
		data.Category = ctx.Param("category")[1:]
		data.Categories = model.GetCategories()

		// 填充分页信息
		paging := ctx.MustGet("paging").(middleware.Paging)
		paging.Total = model.CountBooks(data.Category)/step + 1
		ctx.Set("paging", paging)

		// 填充图书信息
		data.Books = model.GetBooks(data.Category, uint64((paging.Cur-1)*step), uint64(step))
		ctx.Set("tpl_data", data)

		// 设置模板
		files := []string{"layout.html", "home.html", "navbar-guest.html"}
		if sessions.Default(ctx).Get("userID") != nil {
			files[2] = "navbar.html"
		}
		ctx.Set("tpl_files", files)
	})
}
