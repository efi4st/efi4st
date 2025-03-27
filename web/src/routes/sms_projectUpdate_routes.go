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

func SMSprojectUpdate(ctx iris.Context) {
	id := ctx.Params().Get("id")
	projectID, err := strconv.Atoi(id)

	if err != nil {
		ctx.ViewData("error", "Error: converting project ID!")
		ctx.View("sms_showProjectUpdate.html")
		return
	}

	// Projekt-Infos abrufen
	projectInfo := dbprovider.GetDBManager().GetSMSProjectInfo(projectID)
	ctx.ViewData("projectInfo", projectInfo)

	// Geräte + Software abrufen
	deviceSoftwareList, notCleanSystem, err := dbprovider.GetDBManager().GetDevicesAndSoftwareForProject(projectID)
	if err != nil {
		ctx.ViewData("error", "Error fetching device/software list!")
		ctx.View("sms_showProjectUpdate.html")
		return
	}

	ctx.ViewData("deviceSoftwareList", deviceSoftwareList)
	ctx.ViewData("notCleanSystem", notCleanSystem) // 🆕 Flag für die View setzen

	ctx.View("sms_showProjectUpdate.html")
}