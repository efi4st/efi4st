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
)

// GET
func CreateSMSSecurityReportLink(ctx iris.Context) {
	// Hole die ID und den Typ des verlinkten Objekts aus den Parametern
	linkedObjectId := ctx.Params().Get("linkedObjectId")
	linkedObjectType := ctx.Params().Get("linkedObjectType")

	// Konvertiere die ID in einen Integer
	objectId, err := strconv.Atoi(linkedObjectId)
	ctx.ViewData("error", "")
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing linkedObjectId!")
	}

	// Hole die Liste der SecurityReports aus der Datenbank
	securityReports, err := dbprovider.GetDBManager().GetAllSMSSecurityReports()

	// Fehlerbehandlung, wenn die ID nicht korrekt geparsed werden kann
	ctx.ViewData("error", "")
	if err != nil {
		ctx.ViewData("error", "Error: Error retrieving security Reports!")
	}

	// Übergebe die Object ID + Type und die Report-Liste an die View
	ctx.ViewData("objectId", objectId)
	ctx.ViewData("objectType", linkedObjectType)
	ctx.ViewData("securityReportList", securityReports)

	// Lade das entsprechende HTML-Template
	ctx.View("sms_createSecurityReportLink.html")
}

// POST
func AddSMSSecurityReportLink(ctx iris.Context) {
	// Hole die Report ID, die Objekt-ID und den Typ aus den POST-Daten
	reportId := ctx.PostValue("ReportId")
	linkedObjectId := ctx.PostValue("LinkedObjectId")
	linkedObjectType := ctx.PostValue("LinkedObjectType")

	// Konvertiere die IDs in Integer
	rId, err := strconv.Atoi(reportId)
	lId, err := strconv.Atoi(linkedObjectId)
	ctx.ViewData("error", "")
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing report or linked object ID!")
	}

	// Füge die Verknüpfung in der Datenbank hinzu
	err = dbprovider.GetDBManager().AddReportLink(rId, lId, linkedObjectType)
	if err != nil {
		ctx.ViewData("error", "Error: Could not add report link!")
	}

	// Erfolgreich hinzugefügt: Redirect zur Objekt-Seite
	ctx.Redirect(fmt.Sprintf("/sms_securityReportLink/linked/%d/%s", lId, linkedObjectType))
}

func RemoveSMSSecurityReportLink(ctx iris.Context) {
	// Hole die Report ID, die Objekt-ID und den Typ aus den Parametern
	reportId := ctx.Params().Get("reportId")
	linkedObjectId := ctx.Params().Get("linkedObjectId")
	linkedObjectType := ctx.Params().Get("linkedObjectType")

	// Konvertiere die IDs in Integer
	rId, err := strconv.Atoi(reportId)
	lId, err := strconv.Atoi(linkedObjectId)
	ctx.ViewData("error", "")
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing report or linked object ID!")
	}

	// Entferne die Verknüpfung in der Datenbank
	err = dbprovider.GetDBManager().RemoveReportLink(rId, lId, linkedObjectType)
	if err != nil {
		ctx.ViewData("error", "Error: Could not remove report link!")
	}

	// Erfolgreich entfernt: Redirect zur Objekt-Seite
	ctx.Redirect(fmt.Sprintf("/sms_securityReportLink/linked/%d/%s", lId, linkedObjectType))
}

func RedirectToLinkedObject(ctx iris.Context) {
	// Extrahiere die Parameter aus der URL
	linkedObjectId := ctx.Params().Get("linkedObjectId")
	linkedObjectType := ctx.Params().Get("linkedObjectType")

	// Debug: Parameter ausgeben
	fmt.Printf("linkedObjectId: %s, linkedObjectType: %s\n", linkedObjectId, linkedObjectType)

	// Fehlerbehandlung für die ID
	lId, err := strconv.Atoi(linkedObjectId)
	if err != nil {
		fmt.Println("Error: linkedObjectId is not a valid integer:", err)
		ctx.ViewData("error", "Error: Invalid linkedObjectId!")
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}

	// Debug: Überprüfe den Typ
	fmt.Printf("Processing linkedObjectType: %s\n", linkedObjectType)

	// Basierend auf dem Typ zu einer spezifischen Seite weiterleiten
	var redirectURL string
	switch linkedObjectType {
	case "sms_device":
		redirectURL = fmt.Sprintf("/sms_devices/show/%d", lId)
	case "sms_software":
		redirectURL = fmt.Sprintf("/sms_softwares/show/%d", lId)
	case "sms_system":
		redirectURL = fmt.Sprintf("/sms_systems/show/%d", lId)
	default:
		// Ungültiger Typ
		fmt.Println("Error: linkedObjectType is invalid")
		ctx.ViewData("error", "Error: Invalid linkedObjectType!")
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}

	// Debug: Finaler Redirect-Link
	fmt.Printf("Redirecting to URL: %s\n", redirectURL)

	// Weiterleitung zur ermittelten URL
	ctx.Redirect(redirectURL)
}