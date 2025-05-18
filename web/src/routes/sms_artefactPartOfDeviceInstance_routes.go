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
func CreateSMSArtefactPartOfDeviceInstance(ctx iris.Context) {
	id := ctx.Params().Get("id")
	deviceInstanceID, err := strconv.Atoi(id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing deviceInstance ID!")
		deviceInstanceID = -1
	}

	artefacts := dbprovider.GetDBManager().GetSMSArtefact()
	ctx.ViewData("deviceInstanceId", deviceInstanceID)
	ctx.ViewData("artefactList", artefacts)
	ctx.ViewData("error", "")
	ctx.View("sms_createArtefactPartOfDeviceInstance.html")
}

// POST
func AddSMSArtefactPartOfDeviceInstance(ctx iris.Context) {
	artefactIDs := ctx.PostValue("Artefact_ids")
	deviceInstanceIDStr := ctx.PostValue("DeviceInstance_id")
	additionalInfo := ctx.PostValue("AdditionalInfo")

	deviceInstanceID, err := strconv.Atoi(deviceInstanceIDStr)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing DeviceInstance_id!")
		ShowSMSDeviceInstance(ctx)
		return
	}

	idList := strings.Split(artefactIDs, ",")
	for _, artefactIDStr := range idList {
		artefactIDStr = strings.TrimSpace(artefactIDStr)
		if artefactIDStr == "" {
			continue
		}

		artefactID, err := strconv.Atoi(artefactIDStr)
		if err != nil {
			ctx.ViewData("error", fmt.Sprintf("Error parsing ArtefactId: %s", artefactIDStr))
			continue
		}

		err = dbprovider.GetDBManager().AddSMSArtefactPartOfDeviceInstance(deviceInstanceID, artefactID, additionalInfo)
		if err != nil {
			ctx.ViewData("error", fmt.Sprintf("Error adding artefact %d to device instance %d: %v", artefactID, deviceInstanceID, err))
			// Du kannst hier `return` machen, wenn du beim ersten Fehler abbrechen willst
		}
	}

	ctx.Params().Set("id", deviceInstanceIDStr)
	ShowSMSDeviceInstance(ctx)
}

func RemoveSMSArtefactPartOfDeviceInstance(ctx iris.Context) {
	deviceInstanceID := ctx.Params().Get("deviceInstanceId")
	artefactID := ctx.Params().Get("artefactId")

	dID, err1 := strconv.Atoi(deviceInstanceID)
	aID, err2 := strconv.Atoi(artefactID)

	if err1 != nil || err2 != nil {
		ctx.ViewData("error", "Error: Failed to parse ID parameters!")
		SMSProjects(ctx)
		return
	}

	err := dbprovider.GetDBManager().RemoveSMSArtefactPartOfDeviceInstance(dID, aID)
	if err != nil {
		ctx.ViewData("error", "Error: Could not remove artefact from device instance!")
	}

	ctx.Params().Set("id", deviceInstanceID)
	ShowSMSDeviceInstance(ctx)
}