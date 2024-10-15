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
func CreateSMSIssueAffectedDevice(ctx iris.Context) {

	id := ctx.Params().Get("id")
	i, err := strconv.Atoi(id)
	devices := dbprovider.GetDBManager().GetSMSDevice()
	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing device id!")
	}

	ctx.ViewData("issueId", i)
	ctx.ViewData("deviceList", devices)
	ctx.View("sms_createIssueAffectedDevice.html")
}

// POST
func AddSMSIssueAffectedDevice(ctx iris.Context) {

	device_ids, err := ctx.PostValueMany("Device_ids")
	issue_id := ctx.PostValue("Issue_id")
	additionalInfo := ctx.PostValue("AdditionalInfo")
	//confirmed := ctx.PostValue("Confirmed")

	iI, err := strconv.Atoi(issue_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing issue_id!")
	}

	for index, deviceId := range strings.Split(device_ids,","){
		iD, err := strconv.Atoi(deviceId)
		if err != nil {
			ctx.ViewData("error", "Error: Error parsing device_id!")
		}
		err = dbprovider.GetDBManager().AddSMSIssueAffectedDevice(iD, iI, additionalInfo, true)
		fmt.Printf(strconv.Itoa(index))
		ctx.ViewData("error", "")
		if err !=nil {
			ctx.ViewData("error", "Error: Not able to add issue device link!")
		}
	}

	ctx.Params().Set("id", issue_id)
	ShowSMSIssue(ctx)
}

func RemoveSMSIssueAffectedDevice(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveSMSIssueAffectedDevice(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing issue device link!")
	}

	SMSProjects(ctx)
}