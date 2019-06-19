/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (
	"../dbprovider"
	"github.com/kataras/iris"
	"strconv"
	_ "github.com/go-sql-driver/mysql"

)

func Projects(ctx iris.Context) {

	projects := dbprovider.GetDBManager().GetProjects()

	ctx.ViewData("error", "")

	if len(projects) < 1 {
		ctx.ViewData("error", "Error: No projects available. Add one!")
	}

	ctx.ViewData("projectList", projects)
	ctx.View("projects.html")
}

// GET
func CreateProject(ctx iris.Context) {

	ctx.View("createProject.html")
}

// POST
func AddProject(ctx iris.Context) {

	name := ctx.PostValue("Name")

	err := dbprovider.GetDBManager().AddProject(name)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Not able to add project!")
	}

	projects := dbprovider.GetDBManager().GetProjects()
	ctx.ViewData("projectList", projects)
	ctx.View("projects.html")
}

func ShowProject(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing project Id!")
	}

	project := dbprovider.GetDBManager().GetProjectInfo(i)
	firmwareList := dbprovider.GetDBManager().GetFirmwareListForProject(i)

	ctx.ViewData("firmwareList", firmwareList)
	ctx.ViewData("project", project)
	ctx.View("showProject.html")
}

func RemoveProject(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveProject(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing project!")
	}

	projects := dbprovider.GetDBManager().GetProjects()

	ctx.ViewData("projectList", projects)
	ctx.View("projects.html")
}