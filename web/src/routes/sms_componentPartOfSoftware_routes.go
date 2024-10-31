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
func CreateSMSComponentPartOfSoftware(ctx iris.Context) {

	id := ctx.Params().Get("id")
	i, err := strconv.Atoi(id)
	components := dbprovider.GetDBManager().GetSMSComponent()
	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing software id!")
	}

	ctx.ViewData("softwareId", i)
	ctx.ViewData("componentList", components)
	ctx.View("sms_createComponentPartOfSoftware.html")
}

// POST
func AddSMSComponentPartOfSoftware(ctx iris.Context) {

	component_ids, err := ctx.PostValueMany("Component_ids")
	software_id := ctx.PostValue("Software_id")
	additionalInfo := ctx.PostValue("AdditionalInfo")

	iI, err := strconv.Atoi(software_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing software_id!")
	}

	for index, componentId := range strings.Split(component_ids,","){
		iD, err := strconv.Atoi(componentId)
		if err != nil {
			ctx.ViewData("error", "Error: Error parsing componentId!")
		}
		err = dbprovider.GetDBManager().AddSMSComponentPartOfSoftware(iI, iD, additionalInfo)
		fmt.Printf(strconv.Itoa(index))
		ctx.ViewData("error", "")
		if err !=nil {
			ctx.ViewData("error", "Error: Not able to add component software link!")
		}
	}

	ctx.Params().Set("id", software_id)
	ShowSMSSoftware(ctx)
}

func RemoveSMSComponentPartOfSoftware(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveSMSComponentPartOfSoftware(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing component software link!")
	}

	SMSProjects(ctx)
}