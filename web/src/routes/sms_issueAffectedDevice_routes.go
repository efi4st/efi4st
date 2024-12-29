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
	// Hole die Device IDs aus den POST-Daten
	device_ids, err := ctx.PostValueMany("Device_ids")
	issue_id := ctx.PostValue("Issue_id")
	additionalInfo := ctx.PostValue("AdditionalInfo")

	// Konvertiere die Issue ID in eine Zahl
	iI, err := strconv.Atoi(issue_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing issue_id!")
		ctx.Redirect(fmt.Sprintf("/sms_issues/show/%d", iI))
		return
	}

	// Gehe jede Device ID durch und füge sie der Issue hinzu
	for _, deviceId := range strings.Split(device_ids, ",") {
		iD, err := strconv.Atoi(deviceId)
		if err != nil {
			ctx.ViewData("error", "Error: Error parsing device_id!")
			continue
		}
		err = dbprovider.GetDBManager().AddSMSIssueAffectedDevice(iD, iI, additionalInfo, true)
		if err != nil {
			ctx.ViewData("error", "Error: Not able to add issue-device link!")
		}
	}

	// Immer zurück zur Issue-Seite leiten
	ctx.Redirect(fmt.Sprintf("/sms_issues/show/%d", iI))
}

func RemoveSMSIssueAffectedDevice(ctx iris.Context) {
	// Hole sowohl die Issue ID als auch die Device ID aus den Parametern
	issueId := ctx.Params().Get("issueId")
	deviceId := ctx.Params().Get("deviceId")

	// Konvertiere die IDs in Integer
	issue, err := strconv.Atoi(issueId)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing issue id!")
		ctx.Redirect(fmt.Sprintf("/sms_issues/show/%d", issue))
		return
	}

	device, err := strconv.Atoi(deviceId)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing device id!")
		ctx.Redirect(fmt.Sprintf("/sms_issues/show/%d", issue))
		return
	}

	// Entferne die Issue-Device-Verknüpfung
	err = dbprovider.GetDBManager().RemoveSMSIssueAffectedDevice(device, issue)

	// Fehlerbehandlung bei der Entfernung
	if err != nil {
		ctx.ViewData("error", "Error: Error removing issue-device link!")
	}

	// Immer zurück zur Issue-Seite leiten
	ctx.Redirect(fmt.Sprintf("/sms_issues/show/%d", issue))
}