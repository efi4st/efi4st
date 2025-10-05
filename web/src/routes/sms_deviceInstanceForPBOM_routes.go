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

func CreateDeviceInstanceForPBOM(ctx iris.Context) {
	pbomID, _ := strconv.Atoi(ctx.Params().Get("projectBOM_id"))
	if pbomID <= 0 { ctx.Redirect("/sms_projects"); return }

	// Lade PBOM (für Projekt-Redirect), optional System/Design/Variante anzeigen
	rec := dbprovider.GetDBManager().GetProjectBOMByID(pbomID)
	if rec == nil { ctx.Redirect("/sms_projects"); return }

	// Zeige dein bestehendes Create-Form für DeviceInstances – aber mit Hidden PBOM
	ctx.ViewData("ProjectBOMID", pbomID)
	ctx.ViewData("ProjectID", rec.ProjectID)
	ctx.View("sms_createDeviceInstance_for_pbom.html")
}

func LinkExistingDeviceInstanceToPBOM(ctx iris.Context) {
	pbomID, _ := strconv.Atoi(ctx.PostValue("ProjectBOMID"))
	diID,  _  := strconv.Atoi(ctx.PostValue("DeviceInstanceID"))
	note := ctx.PostValue("AdditionalInfo")

	if pbomID <= 0 || diID <= 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.WriteString("invalid ids")
		return
	}

	// PBOM holen (für project_id & Redirect)
	rec := dbprovider.GetDBManager().GetProjectBOMByID(pbomID)
	if rec == nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.WriteString("unknown pbom")
		return
	}

	// **Reuse deiner bestehenden Methode**
	di := dbprovider.GetDBManager().GetSMSDeviceInstanceInfo(diID)
	if di == nil || di.Project_id() != rec.ProjectID {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.WriteString("device does not belong to this project")
		return
	}

	if err := dbprovider.GetDBManager().LinkDeviceInstanceToPBOM(pbomID, diID, note); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.WriteString("link failed: " + err.Error())
		return
	}

	// hübscher Redirect: direkt zur Geräte-Box des Systems springen
	anchor := fmt.Sprintf("#pbom-%d-devices", pbomID)
	ctx.Redirect(fmt.Sprintf("/sms_projects/show/%d%s", rec.ProjectID, anchor))
}

func UnlinkDeviceInstanceFromPBOM(ctx iris.Context) {
	pbomID, _ := strconv.Atoi(ctx.PostValue("ProjectBOMID"))
	diID,  _ := strconv.Atoi(ctx.PostValue("DeviceInstanceID"))
	if pbomID <= 0 || diID <= 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.WriteString("invalid ids")
		return
	}
	if err := dbprovider.GetDBManager().UnlinkDeviceInstanceFromPBOM(pbomID, diID); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.WriteString("unlink failed: "+err.Error())
		return
	}
	prj := dbprovider.GetDBManager().GetProjectBOMByID(pbomID)
	ctx.Redirect(fmt.Sprintf("/sms_projects/show/%d", prj.ProjectID))
}


// GET: Formular anzeigen – Devices des Systems als Dropdown
func CreateDeviceInstanceForPBOMForm(ctx iris.Context) {
	pbomID, _ := strconv.Atoi(ctx.Params().Get("projectBOM_id"))
	if pbomID <= 0 { ctx.Redirect("/sms_projects"); return }

	pbom := dbprovider.GetDBManager().GetProjectBOMByID(pbomID)
	if pbom == nil { ctx.Redirect("/sms_projects"); return }

	// Devices, die zu diesem System gehören
	devChoices := dbprovider.GetDBManager().GetDevicesForSystem(pbom.SystemID)

	ctx.ViewData("ProjectBOMID", pbomID)
	ctx.ViewData("ProjectID",    pbom.ProjectID)
	ctx.ViewData("SystemID",     pbom.SystemID)
	ctx.ViewData("SystemLabel",  fmt.Sprintf("%s v%s", pbom.SystemType, pbom.SystemVersion))
	ctx.ViewData("DeviceChoices", devChoices)
	ctx.ViewData("Today", time.Now().Format("2006-01-02"))

	ctx.View("sms_createDeviceInstance_for_pbom.html")
}

// POST: anlegen + mappen + zurück ins Projekt
func CreateDeviceInstanceForPBOMSubmit(ctx iris.Context) {
	pbomID, _ := strconv.Atoi(ctx.PostValue("ProjectBOMID"))
	projectID, _ := strconv.Atoi(ctx.PostValue("ProjectID"))
	deviceID, _ := strconv.Atoi(ctx.PostValue("DeviceID"))

	serial       := ctx.PostValue("Serialnumber")
	provisioner  := ctx.PostValue("Provisioner")
	configuration:= ctx.PostValue("Configuration")
	date         := ctx.PostValue("Date")
	note         := ctx.PostValue("AdditionalInfo")

	if pbomID == 0 || projectID == 0 || deviceID == 0 || date == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Missing required fields.")
		return
	}

	// 1) DeviceInstance anlegen
	newID, err := dbprovider.GetDBManager().CreateDeviceInstance(projectID, deviceID, serial, provisioner, configuration, date)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Create device instance failed: " + err.Error())
		return
	}

	// 2) Mapping auf PBOM
	if err := dbprovider.GetDBManager().LinkDeviceInstanceToPBOM(pbomID, int(newID), note); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Link to PBOM failed: " + err.Error())
		return
	}

	// 3) Zurück zur Projektseite
	ctx.Redirect(fmt.Sprintf("/sms_projects/show/%d", projectID))
}

