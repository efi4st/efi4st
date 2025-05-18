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
	"strings"
)

// GET
func CreateSMSArtefactPartOfSystem(ctx iris.Context) {
	id := ctx.Params().Get("id")
	systemID, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing system id!")
	}

	artefacts := dbprovider.GetDBManager().GetSMSArtefact()

	ctx.ViewData("systemId", systemID)
	ctx.ViewData("artefactList", artefacts)
	ctx.View("sms_createArtefactPartOfSystem.html")
}


// POST
func AddSMSArtefactPartOfSystem(ctx iris.Context) {
	artefact_ids := ctx.PostValue("Artefact_ids")
	system_id := ctx.PostValue("System_id")
	additionalInfo := ctx.PostValue("AdditionalInfo")

	sysID, err := strconv.Atoi(system_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing system_id!")
	}

	for _, artefactId := range strings.Split(artefact_ids, ",") {
		artID, err := strconv.Atoi(artefactId)
		if err != nil {
			ctx.ViewData("error", "Error: Error parsing artefactId!")
			continue
		}
		err = dbprovider.GetDBManager().AddSMSArtefactPartOfSystem(sysID, artID, additionalInfo)
		if err != nil {
			ctx.ViewData("error", "Error: Not able to add artefact-system link!")
		}
	}

	ctx.Params().Set("id", system_id)
	ShowSMSSystem(ctx)
}

func RemoveSMSArtefactPartOfSystem(ctx iris.Context) {
	systemId := ctx.URLParam("system_id")
	artefactId := ctx.URLParam("artefact_id")

	sysID, err := strconv.Atoi(systemId)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing system_id!")
		SMSProjects(ctx)
		return
	}

	artID, err := strconv.Atoi(artefactId)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing artefact_id!")
		SMSProjects(ctx)
		return
	}

	err = dbprovider.GetDBManager().RemoveSMSArtefactPartOfSystem(sysID, artID)
	if err != nil {
		ctx.ViewData("error", "Error: Error removing artefact-system link!")
	}

	ctx.Params().Set("id", systemId)
	ShowSMSSystem(ctx)
}