/**
 * Author:    Admiral Helmut
 * Created:   18.03.2021
 *
 * (C)
 **/

package routes

import (
	"github.com/efi4st/efi4st/dbprovider"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"strconv"
)

func Tools(ctx iris.Context) {

	tools := dbprovider.GetDBManager().GetAnalysisTools()

	ctx.ViewData("error", "")

	if len(tools) < 1 {
		ctx.ViewData("error", "Error: No tools available. Add one!")
	}

	ctx.ViewData("toolList", tools)
	ctx.View("tools.html")
}

// GET
func CreateTool(ctx iris.Context) {

	ctx.View("createTool.html")
}

// POST
func AddTool(ctx iris.Context) {

	name := ctx.PostValue("Name")
	callPattern := ctx.PostValue("CallPattern")

	err := dbprovider.GetDBManager().AddAnalysisTool(name, callPattern)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Not able to add tool!")
	}
	tools := dbprovider.GetDBManager().GetAnalysisTools()
	ctx.ViewData("toolList", tools)
	ctx.View("tools.html")
}

func ShowTool(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing tool Id!")
	}

	tool := dbprovider.GetDBManager().GetAnalysisToolInfo(i)

	ctx.ViewData("tool", tool)
	ctx.View("showTool.html")
}

func RemoveTool(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveAnalysisTool(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing tool!")
	}

	tools := dbprovider.GetDBManager().GetAnalysisTools()

	ctx.ViewData("toolList", tools)
	ctx.View("tools.html")
}
