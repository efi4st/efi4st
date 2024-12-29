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
	"strings"
)

// GET
func CreateSMSIssueAffectedSoftware(ctx iris.Context) {

	id := ctx.Params().Get("id")
	i, err := strconv.Atoi(id)
	softwares := dbprovider.GetDBManager().GetSMSSoftware()
	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing issue id!")
	}

	ctx.ViewData("issueId", i)
	ctx.ViewData("softwareList", softwares)
	ctx.View("sms_createIssueAffectedSoftware.html")
}

// POST
func AddSMSIssueAffectedSoftware(ctx iris.Context) {
	software_ids, err := ctx.PostValueMany("Software_ids")
	issue_id := ctx.PostValue("Issue_id")
	additionalInfo := ctx.PostValue("AdditionalInfo")

	iI, err := strconv.Atoi(issue_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing issue_id!")
		ctx.Redirect(fmt.Sprintf("/sms_issueAffectedSoftware/createSMSIssueAffectedSoftware/%d", iI))
		return
	}

	for _, softwareId := range strings.Split(software_ids, ",") {
		iD, err := strconv.Atoi(softwareId)
		if err != nil {
			ctx.ViewData("error", "Error: Error parsing softwareId!")
			ctx.Redirect(fmt.Sprintf("/sms_issueAffectedSoftware/createSMSIssueAffectedSoftware/%d", iI))
			return
		}
		err = dbprovider.GetDBManager().AddSMSIssueAffectedSoftware(iD, iI, additionalInfo, true)
		if err != nil {
			ctx.ViewData("error", "Error: Not able to add issue software link!")
			ctx.Redirect(fmt.Sprintf("/sms_issueAffectedSoftware/createSMSIssueAffectedSoftware/%d", iI))
			return
		}
	}

	// Nach erfolgreicher Operation umleiten (GET-Request)
	ctx.Redirect(fmt.Sprintf("/sms_issues/show/%d", iI))
}

func RemoveSMSIssueAffectedSoftware(ctx iris.Context) {
	software_id := ctx.Params().Get("software_id")
	issue_id := ctx.Params().Get("issue_id")

	sId, err := strconv.Atoi(software_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing software_id!")
		ctx.Redirect(fmt.Sprintf("/sms_issues/show/%s", issue_id))
		return
	}

	iId, err := strconv.Atoi(issue_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing issue_id!")
		ctx.Redirect(fmt.Sprintf("/sms_issues/show/%s", issue_id))
		return
	}

	err = dbprovider.GetDBManager().RemoveSMSIssueAffectedSoftware(sId, iId)
	if err != nil {
		ctx.ViewData("error", "Error: Error removing issue software link!")
	}

	// Nach erfolgreicher Löschung umleiten
	ctx.Redirect(fmt.Sprintf("/sms_issues/show/%d", iId))
}