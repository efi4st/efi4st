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

func SMSProjectSettings(ctx iris.Context) {
	// Hole alle ProjectSettings aus der Datenbank
	projectSettings, err := dbprovider.GetDBManager().GetProjectSettings()
	if err != nil {
		ctx.ViewData("error", "Error: Error getting settings!")
	}
	ctx.ViewData("error", "")

	// Falls keine Einstellungen vorhanden sind, setze eine Fehlermeldung
	if len(projectSettings) < 1 {
		ctx.ViewData("error", "Error: No project settings available. Add one!")
	}

	// Füge die Liste der Einstellungen zur View hinzu
	ctx.ViewData("projectSettingsList", projectSettings)
	ctx.View("sms_projectSettings.html")
}

// GET
func CreateSMSProjectSettings(ctx iris.Context) {
	// Lade die HTML-View für das Erstellen einer neuen Projekteinstellung
	ctx.View("sms_createProjectSetting.html")
}

// POST
func AddSMSProjectSetting(ctx iris.Context) {
	// Hole die POST-Daten
	keyName := ctx.PostValue("key_name")
	valueType := ctx.PostValue("value_type")
	defaultValue := ctx.PostValue("default_value")

	// Überprüfe, ob der value_type ein gültiger ENUM-Wert ist
	validTypes := map[string]bool{
		"string": true,
		"int":    true,
		"boolean": true,
		"json":    true,
	}

	// Überprüfe, ob der valueType gültig ist
	if !validTypes[valueType] {
		ctx.ViewData("error", "Invalid value_type provided. Must be 'string', 'int', 'boolean', or 'json'.")
		ctx.View("sms_createProjectSetting.html")
		return
	}

	// Füge das ProjectSetting hinzu
	err := dbprovider.GetDBManager().AddProjectSetting(keyName, valueType, defaultValue)
	if err != nil {
		ctx.ViewData("error", "Error adding project setting.")
	} else {
		ctx.Redirect("/sms_projectSettings")
	}
}

func RemoveSMSProjectSettings(ctx iris.Context) {
	// ID aus der URL-Parameter holen
	id := ctx.Params().Get("id")

	// In Integer umwandeln
	i, err := strconv.Atoi(id)
	if err != nil {
		ctx.ViewData("error", "Error: Invalid project setting ID!")
		ctx.View("sms_projectSettings.html")
		return
	}

	// Löschen des Eintrags aus der Datenbank
	err = dbprovider.GetDBManager().DeleteProjectSetting(i)

	// Fehlerbehandlung
	ctx.ViewData("error", "")
	if err != nil {
		ctx.ViewData("error", "Error: Error removing project setting!")
	}

	// Aktualisierte Liste der Einstellungen abrufen
	projectSettings, err := dbprovider.GetDBManager().GetProjectSettings()
	if err != nil {
		ctx.ViewData("error", "Error: Error getting settings!")
	}
	ctx.ViewData("projectSettingsList", projectSettings)

	// View rendern
	ctx.View("sms_projectSettings.html")
}