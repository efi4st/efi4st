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

func SMSComponent(ctx iris.Context) {

	components := dbprovider.GetDBManager().GetSMSComponent()
	ctx.ViewData("error", "")

	if len(components) < 1 {
		ctx.ViewData("error", "Error: No components available. Add one!")
	}
	ctx.ViewData("componentList", components)
	ctx.View("sms_components.html")
}

// GET
func CreateSMSComponent(ctx iris.Context) {

	ctx.View("sms_createComponent.html")
}

// POST
func AddSMSComponent(ctx iris.Context) {

	componentName := ctx.PostValue("ComponentName")
	componentType := ctx.PostValue("ComponentType")
	version := ctx.PostValue("Version")
	license := ctx.PostValue("License")
	thirdParty := ctx.PostValue("ThirdParty")
	thirdPartyBool, err := strconv.ParseBool(thirdParty)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing thirdParty Bool!")
	}
	releaseNote := ctx.PostValue("ReleaseNote")

	err = dbprovider.GetDBManager().AddSMSComponent(componentName, componentType, version, license, thirdPartyBool, releaseNote)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Not able to add component!")
	}
	components := dbprovider.GetDBManager().GetSMSComponent()
	ctx.ViewData("componentList", components)
	ctx.View("sms_components.html")
}

func ShowSMSComponent(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing component Id!")
	}

	component := dbprovider.GetDBManager().GetSMSComponentInfo(i)
	applicationsParentsOfComponent := dbprovider.GetDBManager().GetSMSComponentPartOfSoftwareForComponent(i)
	issuesForThisComponent, err := dbprovider.GetDBManager().GetSMSIssuesForComponent(i)

	ctx.ViewData("applicationsParentsOfComponent", applicationsParentsOfComponent)
	ctx.ViewData("component", component)
	ctx.ViewData("issuesForThisComponent", issuesForThisComponent)
	ctx.View("sms_showComponent.html")
}

func RemoveSMSComponent(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveSMSComponent(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing component!")
	}

	components := dbprovider.GetDBManager().GetSMSComponent()

	ctx.ViewData("componentList", components)
	ctx.View("sms_components.html")
}