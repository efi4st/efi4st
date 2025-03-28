/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (
	"github.com/efi4st/efi4st/classes"
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

	projectInfo := dbprovider.GetDBManager().GetSMSProjectInfo(projectID)
	ctx.ViewData("projectInfo", projectInfo)

	systemTypeMap, notCleanSystem, err := dbprovider.GetDBManager().GetDevicesAndSoftwareForProject(projectID)
	if err != nil {
		ctx.ViewData("error", "Error fetching device/software list!")
		ctx.View("sms_showProjectUpdate.html")
		return
	}

	convertedMap := make(map[string][]classes.DeviceSoftwareInfo)
	for key, value := range systemTypeMap {
		convertedMap[strconv.Itoa(key)] = value // int → string
	}

	ctx.ViewData("systemTypeMap", convertedMap)
	ctx.ViewData("notCleanSystem", notCleanSystem) // Für den "Not a clean System" Hinweis

	ctx.View("sms_showProjectUpdate.html")
}