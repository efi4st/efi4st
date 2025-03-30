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
)

type KeyValue struct {
	Key   int
	Value string
}

type KeyBool struct {
	Key   int
	Value bool
}

func SMSprojectUpdate(ctx iris.Context) {
	id := ctx.Params().Get("id")
	projectID, err := strconv.Atoi(id)
	if err != nil {
		ctx.ViewData("error", "Error: converting project ID!")
		ctx.View("sms_showProjectUpdate.html")
		return
	}

	// Projektinformationen abrufen
	projectInfo := dbprovider.GetDBManager().GetSMSProjectInfo(projectID)
	ctx.ViewData("projectInfo", projectInfo)

	// Geräte, Software und Systemtypen abrufen
	systemTypeMap, _, err := dbprovider.GetDBManager().GetDevicesAndSoftwareForProject(projectID)
	if err != nil {
		ctx.ViewData("error", "Error fetching device/software list!")
		ctx.View("sms_showProjectUpdate.html")
		return
	}

	// Listen für Template erstellen
	var systemTypeNameList []KeyValue
	var systemTypeCleanList []KeyBool

	systemVersionsMap, err := dbprovider.GetDBManager().GetMostCommonSystemVersionForSystemType(projectID)
	if err != nil {
		fmt.Println("⚠️ Fehler beim Abrufen der häufigsten Systemversionen:", err)
	}

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

		// SystemType-Clean-Status speichern
		systemTypeCleanList = append(systemTypeCleanList, KeyBool{Key: systemTypeID, Value: isClean})

		// SystemType-Namen abrufen
		systemTypeName, err := dbprovider.GetDBManager().GetSystemTypeName(systemTypeID)
		if err != nil {
			fmt.Println("⚠️ Fehler beim Abrufen des SystemType-Namens für ID", systemTypeID, err)
			systemTypeName = fmt.Sprintf("Unknown (ID %d)", systemTypeID) // Fallback
		}
		systemTypeNameList = append(systemTypeNameList, KeyValue{Key: systemTypeID, Value: systemTypeName})
	}

	fmt.Println("DEBUG: SystemTypeNameList:")
	for _, item := range systemTypeNameList {
		fmt.Printf("  - Key: %d, Value: %s\n", item.Key, item.Value)
	}

	fmt.Println("DEBUG: SystemTypeCleanList:")
	for _, item := range systemTypeCleanList {
		fmt.Printf("  - Key: %d, Value: %v\n", item.Key, item.Value)
	}

	// Daten an Template übergeben
	ctx.ViewData("systemTypeMap", systemTypeMap)
	ctx.ViewData("systemTypeNameList", systemTypeNameList)
	ctx.ViewData("systemTypeCleanList", systemTypeCleanList)

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