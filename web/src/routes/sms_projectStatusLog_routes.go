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

// GET /projectstatus/create/{project_id}
func CreateSMSProjectStatusLog(ctx iris.Context) {
	projectID := ctx.Params().Get("project_id")
	ctx.ViewData("projectId", projectID)
	ctx.View("sms_createProjectStatusLog.html") // HTML-View für Eingabe
}

// POST /projectstatus/add
func AddSMSProjectStatusLog(ctx iris.Context) {
	projectID := ctx.PostValue("ProjectID")
	status := ctx.PostValue("Status")
	note := ctx.PostValue("Note")
	accessGroup := ctx.PostValue("AccessGroup")

	i, err := strconv.Atoi(projectID)
	if err != nil {
		ctx.ViewData("error", "Fehler beim Parsen der ProjectID")
		ctx.View("sms_createProjectStatusLog.html")
		return
	}

	err = dbprovider.GetDBManager().AddSMSProjectStatus(i, status, note, accessGroup)

	if err != nil {
		ctx.ViewData("error", "Fehler beim Speichern des Status-Logs")
		ctx.View("sms_createProjectStatusLog.html")
		return
	}

	// Redirect zurück zur Projektseite oder Statusliste
	ctx.Params().Set("id", projectID)
	ShowSMSProject(ctx)
}

// GET /projectstatus/show/{project_id}
func ShowSMSProjectStatusLog(ctx iris.Context) {
	id := ctx.Params().Get("id")
	projectID, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err != nil {
		ctx.ViewData("error", "Fehler beim Parsen der ProjectID")
		ctx.View("sms_showProjectStatusLog.html")
		return
	}

	statusLogs := dbprovider.GetDBManager().GetSMSProjectStatusLogsForProject(projectID)

	ctx.ViewData("statusLogs", statusLogs)
	ctx.ViewData("projectId", projectID)
	ctx.View("sms_showProjectStatusLog.html") // View mit Tabelle
}