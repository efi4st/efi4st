package routes

import (

	"github.com/kataras/iris"
)

func Index(ctx iris.Context) {
	ctx.View("index.html")
}