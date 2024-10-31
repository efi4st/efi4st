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
func CreateSMSReleaseNote(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.ViewData("deviceId", id)
	ctx.View("sms_createReleaseNote.html")
}


// POST
func AddSMSReleaseNote(ctx iris.Context) {

	device_id := ctx.PostValue("Device_id")
	releaseNoteType := ctx.PostValue("ReleaseNoteType")
	details := ctx.PostValue("Details")

	i, err := strconv.Atoi(device_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing device_id!")
	}

	err = dbprovider.GetDBManager().AddSMSReleaseNote(i, releaseNoteType, details)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Not able to add release Note!")
	}

	ctx.Params().Set("id", device_id)
	ShowSMSDevice(ctx)
}

func ShowSMSReleaseNote(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing releaseNote Id!")
	}

	releaseNote := dbprovider.GetDBManager().GetSMSReleaseNoteInfo(i)

	ctx.ViewData("releaseNote", releaseNote)
	ctx.View("sms_showReleaseNote.html")
}

