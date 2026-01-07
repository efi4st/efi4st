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



type SystemTypeGroupView struct {
	SystemTypeID     int
	SystemTypeName   string
	IsClean          bool
	DevicesWithSW    []classes.DeviceUpdateView
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
	projectID, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.ViewData("error", "Fehler beim Konvertieren der Projekt-ID!")
		ctx.View("sms_showProjectUpdate.html")
		return
	}

	projectInfo := dbprovider.GetDBManager().GetSMSProjectInfo(projectID)
	ctx.ViewData("projectInfo", projectInfo)

	// Systeme im Projekt (pro PBOM)
	systemList := dbprovider.GetDBManager().GetSMSProjectBOMForProject(projectID)

	// Updates
	allUpdates, err := dbprovider.GetDBManager().GetSMSUpdateDetailsForProject(projectID)
	if err != nil {
		ctx.ViewData("error", "Fehler beim Holen der Updates!")
		ctx.View("sms_showProjectUpdate.html")
		return
	}

	// Selected update context
	selectedSystemID, _ := strconv.Atoi(ctx.URLParam("system_id"))
	selectedUpdateID, _ := strconv.Atoi(ctx.URLParam("update_id"))

	var selectedUpdate *classes.Sms_UpdateDetails
	if selectedUpdateID > 0 {
		for i := range allUpdates {
			if allUpdates[i].ID == selectedUpdateID {
				selectedUpdate = &allUpdates[i]
				break
			}
		}
	}

	// Zielversionen aus Update (wie bisher)
	var devicesForUpdate []classes.DeviceSoftwareVersion
	var softwareForUpdate []classes.DeviceSoftwareVersion
	if selectedUpdate != nil && selectedUpdate.ToSystemID > 0 {
		devicesForUpdate, _ = dbprovider.GetDBManager().GetDevicesBySystemID(selectedUpdate.ToSystemID)
		softwareForUpdate, _ = dbprovider.GetDBManager().GetSoftwareBySystemID(selectedUpdate.ToSystemID)

		deviceMap := make(map[int]*classes.DeviceSoftwareVersion)
		for i := range devicesForUpdate {
			deviceMap[devicesForUpdate[i].DeviceID] = &devicesForUpdate[i]
		}
		for _, swDevice := range softwareForUpdate {
			if dev, found := deviceMap[swDevice.DeviceID]; found {
				dev.SoftwareList = append(dev.SoftwareList, swDevice.SoftwareList...)
			}
		}
	}

	// Blocks bauen
	blocks := make([]classes.SystemUpdateBlock, 0, len(systemList))
	for _, s := range systemList {
		devicesWithSW, err := dbprovider.GetDBManager().GetDevicesAndSoftwareForProjectBOM(s.ProjectBOMID)
		if err != nil {
			fmt.Println("⚠️ GetDevicesAndSoftwareForProjectBOM:", err)
		}

		available := make([]classes.Sms_UpdateDetails, 0, 16)
		for _, upd := range allUpdates {
			if upd.FromSystemID == s.SystemID {
				available = append(available, upd)
			}
		}

		blocks = append(blocks, classes.SystemUpdateBlock{
			ProjectBOMID:     s.ProjectBOMID,
			SystemID:         s.SystemID,
			SystemTypeName:   s.SystemType,
			SystemVersion:    s.SystemVersion,
			DevicesWithSW:    devicesWithSW,
			AvailableUpdates: available,
		})
	}

	ctx.ViewData("systemBlocks", blocks)
	ctx.ViewData("selectedSystemID", selectedSystemID)
	ctx.ViewData("selectedUpdate", selectedUpdate)
	ctx.ViewData("devicesForUpdate", devicesForUpdate)
	ctx.ViewData("softwareForUpdate", softwareForUpdate)

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