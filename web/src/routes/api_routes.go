/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (

	"github.com/kataras/iris"
	"../dbprovider"
)

func Index(ctx iris.Context) {
	projects := dbprovider.GetDBManager().GetProjects()

	ctx.ViewData("projectList", projects)
	ctx.View("index.html")
}

func Documentation(ctx iris.Context) {
	ctx.View("documentation.html")
}