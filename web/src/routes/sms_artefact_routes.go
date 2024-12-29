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

func SMSArtefact(ctx iris.Context) {

	artefacts := dbprovider.GetDBManager().GetSMSArtefact()
	ctx.ViewData("error", "")

	if len(artefacts) < 1 {
		ctx.ViewData("error", "Error: No artefacts available. Add one!")
	}
	ctx.ViewData("artefactList", artefacts)
	ctx.View("sms_artefacts.html")
}

// GET
func CreateSMSArtefact(ctx iris.Context) {

	artefactTypes := dbprovider.GetDBManager().GetSMSArtefactTypes()

	ctx.ViewData("artefactList", artefactTypes)
	ctx.View("sms_createArtefact.html")
}

// POST
func AddSMSArtefact(ctx iris.Context) {

	artefacttypeId := ctx.PostValue("ArtefacttypeId")
	name := ctx.PostValue("Name")
	version := ctx.PostValue("Version")

	i, err := strconv.Atoi(artefacttypeId)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing artefacttypeId!")
	}

	err = dbprovider.GetDBManager().AddSMSArtefact(i, name, version)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Not able to add artefact!")
	}
	artefacts := dbprovider.GetDBManager().GetSMSArtefact()
	ctx.ViewData("artefactList", artefacts)
	ctx.View("sms_artefacts.html")
}

func ShowSMSArtefact(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing artefact Id!")
	}

	artefact := dbprovider.GetDBManager().GetSMSArtefactInfo(i)
	devicesParentsOfArtefact := dbprovider.GetDBManager().GetSMSArtefactPartOfDeviceForArtefact(i)

	ctx.ViewData("devicesParentsOfArtefact", devicesParentsOfArtefact)
	ctx.ViewData("artefact", artefact)
	ctx.View("sms_showArtefact.html")
}

func RemoveSMSArtefact(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveSMSArtefact(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing artefact!")
	}

	artefacts := dbprovider.GetDBManager().GetSMSArtefact()

	ctx.ViewData("artefactList", artefacts)
	ctx.View("sms_artefacts.html")
}