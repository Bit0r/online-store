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

func setupOrder() {
	router.GET("/order/:id", middleware.AuthUserRedirect)
	router.POST("/order", handleAddOrder)
	router.PUT("/order")
}

func handleAddOrder(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userID, _ := session.Get("userID").(uint64)

	addressID, _ := strconv.Atoi(ctx.PostForm("addressID"))

	booksID := []uint64{}
	for _, v := range ctx.PostFormArray("booksID") {
		bookID, _ := strconv.ParseUint(v, 0, 64)
		booksID = append(booksID, bookID)
	}

	orderID, err := model.AddOrder(userID, uint(addressID), booksID)
	if err != nil {
		log.Println(err)
		ctx.Status(http.StatusInternalServerError)
	} else {
		ctx.Redirect(http.StatusSeeOther, "/order/"+strconv.Itoa(int(orderID)))
	}
}
