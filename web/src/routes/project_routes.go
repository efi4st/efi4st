/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (
	"database/sql"
	"github.com/kataras/iris"
	"../dbprovider"
	"strings"
	"time"
	"../classes"
)

func Projects(ctx iris.Context) {

	projects := dbprovider.GetDBManager().GetProjects()

	ctx.ViewData("projectList", projects)
	ctx.View("projects.html")
}

func CreateProject(ctx iris.Context) {

	//name := ctx.PostValue("Name")

	//var project = classes.NewProject(, dbName, dbUploads, dbDate)

	//createRoom := &sc.Room{id, publicKey, networkAddress,department ,building ,strings.Split(time.Now().String(), ".")[0], true}

	ctx.Redirect("/projects")
}

func AddProject(ctx iris.Context) {

	projects := dbprovider.GetDBManager().GetProjects()

	ctx.ViewData("projectList", projects)
	ctx.View("projects.html")
}