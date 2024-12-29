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
func CreateSMSIssueAffectedArtefact(ctx iris.Context) {
	// Hole die Issue ID aus den Parametern
	id := ctx.Params().Get("id")
	i, err := strconv.Atoi(id)

	// Hole die Liste der Artefacts aus der Datenbank
	artefacts := dbprovider.GetDBManager().GetSMSArtefact()

	// Fehlerbehandlung, wenn die ID nicht korrekt geparsed werden kann
	ctx.ViewData("error", "")
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing issue id!")
	}

	// Übergebe die Issue ID und die Artefact-Liste an die View
	ctx.ViewData("issueId", i)
	ctx.ViewData("artefactList", artefacts)

	// Lade das entsprechende HTML-Template
	ctx.View("sms_createIssueAffectedArtefact.html")
}

// POST
func AddSMSIssueAffectedArtefact(ctx iris.Context) {

	// Hole die Artefakt IDs aus den POST-Daten
	artefact_ids, err := ctx.PostValueMany("Artefact_ids")
	issue_id := ctx.PostValue("Issue_id")
	additionalInfo := ctx.PostValue("AdditionalInfo")
	//confirmed := ctx.PostValue("Confirmed")

	// Konvertiere die Issue ID in eine Zahl
	iI, err := strconv.Atoi(issue_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing issue_id!")
	}

	// Gehe jede Artefakt ID durch und füge sie der Issue hinzu
	for index, artefactId := range strings.Split(artefact_ids, ",") {
		iD, err := strconv.Atoi(artefactId)
		if err != nil {
			ctx.ViewData("error", "Error: Error parsing artefactId!")
		}
		// Füge die Issue-Artefakt-Verknüpfung hinzu
		err = dbprovider.GetDBManager().AddSMSIssueAffectedArtefact(iD, iI, additionalInfo, true)
		fmt.Printf(strconv.Itoa(index))
		ctx.ViewData("error", "")
		if err != nil {
			ctx.ViewData("error", "Error: Not able to add issue-artefact link!")
		}
	}

	// Erfolgreicher Abschluss: Redirect zu ShowSMSIssue
	ctx.Redirect(fmt.Sprintf("/sms_issues/show/%d", iI))
}

func RemoveSMSIssueAffectedArtefact(ctx iris.Context) {
	// Hole sowohl die Issue ID als auch die Artefact ID aus den Parametern
	issueId := ctx.Params().Get("issueId")
	artefactId := ctx.Params().Get("artefactId")

	// Konvertiere die IDs in Integer
	issue, err := strconv.Atoi(issueId)
	artefact, err := strconv.Atoi(artefactId)

	// Fehlerbehandlung
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing issue or artefact id!")
	}

	// Entferne die Issue-Artefact-Verknüpfung
	err = dbprovider.GetDBManager().RemoveSMSIssueAffectedArtefact(artefact, issue)

	// Fehlerbehandlung bei der Entfernung
	ctx.ViewData("error", "")
	if err != nil {
		ctx.ViewData("error", "Error: Error removing issue-artefact link!")
	}

	// Immer zurück zur Issue-Seite leiten
	ctx.Redirect(fmt.Sprintf("/sms_issues/show/%d", issue))
}