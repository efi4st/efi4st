/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (
	"encoding/base64"
	"fmt"
	"github.com/efi4st/efi4st/classes"
	"github.com/efi4st/efi4st/dbprovider"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"io"
	"strconv"
)

func SMSHardwareDesigns(ctx iris.Context) {
	designs := dbprovider.GetDBManager().GetAllSMSHardwareDesigns()
	ctx.ViewData("error", "")

	if len(designs) < 1 {
		ctx.ViewData("error", "Error: No hardware designs available. Add one!")
	}
	ctx.ViewData("hardwaredesignList", designs)
	ctx.View("sms_hardwareDesigns.html") // LIST VIEW
}

func CreateSMSHardwareDesign(ctx iris.Context) {
	ctx.ViewData("error", "")
	ctx.View("sms_createHardwareDesign.html") // CREATE FORM
}

func AddSMSHardwareDesign(ctx iris.Context) {
	name := ctx.PostValue("Name")
	version := ctx.PostValue("Version")
	date := ctx.PostValue("Date") // Format: YYYY-MM-DD
	description := ctx.PostValue("Description")
	author := ctx.PostValue("Author")
	revisionNote := ctx.PostValue("RevisionNote")
	documentNumber := ctx.PostValue("DocumentNumber")
	isApproved := ctx.PostValue("IsApproved") == "on"

	// Datei als []byte lesen (optional)
	_, fileHeader, err := ctx.FormFile("Image")
	var imageBytes []byte

	if err == nil && fileHeader != nil {
		file, err := fileHeader.Open()
		if err == nil {
			defer file.Close()
			imageBytes, _ = io.ReadAll(file)
		}
	}

	fmt.Println("Bildgröße (Bytes):", len(imageBytes))


	design := &classes.Sms_HardwareDesign{
		Name:           name,
		Version:        version,
		Date:           date,
		Description:    description,
		Author:         author,
		RevisionNote:   revisionNote,
		DocumentNumber: documentNumber,
		IsApproved:     isApproved,
		Image:          imageBytes,
	}

	err = dbprovider.GetDBManager().AddSMSHardwareDesign(design)
	if err != nil {
		ctx.ViewData("error", "Error: Could not save hardware design!")
	} else {
		ctx.ViewData("error", "")
	}

	// Liste neu laden
	designs := dbprovider.GetDBManager().GetAllSMSHardwareDesigns()
	ctx.ViewData("hardwaredesignList", designs)
	ctx.View("sms_hardwareDesigns.html")
}

func ShowSMSHardwareDesign(ctx iris.Context) {
	id := ctx.Params().Get("id")
	i, err := strconv.Atoi(id)
	ctx.ViewData("error", "")

	if err != nil {
		ctx.ViewData("error", "Error: ID parsing failed!")
		ctx.View("sms_hardwaredesigns.html")
		return
	}

	design := dbprovider.GetDBManager().GetSMSHardwareDesignByID(i)
	if design == nil {
		ctx.ViewData("error", "Error: Design not found!")
		ctx.View("sms_hardwaredesigns.html")
		return
	}

	// ⚠️ Falls Bild vorhanden: In Base64 konvertieren
	if len(design.Image) > 0 {
		design.ImageBase64 = base64.StdEncoding.EncodeToString(design.Image)
	}

	// ➕ Varianten-Vorschau laden (nur aktive), auf max. 5 begrenzen
	variants := dbprovider.GetDBManager().GetVariantsForHardwareDesign(i, true) // onlyActive=true
	if len(variants) > 5 {
		variants = variants[:5]
	}
	ctx.ViewData("variantsPreview", variants)

	ctx.ViewData("hardwareDesign", design)
	ctx.View("sms_showHardwareDesign.html")
}

func RemoveSMSHardwareDesign(ctx iris.Context) {
	id := ctx.Params().Get("id")
	i, err := strconv.Atoi(id)
	ctx.ViewData("error", "")

	if err == nil {
		// Vorher sicherheitshalber Mapping löschen
		_ = dbprovider.GetDBManager().DeleteSMSHardwareDesignMappingsByDesignID(i)
		err = dbprovider.GetDBManager().DeleteSMSHardwareDesignByID(i)
	}

	if err != nil {
		ctx.ViewData("error", "Error: Could not delete design!")
	}

	designs := dbprovider.GetDBManager().GetAllSMSHardwareDesigns()
	ctx.ViewData("hardwaredesignList", designs)
	ctx.View("sms_hardwareDesigns.html")
}

func CreateSMSHardwareDesignPartOfSystem(ctx iris.Context) {
	idParam := ctx.Params().Get("id")
	systemID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.ViewData("error", "Invalid system ID")
		ctx.View("sms_hardwaredesigns.html")
		return
	}
	designs := dbprovider.GetDBManager().GetAllSMSHardwareDesigns()
	ctx.ViewData("systemID", systemID)
	ctx.ViewData("hardwareDesignList", designs)
	ctx.View("sms_createHardwareDesignPartOfSystem.html")
}

func AddSMSHardwareDesignPartOfSystem(ctx iris.Context) {
	systemIDStr := ctx.PostValue("SystemID")
	designIDStr := ctx.PostValue("HardwareDesignID")
	additionalInfo := ctx.PostValue("AdditionalInfo")
	isDefault := ctx.PostValue("IsDefault") == "on"

	// 'recommended' | 'compatible' | 'deprecated'
	compatStatus := ctx.PostValue("CompatibilityStatus")
	switch compatStatus {
	case "recommended", "compatible", "deprecated":
	default:
		compatStatus = "compatible"
	}
	// Default erzwingt recommended
	if isDefault {
		compatStatus = "recommended"
	}

	systemID, err1 := strconv.Atoi(systemIDStr)
	designID, err2 := strconv.Atoi(designIDStr)
	if err1 != nil || err2 != nil {
		ctx.ViewData("error", "Invalid input values")
		ctx.Redirect("/sms_systems/show/" + systemIDStr)
		return
	}

	err := dbprovider.GetDBManager().AddSMSHardwareDesignMappingWithFlags(systemID, designID, additionalInfo, isDefault, compatStatus)
	if err != nil {
		ctx.ViewData("error", "Failed to link hardware design to system: "+err.Error())
	}

	ctx.Redirect("/sms_systems/show/" + systemIDStr)
}

// Handler
func SetDefaultHardwareDesign(ctx iris.Context) {
	systemIDStr := ctx.URLParam("system_id")
	designIDStr := ctx.URLParam("hardwaredesign_id")

	systemID, err1 := strconv.Atoi(systemIDStr)
	designID, err2 := strconv.Atoi(designIDStr)
	if err1 != nil || err2 != nil {
		ctx.ViewData("error", "Invalid parameters")
		ctx.Redirect("/sms_systems/show/" + systemIDStr)
		return
	}

	if err := dbprovider.GetDBManager().SetDefaultHardwareDesign(systemID, designID); err != nil {
		ctx.ViewData("error", "Failed to set default: "+err.Error())
	}

	ctx.Redirect("/sms_systems/show/" + systemIDStr)
}

func RemoveSMSHardwareDesignPartOfSystem(ctx iris.Context) {
	systemIDStr := ctx.URLParam("system_id")
	designIDStr := ctx.URLParam("hardwaredesign_id")

	systemID, err1 := strconv.Atoi(systemIDStr)
	designID, err2 := strconv.Atoi(designIDStr)

	if err1 != nil || err2 != nil {
		ctx.ViewData("error", "Invalid parameters")
		ctx.Redirect("/sms_systems/show/" + systemIDStr)
		return
	}

	err := dbprovider.GetDBManager().DeleteSMSHardwareDesignMapping(systemID, designID)
	if err != nil {
		ctx.ViewData("error", "Failed to unlink hardware design from system")
	}

	ctx.Redirect("/sms_systems/show/" + systemIDStr)
}

