package services

import (
	"log"
	"net/http"
	"strconv"

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

func EditAddress(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	ctx.Set("tpl_files", []string{"layout.html", "navbar.html", "address-editor.html"})

	var address model.Address
	if id == 0 {
		// 如果是新增地址，则不需要查询数据库
		return
	}

	address, err = model.GetAddress(uint64(id))
	if err != nil || address.UserID != ctx.GetUint64("userID") {
		// 如果地址不存在或者不属于当前用户，则返回错误
		ctx.Set("tpl_files", nil)
		ctx.Status(http.StatusBadRequest)
		return
	}
	ctx.Set("tpl_data", address)
}

func ReplaceAddress(ctx *gin.Context) {
	address := model.Address{UserID: ctx.GetUint64("userID")}
	err := ctx.ShouldBind(&address)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	if address.ID == 0 {
		if model.CountAddresses(address.UserID) >= 20 {
			// 如果用户的地址数量已经达到20个，则不允许新增
			ctx.Status(http.StatusBadRequest)
			return
		}

		_, err := model.AddAddress(address)
		if err != nil {
			// 如果新增地址失败，则返回错误
			log.Println(err)
			ctx.Status(http.StatusInternalServerError)
		}
		ctx.Redirect(http.StatusSeeOther, "/user/address")
		return
	}

	// 如果是更新地址，则需要查询数据库
	addressOld, err := model.GetAddress(address.ID)
	if err != nil || addressOld.UserID != address.UserID {
		// 如果地址不存在或者不属于当前用户，则返回错误
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = model.UpdateAddress(address)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Redirect(http.StatusSeeOther, "/user/address")
}
