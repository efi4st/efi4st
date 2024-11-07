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

func SMSSystems(ctx iris.Context) {

	systems := dbprovider.GetDBManager().GetSMSSystems()
	ctx.ViewData("error", "")

	if len(systems) < 1 {
		ctx.ViewData("error", "Error: No projects available. Add one!")
	}
	ctx.ViewData("systemList", systems)
	ctx.View("sms_systems.html")
}

// GET
func CreateSMSSystem(ctx iris.Context) {

	systemTypes := dbprovider.GetDBManager().GetSMSSystemTypes()

	ctx.ViewData("typeList", systemTypes)
	ctx.View("sms_createSystem.html")
}


// POST
func AddSMSSystem(ctx iris.Context) {

	systemtypeId := ctx.PostValue("SystemtypeId")
	version := ctx.PostValue("Version")

	i, err := strconv.Atoi(systemtypeId)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing projecttypeId!")
	}

	err = dbprovider.GetDBManager().AddSMSSystem(i, version, time.Now().String())

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Not able to add project!")
	}
	systems := dbprovider.GetDBManager().GetSMSSystems()
	ctx.ViewData("systemList", systems)
	ctx.View("sms_systems.html")
}

func ShowSMSSystem(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing project Id!")
	}

	system := dbprovider.GetDBManager().GetSMSSystemInfo(i)
	devicesUnderSystem := dbprovider.GetDBManager().GetSMSDevicePartOfSystemForSystem(i)

	ctx.ViewData("devicesUnderSystem", devicesUnderSystem)
	ctx.ViewData("system", system)
	ctx.View("sms_showSystem.html")
}

func RemoveSMSSystem(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveSMSSystem(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing project!")
	}

	systems := dbprovider.GetDBManager().GetSMSSystems()

	ctx.ViewData("systemList", systems)
	ctx.View("sms_systems.html")
}