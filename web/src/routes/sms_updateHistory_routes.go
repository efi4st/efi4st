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
func CreateSMSUpdateHistory(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.ViewData("deviceInstanceId", id)
	ctx.View("sms_createUpdateLog.html")
}


// POST
func AddSMSUpdateHistory(ctx iris.Context) {

	deviceInstance_id := ctx.PostValue("DeviceInstance_id")
	user := ctx.PostValue("User")
	updateType := ctx.PostValue("UpdateType")
	description := ctx.PostValue("Description")

	i, err := strconv.Atoi(deviceInstance_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing deviceInstance_id!")
	}

	err = dbprovider.GetDBManager().AddSMSUpdateHistory(i, user, updateType, description)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Not able to add Log entry!")
	}

	ctx.Params().Set("id", deviceInstance_id)
	ShowSMSDeviceInstance(ctx)
}

func ShowSMSUpdateHistory(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing project Id!")
	}

	updateLog := dbprovider.GetDBManager().GetSMSUdateHistoryInfo(i)

	ctx.ViewData("updateLog", updateLog)
	ctx.View("sms_showUpdateHistory.html")
}

