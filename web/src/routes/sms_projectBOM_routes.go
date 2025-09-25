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
	"strconv"
)

// GET
func CreateSMSProjectBOMForProject(ctx iris.Context) {
	// project_id aus path ODER query
	pidStr := ctx.Params().Get("project_id")
	if pidStr == "" {
		pidStr = ctx.URLParam("project_id")
	}
	projectID, err := strconv.Atoi(pidStr)
	if err != nil || projectID <= 0 {
		ctx.ViewData("error", "Invalid project_id")
		ctx.Redirect("/sms_projects")
		return
	}

	// Systeme für Auswahl
	systemList := dbprovider.GetDBManager().GetAllSystemsMinimal()
	ctx.ViewData("systemList", systemList)
	ctx.ViewData("projectId", projectID)

	// optional vorgewähltes System
	selectedSystemID, _ := strconv.Atoi(ctx.URLParamDefault("system_id", "0"))
	if selectedSystemID > 0 {
		designs := dbprovider.GetDBManager().GetCompatibleHardwareDesignsForSystem(selectedSystemID)
		// Logging hilft beim Debuggen
		fmt.Printf("[PBOM] selectedSystemID=%d, compatible designs=%d\n", selectedSystemID, len(designs))

		selectedDesignID := 0
		if def := dbprovider.GetDBManager().GetDefaultHardwareDesignForSystem(selectedSystemID); def != nil {
			selectedDesignID = def.HardwareDesignID
		} else if len(designs) > 0 {
			selectedDesignID = designs[0].HardwareDesignID
		}

		variants := []classes.Sms_HardwareDesignVariant{}
		if selectedDesignID != 0 {
			variants = dbprovider.GetDBManager().GetVariantsForHardwareDesign(selectedDesignID, true)
		}

		ctx.ViewData("selectedSystemID", selectedSystemID)
		ctx.ViewData("designs", designs)
		ctx.ViewData("selectedDesignID", selectedDesignID)
		ctx.ViewData("variants", variants)
	}

	ctx.View("sms_createProjectBOMForProject.html")
}

// POST: Insert
func AddSMSProjectBOM(ctx iris.Context) {
	projectID, _ := strconv.Atoi(ctx.PostValue("Project_id"))
	systemID,  _ := strconv.Atoi(ctx.PostValue("System_id"))
	designID,  _ := strconv.Atoi(ctx.PostValueDefault("HardwareDesignID", "0"))
	variantID, _ := strconv.Atoi(ctx.PostValueDefault("HardwareDesignVariantID", "0"))
	orderNumber    := ctx.PostValue("OrderNumber")
	additionalInfo := ctx.PostValue("AdditionalInfo")

	log.Printf("[PBOM] POST project=%d system=%d design=%d variant=%d order=%q info=%q",
		projectID, systemID, designID, variantID, orderNumber, additionalInfo)

	// einfache Validierung
	if projectID == 0 || systemID == 0 || designID == 0 || variantID == 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Bitte alle Dropdowns auswählen (Project/System/Design/Variante).")
		return
	}

	if err := dbprovider.GetDBManager().AddSMSProjectBOM(projectID, systemID, designID, variantID, orderNumber, additionalInfo); err != nil {
		// **Fehler nicht weg-redirecten**, sondern zeigen + kurz Diagnose-Hinweis
		log.Printf("[PBOM] insert error: %v", err)
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Insert failed: " + err.Error() +
			"\nHinweis: Existiert ein Mapping in sms_hardwaredesignPartOfSystem für (system_id, hardwaredesign_id)? " +
			"Und gehört die Variante zum Design?")
		return
	}

	// Erfolg → zurück zur Projektseite
	ctx.Redirect(fmt.Sprintf("/sms_projects/show/%d", projectID))
}

// JSON: aktive Varianten zu einem Design
func GetVariantsForDesignJSON(ctx iris.Context) {
	designID, _ := strconv.Atoi(ctx.URLParam("hardwaredesign_id"))
	variants := dbprovider.GetDBManager().GetVariantsForHardwareDesign(designID, true)
	type item struct {
		ID   int    `json:"id"`
		Code string `json:"code"`
		Name string `json:"name"`
	}
	resp := make([]item, 0, len(variants))
	for _, v := range variants {
		resp = append(resp, item{ID: v.HardwareDesignVariantID, Code: v.Code, Name: v.Name})
	}
	_ = ctx.JSON(resp)
}

func DeleteSMSProjectBOM(ctx iris.Context) {
	idStr := ctx.Params().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.WriteString("invalid projectBOM id")
		return
	}

	// Projekt-ID für Redirect ermitteln
	rec := dbprovider.GetDBManager().GetProjectBOMByID(id)
	projectID := 0
	if rec != nil {
		projectID = rec.ProjectID
	}

	if err := dbprovider.GetDBManager().RemoveSMSProjectBOM(id); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.WriteString("delete failed: " + err.Error())
		return
	}

	// zurück zum Projekt (oder Fallback)
	if projectID > 0 {
		ctx.Redirect(fmt.Sprintf("/sms_projects/show/%d", projectID))
	} else {
		ctx.Redirect("/sms_projects")
	}
}