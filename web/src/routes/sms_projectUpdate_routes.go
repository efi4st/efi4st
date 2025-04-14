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

type KeyValue struct {
	Key   int
	Value string
}

type KeyBool struct {
	Key   int
	Value bool
}

type SystemTypeUpdate struct {
	SystemTypeID   int
	SystemTypeName string
	Updates        []classes.Sms_UpdateDetails
}

func SMSprojectUpdate(ctx iris.Context) {
	// Projekt-ID aus den URL-Parametern holen
	id := ctx.Params().Get("id")
	projectID, err := strconv.Atoi(id)
	if err != nil {
		ctx.ViewData("error", "Fehler beim Konvertieren der Projekt-ID!")
		ctx.View("sms_showProjectUpdate.html")
		return
	}

	// Projektinformationen abrufen
	projectInfo := dbprovider.GetDBManager().GetSMSProjectInfo(projectID)
	ctx.ViewData("projectInfo", projectInfo)

	// Geräte, Software und Systemtypen für das Projekt holen
	systemTypeMap, _, err := dbprovider.GetDBManager().GetDevicesAndSoftwareForProject(projectID)
	if err != nil {
		ctx.ViewData("error", "Fehler beim Abrufen der Geräte-/Software-Liste!")
		ctx.View("sms_showProjectUpdate.html")
		return
	}

	// Listen für die Template-Daten erstellen
	var systemTypeNameList []KeyValue
	var systemTypeCleanList []KeyBool

	// Häufigste Versionen der Systemtypen für das Projekt holen
	systemVersionsMap, err := dbprovider.GetDBManager().GetMostCommonSystemVersionForSystemType(projectID)
	if err != nil {
		fmt.Println("⚠️ Fehler beim Abrufen der häufigsten Systemversionen:", err)
	}

	// Clean-Status und Systemnamen für jeden Systemtyp ermitteln
	for systemTypeID, devices := range systemTypeMap {
		isClean := true
		if mostCommonVersion, found := systemVersionsMap[systemTypeID]; found {
			for _, ds := range devices {
				if !contains(ds.SystemVersions, mostCommonVersion) {
					isClean = false
					break
				}
			}
		}

		systemTypeCleanList = append(systemTypeCleanList, KeyBool{Key: systemTypeID, Value: isClean})

		// Systemtyp-Namen holen
		systemTypeName, err := dbprovider.GetDBManager().GetSystemTypeName(systemTypeID)
		if err != nil {
			fmt.Println("⚠️ Fehler beim Abrufen des SystemType-Namens für ID", systemTypeID, err)
			systemTypeName = fmt.Sprintf("Unbekannt (ID %d)", systemTypeID)
		}
		systemTypeNameList = append(systemTypeNameList, KeyValue{Key: systemTypeID, Value: systemTypeName})
	}

	// Alle Updates für das Projekt holen
	allUpdates, err := dbprovider.GetDBManager().GetSMSUpdateDetailsForProject(projectID)
	if err != nil {
		fmt.Println("⚠️ Fehler beim Holen der Update-Details:", err)
		ctx.ViewData("error", "Fehler beim Holen der Updates!")
		ctx.View("sms_showProjectUpdate.html")
		return
	}

	// Updates nach SystemTypeID gruppieren
	updateMap := make(map[int][]classes.Sms_UpdateDetails)

	for _, update := range allUpdates {
		updateMap[update.ToSystemTypeID] = append(updateMap[update.ToSystemTypeID], update)
	}

	// Finales Mapping der Updates je SystemTypeID erstellen
	var systemTypeUpdates []SystemTypeUpdate
	for _, kv := range systemTypeNameList {
		systemTypeID := kv.Key
		systemTypeName := kv.Value

		// SystemTypeUpdate für diesen Systemtyp erstellen
		systemTypeUpdates = append(systemTypeUpdates, SystemTypeUpdate{
			SystemTypeID:   systemTypeID,
			SystemTypeName: systemTypeName,
			Updates:        updateMap[systemTypeID], // Updates für diesen Systemtyp
		})
	}

	// Debugging: Ausgeben der systemTypeUpdates zur Überprüfung
	fmt.Println("SystemTypeUpdates:", systemTypeUpdates)
	ctx.ViewData("systemTypeUpdates", systemTypeUpdates)
	ctx.ViewData("systemTypeMap", systemTypeMap)
	ctx.ViewData("systemTypeNameList", systemTypeNameList)
	ctx.ViewData("systemTypeCleanList", systemTypeCleanList)

	// Template anzeigen
	ctx.View("sms_showProjectUpdate.html")
}

func contains(versions []string, version string) bool {
	for _, v := range versions {
		if v == version {
			return true
		}
	}
	return false
}