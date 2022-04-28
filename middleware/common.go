package middleware

import (
	"html/template"
	"log"
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
		ctx.Set("paging", nil)
		return
	}

	// 设置当前页
	if paging.Cur > paging.Total {
		paging.Cur = paging.Total
	}
	paging.Pre, paging.Next = paging.Cur-1, paging.Cur+1
	ctx.Set("paging", paging)
}

var tplRoot = "template/"

func TemplateExecute(ctx *gin.Context) {
	ctx.Next()

	_, ok := ctx.Get("tpl_files")
	if !ok {
		return
	}

	files := ctx.GetStringSlice("tpl_files")
	files = append(files, "pagination.html")
	for idx, file := range files {
		files[idx] = tplRoot + file
	}
	tpl, _ := template.ParseFiles(files...)

	data := gin.H{}
	tpl_data, ok := ctx.Get("tpl_data")
	if !ok {
		err := tpl.Execute(ctx.Writer, nil)
		if err != nil {
			log.Println(err)
		}
		return
	}
	data["Data"] = tpl_data

	paging, ok := ctx.MustGet("paging").(Paging)
	if ok {
		data["Paging"] = paging
	}

	sudo, ok := ctx.Get("sudo")
	if ok {
		data["Sudo"], _ = sudo.(string)
	}

	err := tpl.Execute(ctx.Writer, data)
	if err != nil {
		log.Println(err)
	}
}
