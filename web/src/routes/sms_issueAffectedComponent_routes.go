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
func CreateSMSIssueAffectedComponent(ctx iris.Context) {
	// Hole die Issue ID aus den Parametern
	id := ctx.Params().Get("id")
	i, err := strconv.Atoi(id)

	// Hole die Liste der Components aus der Datenbank
	components := dbprovider.GetDBManager().GetSMSComponent()

	// Fehlerbehandlung, wenn die ID nicht korrekt geparsed werden kann
	ctx.ViewData("error", "")
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing issue id!")
	}

	// Übergebe die Issue ID und die Component-Liste an die View
	ctx.ViewData("issueId", i)
	ctx.ViewData("componentList", components)

	// Lade das entsprechende HTML-Template
	ctx.View("sms_createIssueAffectedComponent.html")
}

// POST
func AddSMSIssueAffectedComponent(ctx iris.Context) {
	// Hole die Component IDs aus den POST-Daten
	component_ids, err := ctx.PostValueMany("Component_ids")
	issue_id := ctx.PostValue("Issue_id")
	additionalInfo := ctx.PostValue("AdditionalInfo")

	// Konvertiere die Issue ID in eine Zahl
	iI, err := strconv.Atoi(issue_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing issue_id!")
		ctx.Redirect(fmt.Sprintf("/sms_issueAffectedComponent/create/%d", iI))
		return
	}

	// Gehe jede Component ID durch und füge sie der Issue hinzu
	for _, componentId := range strings.Split(component_ids, ",") {
		iD, err := strconv.Atoi(componentId)
		if err != nil {
			ctx.ViewData("error", "Error: Error parsing componentId!")
			ctx.Redirect(fmt.Sprintf("/sms_issueAffectedComponent/create/%d", iI))
			return
		}
		// Füge die Issue-Component-Verknüpfung hinzu
		err = dbprovider.GetDBManager().AddSMSIssueAffectedComponent(iD, iI, additionalInfo, true)
		if err != nil {
			ctx.ViewData("error", "Error: Not able to add issue-component link!")
			ctx.Redirect(fmt.Sprintf("/sms_issueAffectedComponent/create/%d", iI))
			return
		}
	}

	// Erfolgreicher Abschluss: Redirect zu ShowSMSIssue
	ctx.Redirect(fmt.Sprintf("/sms_issues/show/%d", iI))
}

func RemoveSMSIssueAffectedComponent(ctx iris.Context) {
	// Hole sowohl die Issue ID als auch die Component ID aus den Parametern
	issueId := ctx.Params().Get("issueId")
	componentId := ctx.Params().Get("componentId")

	// Konvertiere die IDs in Integer
	issue, err := strconv.Atoi(issueId)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing issue id!")
		ctx.Redirect(fmt.Sprintf("/sms_issues/show/%d", issue))
		return
	}

	component, err := strconv.Atoi(componentId)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing component id!")
		ctx.Redirect(fmt.Sprintf("/sms_issues/show/%d", issue))
		return
	}

	// Entferne die Issue-Component-Verknüpfung
	err = dbprovider.GetDBManager().RemoveSMSIssueAffectedComponent(component, issue)

	// Fehlerbehandlung bei der Entfernung
	if err != nil {
		ctx.ViewData("error", "Error: Error removing issue-component link!")
	}

	// Immer zurück zur Issue-Seite leiten
	ctx.Redirect(fmt.Sprintf("/sms_issues/show/%d", issue))
}