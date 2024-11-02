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
func CreateSMSSoftwarePartOfDevice(ctx iris.Context) {

	id := ctx.Params().Get("id")
	i, err := strconv.Atoi(id)
	softwares := dbprovider.GetDBManager().GetSMSSoftware()
	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing device id!")
	}

	ctx.ViewData("deviceId", i)
	ctx.ViewData("softwareList", softwares)
	ctx.View("sms_createSoftwarePartOfDevice.html")
}

// POST
func AddSMSSoftwarePartOfDevice(ctx iris.Context) {

	software_ids, err := ctx.PostValueMany("Software_ids")
	device_id := ctx.PostValue("Device_id")
	additionalInfo := ctx.PostValue("AdditionalInfo")

	iI, err := strconv.Atoi(device_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing device_id!")
	}

	for index, softwareId := range strings.Split(software_ids,","){
		iD, err := strconv.Atoi(softwareId)
		if err != nil {
			ctx.ViewData("error", "Error: Error parsing softwareId!")
		}
		err = dbprovider.GetDBManager().AddSMSSoftwarePartOfDevice(iI, iD, additionalInfo)
		fmt.Printf(strconv.Itoa(index))
		ctx.ViewData("error", "")
		if err !=nil {
			ctx.ViewData("error", "Error: Not able to add software device link!")
		}
	}

	ctx.Params().Set("id", device_id)
	ShowSMSDevice(ctx)
}

func RemoveSMSSoftwarePartOfDevice(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveSMSSoftwarePartOfDevice(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing software device link!")
	}

	SMSProjects(ctx)
}