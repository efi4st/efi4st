/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (
	"github.com/efi4st/efi4st/classes"
	"github.com/efi4st/efi4st/dbprovider"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"strconv"
)

func SMSUpdates(ctx iris.Context) {
	updates, err := dbprovider.GetDBManager().GetAllSMSUpdates()
	ctx.ViewData("error", "")

	if err != nil || len(updates) < 1 {
		ctx.ViewData("error", "Error: No updates available!")
	}

	ctx.ViewData("updateList", updates)
	ctx.View("sms_updates.html")
}

// GET
func CreateSMSUpdate(ctx iris.Context) {
	// Hier laden wir die Systeme, die im Dropdown angezeigt werden sollen
	systems, err := dbprovider.GetDBManager().GetAllSystems()
	if err != nil {
		ctx.ViewData("error", "Error: Unable to load systems!")
		ctx.View("sms_createUpdate.html")
		return
	}

	// Wir übergeben die Systeme zur Anzeige im Dropdown
	ctx.ViewData("systems", systems)
	ctx.View("sms_createUpdate.html")
}

// POST
// POST Route zum Hinzufügen eines neuen Updates
func AddSMSUpdate(ctx iris.Context) {
	fromSystemID, _ := strconv.Atoi(ctx.PostValue("from_system_id"))
	toSystemID, _ := strconv.Atoi(ctx.PostValue("to_system_id"))
	mandatorySystemID, _ := strconv.Atoi(ctx.PostValue("mandatory_system_id"))
	updateType := ctx.PostValue("update_type")
	projectName := ctx.PostValue("project_name")
	externalIssueLink := ctx.PostValue("external_issue_link")
	additionalInfo := ctx.PostValueTrim("additional_info")
	isApproved := ctx.PostValue("is_approved") == "on"

	// Füge das Update zur Datenbank hinzu
	err := dbprovider.GetDBManager().AddSMSUpdate(fromSystemID, toSystemID, mandatorySystemID, updateType, additionalInfo, isApproved, externalIssueLink, projectName)
	ctx.ViewData("error", "")

	if err != nil {
		ctx.ViewData("error", "Error: Unable to add update!")
	}

	// Lade alle Updates und zeige sie an
	updates, _ := dbprovider.GetDBManager().GetAllSMSUpdates()
	ctx.ViewData("updateList", updates)
	ctx.View("sms_updates.html")
}
func ShowSMSUpdateDetails(ctx iris.Context) {
	updateID, err := strconv.Atoi(ctx.Params().Get("id"))
	if err != nil {
		ctx.ViewData("error", "Invalid update ID!")
		ctx.View("error.html")
		return
	}

	update, err := dbprovider.GetDBManager().GetSMSUpdateByID(updateID)
	if err != nil {
		ctx.ViewData("error", "Update not found!")
		ctx.View("error.html")
		return
	}

	ctx.ViewData("update", update)
	ctx.View("sms_showUpdate.html")
}

func RemoveSMSUpdate(ctx iris.Context) {
	id, _ := strconv.Atoi(ctx.Params().Get("id"))

	err := dbprovider.GetDBManager().DeleteSMSUpdate(id)
	ctx.ViewData("error", "")

	if err != nil {
		ctx.ViewData("error", "Error: Unable to remove update!")
	}

	updates, _ := dbprovider.GetDBManager().GetAllSMSUpdates()
	ctx.ViewData("updateList", updates)
	ctx.View("sms_updates.html")
}

func SMSUpdateEditPost(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.ViewData("error", "Invalid Update ID")
		ctx.Redirect("/sms_updates")
		return
	}

	// Checkbox "is_approved"
	isApproved := ctx.PostValue("is_approved") == "true"

	update := classes.Sms_UpdateDetails{
		ID:                  id,
		FromSystemID:        ctx.PostValueIntDefault("from_system_id", 0),
		ToSystemID:          ctx.PostValueIntDefault("to_system_id", 0),
		MandatorySystemID:   ctx.PostValueIntDefault("mandatory_system_id", 0),
		UpdateType:          ctx.PostValue("update_type"),
		AdditionalInfo:      ctx.PostValue("additional_info"),
		IsApproved:          isApproved,
		IssueLink:           ctx.PostValue("issue_link"),
		ProjectName:         ctx.PostValue("project_name"),
	}

	err = dbprovider.GetDBManager().UpdateSMSUpdate(update)
	if err != nil {
		ctx.ViewData("error", "Error updating SMS update")
		ctx.Redirect("/sms_updates/edit/" + strconv.Itoa(id) + "?error=Update%20failed")
		return
	}

	ctx.Redirect("/sms_updates/show/" + strconv.Itoa(id))
}

func EditSMSUpdateForm(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.ViewData("error", "Invalid Update ID")
		ctx.Redirect("/sms_updates")
		return
	}

	// Holt das Update aus der DB
	update, err := dbprovider.GetDBManager().GetSMSUpdateByID(id)
	if err != nil {
		ctx.ViewData("error", "Update not found")
		ctx.Redirect("/sms_updates")
		return
	}

	// Mögliche Update-Typen
	updateTypes := []string{"security", "bugfix", "feature", "maintenance"}

	// Daten für das Template setzen
	ctx.ViewData("update", update)
	ctx.ViewData("updateTypes", updateTypes)

	systems, err := dbprovider.GetDBManager().GetAllSystems()
	if err != nil {
		ctx.ViewData("error", "Could not load systems")
		ctx.Redirect("/sms_updates")
		return
	}
	ctx.ViewData("systems", systems)

	// Template rendern
	ctx.View("sms_editUpdate.html")
}