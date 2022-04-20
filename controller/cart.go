package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Bit0r/online-store/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func setupCart() {
	cartGroup := router.Group("/cart", auth)
	cartGroup.GET("/add/:id", handleAddCart)
}

func handleAddCart(ctx *gin.Context) {
	session := sessions.Default(ctx)
	id, _ := session.Get("userID").(uint64)

	bookID, _ := strconv.ParseUint(ctx.Param("id"), 0, 64)
	err := model.AddCartItem(id, bookID)
	if err != nil {
		log.Fatal(err)
		ctx.Status(http.StatusInternalServerError)
	}
}
