package middleware

import (
	"html/template"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Paging struct {
	Pre, Cur, Next, Total uint64
}

func Pagination(ctx *gin.Context) {
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

var tplRoot = "template/"

func TemplateExecute(ctx *gin.Context) {
	ctx.Next()

	_, ok := ctx.Get("tpl_files")
	if !ok {
		return
	}

	files := ctx.GetStringSlice("tpl_files")
	for idx, file := range files {
		files[idx] = tplRoot + file
	}
	tpl, _ := template.ParseFiles(files...)

	data, _ := ctx.Get("tpl_data")
	tpl.Execute(ctx.Writer, data)
}
