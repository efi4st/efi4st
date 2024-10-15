/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (
	"fmt"
	"github.com/efi4st/efi4st/dbprovider"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"strconv"
)

func SMSProjects(ctx iris.Context) {

	projects := dbprovider.GetDBManager().GetSMSProjects()
	fmt.Printf("->"+string(len(projects)))
	ctx.ViewData("error", "")

	if len(projects) < 1 {
		ctx.ViewData("error", "Error: No projects available. Add one!")
	}
	fmt.Printf(strconv.Itoa(len(projects)))
	fmt.Printf(projects[0].Name())
	ctx.ViewData("projectList", projects)
	ctx.View("sms_projects.html")
}

// GET
func CreateSMSProject(ctx iris.Context) {

	projectTypes := dbprovider.GetDBManager().GetSMSProjectTypes()

	ctx.ViewData("typeList", projectTypes)
	ctx.View("sms_createProject.html")
}


// POST
func AddSMSProject(ctx iris.Context) {

	projectName := ctx.PostValue("ProjectName")
	customer := ctx.PostValue("Customer")
	projecttypeId := ctx.PostValue("ProjecttypeId")
	reference := ctx.PostValue("Reference")

	i, err := strconv.Atoi(projecttypeId)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing projecttypeId!")
	}

	err = dbprovider.GetDBManager().AddSMSProject(projectName, customer, i, reference)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Not able to add project!")
	}
	projects := dbprovider.GetDBManager().GetSMSProjects()
	ctx.ViewData("projectList", projects)
	ctx.View("sms_projects.html")
}

func ShowSMSProject(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing project Id!")
	}

	project := dbprovider.GetDBManager().GetSMSProjectInfo(i)
	deviceInstanceList := dbprovider.GetDBManager().GetDeviceInstanceListForProject(i)

	ctx.ViewData("deviceInstanceList", deviceInstanceList)
	ctx.ViewData("project", project)
	ctx.View("sms_showProject.html")
}

func RemoveSMSProject(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveSMSProject(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing project!")
	}

	projects := dbprovider.GetDBManager().GetSMSProjects()

	ctx.ViewData("projectList", projects)
	ctx.View("sms_projects.html")
}