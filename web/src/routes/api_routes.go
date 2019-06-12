/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (

	"github.com/kataras/iris"
)

func Index(ctx iris.Context) {
	ctx.View("index.html")
}