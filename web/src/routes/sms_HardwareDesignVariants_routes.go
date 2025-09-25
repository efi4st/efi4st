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
	"strconv"
)

// LISTE: Varianten eines Designs
func SMSHardwareDesignVariants(ctx iris.Context) {
	designIDStr := ctx.Params().Get("id")
	all := ctx.URLParamDefault("all", "0") == "1"
	designID, err := strconv.Atoi(designIDStr)
	if err != nil {
		ctx.ViewData("error", "Invalid hardware design ID")
		ctx.Redirect("/sms_hardwaredesigns")
		return
	}

	design := dbprovider.GetDBManager().GetSMSHardwareDesignByID(designID)
	if design == nil {
		ctx.ViewData("error", "Design not found")
		ctx.Redirect("/sms_hardwaredesigns")
		return
	}

	variants := dbprovider.GetDBManager().GetVariantsForHardwareDesign(designID, !all)

	ctx.ViewData("error", "")
	ctx.ViewData("design", design)
	ctx.ViewData("variants", variants)
	ctx.ViewData("showAll", all)
	ctx.View("sms_showHardwareDesignVariant.html")
}

// CREATE FORM
func CreateSMSHardwareDesignVariant(ctx iris.Context) {
	designIDStr := ctx.Params().Get("id")
	designID, err := strconv.Atoi(designIDStr)
	if err != nil {
		ctx.ViewData("error", "Invalid hardware design ID")
		ctx.Redirect("/sms_hardwaredesigns")
		return
	}
	design := dbprovider.GetDBManager().GetSMSHardwareDesignByID(designID)
	if design == nil {
		ctx.ViewData("error", "Design not found")
		ctx.Redirect("/sms_hardwaredesigns")
		return
	}
	ctx.ViewData("error", "")
	ctx.ViewData("design", design)
	ctx.View("sms_createHardwareDesignVariant.html")
}

// POST: anlegen
func AddSMSHardwareDesignVariant(ctx iris.Context) {
	designIDStr := ctx.Params().Get("id")
	designID, err := strconv.Atoi(designIDStr)
	if err != nil {
		ctx.ViewData("error", "Invalid hardware design ID")
		ctx.Redirect("/sms_hardwaredesigns")
		return
	}

	code := ctx.PostValue("Code")
	name := ctx.PostValue("Name")
	desc := ctx.PostValue("Description")
	spec := ctx.PostValue("Spec") // JSON optional
	isActive := ctx.PostValue("IsActive") == "on"

	v := &classes.Sms_HardwareDesignVariant{
		HardwareDesignID: designID,
		Code:             code,
		Name:             name,
		Description:      desc,
		Spec:             spec,
		IsActive:         isActive,
	}

	if err := dbprovider.GetDBManager().AddVariant(v); err != nil {
		ctx.ViewData("error", "Error: could not create variant — "+err.Error())
		// wieder Formular anzeigen
		design := dbprovider.GetDBManager().GetSMSHardwareDesignByID(designID)
		ctx.ViewData("design", design)
		ctx.View("sms_createHardwareDesignVariant.html")
		return
	}

	ctx.Redirect(fmt.Sprintf("/sms_hardwaredesigns/%d/variants", designID))
}

// EDIT FORM
func EditSMSHardwareDesignVariant(ctx iris.Context) {
	variantIDStr := ctx.Params().Get("variant_id")
	variantID, err := strconv.Atoi(variantIDStr)
	if err != nil {
		ctx.ViewData("error", "Invalid variant ID")
		ctx.Redirect("/sms_hardwaredesigns")
		return
	}
	variant := dbprovider.GetDBManager().GetVariantByID(variantID)
	if variant == nil {
		ctx.ViewData("error", "Variant not found")
		ctx.Redirect("/sms_hardwaredesigns")
		return
	}
	design := dbprovider.GetDBManager().GetSMSHardwareDesignByID(variant.HardwareDesignID)
	ctx.ViewData("error", "")
	ctx.ViewData("design", design)
	ctx.ViewData("variant", variant)
	ctx.View("sms_editHardwareDesignVariant.html")
}

// POST: update
func UpdateSMSHardwareDesignVariant(ctx iris.Context) {
	variantIDStr := ctx.Params().Get("variant_id")
	variantID, err := strconv.Atoi(variantIDStr)
	if err != nil {
		ctx.ViewData("error", "Invalid variant ID")
		ctx.Redirect("/sms_hardwaredesigns")
		return
	}
	// alte Daten holen (für Redirect)
	vOld := dbprovider.GetDBManager().GetVariantByID(variantID)
	if vOld == nil {
		ctx.ViewData("error", "Variant not found")
		ctx.Redirect("/sms_hardwaredesigns")
		return
	}

	v := &classes.Sms_HardwareDesignVariant{
		HardwareDesignVariantID: variantID,
		HardwareDesignID:        vOld.HardwareDesignID,
		Code:                    ctx.PostValue("Code"),
		Name:                    ctx.PostValue("Name"),
		Description:             ctx.PostValue("Description"),
		Spec:                    ctx.PostValue("Spec"),
		IsActive:                ctx.PostValue("IsActive") == "on",
	}
	if err := dbprovider.GetDBManager().UpdateVariant(v); err != nil {
		ctx.ViewData("error", "Error: could not update variant — "+err.Error())
		design := dbprovider.GetDBManager().GetSMSHardwareDesignByID(v.HardwareDesignID)
		ctx.ViewData("design", design)
		ctx.ViewData("variant", v)
		ctx.View("sms_editHardwareDesignVariant.html")
		return
	}
	ctx.Redirect(fmt.Sprintf("/sms_hardwaredesigns/%d/variants", v.HardwareDesignID))
}

// Toggle aktiv/inaktiv
func ToggleSMSHardwareDesignVariantActive(ctx iris.Context) {
	variantIDStr := ctx.Params().Get("variant_id")
	variantID, err := strconv.Atoi(variantIDStr)
	if err != nil {
		ctx.Redirect("/sms_hardwaredesigns")
		return
	}
	active := ctx.URLParamDefault("active", "1") == "1"

	_ = dbprovider.GetDBManager().SetVariantActiveFlag(variantID, active)

	// zurück zur Variantenliste des Designs
	v := dbprovider.GetDBManager().GetVariantByID(variantID)
	if v != nil {
		ctx.Redirect(fmt.Sprintf("/sms_hardwaredesigns/%d/variants?all=1", v.HardwareDesignID))
		return
	}
	ctx.Redirect("/sms_hardwaredesigns")
}

// DELETE
func RemoveSMSHardwareDesignVariant(ctx iris.Context) {
	variantIDStr := ctx.Params().Get("variant_id")
	variantID, err := strconv.Atoi(variantIDStr)
	if err != nil {
		ctx.Redirect("/sms_hardwaredesigns")
		return
	}
	v := dbprovider.GetDBManager().GetVariantByID(variantID)
	if v == nil {
		ctx.Redirect("/sms_hardwaredesigns")
		return
	}
	if err := dbprovider.GetDBManager().DeleteVariantByID(variantID); err != nil {
		// z.B. FK-Restrict weil in projectBOM genutzt
		ctx.ViewData("error", "Variante konnte nicht gelöscht werden: "+err.Error())
		design := dbprovider.GetDBManager().GetSMSHardwareDesignByID(v.HardwareDesignID)
		variants := dbprovider.GetDBManager().GetVariantsForHardwareDesign(v.HardwareDesignID, false)
		ctx.ViewData("design", design)
		ctx.ViewData("variants", variants)
		ctx.ViewData("showAll", true)
		ctx.View("sms_showHardwareDesignVariant.html")
		return
	}
	ctx.Redirect(fmt.Sprintf("/sms_hardwaredesigns/%d/variants?all=1", v.HardwareDesignID))
}