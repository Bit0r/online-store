package services

import (
	"net/http"

	"github.com/Bit0r/online-store/model"
	"github.com/Bit0r/online-store/view"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

func ShowUsers(ctx *gin.Context) {
	isAdmin := ctx.Query("role") == "admin"
	step := uint64(10)

	paging := ctx.MustGet("paging").(view.Paging)
	count := model.CountUsers()
	paging.Total = count/step + 1
	ctx.Set("paging", paging)

	limit := model.Limit{Offset: (paging.Cur - 1) * step, Count: step}
	users, err := model.GetUsers(isAdmin, limit)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Set("tpl_files", []string{"layout.html", "navbar.html", "users.html"})
	ctx.Set("tpl_data", gin.H{
		"Users":   users,
		"HasPriv": lo.Contains[string],
		"IsAdmin": isAdmin,
	})
}
