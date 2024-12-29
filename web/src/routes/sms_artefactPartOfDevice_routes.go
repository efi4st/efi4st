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
func CreateSMSArtefactPartOfDevice(ctx iris.Context) {

	id := ctx.Params().Get("id")
	i, err := strconv.Atoi(id)
	artefacts := dbprovider.GetDBManager().GetSMSArtefact()
	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing device id!")
	}

	ctx.ViewData("deviceId", i)
	ctx.ViewData("artefactList", artefacts)
	ctx.View("sms_createArtefactPartOfDevice.html")
}

// POST
func AddSMSArtefactPartOfDevice(ctx iris.Context) {

	artefact_ids, err := ctx.PostValueMany("Artefact_ids")
	device_id := ctx.PostValue("Device_id")
	additionalInfo := ctx.PostValue("AdditionalInfo")

	iI, err := strconv.Atoi(device_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing device_id!")
	}

	for index, artefactId := range strings.Split(artefact_ids,","){
		iD, err := strconv.Atoi(artefactId)
		if err != nil {
			ctx.ViewData("error", "Error: Error parsing artefactId!")
		}
		err = dbprovider.GetDBManager().AddSMSArtefactPartOfDevice(iI, iD, additionalInfo)
		fmt.Printf(strconv.Itoa(index))
		ctx.ViewData("error", "")
		if err !=nil {
			ctx.ViewData("error", "Error: Not able to add artefact device link!")
		}
	}

	ctx.Params().Set("id", device_id)
	ShowSMSDevice(ctx)
}

func RemoveSMSArtefactPartOfDevice(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveSMSArtefactPartOfDevice(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing artefact device link!")
	}

	SMSProjects(ctx)
}