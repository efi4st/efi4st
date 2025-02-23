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

func SMSProjects(ctx iris.Context) {

	projects := dbprovider.GetDBManager().GetSMSProjects()
	fmt.Printf("->"+string(len(projects)))
	ctx.ViewData("error", "")

	if len(projects) < 1 {
		ctx.ViewData("error", "Error: No projects available. Add one!")
	}
	fmt.Printf(strconv.Itoa(len(projects)))
	fmt.Printf(projects[0].Name())
	ctx.ViewData("projectList", projects)
	ctx.View("sms_projects.html")
}

// GET
func CreateSMSProject(ctx iris.Context) {
	// Projekt-Typen aus der Datenbank abrufen
	projectTypes := dbprovider.GetDBManager().GetSMSProjectTypes()

	// Alle verfügbaren ProjectSettings abrufen
	projectSettings, err := dbprovider.GetDBManager().GetProjectSettings()
	if err != nil {
		ctx.ViewData("error", "Error loading project settings!")
	}

	// Daten an die View übergeben
	ctx.ViewData("typeList", projectTypes)           // Projekt-Typen für Drop-down
	ctx.ViewData("projectSettingsList", projectSettings) // Verfügbare Einstellungen

	// HTML-Seite rendern
	ctx.View("sms_createProject.html")
}

// POST
func AddSMSProject(ctx iris.Context) {
	projectName := ctx.PostValue("ProjectName")
	customer := ctx.PostValue("Customer")
	projecttypeId := ctx.PostValue("ProjecttypeId")
	reference := ctx.PostValue("Reference")

	// Projekt-Typ ID konvertieren
	iProjectType, err := strconv.Atoi(projecttypeId)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing projecttypeId!")
		ctx.View("sms_createProject.html")
		return
	}

	// Neues Projekt erstellen und ID erhalten
	projectID, err := dbprovider.GetDBManager().AddSMSProject(projectName, customer, iProjectType, reference)
	if err != nil {
		ctx.ViewData("error", "Error: Not able to add project!")
		ctx.View("sms_createProject.html")
		return
	}

	// Alle ausgewählten Settings abrufen
	selectedSettingsStr, err := ctx.PostValueMany("selectedSettings")
	if err != nil {
		ctx.ViewData("error", "Error retrieving selected settings!")
		return
	}

	// Die ausgewählten Setting-IDs in eine Liste umwandeln (angenommen, sie sind durch Kommas getrennt)
	selectedSettings := strings.Split(selectedSettingsStr, ",")

	for _, settingID := range selectedSettings {
		// Die SettingID in int konvertieren
		iSettingID, err := strconv.Atoi(settingID)
		if err != nil {
			fmt.Println("Error parsing setting ID:", err)
			continue
		}

		// Benutzerdefinierten Wert abrufen (falls eingegeben)
		settingValue := ctx.PostValue(fmt.Sprintf("SettingValue_%d", iSettingID))

		// Falls kein Wert eingegeben wurde, den Default-Wert setzen
		if settingValue == "" {
			defaultVal, err := dbprovider.GetDBManager().GetProjectSettingDefaultValue(iSettingID)
			if err != nil {
				fmt.Println("Error retrieving default value:", err)
				continue
			}
			settingValue = defaultVal
		}

		// ProjectSettingLink speichern
		err = dbprovider.GetDBManager().AddProjectSettingLink(projectID, iSettingID, settingValue)
		if err != nil {
			fmt.Println("Error adding project setting link:", err)
		}
	}

	// Erfolgreicher Abschluss -> Projektliste neu laden
	ctx.ViewData("error", "")
	projects := dbprovider.GetDBManager().GetSMSProjects()
	ctx.ViewData("projectList", projects)
	ctx.View("sms_projects.html")
}

func ShowSMSProject(ctx iris.Context) {
	// Hole die Projekt-ID aus den URL-Parametern
	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing project Id!")
		return
	}

	// Hole Projektinformationen und andere Daten
	project := dbprovider.GetDBManager().GetSMSProjectInfo(i)
	deviceInstanceList := dbprovider.GetDBManager().GetDeviceInstanceListForProject(i)
	systemList := dbprovider.GetDBManager().GetSMSProjectBOMForProject(i)
	issuesForThisProject, err := dbprovider.GetDBManager().GetSMSIssuesForProject(i)

	// Hole die verlinkten Settings für das Projekt
	projectSettings, err := dbprovider.GetDBManager().GetLinkedProjectSettings(i)
	if err != nil {
		ctx.ViewData("error", "Error: Could not retrieve project settings!")
		return
	}

	// Übergebe alle Daten an die View
	ctx.ViewData("deviceInstanceList", deviceInstanceList)
	ctx.ViewData("systemList", systemList)
	ctx.ViewData("issuesForThisProject", issuesForThisProject)
	ctx.ViewData("project", project)
	ctx.ViewData("projectSettings", projectSettings)

	// Lade die View
	ctx.View("sms_showProject.html")
}

func RemoveSMSProject(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveSMSProject(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing project!")
	}

	projects := dbprovider.GetDBManager().GetSMSProjects()

	ctx.ViewData("projectList", projects)
	ctx.View("sms_projects.html")
}

func SMSProjectIPs(ctx iris.Context) {
	// Projekt-ID aus URL holen
	projectIDStr := ctx.Params().Get("project_id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		ctx.ViewData("error", "Invalid project ID!")
		ctx.View("sms_projectIPs.html")
		return
	}

	// IPs abrufen
	ipList, err := dbprovider.GetDBManager().GetIPsForProject(projectID)
	if err != nil {
		ctx.ViewData("error", "Error retrieving IPs for project!")
		ctx.View("sms_projectIPs.html")
		return
	}

	// Projekt-ID und IP-Liste an die View übergeben
	ctx.ViewData("projectID", projectID)
	ctx.ViewData("ipList", ipList)
	ctx.View("sms_projectIPs.html")
}