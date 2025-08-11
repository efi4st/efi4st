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
	"strings"
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

	inst := dbprovider.GetDBManager().GetChecklistInstanceByID(id)
	if inst == nil {
		ctx.ViewData("error", "Checklist instance not found")
		ctx.View("sms_checklistTemplates.html")
		return
	}

	items := dbprovider.GetDBManager().GetChecklistItemInstancesWithDefinition(id)

	// Projekt-Kontext: passende DeviceInstances listen
	if inst.ProjectID != nil {
		for i := range items {
			if items[i].CheckDefinitionID != nil && items[i].DeviceTypeID != nil {
				candidates := dbprovider.GetDBManager().
					GetDeviceInstancesForProjectAndDeviceType(*inst.ProjectID, *items[i].DeviceTypeID)

				var matches []classes.MatchingDevice
				for _, c := range candidates {
					if matchesApplicable(c.DeviceVersion, items[i].ApplicableVersions) {
						matches = append(matches, c)
					}
				}
				items[i].MatchingDevices = matches
			}
			// ▼ Für Projekt-Kontext keine Device-Anwendbarkeit: explizit "none"
			items[i].AppliesToThisDevice = nil
			items[i].AppliesToThisDeviceStr = "none"
		}
	}

	// Device-Kontext: diesen Device-Typ + Version prüfen
	if inst.DeviceID != nil {
		dbasic, err := dbprovider.GetDBManager().GetDeviceBasicByID(*inst.DeviceID)
		if err == nil && dbasic != nil {
			for i := range items {
				items[i].DeviceContextTypeName = dbasic.DeviceType
				items[i].DeviceContextVersion  = dbasic.Version

				if items[i].CheckDefinitionID != nil {
					var typeOK bool
					if items[i].DeviceTypeID != nil {
						typeOK = (*items[i].DeviceTypeID == dbasic.DeviceTypeID)
					} else if items[i].DeviceTypeName != "" {
						typeOK = strings.EqualFold(items[i].DeviceTypeName, dbasic.DeviceType)
					} else {
						typeOK = false
					}
					versOK := matchesApplicable(dbasic.Version, items[i].ApplicableVersions)

					res := typeOK && versOK
					items[i].AppliesToThisDevice = &res
					// ▼ NEU: String-Repräsentation setzen
					if res {
						items[i].AppliesToThisDeviceStr = "true"
					} else {
						items[i].AppliesToThisDeviceStr = "false"
					}

					log.Printf("Check apply? inst=%d item=%d typeOK=%v (need %v/%s have %d/%s) versOK=%v (need '%s' have '%s')",
						inst.ChecklistInstanceID, items[i].ChecklistItemInstanceID,
						typeOK,
						valueOr(items[i].DeviceTypeID, 0), items[i].DeviceTypeName,
						dbasic.DeviceTypeID, dbasic.DeviceType,
						versOK,
						items[i].ApplicableVersions, dbasic.Version,
					)
				} else {
					// ▼ NEU: explizit „none“ setzen, wenn keine Definition
					items[i].AppliesToThisDevice = nil
					items[i].AppliesToThisDeviceStr = "none"
				}
			}
		}
	}

	ctx.ViewData("instance", inst)
	ctx.ViewData("items", items)
	ctx.View("sms_showChecklistInstance.html")
}

// kleiner helper:
func valueOr(p *int, def int) int {
	if p == nil { return def }
	return *p
}

// 'all' oder CSV-Liste exakter Versionsstrings (einfacher Ansatz)
func matchesApplicable(version, applicable string) bool {
	if applicable == "" || strings.EqualFold(applicable, "all") {
		return true
	}
	for _, tok := range strings.Split(applicable, ",") {
		if strings.EqualFold(strings.TrimSpace(tok), strings.TrimSpace(version)) {
			return true
		}
	}
	return false
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

// POST: System-Checkliste instanziieren
func GenerateChecklistInstanceForSystem(ctx iris.Context) {
	systemID, _ := strconv.Atoi(ctx.Params().Get("system_id"))
	templateID, _ := strconv.Atoi(ctx.PostValue("ChecklistTemplateID"))

	_, err := dbprovider.GetDBManager().AddChecklistInstanceForSystem(templateID, systemID, "system")
	if err != nil {
		log.Printf("❌ System-ChecklistInstance failed (sys=%d, tmpl=%d): %v", systemID, templateID, err)
		ctx.ViewData("error", fmt.Sprintf("Fehler beim Erzeugen der System-Checkliste: %v", err))
		// Zur System-Seite zurück
		ctx.Redirect(fmt.Sprintf("/sms_systems/show/%d", systemID))
		return
	}
	ctx.Redirect(fmt.Sprintf("/sms_systems/show/%d", systemID))
}