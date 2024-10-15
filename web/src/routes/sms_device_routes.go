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
	"time"
)

func SMSDevice(ctx iris.Context) {

	devices := dbprovider.GetDBManager().GetSMSDevice()
	ctx.ViewData("error", "")

	if len(devices) < 1 {
		ctx.ViewData("error", "Error: No devices available. Add one!")
	}
	ctx.ViewData("deviceList", devices)
	ctx.View("sms_devices.html")
}

// GET
func CreateSMSDevice(ctx iris.Context) {

	deviceTypes := dbprovider.GetDBManager().GetSMSDeviceTypes()

	ctx.ViewData("typeList", deviceTypes)
	ctx.View("sms_createDevice.html")
}

// POST
func AddSMSDevice(ctx iris.Context) {

	devicetypeId := ctx.PostValue("DevicetypeId")
	version := ctx.PostValue("Version")

	i, err := strconv.Atoi(devicetypeId)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing devicetypeId!")
	}

	err = dbprovider.GetDBManager().AddSMSDevice(i, version, time.Now().String())

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Not able to add device!")
	}
	devices := dbprovider.GetDBManager().GetSMSDevice()
	ctx.ViewData("deviceList", devices)
	ctx.View("sms_devices.html")
}

func ShowSMSDevice(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing device Id!")
	}

	device := dbprovider.GetDBManager().GetSMSDeviceInfo(i)

	ctx.ViewData("device", device)
	ctx.View("sms_showDevice.html")
}

func RemoveSMSDevice(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveSMSDevice(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing device!")
	}

	devices := dbprovider.GetDBManager().GetSMSDevice()

	ctx.ViewData("deviceList", devices)
	ctx.View("sms_devices.html")
}