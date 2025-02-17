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
func CreateSMSProjectSettingsLink(ctx iris.Context) {
	// Hole die Projekt-ID aus den Parametern
	id := ctx.Params().Get("id")
	projectID, err := strconv.Atoi(id)
	if err != nil {
		ctx.ViewData("error", "Error: Invalid project ID!")
		ctx.View("sms_createProjectSettingsLink.html")
		return
	}

	// Hole die Liste der verfügbaren Einstellungen aus der Datenbank
	settings, err := dbprovider.GetDBManager().GetAvailableProjectSettings(projectID)
	if err != nil {
		ctx.ViewData("error", "Error: Could not retrieve project settings!")
		ctx.View("sms_createProjectSettingsLink.html")
		return
	}
	fmt.Println("Available Project Settings:", settings)
	// Übergebe die Projekt-ID und die verfügbaren Einstellungen an die View
	ctx.ViewData("projectId", projectID)
	ctx.ViewData("settingsList", settings)

	// Lade das entsprechende HTML-Template
	ctx.View("sms_createProjectSettingsLink.html")
}

// POST
func AddSMSProjectSettingsLink(ctx iris.Context) {
	// Debug: Alle POST-Werte ausgeben
	fmt.Println("POST values received:")
	settingsIDsStr := ctx.PostValue("Settings_ids[]")
	fmt.Println("Raw Settings_ids:", settingsIDsStr) // Debug-Ausgabe der raw Daten
	settingsIDs, err := ctx.PostValueMany("Settings_ids[]")
	fmt.Println("Raw Settings_ids:", settingsIDs) // Debugging der Rohdaten

	projectID := ctx.PostValue("Project_id")

	// Konvertiere die Projekt-ID in eine Zahl
	pID, err := strconv.Atoi(projectID)
	if err != nil {
		fmt.Println("Error: Invalid project ID!", err)
		ctx.ViewData("error", "Error: Invalid project ID!")
		ctx.View("sms_projectSettings.html")
		return
	}

	// Step 1: Split the comma-separated string into a slice of strings
	listOfSettings := strings.Split(settingsIDs, ",")
	fmt.Println("Parsed Settings_ids:", settingsIDs)

	// Step 2: Iterate through each settingID and convert to int
	for _, settingIDStr := range listOfSettings {
		// Konvertiere jeden String in int
		settingID, err := strconv.Atoi(settingIDStr)
		if err != nil {
			fmt.Println("Error converting settingID:", err)
			continue
		}
		settingValue := ctx.PostValue(fmt.Sprintf("SettingValue_%d", settingID))

		err = dbprovider.GetDBManager().AddProjectSettingLink(pID, settingID, settingValue) // Füge das Setting hinzu
		if err != nil {
			fmt.Println("Error: Could not add setting to project!", err)
			ctx.ViewData("error", "Error: Could not add setting to project!")
		}
	}

	// Erfolgreich hinzugefügt: Redirect zur Projektseite
	fmt.Println("Redirecting to project view:", fmt.Sprintf("/sms_projects/show/%d", pID))
	ctx.Redirect(fmt.Sprintf("/sms_projects/show/%d", pID))
}

func RemoveSMSProjectSettingsLink(ctx iris.Context) {
	// Hole sowohl die Project ID als auch die Setting ID aus den Parametern
	projectID := ctx.Params().Get("project_id")
	settingID := ctx.Params().Get("setting_id")

	// Konvertiere die IDs in Integer
	pID, err := strconv.Atoi(projectID)
	sID, err := strconv.Atoi(settingID)

	// Fehlerbehandlung, falls die IDs nicht richtig konvertiert werden können
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing project_id or setting_id!")
	}

	// Entferne den Link zwischen Project und Setting
	err = dbprovider.GetDBManager().DeleteProjectSettingLink(pID, sID)

	// Fehlerbehandlung, falls beim Löschen etwas schiefgeht
	ctx.ViewData("error", "")
	if err != nil {
		ctx.ViewData("error", "Error: Error removing project-setting link!")
	}

	// Weiterleitung zur Projekt-Seite oder einer Liste von Projekt-Einstellungen
	ctx.Redirect(fmt.Sprintf("/sms_projectSettings/show/%d", pID))
}