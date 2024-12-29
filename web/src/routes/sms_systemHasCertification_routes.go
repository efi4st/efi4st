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
	"strings"
)

// GET
func CreateSMSSystemHasCertification(ctx iris.Context) {
	// Hole die System-ID aus den Routenparametern
	id := ctx.Params().Get("id")
	systemID, err := strconv.Atoi(id)

	// Fehlerbehandlung bei der Umwandlung der ID
	ctx.ViewData("error", "")
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing system ID!")
	}

	// Hole die Liste der verfügbaren Zertifizierungen aus der Datenbank
	certifications := dbprovider.GetDBManager().GetSMSCertification()

	// Übergib die System-ID und die Liste der Zertifizierungen an die View
	ctx.ViewData("systemId", systemID)
	ctx.ViewData("certificationList", certifications)

	// Rendere die View
	ctx.View("sms_createSystemHasCertification.html")
}

// POST
func AddSMSSystemHasCertification(ctx iris.Context) {
	// Hole die Certification-IDs und die System-ID aus dem POST-Request
	certification_ids := ctx.PostValue("Certification_ids")
	system_id := ctx.PostValue("System_id")
	additionalInfo := ctx.PostValue("AdditionalInfo")

	// Konvertiere die System-ID von String zu Integer
	systemID, err := strconv.Atoi(system_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing system_id!")
		ctx.Redirect("/systems") // Fallback
		return
	}

	// Iteriere durch die Certification-IDs
	for _, certificationId := range strings.Split(certification_ids, ",") {
		certificationID, err := strconv.Atoi(certificationId)
		if err != nil {
			ctx.ViewData("error", "Error: Error parsing certificationId!")
			continue
		}

		// Füge die Verknüpfung in die Datenbank ein
		err = dbprovider.GetDBManager().AddSystemHasCertification(systemID, certificationID, additionalInfo)
		if err != nil {
			ctx.ViewData("error", "Error: Not able to add system-certification link!")
			continue
		}
	}

	// Erfolgreiche Verarbeitung: Zeige die System-Detailansicht erneut
	ctx.Params().Set("id", system_id)
	ShowSMSSystem(ctx) // Funktion, die die System-Details erneut rendert
}

func RemoveSMSSystemHasCertification(ctx iris.Context) {
	// Hole die System-ID und die Certification-ID aus den Routenparametern
	systemIDParam := ctx.Params().Get("systemId")
	certificationIDParam := ctx.Params().Get("certificationId")

	// Konvertiere die IDs von String zu Integer
	systemID, err := strconv.Atoi(systemIDParam)
	certificationID, certErr := strconv.Atoi(certificationIDParam)

	// Fehlerbehandlung bei der Konvertierung
	ctx.ViewData("error", "")
	if err != nil || certErr != nil {
		ctx.ViewData("error", "Error: Error parsing system or certification ID!")
	} else {
		// Lösche die Verknüpfung in der Datenbank
		err = dbprovider.GetDBManager().RemoveSystemHasCertification(systemID, certificationID)

		// Falls ein Fehler auftritt, Fehler in die View-Daten setzen
		if err != nil {
			ctx.ViewData("error", "Error: Error removing system certification link!")
		}
	}

	// Leite zurück zur vorherigen Seite (Referer-Header)
	referer := ctx.Request().Header.Get("Referer")
	if referer != "" {
		ctx.Redirect(referer)
	} else {
		// Fallback: Zur Übersicht der Systeme
		SMSProjects(ctx)
	}
}