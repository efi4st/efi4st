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
	"io"
	"log"
	"mime"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func SMSSecurityReports(ctx iris.Context) {
	reports, err := dbprovider.GetDBManager().GetAllSMSSecurityReports()
	ctx.ViewData("error", "")

	if err != nil {
		ctx.ViewData("error", "Error: Unable to fetch reports!")
	} else if len(reports) < 1 {
		ctx.ViewData("error", "Error: No reports available. Add one!")
	}

	ctx.ViewData("reportList", reports)
	ctx.View("sms_securityReports.html")
}

// GET
func CreateSMSSecurityReport(ctx iris.Context) {
	ctx.View("sms_createSecurityReport.html")
}

// POST
func AddSMSSecurityReport(ctx iris.Context) {
	reportName := ctx.PostValue("ReportName")
	scannerName := ctx.PostValue("ScannerName")
	scannerVersion := ctx.PostValue("ScannerVersion")
	creationDate := ctx.PostValue("CreationDate")
	uploadedBy := ctx.PostValue("UploadedBy")
	scanScope := ctx.PostValue("ScanScope")
	vulnerabilityCount, errVuln := strconv.Atoi(ctx.PostValue("VulnerabilityCount"))
	componentCount, errComp := strconv.Atoi(ctx.PostValue("ComponentCount"))

	ctx.ViewData("error", "")

	if errVuln != nil || errComp != nil {
		ctx.ViewData("error", "Error: Invalid count values!")
	} else {
		creationTime, errDate := time.Parse("2006-01-02T15:04", creationDate)
		if errDate != nil {
			fmt.Println("Error:", errDate)
			ctx.ViewData("error", "Error: Invalid date format!")
		} else {
			err := dbprovider.GetDBManager().AddSMSSecurityReport(
				reportName, scannerName, scannerVersion, creationTime,
				uploadedBy, scanScope, vulnerabilityCount, componentCount,
			)
			if err != nil {
				ctx.ViewData("error", "Error: Unable to add report!")
			}
		}
	}

	reports, _ := dbprovider.GetDBManager().GetAllSMSSecurityReports()
	ctx.ViewData("reportList", reports)
	ctx.View("sms_securityReports.html")
}

func ShowSMSSecurityReport(ctx iris.Context) {
	// Holen der Report ID aus der URL
	reportID, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.ViewData("error", "Invalid report ID")
		ctx.Redirect("/sms_securityReports")
		return
	}

	// Report aus der DB holen
	report, err := dbprovider.GetDBManager().GetSMSSecurityReportByID(reportID)
	if err != nil {
		ctx.ViewData("error", "Report not found")
		ctx.Redirect("/sms_securityReports")
		return
	}

	// Dateinamen mit der bestehenden Funktion abrufen
	reportFilename, err := dbprovider.GetDBManager().GetReportFilename(reportID)
	if err != nil {
		log.Printf("Error retrieving filename for report ID %d: %v", reportID, err)
		ctx.ViewData("error", "Error retrieving report file")
		reportFilename = "" // Wenn ein Fehler auftritt, bleibt das Feld leer
	}

	// Die neue Datenklasse füllen
	reportDetail := classes.SecurityReportDetail{
		Report:        *report,
		ReportFilename: reportFilename,
	}

	// Setze die View-Daten
	ctx.ViewData("reportDetail", reportDetail)

	// Lade die View für die Report-Daten
	ctx.View("sms_showSecurityReport.html")
}

func RemoveSMSSecurityReport(ctx iris.Context) {
	id := ctx.Params().Get("id")
	reportID, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err != nil {
		ctx.ViewData("error", "Error: Invalid report ID!")
	} else {
		err = dbprovider.GetDBManager().RemoveSMSSecurityReport(reportID)
		if err != nil {
			ctx.ViewData("error", "Error: Unable to remove report!")
		}
	}

	reports, _ := dbprovider.GetDBManager().GetAllSMSSecurityReports()
	ctx.ViewData("reportList", reports)
	ctx.View("sms_securityReports.html")
}

func UploadSecurityReportFile(ctx iris.Context) {
	reportID, err := ctx.URLParamInt("report_id")
	if err != nil {
		log.Printf("Invalid report ID: %v", err)
		ctx.ViewData("error", "Error: Invalid report ID!")
		ctx.Redirect(fmt.Sprintf("/sms_securityReports/show/%d", reportID))
		return
	}

	log.Printf("Starting file upload for report ID %d", reportID)

	// Prüfen, ob bereits eine Datei existiert
	oldFilename, err := dbprovider.GetDBManager().GetReportFilename(reportID)
	if err != nil {
		log.Printf("Error retrieving existing file entry: %v", err)
		ctx.ViewData("error", "Error: Unable to retrieve existing file!")
		ctx.Redirect(fmt.Sprintf("/sms_securityReports/show/%d", reportID))
		return
	}

	// Datei aus Upload abrufen
	file, info, err := ctx.FormFile("report_file")
	if err != nil {
		log.Printf("Failed to get uploaded file: %v", err)
		ctx.ViewData("error", "Error: Failed to get uploaded file!")
		ctx.Redirect(fmt.Sprintf("/sms_securityReports/show/%d", reportID))
		return
	}
	defer file.Close()

	// Neuen Dateinamen mit Timestamp generieren
	timestamp := time.Now().Format("20060102_150405")
	newFilename := fmt.Sprintf("uploads/reports/report_%d_%s_%s", reportID, timestamp, info.Filename)

	log.Printf("New filename will be: %s", newFilename)

	// Falls alte Datei existiert, löschen
	if oldFilename != "" {
		err := os.Remove(oldFilename)
		if err != nil && !os.IsNotExist(err) {
			log.Printf("Failed to delete old file %s: %v", oldFilename, err)
			ctx.ViewData("error", "Error: Unable to remove old report file!")
			ctx.Redirect(fmt.Sprintf("/sms_securityReports/show/%d", reportID))
			return
		}
		log.Printf("Old file %s deleted successfully", oldFilename)
	}

	// Neue Datei speichern
	outFile, err := os.Create(newFilename)
	if err != nil {
		log.Printf("Failed to save new file: %v", err)
		ctx.ViewData("error", "Error: Unable to save new report file!")
		ctx.Redirect(fmt.Sprintf("/sms_securityReports/show/%d", reportID))
		return
	}
	defer outFile.Close()
	_, err = io.Copy(outFile, file)
	if err != nil {
		log.Printf("Failed to write to file: %v", err)
		ctx.ViewData("error", "Error: Unable to write report file!")
		ctx.Redirect(fmt.Sprintf("/sms_securityReports/show/%d", reportID))
		return
	}

	// Datenbankeintrag aktualisieren
	err = dbprovider.GetDBManager().UpdateReportFilename(reportID, newFilename)
	if err != nil {
		log.Printf("Failed to update database: %v", err)
		ctx.ViewData("error", "Error: Failed to update database entry!")
		ctx.Redirect(fmt.Sprintf("/sms_securityReports/show/%d", reportID))
		return
	}

	log.Printf("File upload successful for report ID %d", reportID)
	ctx.Redirect(fmt.Sprintf("/sms_securityReports/show/%d", reportID))
}

func ViewSecurityReport(ctx iris.Context) {
	reportIDStr := ctx.Params().GetString("report_id")
	reportID, err := strconv.Atoi(reportIDStr)
	if err != nil {
		log.Printf("Invalid report ID: %v", err)
		ctx.ViewData("error", "Error: Invalid report ID!")
		ctx.Redirect(fmt.Sprintf("/sms_securityReports/show/%d", reportID))
		return
	}

	log.Printf("Retrieving file for report ID %d", reportID)

	filename, err := dbprovider.GetDBManager().GetReportFilename(reportID)
	if err != nil || filename == "" {
		log.Printf("Report file not found for ID %d", reportID)
		ctx.ViewData("error", "Error: Report file not found!")
		ctx.Redirect(fmt.Sprintf("/sms_securityReports/show/%d", reportID))
		return
	}

	log.Printf("Redirecting to report file: %s", filename)
	ctx.Redirect("/" + filename)
}

func GetSecurityReportFile(ctx iris.Context) {
	requestPath := ctx.Params().Get("file")
	filePath := "./uploads/reports/" + requestPath

	// Erlaubte Dateiendungen
	allowedExtensions := []string{".html", ".htm", ".pdf", ".txt", ".zip"}
	ext := strings.ToLower(filepath.Ext(requestPath))

	allowed := false
	for _, e := range allowedExtensions {
		if ext == e {
			allowed = true
			break
		}
	}

	if !allowed {
		log.Printf("⚠ Blocked file access: %s", requestPath)
		ctx.StatusCode(iris.StatusForbidden)
		ctx.WriteString("Access denied!")
		return
	}

	// Content-Type setzen
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	ctx.ContentType(contentType)

	// HTML-Dateien inline anzeigen, andere als Download
	if ext == ".html" || ext == ".htm" {
		ctx.Header("Content-Disposition", "inline")
	} else {
		ctx.Header("Content-Disposition", "attachment; filename="+requestPath)
	}

	// Datei senden
	ctx.SendFile(filePath, requestPath)
}

func getContentType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".html", ".htm":
		return "text/html"
	case ".pdf":
		return "application/pdf"
	case ".txt":
		return "text/plain"
	default:
		return "application/octet-stream" // Standard: Download
	}
}