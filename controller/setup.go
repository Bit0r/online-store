package controller

import (
	"net/http"

	"github.com/Bit0r/online-store/conf"
	"github.com/Bit0r/online-store/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

var (
	router    = gin.Default()
	uploadDir = conf.Get("website", "online_store_website", "upload_dir").(string)
)

func Setup() {
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("Ju8AbyXfnjoMktzh"))

	router.Use(
		middleware.TemplateExecute,
		middleware.Pagination,
		sessions.Sessions("gin-session", store),
	)

	router.Static("/js", "js/")
	router.Static("/cover", uploadDir+"/cover/")

	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/index")
	})

	setupUser()
	setupBooks()
	setupCart()
	setupOrder()
	setupAdmin()
	setupBook()

	router.Run()
}
