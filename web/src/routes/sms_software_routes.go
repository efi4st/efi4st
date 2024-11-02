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

func SMSSoftware(ctx iris.Context) {

	softwares := dbprovider.GetDBManager().GetSMSSoftware()
	ctx.ViewData("error", "")

	if len(softwares) < 1 {
		ctx.ViewData("error", "Error: No softwares available. Add one!")
	}
	ctx.ViewData("softwareList", softwares)
	ctx.View("sms_softwares.html")
}

// GET
func CreateSMSSoftware(ctx iris.Context) {

	softwareTypes := dbprovider.GetDBManager().GetSMSSoftwareTypes()

	ctx.ViewData("typeList", softwareTypes)
	ctx.View("sms_createSoftware.html")
}

// POST
func AddSMSSoftware(ctx iris.Context) {

	softwaretypeId := ctx.PostValue("SoftwaretypeId")
	version := ctx.PostValue("Version")
	license := ctx.PostValue("License")
	thirdParty := ctx.PostValue("ThirdParty")
	thirdPartyBool, err := strconv.ParseBool(thirdParty)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing thirdParty Bool!")
	}
	releaseNote := ctx.PostValue("ReleaseNote")

	i, err := strconv.Atoi(softwaretypeId)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing softwaretypeId!")
	}

	err = dbprovider.GetDBManager().AddSMSSoftware(i, version, license, thirdPartyBool, releaseNote)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Not able to add software!")
	}
	softwares := dbprovider.GetDBManager().GetSMSSoftware()
	ctx.ViewData("softwareList", softwares)
	ctx.View("sms_softwares.html")
}

func ShowSMSSoftware(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing software Id!")
	}

	software := dbprovider.GetDBManager().GetSMSSoftwareInfo(i)
	componentsUnderSoftware := dbprovider.GetDBManager().GetSMSComponentPartOfSoftwareForSoftware(i)
	devicesParentsOfSoftware := dbprovider.GetDBManager().GetSMSSoftwarePartOfDeviceForSoftware(i)

	ctx.ViewData("devicesParentsOfSoftware", devicesParentsOfSoftware)
	ctx.ViewData("componentsUnderSoftware", componentsUnderSoftware)
	ctx.ViewData("software", software)
	ctx.View("sms_showSoftware.html")
}

func RemoveSMSSoftware(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveSMSSoftware(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing software!")
	}

	softwares := dbprovider.GetDBManager().GetSMSSoftware()

	ctx.ViewData("softwareList", softwares)
	ctx.View("sms_softwares.html")
}