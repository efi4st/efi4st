/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (
	"fmt"
	"github.com/efi4st/efi4st/classes"
	"github.com/efi4st/efi4st/dbprovider"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"strconv"
	"time"
)

// LIST VIEW
func SMSUpdateCenters(ctx iris.Context) {
	centers, err := dbprovider.GetDBManager().GetAllSMSUpdateCenters()
	ctx.ViewData("error", "")

	if err != nil || len(centers) < 1 {
		ctx.ViewData("error", "Error: No update centers available!")
	}

	ctx.ViewData("centerList", centers)
	ctx.View("sms_updateCenter.html")
}

// CREATE FORM
func CreateSMSUpdateCenter(ctx iris.Context) {
	projects:= dbprovider.GetDBManager().GetSMSProjects()

	ctx.ViewData("projectList", projects)
	ctx.View("sms_createUpdateCenter.html")
}

// HANDLE POST
func AddSMSUpdateCenter(ctx iris.Context) {
	projectID, _ := strconv.Atoi(ctx.PostValue("ProjectID"))
	updaterID, _ := strconv.Atoi(ctx.PostValue("UpdaterID"))
	updaterType := ctx.PostValue("UpdaterType")
	version := ctx.PostValue("Version")
	environment := ctx.PostValue("Environment")
	status := ctx.PostValue("Status")
	description := ctx.PostValueTrim("Description")
	note := ctx.PostValueTrim("Note")
	owner := ctx.PostValue("Owner")

	err := dbprovider.GetDBManager().AddSMSUpdateCenter(projectID, updaterID, updaterType, version, environment, status, description, note, owner)
	ctx.ViewData("error", "")

	if err != nil {
		fmt.Println("AddSMSUpdateCenter ERROR:", err)
		ctx.ViewData("error", "Error: Unable to add update center!")
	}

	centers, _ := dbprovider.GetDBManager().GetAllSMSUpdateCenters()
	ctx.ViewData("centerList", centers)
	ctx.View("sms_updateCenter.html")
}

// SHOW ONE
func ShowSMSUpdateCenter(ctx iris.Context) {
	id, err := strconv.Atoi(ctx.Params().Get("id"))
	ctx.ViewData("error", "")

	if err != nil {
		ctx.ViewData("error", "Error: Invalid update center ID!")
	}

	center, err := dbprovider.GetDBManager().GetSMSUpdateCenterByID(id)
	if err != nil {
		ctx.ViewData("error", "Error: Update center not found!")
	}

	ctx.ViewData("center", center)
	ctx.View("sms_showUpdateCenter.html")
}

// DELETE
func RemoveSMSUpdateCenter(ctx iris.Context) {
	id, _ := strconv.Atoi(ctx.Params().Get("id"))

	err := dbprovider.GetDBManager().DeleteSMSUpdateCenter(id)
	ctx.ViewData("error", "")

	if err != nil {
		ctx.ViewData("error", "Error: Unable to remove update center!")
	}

	centers, _ := dbprovider.GetDBManager().GetAllSMSUpdateCenters()
	ctx.ViewData("centerList", centers)
	ctx.View("sms_updateCenter.html")
}

func PingSMSUpdateCenter(ctx iris.Context) {
	id, _ := strconv.Atoi(ctx.Params().Get("id"))
	now := time.Now()

	err := dbprovider.GetDBManager().UpdateSMSUpdateCenterLastContact(id, &now)

	if err != nil {
		ctx.JSON(iris.Map{"status": "error", "message": "Could not update last contact"})
		return
	}

	ctx.JSON(iris.Map{"status": "ok", "message": "Last contact updated"})
}

func EditSMSUpdateCenter(ctx iris.Context) {
	id, err := strconv.Atoi(ctx.Params().Get("id"))
	ctx.ViewData("error", "")

	if err != nil {
		ctx.ViewData("error", "Error: Invalid update center ID!")
		ctx.View("sms_updateCenter.html")
		return
	}

	center, err := dbprovider.GetDBManager().GetSMSUpdateCenterByID(id)
	if err != nil {
		ctx.ViewData("error", "Error: Update center not found!")
		ctx.View("sms_updateCenter.html")
		return
	}

	projects := dbprovider.GetDBManager().GetSMSProjects()

	ctx.ViewData("projectList", projects)
	ctx.ViewData("updateCenter", center)
	ctx.View("sms_editUpdateCenter.html")
}

func UpdateSMSUpdateCenter(ctx iris.Context) {
	ctx.ViewData("error", "")

	// ID kommt aus der Route: /sms_update_centers/update/{id}
	idStr := ctx.Params().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.ViewData("error", "Invalid ID provided.")
		ctx.View("sms_updateCenter.html")
		return
	}

	// Rest aus dem Form
	projectID, err := strconv.Atoi(ctx.PostValue("project_id"))
	if err != nil {
		ctx.ViewData("error", "Invalid Project ID.")
		ctx.View("sms_updateCenter.html")
		return
	}

	updaterID, err := strconv.Atoi(ctx.PostValue("updater_id"))
	if err != nil {
		ctx.ViewData("error", "Invalid Updater ID.")
		ctx.View("sms_updateCenter.html")
		return
	}

	center := classes.Sms_UpdateCenter{
		ID:          id,
		ProjectID:   projectID,
		UpdaterID:   updaterID,
		UpdaterType: ctx.PostValue("updater_type"),
		Version:     ctx.PostValue("version"),
		Environment: ctx.PostValue("environment"),
		Status:      ctx.PostValue("status"),
		Description: ctx.PostValue("description"),
		Note:        ctx.PostValue("note"),
		Owner:       ctx.PostValue("owner"),
	}

	err = dbprovider.GetDBManager().UpdateSMSUpdateCenter(center)
	if err != nil {
		ctx.ViewData("error", "Error: Unable to update update center!")
		centers, _ := dbprovider.GetDBManager().GetAllSMSUpdateCenters()
		ctx.ViewData("centerList", centers)
		ctx.View("sms_updateCenter.html")
		return
	}

	// Erfolgreich aktualisiert – weiterleiten
	ctx.Redirect("/sms_update_centers")
}