/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (
	"encoding/csv"
	"fmt"
	"github.com/efi4st/efi4st/dbprovider"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"gopkg.in/yaml.v3"
	"log"
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
	log.Printf("[PBOM] project=%d list len=%d", i, len(systemList))

	var currentSystemVersion string
	if len(systemList) > 0 {
		currentSystemVersion = systemList[0].SystemVersion // statt .Tmp()
	} else {
		currentSystemVersion = ""
		log.Println("Warnung: Kein System mit diesem Projekt verlinkt (systemList ist leer)")
	}

	issuesForThisProject, err := dbprovider.GetDBManager().GetSMSIssuesForProject(i)

	for i := range deviceInstanceList {
		dbprovider.GetDBManager().EnrichDeviceInstanceWithSystemInfo(&deviceInstanceList[i], currentSystemVersion)
	}

	// Hole die verlinkten Settings für das Projekt
	projectSettings, err := dbprovider.GetDBManager().GetLinkedProjectSettings(i)
	if err != nil {
		ctx.ViewData("error", "Error: Could not retrieve project settings!")
		return
	}

	statusLogs := dbprovider.GetDBManager().GetSMSProjectStatusLogsForProject(i)
	checklistTemplates := dbprovider.GetDBManager().GetAllChecklistTemplates()
	checklistInstances := dbprovider.GetDBManager().GetChecklistInstancesForProject(i)

	// Übergebe alle Daten an die View
	ctx.ViewData("checklistTemplates", checklistTemplates)
	ctx.ViewData("checklistInstances", checklistInstances)
	ctx.ViewData("statusLogs", statusLogs)
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

func SMSExportProjectIPsCSV(ctx iris.Context) {
	// Projekt-ID aus URL holen
	projectIDStr := ctx.Params().Get("project_id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Invalid project ID")
		return
	}

	// IPs abrufen
	ipList, err := dbprovider.GetDBManager().GetIPsForProject(projectID)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("Error retrieving IPs for project")
		return
	}

	// CSV-Header setzen
	ctx.ResponseWriter().Header().Set("Content-Type", "text/csv")
	ctx.ResponseWriter().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=project_%d_ips.csv", projectID))

	// CSV-Writer erstellen
	writer := csv.NewWriter(ctx.ResponseWriter())
	defer writer.Flush()

	// Spalten-Header schreiben
	header := []string{"IPAddress", "ApplicableVersions", "VLANID", "Description", "DeviceType", "InstanceCount", "Versions"}
	if err := writer.Write(header); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("Error writing CSV header")
		return
	}

	// Daten in CSV schreiben
	for _, ip := range ipList {
		record := []string{
			ip.IPAddress,
			ip.ApplicableVersions,
			"",
			"",
			ip.DeviceType,
			strconv.Itoa(ip.InstanceCount),
			ip.Versions,
		}

		// VLANID prüfen und konvertieren
		if ip.VLANID != nil {
			record[2] = strconv.Itoa(*ip.VLANID) // Dereferenzieren, um den int-Wert zu erhalten
		}

		// Description prüfen
		if ip.Description != nil {
			record[3] = *ip.Description
		}

		if err := writer.Write(record); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString("Error writing CSV data")
			return
		}
	}
}

func SMSExportProjectIPsCSVCustomer(ctx iris.Context) {
	// Projekt-ID aus URL holen
	projectIDStr := ctx.Params().Get("project_id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Invalid project ID!")
		return
	}

	// IPs abrufen
	ipList, err := dbprovider.GetDBManager().GetIPsForProject(projectID)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("Error retrieving IPs for project!")
		return
	}

	// CSV-Header setzen
	ctx.ResponseWriter().Header().Set("Content-Type", "text/csv")
	ctx.ResponseWriter().Header().Set("Content-Disposition", "attachment; filename=project_ips_minimal.csv")

	// CSV-Schreiber initialisieren
	writer := csv.NewWriter(ctx.ResponseWriter())
	defer writer.Flush()

	// Header-Zeile schreiben
	writer.Write([]string{"IP Address", "Device Type", "Description"})

	// Daten schreiben
	for _, ip := range ipList {
		record := []string{
			ip.IPAddress,
			ip.DeviceType,
		}

		// Description prüfen
		if ip.Description != nil {
			record = append(record, *ip.Description) // Dereferenzieren, falls vorhanden
		} else {
			record = append(record, "") // Falls nil, leeres Feld einfügen
		}

		// Zeile schreiben
		if err := writer.Write(record); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString("Error writing CSV data")
			return
		}
	}
}

func SMSProjectCheckList(ctx iris.Context) {
	// Projekt-ID aus URL holen
	projectID, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.ViewData("error", "Invalid project ID!")
		ctx.View("sms_projectChecks.html")
		return
	}

	// Check-Typ aus der URL holen (Standard: "all", falls keiner übergeben wurde)
	checkType := ctx.Params().GetStringDefault("check_type", "all")

	// Checkliste für das Projekt abrufen (mit Filter auf Check-Typ)
	checkList, err := dbprovider.GetDBManager().GetChecksForProject(projectID, checkType)
	if err != nil {
		ctx.ViewData("error", "Error retrieving checks for project!")
		ctx.View("sms_projectChecks.html")
		return
	}

	// Projekt-ID, Checkliste und den gewählten Check-Typ an die View übergeben
	ctx.ViewData("projectID", projectID)
	ctx.ViewData("projectChecks", checkList)
	ctx.ViewData("selectedCheckType", checkType) // Falls du das für die UI brauchst
	ctx.View("sms_projectChecks.html")
}


func SMSExportProjectStructureCSV(ctx iris.Context) {
	projectIDStr := ctx.Params().Get("project_id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Invalid project ID")
		return
	}

	// Projektstruktur laden
	structure := dbprovider.GetDBManager().GetProjectStructure(projectID)

	// CSV Header
	ctx.ResponseWriter().Header().Set("Content-Type", "text/csv")
	ctx.ResponseWriter().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=project_%d_structure.csv", projectID))

	writer := csv.NewWriter(ctx.ResponseWriter())
	defer writer.Flush()

	// Spalten-Header
	header := []string{
		"DeviceSerial", "DeviceType", "DeviceVersion",
		"SoftwareType", "SoftwareVersion",
		"ComponentName", "ComponentVersion",
	}
	if err := writer.Write(header); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("Failed to write CSV header")
		return
	}

	// Daten flatten und schreiben
	for _, device := range structure {
		if len(device.Software) == 0 {
			// Gerät ohne Software, leere Spalte
			record := []string{
				device.SerialNumber,
				device.DeviceType,
				device.DeviceVersion,
				"", "", "", "",
			}
			writer.Write(record)
			continue
		}

		for _, sw := range device.Software {
			if len(sw.Components) == 0 {
				// Software ohne Komponenten
				record := []string{
					device.SerialNumber,
					device.DeviceType,
					device.DeviceVersion,
					sw.Type,
					sw.Version,
					"", "",
				}
				writer.Write(record)
				continue
			}

			for _, comp := range sw.Components {
				record := []string{
					device.SerialNumber,
					device.DeviceType,
					device.DeviceVersion,
					sw.Type,
					sw.Version,
					comp.Name,
					comp.Version,
				}
				writer.Write(record)
			}
		}
	}
}

func SMSExportProjectStructureYAML(ctx iris.Context) {
	projectIDStr := ctx.Params().Get("project_id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Invalid project ID")
		return
	}

	// Datenstruktur abrufen
	structure := dbprovider.GetDBManager().GetProjectStructure(projectID)

	// YAML serialisieren
	yamlData, err := yaml.Marshal(structure)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("Error generating YAML")
		return
	}

	// Header setzen und Datei senden
	ctx.ResponseWriter().Header().Set("Content-Type", "application/x-yaml")
	ctx.ResponseWriter().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=project_%d_structure.yaml", projectID))
	ctx.Write(yamlData)
}
