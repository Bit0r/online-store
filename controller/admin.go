package controller

import "github.com/Bit0r/online-store/middleware"

func setupAdmin() {
	adminGroup := router.Group("/admin")
	adminGroup.GET("/orders", middleware.AuthUserRedirect, getOrders(true))
}
