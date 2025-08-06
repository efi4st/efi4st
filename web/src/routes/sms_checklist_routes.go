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
	"log"
	"net/http"
	"strconv"
)

/////////////
//
// ChecklistTemplates
//
/////////////
func SMSChecklistTemplates(ctx iris.Context) {
	templates := dbprovider.GetDBManager().GetAllChecklistTemplates()
	ctx.ViewData("templateList", templates)
	ctx.View("sms_checklistTemplates.html")
}

func CreateSMSChecklistTemplate(ctx iris.Context) {
	ctx.View("sms_createChecklistTemplate.html")
}

func AddSMSChecklistTemplate(ctx iris.Context) {
	name := ctx.PostValue("Name")
	description := ctx.PostValue("Description")

	err := dbprovider.GetDBManager().AddChecklistTemplate(&classes.Sms_ChecklistTemplate{
		Name: name,
		Description: description,
	})
	if err != nil {
		ctx.ViewData("error", "Error: Could not create checklist template")
	}
	templates := dbprovider.GetDBManager().GetAllChecklistTemplates()
	ctx.ViewData("templateList", templates)
	ctx.View("sms_checklistTemplates.html")
}

func ShowSMSChecklistTemplate(ctx iris.Context) {
	idStr := ctx.Params().Get("id")
	id, _ := strconv.Atoi(idStr)

	template := dbprovider.GetDBManager().GetChecklistTemplateByID(id)
	items := dbprovider.GetDBManager().GetChecklistTemplateItems(id)
	artefactTypes := dbprovider.GetDBManager().GetSMSArtefactTypes()

	checkDefs, err := dbprovider.GetDBManager().GetAllDeviceCheckDefinitions()
	if err != nil {
		log.Println("Fehler beim Abrufen der Check Definitions:", err)
	}

	ctx.ViewData("template", template)
	ctx.ViewData("items", items)
	ctx.ViewData("artefactTypes", artefactTypes)
	ctx.ViewData("checkDefinitions", checkDefs)
	ctx.View("sms_showChecklistTemplate.html")
}

func RemoveSMSChecklistTemplate(ctx iris.Context) {
	idStr := ctx.Params().Get("id")
	id, _ := strconv.Atoi(idStr)

	_ = dbprovider.GetDBManager().DeleteChecklistTemplateByID(id)

	templates := dbprovider.GetDBManager().GetAllChecklistTemplates()
	ctx.ViewData("templateList", templates)
	ctx.View("sms_checklistTemplates.html")
}

/////////////
//
// ChecklistTemplateItem
//
/////////////
func AddSMSChecklistTemplateItem(ctx iris.Context) {
	templateID, _ := strconv.Atoi(ctx.PostValue("ChecklistTemplateID"))
	checkDefinitionID := ctx.PostValueIntDefault("CheckDefinitionID", 0)
	artefactTypeID := ctx.PostValueIntDefault("ArtefactTypeID", 0)
	scope := ctx.PostValue("TargetScope")
	expected := ctx.PostValue("ExpectedValue")
	optional := ctx.PostValue("Optional") == "on"

	var checkIDPtr *int = nil
	if checkDefinitionID != 0 {
		checkIDPtr = &checkDefinitionID
	}
	var artefactIDPtr *int = nil
	if artefactTypeID != 0 {
		artefactIDPtr = &artefactTypeID
	}

	item := &classes.Sms_ChecklistTemplateItem{
		ChecklistTemplateID: templateID,
		CheckDefinitionID:   checkIDPtr,
		ArtefactTypeID:      artefactIDPtr,
		TargetScope:         scope,
		ExpectedValue:       expected,
		Optional:            optional,
	}

	err := dbprovider.GetDBManager().AddChecklistTemplateItem(item)
	if err != nil {
		log.Println("Fehler beim Hinzufügen des Checklist-Items:", err)
		ctx.ViewData("error", "Fehler beim Hinzufügen des Checklist-Items: "+err.Error())

		// Wiederlade die Seite mit Template + Items zur Anzeige des Fehlers
		template := dbprovider.GetDBManager().GetChecklistTemplateByID(templateID)
		items := dbprovider.GetDBManager().GetChecklistTemplateItems(templateID)

		ctx.ViewData("template", template)
		ctx.ViewData("items", items)
		ctx.View("sms_showChecklistTemplate.html")
		return
	}

	// Wenn kein Fehler: Redirect zur Übersicht
	http.Redirect(ctx.ResponseWriter(), ctx.Request(), fmt.Sprintf("/sms_checklistTemplate/show/%d", templateID), http.StatusFound)
}

func RemoveSMSChecklistTemplateItem(ctx iris.Context) {
	id, _ := strconv.Atoi(ctx.Params().Get("id"))
	templateID, _ := strconv.Atoi(ctx.URLParam("template_id"))
	_ = dbprovider.GetDBManager().DeleteChecklistTemplateItemByID(id)
	ctx.Redirect(fmt.Sprintf("/sms_checklistTemplate/show/%d", templateID))
}

///////////////
//
//// ChecklistInstance
//
//////////////

func GenerateChecklistInstanceForProject(ctx iris.Context) {
	projectID, _ := strconv.Atoi(ctx.Params().Get("project_id"))
	templateID, _ := strconv.Atoi(ctx.PostValue("ChecklistTemplateID"))

	instance := &classes.Sms_ChecklistInstance{
		ProjectID:           &projectID,
		DeviceID:            nil,
		ChecklistTemplateID: templateID,
		GeneratedBy:         "system", // ← Beispielwert
		Status:              "open",
	}

	err := dbprovider.GetDBManager().AddChecklistInstance(instance)
	if err != nil {
		log.Printf("❌ Fehler beim Instanzieren der Checklist (Template %d, Project %d): %v", templateID, projectID, err)
		ctx.ViewData("error", fmt.Sprintf("Fehler beim Erzeugen der Checkliste: %v", err))

		// Lade Projekt erneut, damit wir zur Ursprungsseite zurückkehren können
		project := dbprovider.GetDBManager().GetSMSProjectInfo(projectID)
		ctx.ViewData("project", project)

		// Optional: zeige Fehlermeldung direkt in Projekt-Ansicht oder eigene Error-View
		ctx.View("sms_showProject.html")
		return
	}

	ctx.Redirect(fmt.Sprintf("/sms_projects/show/%d", projectID))
}

func GenerateChecklistInstanceForDevice(ctx iris.Context) {
	deviceID, _ := strconv.Atoi(ctx.Params().Get("device_id"))
	templateID, _ := strconv.Atoi(ctx.PostValue("ChecklistTemplateID"))

	err := dbprovider.GetDBManager().AddChecklistInstance(&classes.Sms_ChecklistInstance{
		ProjectID:           nil,
		DeviceID:            &deviceID,
		ChecklistTemplateID: templateID,
		Status:              "open",
	})
	if err != nil {
		ctx.ViewData("error", "Failed to create checklist instance")
	}
	ctx.Redirect(fmt.Sprintf("/sms_devices/show/%d", deviceID))
}

func ShowChecklistInstance(ctx iris.Context) {
	id, _ := strconv.Atoi(ctx.Params().Get("id"))
	instance := dbprovider.GetDBManager().GetChecklistInstanceByID(id)
	items := dbprovider.GetDBManager().GetChecklistItemInstances(id)

	ctx.ViewData("instance", instance)
	ctx.ViewData("items", items)
	ctx.View("sms_showChecklistInstance.html")
}

func DeleteChecklistInstance(ctx iris.Context) {
	id, _ := strconv.Atoi(ctx.Params().Get("id"))
	dbprovider.GetDBManager().DeleteChecklistInstanceByID(id)
	ctx.Redirect("/sms_checklistTemplates")
}

func MarkChecklistInstanceStatus(ctx iris.Context) {
	id, _ := strconv.Atoi(ctx.Params().Get("id"))
	status := ctx.URLParam("status")
	_ = dbprovider.GetDBManager().UpdateChecklistInstanceStatus(id, status)
	ctx.Redirect(fmt.Sprintf("/sms_checklistInstance/show/%d", id))
}


func UpdateChecklistItemInstance(ctx iris.Context) {
	itemID, _ := strconv.Atoi(ctx.PostValue("ChecklistItemInstanceID"))
	checklistInstanceID, _ := strconv.Atoi(ctx.PostValue("ChecklistInstanceID"))
	isOKStr := ctx.PostValue("IsOK")
	actualValue := ctx.PostValue("ActualValue")
	comment := ctx.PostValue("Comment")

	var isOK *bool = nil
	if isOKStr != "" {
		val := isOKStr == "true" || isOKStr == "on"
		isOK = &val
	}

	item := &classes.Sms_ChecklistItemInstance{
		ChecklistItemInstanceID: itemID,
		IsOK:                    isOK,
		ActualValue:             actualValue,
		Comment:                 comment,
	}

	err := dbprovider.GetDBManager().UpdateChecklistItemInstance(item)
	if err != nil {
		ctx.ViewData("error", "Failed to update item")
	}

	ctx.Redirect(fmt.Sprintf("/sms_checklistInstance/show/%d", checklistInstanceID))
}