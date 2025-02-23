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
	"net"
	"strconv"
)

func SMSDeviceIPDefinitions(ctx iris.Context) {
	// Hole alle Device IP Definitions aus der DB
	deviceIPDefinitions, err := dbprovider.GetDBManager().GetAllDeviceIPDefinitions()
	if err != nil {
		ctx.ViewData("error", "Error: Error getting device IP definitions!")
	}
	ctx.ViewData("error", "")

	// Falls keine IP-Definitionen vorhanden sind, setze eine Fehlermeldung
	if len(deviceIPDefinitions) < 1 {
		ctx.ViewData("error", "Error: No device IP definitions available. Add one!")
	}

	// Füge die Liste der Device IP Definitions zur View hinzu
	ctx.ViewData("deviceIPDefinitionsList", deviceIPDefinitions)
	ctx.View("sms_deviceIPDefinitions.html")
}

// GET
func CreateSMSDeviceIPDefinitions(ctx iris.Context) {
	// Hole alle Device-Typen aus der DB (Beispiel-Query)
	deviceTypes := dbprovider.GetDBManager().GetSMSDeviceTypes()

	// Übergebe die Device-Typen an die View
	ctx.ViewData("deviceTypes", deviceTypes)
	ctx.View("sms_createDeviceIPDefinition.html")
}

// POST
func AddSMSDeviceIPDefinition(ctx iris.Context) {
	// Hole die POST-Daten
	deviceTypeID, err := ctx.PostValueInt("device_type_id")
	if err != nil {
		ctx.ViewData("error", "Invalid device type ID.")
		ctx.View("sms_createDeviceIPDefinition.html")
		return
	}
	applicableVersions := ctx.PostValue("applicable_versions")
	ipAddress := ctx.PostValue("ip_address")
	vlanID := ctx.PostValueIntDefault("vlan_id", 0) // Wenn VLAN nicht gesetzt ist, wird 0 verwendet
	description := ctx.PostValue("description")
	filterCondition := ctx.PostValue("filter_condition")

	// Überprüfe, ob die IP-Adresse gültig ist (IPv4/IPv6 Check)
	if net.ParseIP(ipAddress) == nil {
		ctx.ViewData("error", "Invalid IP address format.")
		ctx.View("sms_createDeviceIPDefinition.html")
		return
	}
	fmt.Println(applicableVersions)
	// Füge die Device IP Definition hinzu
	err = dbprovider.GetDBManager().AddDeviceIPDefinition(deviceTypeID, applicableVersions, ipAddress, &vlanID, &description, &filterCondition)
	if err != nil {
		ctx.ViewData("error", "Error adding device IP definition.")
	} else {
		ctx.Redirect("/sms_deviceIPDefinitions")
	}
}

func RemoveSMSDeviceIPDefinition(ctx iris.Context) {
	// ID aus der URL-Parameter holen
	id := ctx.Params().Get("id")

	// In Integer umwandeln
	i, err := strconv.Atoi(id)
	if err != nil {
		ctx.ViewData("error", "Error: Invalid device IP definition ID!")
		ctx.View("sms_deviceIPDefinitions.html")
		return
	}

	// Löschen des Eintrags aus der Datenbank
	err = dbprovider.GetDBManager().DeleteDeviceIPDefinition(i)

	// Fehlerbehandlung
	ctx.ViewData("error", "")
	if err != nil {
		ctx.ViewData("error", "Error: Error removing device IP definition!")
	}

	// Aktualisierte Liste der IP-Definitionen abrufen
	deviceIPDefinitions, err := dbprovider.GetDBManager().GetAllDeviceIPDefinitions()
	if err != nil {
		ctx.ViewData("error", "Error: Error getting device IP definitions!")
	}
	ctx.ViewData("deviceIPDefinitionsList", deviceIPDefinitions)

	// View rendern
	ctx.View("sms_deviceIPDefinitions.html")
}