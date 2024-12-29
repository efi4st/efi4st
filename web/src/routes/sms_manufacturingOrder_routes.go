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

// GET
func CreateSMSManufacturingOrder(ctx iris.Context) {
	id := ctx.Params().Get("id")
	ctx.ViewData("systemId", id)
	ctx.View("sms_createManufacturingOrder.html")
}


// POST
func AddSMSManufacturingOrder(ctx iris.Context) {

	system_id := ctx.PostValue("System_id")
	packageReference := ctx.PostValue("PackageReference")
	description := ctx.PostValue("Description")

	i, err := strconv.Atoi(system_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing system_id!")
	}

	err = dbprovider.GetDBManager().AddSMSManufacturingOrder(i, packageReference, description)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Not able to add ManufacturingOrder!")
	}

	ctx.Params().Set("id", system_id)
	ShowSMSSystem(ctx)
}

func ShowSMSManufacturingOrder(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing manufacturing order Id!")
	}

	manufacturingOrder := dbprovider.GetDBManager().GetSMSManufacturingOrderInfo(i)

	ctx.ViewData("manufacturingOrder", manufacturingOrder)
	ctx.View("sms_showManufacturingOrder.html")
}

