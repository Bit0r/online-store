package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Bit0r/online-store/middleware"
	"github.com/Bit0r/online-store/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func setupCart() {
	cartGroup := router.Group("/cart")
	cartGroup.POST("/book", middleware.AuthUser, handleAddCart)
	cartGroup.GET("/books", middleware.AuthUserRedirect, handleShowCart)
	cartGroup.PUT("/book", middleware.AuthUser, handleUpdateCart)
}

func handleShowCart(ctx *gin.Context) {
	id, _ := sessions.Default(ctx).Get("userID").(uint64)

	files := []string{"layout.html", "navbar.html", "shopping-cart.html"}

	data := struct {
		model.CartBooks
		model.CartAddresses
	}{
		model.GetCartItems(id),
		model.GetCartAddresses(id)}

	ctx.Set("tpl_files", files)
	ctx.Set("tpl_data", data)
}

func handleAddCart(ctx *gin.Context) {
	session := sessions.Default(ctx)
	id, _ := session.Get("userID").(uint64)

	bookID, _ := strconv.ParseUint(ctx.PostForm("id"), 0, 64)
	err := model.AddCartItem(id, bookID)
	if err != nil {
		log.Println(err)
		ctx.Status(http.StatusInternalServerError)
	}
}

func handleUpdateCart(ctx *gin.Context) {
	userID := sessions.Default(ctx).Get("userID").(uint64)
	bookInfo := struct {
		BookID   uint64
		Quantity uint
	}{}
	ctx.Bind(&bookInfo)
	err := model.UpdateCartItem(userID, bookInfo.BookID, bookInfo.Quantity)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}
}
