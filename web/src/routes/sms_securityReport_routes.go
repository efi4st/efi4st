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
	id := ctx.Params().Get("id")
	reportID, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err != nil {
		ctx.ViewData("error", "Error: Invalid report ID!")
	} else {
		report, err := dbprovider.GetDBManager().GetSMSSecurityReportByID(reportID)
		if err != nil {
			ctx.ViewData("error", "Error: Unable to fetch report details!")
		}
		ctx.ViewData("report", report)
	}

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