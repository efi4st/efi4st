/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
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

// GET
func CreateSMSProjectBOMForProject(ctx iris.Context) {

	id := ctx.Params().Get("id")
	i, err := strconv.Atoi(id)
	systems := dbprovider.GetDBManager().GetSMSSystems()
	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing project id!")
	}

	ctx.ViewData("projectId", i)
	ctx.ViewData("systemList", systems)
	ctx.View("sms_createProjectBOMForProject.html")
}

// POST
func AddSMSProjectBOM(ctx iris.Context) {

	project_id := ctx.PostValue("Project_id")
	system_id := ctx.PostValue("System_id")
	orderNumber := ctx.PostValue("OrderNumber")
	additionalInfo := ctx.PostValue("AdditionalInfo")

	iP, err := strconv.Atoi(project_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing project_id!")
	}

	iS, err := strconv.Atoi(system_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing system_id!")
	}

	err = dbprovider.GetDBManager().AddSMSProjectBOM(iP, iS, orderNumber, additionalInfo)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Not able to add system to Project!")
	}

	ctx.Params().Set("id", project_id)
	ShowSMSProject(ctx)
}

func RemoveSMSProjectBOM(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveSMSProjectBOM(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing system from Project!")
	}

	SMSProjects(ctx)
}