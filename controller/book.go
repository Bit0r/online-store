package controller

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Bit0r/online-store/middleware"
	"github.com/Bit0r/online-store/model"
	"github.com/gin-gonic/gin"
)

func setupBook() {
	bookGroup := router.Group("/book")
	bookGroup.GET("/edit/:id",
		middleware.AuthUserRedirect,
		middleware.Permission("book"),
		bookEdit)
}

func bookEdit(ctx *gin.Context) {
	bookID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	book, err := model.GetBook(bookID)
	if err != nil {
		log.Println(err)
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.Set("tpl_files", []string{"layout.html", "navbar.html", "book-editor.html"})
	ctx.Set("tpl_data", struct {
		model.Book
		Categories string
	}{book, strings.Join(model.GetBookCategories(bookID), ",")})
}
