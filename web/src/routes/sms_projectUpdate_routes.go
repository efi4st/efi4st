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

type DeviceUpdateView struct {
	DeviceName              string
	DeviceVersion           string
	UpdateVersion           string
	DeviceCount             int
	IsInvalidSystemVersion  bool
	MostCommonSystemVersion string
	ShortenedSystemVersions string
	SoftwareList            []SoftwareUpdateView
}

type SoftwareUpdateView struct {
	SoftwareName    string
	SoftwareVersion string
	UpdateVersion   string
}

type SystemTypeGroupView struct {
	SystemTypeID     int
	SystemTypeName   string
	IsClean          bool
	DevicesWithSW    []DeviceUpdateView
	AvailableUpdates []classes.Sms_UpdateDetails
}

type DisplayDeviceEntry struct {
	DeviceName              string
	DeviceVersion           string
	DeviceCount             int
	MostCommonSystemVersion string
	IsInvalidSystemVersion  bool
	ShortenedSystemVersions string
	UpdateVersion           string
	SoftwareList            []DisplaySoftwareEntry
}

type DisplaySoftwareEntry struct {
	SoftwareName    string
	SoftwareVersion string
	UpdateVersion   string
}

type DeviceInfo struct {
	ID              int
	DeviceName      string
	DeviceVersion   string
	UpdateVersion   string
	SoftwareList    []SoftwareInfo
	DeviceCount     int
}


// Für Software (innerhalb eines Geräts)
type SoftwareInfo struct {
	SoftwareID       int
	SoftwareName     string
	SoftwareVersion  string
	UpdateVersion    string // ❗️NEU: Ziel-Version aus dem Update-Paket
}

func SMSprojectUpdate(ctx iris.Context) {
	id := ctx.Params().Get("id")
	projectID, err := strconv.Atoi(id)
	if err != nil {
		ctx.ViewData("error", "Fehler beim Konvertieren der Projekt-ID!")
		ctx.View("sms_showProjectUpdate.html")
		return
	}

	projectInfo := dbprovider.GetDBManager().GetSMSProjectInfo(projectID)
	ctx.ViewData("projectInfo", projectInfo)

	// Abruf der Geräte und Software für das Projekt
	systemTypeMap, _, err := dbprovider.GetDBManager().GetDevicesAndSoftwareForProject(projectID)
	if err != nil {
		ctx.ViewData("error", "Fehler beim Abrufen der Geräte-/Software-Liste!")
		ctx.View("sms_showProjectUpdate.html")
		return
	}

	var systemTypeNameList []KeyValue
	var systemTypeCleanList []KeyBool

	systemVersionsMap, err := dbprovider.GetDBManager().GetMostCommonSystemVersionForSystemType(projectID)
	if err != nil {
		fmt.Println("⚠️ Fehler beim Abrufen der häufigsten Systemversionen:", err)
	}

	// Hier holen wir die Systemtypen, um sie später der View zu übergeben
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

		systemTypeName, err := dbprovider.GetDBManager().GetSystemTypeName(systemTypeID)
		if err != nil {
			fmt.Println("⚠️ Fehler beim Abrufen des SystemType-Namens für ID", systemTypeID, err)
			systemTypeName = fmt.Sprintf("Unbekannt (ID %d)", systemTypeID)
		}
		systemTypeNameList = append(systemTypeNameList, KeyValue{Key: systemTypeID, Value: systemTypeName})
	}

	allUpdates, err := dbprovider.GetDBManager().GetSMSUpdateDetailsForProject(projectID)
	if err != nil {
		fmt.Println("⚠️ Fehler beim Holen der Update-Details:", err)
		ctx.ViewData("error", "Fehler beim Holen der Updates!")
		ctx.View("sms_showProjectUpdate.html")
		return
	}

	updateMap := make(map[int][]classes.Sms_UpdateDetails)
	for _, update := range allUpdates {
		updateMap[update.ToSystemTypeID] = append(updateMap[update.ToSystemTypeID], update)
	}

	var systemTypeUpdates []SystemTypeUpdate
	for _, kv := range systemTypeNameList {
		systemTypeID := kv.Key
		systemTypeName := kv.Value

		systemTypeUpdates = append(systemTypeUpdates, SystemTypeUpdate{
			SystemTypeID:   systemTypeID,
			SystemTypeName: systemTypeName,
			Updates:        updateMap[systemTypeID],
		})
	}

	// Abruf der Parameter, falls sie gesetzt sind
	systemTypeIDParam := ctx.URLParam("system_type_id")
	updateIDParam := ctx.URLParam("update_id")

	var selectedSystemTypeID, selectedUpdateID int
	if systemTypeIDParam != "" {
		selectedSystemTypeID, _ = strconv.Atoi(systemTypeIDParam)
	}
	if updateIDParam != "" {
		selectedUpdateID, _ = strconv.Atoi(updateIDParam)
	}

	var selectedUpdate *classes.Sms_UpdateDetails
	if selectedUpdateID > 0 {
		for _, upd := range allUpdates {
			if upd.ID == selectedUpdateID {
				selectedUpdate = &upd
				break
			}
		}
	}

	// Abruf der Geräte- und Softwareversionen für das ausgewählte Update
	var devicesForUpdate []classes.DeviceSoftwareVersion
	var softwareForUpdate []classes.DeviceSoftwareVersion
	if selectedUpdate != nil {
		if selectedUpdate.ToSystemID > 0 {
			fmt.Println("Abruf der Geräte für System-ID:", selectedUpdate.ToSystemID)
			devicesForUpdate, err = dbprovider.GetDBManager().GetDevicesBySystemID(selectedUpdate.ToSystemID)
			if err != nil {
				fmt.Println("⚠️ Fehler beim Abrufen der Geräte:", err)
			}

			softwareForUpdate, err = dbprovider.GetDBManager().GetSoftwareBySystemID(selectedUpdate.ToSystemID)
			if err != nil {
				fmt.Println("⚠️ Fehler beim Abrufen der Software:", err)
			}
		}
	}

	log.Printf("Devices for Update: %+v", devicesForUpdate)
	log.Printf("Software for Update: %+v", softwareForUpdate)

	// Daten an die View übergeben
	ctx.ViewData("systemTypeMap", systemTypeMap)
	ctx.ViewData("systemTypeNameList", systemTypeNameList)
	ctx.ViewData("systemTypeCleanList", systemTypeCleanList)
	ctx.ViewData("systemTypeUpdates", systemTypeUpdates)
	ctx.ViewData("selectedSystemTypeID", selectedSystemTypeID)
	ctx.ViewData("selectedUpdate", selectedUpdate)

	// Geräte und Software an die View weitergeben
	ctx.ViewData("devicesForUpdate", devicesForUpdate)
	ctx.ViewData("softwareForUpdate", softwareForUpdate)

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