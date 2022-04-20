package controller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Paging struct {
	Pre, Cur, Next, Total uint64
}

func auth(ctx *gin.Context) {
	session := sessions.Default(ctx)
	_, ok := session.Get("userID").(uint64)
	if !ok {
		ctx.Status(http.StatusUnauthorized)
		ctx.Abort()
	}
}

func authRedirect(ctx *gin.Context) {
	session := sessions.Default(ctx)
	_, ok := session.Get("userID").(uint64)
	if !ok {
		ctx.Redirect(http.StatusFound, "/user/log-in")
		ctx.Abort()
	}
}

func pagination(ctx *gin.Context) {
	paging := Paging{Total: 0}
	current, err := strconv.Atoi(ctx.Query("page"))
	if err != nil || current == 0 {
		paging.Cur = 1
	} else {
		paging.Cur = uint64(current)
	}
	ctx.Set("paging", paging)

	ctx.Next()

	paging = ctx.MustGet("paging").(Paging)
	if paging.Total == 0 {
		// 确认是否需要分页
		return
	}

	// 设置当前页
	if paging.Cur > paging.Total {
		paging.Cur = paging.Total
	}
	paging.Pre, paging.Next = paging.Cur-1, paging.Cur+1
	ctx.Set("tpl_data", map[string]any{
		"Data":   ctx.MustGet("tpl_data"),
		"Paging": paging,
	})
}

func templateExecute(ctx *gin.Context) {
	ctx.Next()

	files, ok := ctx.Get("tpl_files")
	if !ok {
		return
	}

	tpl, _ := template.ParseFiles(getFiles(files.([]string)...)...)

	data, _ := ctx.Get("tpl_data")
	tpl.Execute(ctx.Writer, data)
}
