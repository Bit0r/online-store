package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

var (
	router  = gin.Default()
	tplRoot = "template/"
)

func Setup() {
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("Ju8AbyXfnjoMktzh"))
	router.Use(sessions.Sessions("gin-session", store), templateExecute, pagination)

	router.Static("/js", "js/")

	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/index")
	})

	setupUser()
	setupBooks()

	router.Run()
}

func getFiles(files ...string) []string {
	for idx, file := range files {
		files[idx] = tplRoot + file
	}
	return files
}
