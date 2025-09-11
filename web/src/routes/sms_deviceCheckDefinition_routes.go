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
	"strings"
)

/////////////////////////////////////////
////	Routen für SMS_DeviceCheckDefinition
/////////////////////////////////////////
// Übersicht der Checks anzeigen
func SMSDeviceCheckDefinitions(ctx iris.Context) {
	deviceCheckDefinitions, err := dbprovider.GetDBManager().GetAllDeviceCheckDefinitions()
	if err != nil {
		ctx.ViewData("error", "Error: Error getting device check definitions!")
	}
	ctx.ViewData("error", "")

	if len(deviceCheckDefinitions) < 1 {
		ctx.ViewData("error", "Error: No device check definitions available. Add one!")
	}

	ctx.ViewData("deviceCheckDefinitionsList", deviceCheckDefinitions)
	ctx.View("sms_deviceCheckDefinitions.html")
}

// GET: Formular für neue Check-Definition
func CreateSMSDeviceCheckDefinitions(ctx iris.Context) {
	deviceTypes := dbprovider.GetDBManager().GetSMSDeviceTypes()

	// Check-Typen Liste
	checkTypes := []string{"SL1 Abnahme", "Basic IBN Check", "Production Check", "Regular Re-Check"}

	// Software-Typen abrufen
	applicationList, err := dbprovider.GetDBManager().GetSoftwareTypesForCheckList()
	if err != nil {
		fmt.Println("Error fetching application list:", err)
		applicationList = []string{} // Fallback auf leere Liste
	}

	artefactTypes := dbprovider.GetDBManager().GetSMSArtefactTypes()      // ⬅️ schon vorhanden bei dir


	// Daten an View übergeben
	ctx.ViewData("deviceTypes", deviceTypes)
	ctx.ViewData("checkTypes", checkTypes)
	ctx.ViewData("applicationList", applicationList) // ✅ Neu hinzugefügt
	ctx.ViewData("artefactTypes", artefactTypes)
	ctx.View("sms_createDeviceCheckDefinition.html")
}

// POST: Neuen Check hinzufügen
func AddSMSDeviceCheckDefinition(ctx iris.Context) {
	deviceTypeID, err := ctx.PostValueInt("device_type_id")
	if err != nil {
		ctx.ViewData("error", "Invalid device type ID.")
		ctx.View("sms_createDeviceCheckDefinition.html")
		return
	}

	applicableVersions := ctx.PostValue("applicable_versions")
	testName := ctx.PostValue("test_name")
	testDescription := ctx.PostValue("test_description")
	explanation := ctx.PostValue("explanation")
	expectedResult := ctx.PostValue("expected_result")
	filterCondition := ctx.PostValue("filter_condition")

	// Mehrfachauswahl für checkType als Liste verarbeiten
	checkTypes, _ := ctx.PostValues("check_type") // Gibt immer []string zurück

	if len(checkTypes) == 0 { // Prüfen, ob mindestens eine Auswahl getroffen wurde
		ctx.ViewData("error", "Error: At least one check type must be selected.")
		ctx.View("sms_createDeviceCheckDefinition.html")
		return
	}

	// CheckTypes als kommagetrennten String speichern
	checkTypesStr := strings.Join(checkTypes, ",") // Umwandlung in String

	err = dbprovider.GetDBManager().AddDeviceCheckDefinition(
		deviceTypeID, applicableVersions, testName, testDescription, &explanation, expectedResult, &filterCondition, checkTypesStr,
	)
	if err != nil {
		ctx.ViewData("error", "Error adding device check definition.")
	} else {
		ctx.Redirect("/sms_deviceCheckDefinitions")
	}
}

// Entfernen einer Check-Definition
func RemoveSMSDeviceCheckDefinition(ctx iris.Context) {
	id := ctx.Params().Get("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		ctx.ViewData("error", "Error: Invalid device check definition ID!")
		ctx.View("sms_deviceCheckDefinitions.html")
		return
	}

	err = dbprovider.GetDBManager().DeleteDeviceCheckDefinition(i)
	ctx.ViewData("error", "")
	if err != nil {
		ctx.ViewData("error", "Error: Error removing device check definition!")
	}

	deviceCheckDefinitions, err := dbprovider.GetDBManager().GetAllDeviceCheckDefinitions()
	if err != nil {
		ctx.ViewData("error", "Error: Error getting device check definitions!")
	}
	ctx.ViewData("deviceCheckDefinitionsList", deviceCheckDefinitions)
	ctx.View("sms_deviceCheckDefinitions.html")
}

func SMSDeviceCheckDetails(ctx iris.Context) {
	checkID, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.ViewData("error", "Invalid Check ID.")
		ctx.Redirect("/sms_deviceCheckDefinitions")
		return
	}

	check, err := dbprovider.GetDBManager().GetDeviceCheckByID(checkID)
	if err != nil {
		ctx.ViewData("error", "Error retrieving check details.")
		ctx.View("error.html")
		return
	}

	ctx.ViewData("check", check)
	ctx.View("sms_showDeviceCheckDefinition.html")
}

func SMSUpdateCheckDefinition(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.ViewData("error", "Invalid Check ID")
		ctx.Redirect("/sms_deviceCheckDefinitions")
		return
	}

	// Holt alle ausgewählten Check Types
	checkTypes, _ := ctx.PostValues("check_type")

	// Wenn keine Check Types ausgewählt wurden
	if len(checkTypes) == 0 {
		ctx.Redirect("/sms_deviceCheckDefinitions/edit/" + strconv.Itoa(id) + "?error=At%20least%20one%20check%20type%20must%20be%20selected.")
		return
	}

	// CheckTypes als kommagetrennten String speichern
	checkTypesStr := strings.Join(checkTypes, ",") // Umwandlung in String

	// Bereite das Update vor
	filterCondition := ctx.PostValue("filter_condition")
	var filterConditionPtr *string
	if filterCondition != "" {
		filterConditionPtr = &filterCondition
	}

	check := classes.Sms_DeviceCheckDefinition{
		ID:                 id, // Hier wird die ID hinzugefügt
		DeviceTypeID:       ctx.PostValueIntDefault("device_type_id", 0),
		ApplicableVersions: ctx.PostValue("applicable_versions"),
		TestName:           ctx.PostValue("test_name"),
		TestDescription:    ctx.PostValue("test_description"),
		Explanation:        ctx.PostValue("explanation"),
		ExpectedResult:     ctx.PostValue("expected_result"),
		FilterCondition:    filterConditionPtr,
		CheckType:          checkTypesStr, // Hier speicherst du die ausgewählten Check Types als String
	}

	// Update in die DB durchführen
	err = dbprovider.GetDBManager().UpdateDeviceCheck(check)
	if err != nil {
		ctx.ViewData("error", "Error updating check")
	}

	// Weiterleitung nach dem Update
	ctx.Redirect("/sms_deviceCheckDefinitions/show/" + strconv.Itoa(id))
}

func SMSEditProjectCheck(ctx iris.Context) {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.ViewData("error", "Invalid Check ID")
		ctx.Redirect("/sms_deviceCheckDefinitions")
		return
	}

	// Holen der Check-Definition basierend auf der ID
	check, err := dbprovider.GetDBManager().GetDeviceCheckByID(id)
	if err != nil {
		ctx.ViewData("error", "Check not found")
		ctx.Redirect("/sms_deviceCheckDefinitions")
		return
	}

	// Fehler aus der URL extrahieren, falls vorhanden
	errorMessage := ctx.URLParam("error")
	if errorMessage != "" {
		ctx.ViewData("error", errorMessage)
	}

	// Hole alle verfügbaren checkTypes (diese kommen eventuell aus der DB)
	checkTypes := []string{"SL1 Abnahme", "Basic IBN Check", "Production Check", "Regular Re-Check"}

	// Wenn checkTypes in der DB als kommagetrennter String gespeichert sind, splitte diesen in eine Liste
	selectedCheckTypes := strings.Split(check.CheckType, ",")

	// Hole die verfügbaren Gerätetypen aus der DB
	deviceTypes := dbprovider.GetDBManager().GetSMSDeviceTypes()

	// Übergabe der Daten an die View
	ctx.ViewData("checkTypes", checkTypes) // Liste der verfügbaren Check-Typen
	ctx.ViewData("selectedCheckTypes", selectedCheckTypes) // Liste der ausgewählten Check-Typen
	ctx.ViewData("check", check) // Die Check-Definition, die bearbeitet werden soll
	ctx.ViewData("deviceTypes", deviceTypes) // Liste der verfügbaren Gerätetypen

	// Ansicht laden
	ctx.View("sms_editDeviceCheckDefinition.html")
}