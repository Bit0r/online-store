package controller

import (
	"github.com/Bit0r/online-store/middleware"
	"github.com/Bit0r/online-store/services"
)

func setupAdmin() {
	adminGroup := router.Group("/admin", middleware.AuthUser)
	adminGroup.GET("/orders", middleware.Permission("order"), getOrders(true))
	adminGroup.GET("/users", middleware.Permission("user"), services.ShowUsers)
}
