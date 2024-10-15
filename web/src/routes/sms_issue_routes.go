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

func SMSIssues(ctx iris.Context) {

	issues := dbprovider.GetDBManager().GetSMSIssues()
	ctx.ViewData("error", "")

	if len(issues) < 1 {
		ctx.ViewData("error", "Error: No issues available. Add one!")
	}
	ctx.ViewData("issueList", issues)
	ctx.View("sms_issues.html")
}

// GET
func CreateSMSIssue(ctx iris.Context) {

	ctx.View("sms_createIssue.html")
}


// POST
func AddSMSIssue(ctx iris.Context) {

	name := ctx.PostValue("Name")
	issueType := ctx.PostValue("IssueType")
	reference := ctx.PostValue("Reference")
	criticality := ctx.PostValue("Criticality")
	cve := ctx.PostValue("Cve")
	description := ctx.PostValue("Description")

	c, err := strconv.Atoi(criticality)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing criticality!")
	}

	err = dbprovider.GetDBManager().AddSMSIssue(name, issueType, reference, c, cve, description)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Not able to add issue!")
	}
	issues := dbprovider.GetDBManager().GetSMSIssues()
	ctx.ViewData("issueList", issues)
	ctx.View("sms_issues.html")
}

func ShowSMSIssue(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing issue Id!")
	}

	issue := dbprovider.GetDBManager().GetSMSIssueInfo(i)
	affectedDevices := dbprovider.GetDBManager().GetSMSIssueAffectedDevicesForIssueID(i)
	affectedDeviceInstancesAndProjects := dbprovider.GetDBManager().GetSMSAffectedDeviceInstancesAndProjects(i)
	solutionsForThisIssue := dbprovider.GetDBManager().GetSMSSolutionsForIssue(i)

	ctx.ViewData("solutionsForThisIssue", solutionsForThisIssue)
	ctx.ViewData("affectedDeviceInstancesAndProjects", affectedDeviceInstancesAndProjects)
	ctx.ViewData("affectedDevices", affectedDevices)
	ctx.ViewData("issue", issue)
	ctx.View("sms_showIssue.html")
}

func RemoveSMSIssue(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveSMSIssue(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing issue!")
	}

	issues := dbprovider.GetDBManager().GetSMSIssues()

	ctx.ViewData("issueList", issues)
	ctx.View("sms_issues.html")
}