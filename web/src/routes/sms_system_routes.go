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
	"time"
)

func SMSSystems(ctx iris.Context) {

	systems := dbprovider.GetDBManager().GetSMSSystems()
	ctx.ViewData("error", "")

	if len(systems) < 1 {
		ctx.ViewData("error", "Error: No projects available. Add one!")
	}
	ctx.ViewData("systemList", systems)
	ctx.View("sms_systems.html")
}

// GET
func CreateSMSSystem(ctx iris.Context) {

	systemTypes := dbprovider.GetDBManager().GetSMSSystemTypes()

	ctx.ViewData("typeList", systemTypes)
	ctx.View("sms_createSystem.html")
}


// POST
func AddSMSSystem(ctx iris.Context) {

	systemtypeId := ctx.PostValue("SystemtypeId")
	version := ctx.PostValue("Version")

	i, err := strconv.Atoi(systemtypeId)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing projecttypeId!")
	}

	err = dbprovider.GetDBManager().AddSMSSystem(i, version, time.Now().String())

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Not able to add project!")
	}
	systems := dbprovider.GetDBManager().GetSMSSystems()
	ctx.ViewData("systemList", systems)
	ctx.View("sms_systems.html")
}

func ShowSMSSystem(ctx iris.Context) {
	id := ctx.Params().Get("id")
	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing system Id!")
		return
	}

	system := dbprovider.GetDBManager().GetSMSSystemInfo(i)
	devicesUnderSystem := dbprovider.GetDBManager().GetSMSDevicePartOfSystemForSystem(i)
	systemManufacturingOrders := dbprovider.GetDBManager().GetSMSManufactoringOrderForSystem(i)
	systemTree := dbprovider.GetDBManager().GetSMSSystemTreeForSystem(i)
	systemHasCertificates, err := dbprovider.GetDBManager().GetCertificationsForSystem(i)
	if err != nil {
		ctx.ViewData("error", "Error: Error matching certificates!")
	}

	deviceIssuesForThisSystem, err := dbprovider.GetDBManager().GetSMSIssuesForSystem(i)
	reportsForThisSystem, err := dbprovider.GetDBManager().GetReportsForLinkedObject(i, "sms_system")

	// 🔽 Artefakte unter dem System abrufen
	artefactsUnderSystem := dbprovider.GetDBManager().GetSMSArtefactPartOfSystemForSystem(i)

	// 🔼 Artefakte als ViewData bereitstellen
	ctx.ViewData("artefactsUnderSystem", artefactsUnderSystem)

	// Bestehende ViewData
	ctx.ViewData("systemTree", systemTree)
	ctx.ViewData("systemManufacturingOrders", systemManufacturingOrders)
	ctx.ViewData("devicesUnderSystem", devicesUnderSystem)
	ctx.ViewData("system", system)
	ctx.ViewData("systemHasCertificates", systemHasCertificates)
	ctx.ViewData("deviceIssuesForThisSystem", deviceIssuesForThisSystem)
	ctx.ViewData("reportsForThisSystem", reportsForThisSystem)

	ctx.View("sms_showSystem.html")
}

func DownloadSystemTreeJSON(ctx iris.Context) {
	// System-ID aus der URL holen
	systemIDStr := ctx.Params().Get("system_id")
	systemID, err := strconv.Atoi(systemIDStr)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Ungültige System-ID")
		return
	}

	// JSON-Daten abrufen
	jsonData, err := dbprovider.GetDBManager().GetSMSSystemTreeAsJSON(systemID)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("Fehler beim Erstellen der JSON-Datei")
		return
	}

	// JSON als Datei zurückgeben
	ctx.ResponseWriter().Header().Set("Content-Type", "application/json")
	ctx.ResponseWriter().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=system_tree_%d.json", systemID))
	ctx.Write(jsonData) // JSON-Daten direkt zurückgeben
}

func RemoveSMSSystem(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveSMSSystem(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing project!")
	}

	systems := dbprovider.GetDBManager().GetSMSSystems()

	ctx.ViewData("systemList", systems)
	ctx.View("sms_systems.html")
}

func ShowSMSReleaseNotesForSystem(ctx iris.Context) {
	id := ctx.Params().Get("id")
	systemID, err := strconv.Atoi(id)

	if err != nil {
		ctx.ViewData("error", "Fehler beim Parsen der System-ID!")
		ctx.View("sms_showSystemReleaseNotes.html")
		return
	}

	// 1. Hole das System-Objekt
	system := dbprovider.GetDBManager().GetSMSSystemTypeForReleaseNotes(systemID)
	if system == nil {
		ctx.ViewData("error", "System nicht gefunden!")
		ctx.View("sms_showSystemReleaseNotes.html")
		return
	}
	fmt.Println(system.Systemtype_id())
	systemtypeIDInt, err := strconv.Atoi(system.Systemtype_id())
	if err != nil {
		ctx.ViewData("error", "Fehler beim Konvertieren des systemtype_id in int")
		ctx.View("error.html")
		return
	}
	releaseNotes := dbprovider.GetDBManager().GetReleaseNotesForSystemUpToVersion(systemtypeIDInt, system.Version())

	// 3. Übergib die Daten an die View
	ctx.ViewData("system", system)
	ctx.ViewData("releaseNotes", releaseNotes)
	ctx.View("sms_showSystemReleaseNotes.html")
}