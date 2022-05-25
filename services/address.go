package services

import (
	"net/http"

	"github.com/Bit0r/online-store/model"
	"github.com/gin-gonic/gin"
)

func GetAddresses(ctx *gin.Context) {
	userID := ctx.GetUint64("userID")
	addresses, err := model.GetAddresses(userID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Set("tpl_files", []string{"layout.html", "navbar.html", "addresses.html"})
	ctx.Set("tpl_data", struct {
		model.Addresses
		Used, Free int
	}{Addresses: addresses,
		Used: len(addresses),
		Free: 20 - len(addresses),
	})
}
