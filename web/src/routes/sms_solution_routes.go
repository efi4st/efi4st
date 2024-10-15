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
func CreateSMSSolution(ctx iris.Context) {

	id := ctx.Params().Get("id")
	i, err := strconv.Atoi(id)
	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing issue id!")
	}
	deviceTypes := dbprovider.GetDBManager().GetSMSDeviceTypes()

	ctx.ViewData("issueId", i)
	ctx.ViewData("typeList", deviceTypes)
	ctx.View("sms_createSolution.html")
}


// POST
func AddSMSSolution(ctx iris.Context) {

	issue_id := ctx.PostValue("Issue_id")
	devicetype_id := ctx.PostValue("Devicetype_id")
	name := ctx.PostValue("Name")
	description := ctx.PostValue("Description")
	reference := ctx.PostValue("Reference")

	i, err := strconv.Atoi(issue_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing issue_id!")
	}
	d, err := strconv.Atoi(devicetype_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing devicetype_id!")
	}

	err = dbprovider.GetDBManager().AddSMSSolution(i, d, name, description, reference)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Not able to add solution!")
	}
	issues := dbprovider.GetDBManager().GetSMSIssues()
	ctx.ViewData("issueList", issues)
	ctx.View("sms_issues.html")
}

func ShowSMSSolution(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing solution Id!")
	}

	solution := dbprovider.GetDBManager().GetSMSSolutionInfo(i)

	ctx.ViewData("solution", solution)
	ctx.View("sms_showSolution.html")
}

func RemoveSMSSolution(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveSMSSolution(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing solution!")
	}

	issues := dbprovider.GetDBManager().GetSMSIssues()

	ctx.ViewData("issueList", issues)
	ctx.View("sms_issues.html")
}