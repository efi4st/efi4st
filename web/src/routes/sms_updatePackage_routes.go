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

func SMSUpdatePackages(ctx iris.Context) {
	packages, err := dbprovider.GetDBManager().GetAllSMSUpdatePackages()
	ctx.ViewData("error", "")

	if err != nil || len(packages) < 1 {
		ctx.ViewData("error", "Error: No update packages available!")
	}

	ctx.ViewData("packageList", packages)
	ctx.View("sms_update_packages.html")
}

// GET
func CreateSMSUpdatePackage(ctx iris.Context) {
	updates, _ := dbprovider.GetDBManager().GetAllSMSUpdates()
	deviceTypes := dbprovider.GetDBManager().GetSMSDeviceTypes()

	ctx.ViewData("updateList", updates)
	ctx.ViewData("deviceTypeList", deviceTypes)
	ctx.View("sms_createUpdatePackage.html")
}

// POST
func AddSMSUpdatePackage(ctx iris.Context) {
	updateID, _ := strconv.Atoi(ctx.PostValue("UpdateID"))
	deviceTypeID, _ := strconv.Atoi(ctx.PostValue("DeviceTypeID"))
	packageIdentifier := ctx.PostValue("PackageIdentifier")
	packageVersion := ctx.PostValue("PackageVersion")
	packageName := ctx.PostValue("PackageName")
	updatePackageFile := ctx.PostValue("UpdatePackageFile")
	creator := ctx.PostValue("Creator")
	packageDescription := ctx.PostValueTrim("PackageDescription")
	isTested := ctx.PostValue("IsTested") == "on"

	err := dbprovider.GetDBManager().AddSMSUpdatePackage(updateID, deviceTypeID, packageIdentifier, packageVersion, packageName, updatePackageFile, creator, &packageDescription, isTested)
	ctx.ViewData("error", "")

	if err != nil {
		ctx.ViewData("error", "Error: Unable to add update package!")
	}

	packages, _ := dbprovider.GetDBManager().GetAllSMSUpdatePackages()
	ctx.ViewData("packageList", packages)
	ctx.View("sms_update_packages.html")
}

func ShowSMSUpdatePackage(ctx iris.Context) {
	id, err := strconv.Atoi(ctx.Params().Get("id"))

	ctx.ViewData("error", "")
	if err != nil {
		ctx.ViewData("error", "Error: Invalid package ID!")
	}

	pkg, err := dbprovider.GetDBManager().GetSMSUpdatePackageByID(id)
	if err != nil {
		ctx.ViewData("error", "Error: Update package not found!")
	}

	ctx.ViewData("package", pkg)
	ctx.View("sms_showUpdatePackage.html")
}

func RemoveSMSUpdatePackage(ctx iris.Context) {
	id, _ := strconv.Atoi(ctx.Params().Get("id"))

	err := dbprovider.GetDBManager().DeleteSMSUpdatePackage(id)
	ctx.ViewData("error", "")

	if err != nil {
		ctx.ViewData("error", "Error: Unable to remove update package!")
	}

	packages, _ := dbprovider.GetDBManager().GetAllSMSUpdatePackages()
	ctx.ViewData("packageList", packages)
	ctx.View("sms_update_packages.html")
}