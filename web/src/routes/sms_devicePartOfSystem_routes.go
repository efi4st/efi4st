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
func CreateSMSDevicePartOfSystem(ctx iris.Context) {

	id := ctx.Params().Get("id")
	i, err := strconv.Atoi(id)
	devices := dbprovider.GetDBManager().GetSMSDevice()
	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing system id!")
	}

	ctx.ViewData("systemId", i)
	ctx.ViewData("deviceList", devices)
	ctx.View("sms_createDevicePartOfSystem.html")
}

// POST
func AddSMSDevicePartOfSystem(ctx iris.Context) {

	device_ids, err := ctx.PostValueMany("Device_ids")
	system_id := ctx.PostValue("System_id")
	additionalInfo := ctx.PostValue("AdditionalInfo")

	iI, err := strconv.Atoi(system_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing system_id!")
	}

	for index, deviceId := range strings.Split(device_ids,","){
		iD, err := strconv.Atoi(deviceId)
		if err != nil {
			ctx.ViewData("error", "Error: Error parsing deviceId!")
		}
		err = dbprovider.GetDBManager().AddSMSDevicePartOfSystem(iI, iD, additionalInfo)
		fmt.Printf(strconv.Itoa(index))
		ctx.ViewData("error", "")
		if err !=nil {
			ctx.ViewData("error", "Error: Not able to add device system link!")
		}
	}

	ctx.Params().Set("id", system_id)
	ShowSMSSystem(ctx)
}

func RemoveSMSDevicePartOfSystem(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveSMSDevicePartOfSystem(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing device system link!")
	}

	SMSProjects(ctx)
}