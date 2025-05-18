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

func SMSDeviceInstance(ctx iris.Context) {

	deviceInstances := dbprovider.GetDBManager().GetSMSDeviceInstances()
	ctx.ViewData("error", "")

	if len(deviceInstances) < 1 {
		ctx.ViewData("error", "Error: No devices available. Add one!")
	}
	ctx.ViewData("deviceInstanceList", deviceInstances)
	ctx.View("sms_deviceInstances.html")
}

// GET
func CreateSMSDeviceInstance(ctx iris.Context) {

	ctx.View("sms_createDeviceInstance.html")
}

// GET
func CreateSMSDeviceInstanceForProject(ctx iris.Context) {

	id := ctx.Params().Get("id")
	fmt.Print("dfsdfsdfsddfsdfd")
	i, err := strconv.Atoi(id)
	devices := dbprovider.GetDBManager().GetSMSDevice()
	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing device id!")
	}

	ctx.ViewData("projectId", i)
	ctx.ViewData("deviceList", devices)
	ctx.View("sms_createDeviceInstanceForProject.html")
}

// POST
func AddSMSDeviceInstance(ctx iris.Context) {

	project_id := ctx.PostValue("Project_id")
	device_id := ctx.PostValue("Device_id")
	serialnumber := ctx.PostValue("Serialnumber")
	provisioner := ctx.PostValue("Provisioner")
	configuration := ctx.PostValue("Configuration")

	fmt.Printf("ProjectID"+project_id)
	fmt.Printf("device:"+device_id)
	fmt.Printf(serialnumber)
	fmt.Printf(provisioner)
	fmt.Printf(configuration)


	iP, err := strconv.Atoi(project_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing project_id!")
	}

	iD, err := strconv.Atoi(device_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing device_id!")
	}

	err = dbprovider.GetDBManager().AddSMSDeviceInstance(iP, iD, serialnumber, provisioner, configuration)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Not able to add device!")
	}

	ctx.Params().Set("id", project_id)
	ShowSMSProject(ctx)
}

func ShowSMSDeviceInstance(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing device Id!")
	}

	deviceInstanceUpdateHistories := dbprovider.GetDBManager().GetSMSUpdateHistoryForDevice(i)
	deviceInstance := dbprovider.GetDBManager().GetSMSDeviceInstanceInfo(i)
	deviceModel := dbprovider.GetDBManager().GetSMSDeviceInfo(deviceInstance.Device_id())
	issuesForThisDeviceInstance, err := dbprovider.GetDBManager().GetSMSIssuesForDeviceInstance(i)
	if err != nil {
		ctx.ViewData("error", "Error: Error getting issues for device instance!")
	}

	availableVersions, err := dbprovider.GetDBManager().GetAllVersionsForDevice(deviceModel.Device_id)
	if err != nil {
		ctx.ViewData("error", "Could not get available versions for this device.")
	}

	artefactsUnderDeviceInstance := dbprovider.GetDBManager().GetSMSArtefactPartOfDeviceInstanceDetailedForDeviceInstance(i)

	ctx.ViewData("artefactsUnderDeviceInstance", artefactsUnderDeviceInstance)
	ctx.ViewData("availableDeviceVersions", availableVersions)
	ctx.ViewData("availableDeviceVersions", availableVersions)
	ctx.ViewData("deviceModel", deviceModel)
	ctx.ViewData("deviceInstanceId", i)
	ctx.ViewData("deviceInstanceUpdateHistories", deviceInstanceUpdateHistories)
	ctx.ViewData("deviceInstance", deviceInstance)
	ctx.ViewData("issuesForThisDeviceInstance", issuesForThisDeviceInstance)
	ctx.View("sms_showDeviceInstance.html")
}

func UpgradeSMSDeviceInstance(ctx iris.Context) {
	idStr := ctx.Params().Get("id")
	newDevIDStr := ctx.PostValue("newDeviceID")

	deviceInstanceID, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.ViewData("error", "Invalid instance ID.")
		ctx.Redirect("/sms_deviceInstances/show/" + idStr)
		return
	}

	newDeviceID, err := strconv.Atoi(newDevIDStr)
	if err != nil {
		ctx.ViewData("error", "Invalid new device ID.")
		ctx.Redirect("/sms_deviceInstances/show/" + idStr)
		return
	}

	// Aktuelle Device-Version ermitteln (vor dem Update)
	oldDevice, err := dbprovider.GetDBManager().GetDeviceForInstance(deviceInstanceID)
	if err != nil {
		ctx.ViewData("error", "Could not fetch current device for update log.")
		ctx.Redirect("/sms_deviceInstances/show/" + idStr)
		return
	}

	// Neue Device-Version ermitteln
	newDevice, err := dbprovider.GetDBManager().GetDeviceByID(newDeviceID)
	if err != nil {
		ctx.ViewData("error", "Failed to fetch new device details.")
		ctx.Redirect("/sms_deviceInstances/show/" + idStr)
		return
	}

	// Update durchführen
	err = dbprovider.GetDBManager().UpgradeDeviceInstance(deviceInstanceID, newDeviceID)
	if err != nil {
		ctx.ViewData("error", "Upgrade failed.")
		ctx.Redirect("/sms_deviceInstances/show/" + idStr)
		return
	}

	// Optional: Benutzer ermitteln
	user := ctx.Values().GetString("user")
	if user == "" {
		user = "system"
	}

	// Beschreibung fürs Log
	description := fmt.Sprintf(
		"Upgraded from version %s (Device ID %d) to version %s (Device ID %d)",
		oldDevice.Version, oldDevice.Device_id,
		newDevice.Version, newDevice.Device_id,
	)

	// Zeitstempel setzen
	now := time.Now().Format("2006-01-02")

	// Log-Eintrag anlegen
	err = dbprovider.GetDBManager().InsertUpdateHistory(deviceInstanceID, user, "DeviceUpgrade", now, description)
	if err != nil {
		ctx.ViewData("error", "Upgrade succeeded, but could not log update history.")
	}

	// Weiterleitung zurück zur Detailansicht
	ctx.Redirect("/sms_deviceInstances/show/" + idStr)
}

func RemoveSMSDeviceInstance(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveSMSDeviceInstances(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing device!")
	}

	deviceInstances := dbprovider.GetDBManager().GetSMSDeviceInstances()

	ctx.ViewData("deviceInstanceList", deviceInstances)
	ctx.View("sms_deviceInstances.html")
}