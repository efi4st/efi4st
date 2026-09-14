/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/efi4st/efi4st/classes"
	"github.com/efi4st/efi4st/dbprovider"
	"github.com/efi4st/efi4st/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"io"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
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

	// update target maps: deviceName -> version, and deviceName -> swName -> swVersion
	updateDev := map[string]string{}
	updateSw := map[string]map[string]string{}

	for _, d := range devicesForUpdate {
		if d.DeviceName != "" && d.DeviceVersion != "" {
			updateDev[d.DeviceName] = d.DeviceVersion
		}
		if _, ok := updateSw[d.DeviceName]; !ok {
			updateSw[d.DeviceName] = map[string]string{}
		}
		for _, sw := range d.SoftwareList {
			if sw.SoftwareName != "" && sw.SoftwareVersion != "" {
				updateSw[d.DeviceName][sw.SoftwareName] = sw.SoftwareVersion
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

	// Alte Projekt-Metadaten wiederverwenden: SystemVersions / belongs-to / clean-check
	legacySystemTypeMap, _, err := dbprovider.GetDBManager().GetDevicesAndSoftwareForProject(projectID)
	if err != nil {
		log.Println("[UPDATE-VIEW] warning: could not load legacy device/system meta:", err)
	} else {
		// flatten nach deviceName|deviceVersion
		projectMetaByKey := make(map[string]classes.DeviceSoftwareInfo)

		for _, list := range legacySystemTypeMap {
			for _, ds := range list {
				key := ds.DeviceName + "|" + ds.DeviceVersion
				if _, exists := projectMetaByKey[key]; !exists {
					projectMetaByKey[key] = ds
				}
			}
		}

		// in die neuen SystemBlocks eintragen
		for bi := range blocks {
			blocks[bi].IsClean = true

			for di := range blocks[bi].DevicesWithSW {
				entry := &blocks[bi].DevicesWithSW[di]
				key := entry.DeviceName + "|" + entry.DeviceVersion

				if meta, ok := projectMetaByKey[key]; ok {
					entry.SystemVersions = append([]string{}, meta.SystemVersions...)
					entry.MostCommonSystemVersion = meta.MostCommonSystemVersion
					entry.ShortenedSystemVersions = meta.ShortenedSystemVersions

					// Wichtiger als "most common": passt dieses Device überhaupt in diesen Block?
					entry.IsInvalidSystemVersion = !contains(meta.SystemVersions, blocks[bi].SystemVersion)
					if entry.IsInvalidSystemVersion {
						blocks[bi].IsClean = false
					}

					// Software bekommt dieselbe Projekt-System-Info mit
					for si := range entry.SoftwareList {
						entry.SoftwareList[si].ShortenedSystemVersions = meta.ShortenedSystemVersions
					}
				}
			}
		}
	}

	// Live report laden (optional)
	var liveState *classes.LiveProjectState

	liveRow, err := dbprovider.GetDBManager().GetLatestLiveReportForProject(projectID)
	if err == nil && liveRow != nil {
		var lr classes.LiveReportV1
		if jsonErr := json.Unmarshal([]byte(liveRow.PayloadJSON), &lr); jsonErr == nil {
			liveState = utils.BuildLiveProjectStateFromReport(lr, liveRow.CreatedAt, liveRow.ReceivedAt)
			ctx.ViewData("liveState", liveState)
		}
	}

	// DB- und Live-Zustand pro konkreter Geräteinstanz setzen.
	// Funktioniert auch dann, wenn kein Live-Report vorhanden ist.
	for bi := range blocks {
		for di := range blocks[bi].DevicesWithSW {
			dev := &blocks[bi].DevicesWithSW[di]

			// Vergleich pro konkreter Instanz.
			for ii := range dev.Instances {
				inst := &dev.Instances[ii]

				overrides, err :=
					dbprovider.GetDBManager().
						GetDeviceInstanceSoftwareOverrides(inst.DeviceInstanceID)

				if err != nil {
					ctx.StopWithError(iris.StatusInternalServerError, err)
					return
				}

				overrideBySoftwaretypeID :=
					make(map[int]classes.Sms_DeviceInstanceSoftwareOverride)

				for _, override := range overrides {
					overrideBySoftwaretypeID[override.SoftwaretypeID] = override
				}

				inst.DBDeviceVersion = dev.DeviceVersion
				inst.LiveDeviceVersion = ""
				inst.FoundInLive = false
				inst.DBLiveMatch = false
				inst.StatusText = ""
				inst.Software = []classes.InstanceSoftwareUpdateView{}

				// Kein Live-Report vorhanden.
				if liveState == nil {
					continue
				}

				liveDevice, found :=
					liveState.DeviceBySerialnumber[inst.Serialnumber]

				if !found {
					inst.StatusText = "not found in Live Report"
					continue
				}

				inst.FoundInLive = true
				inst.LiveDeviceVersion = liveDevice.DeviceVersion
				inst.DBLiveMatch =
					inst.DBDeviceVersion == inst.LiveDeviceVersion

				if !inst.DBLiveMatch {
					inst.StatusText = "Live: " + inst.LiveDeviceVersion
				}

				// Live-Software dieser konkreten Instanz als Map.
				liveSoftwareByName := make(map[string]string)

				for _, liveSoftware := range liveDevice.Software {
					if liveSoftware.Name == "" {
						continue
					}

					liveSoftwareByName[liveSoftware.Name] =
						liveSoftware.Version
				}

				handledSoftwaretypeIDs := make(map[int]bool)

				// Erwartete DB-Software mit Live vergleichen.
				for _, dbSoftware := range dev.SoftwareList {

					handledSoftwaretypeIDs[dbSoftware.SoftwaretypeID] = true

					modelVersion := dbSoftware.SoftwareVersion
					overrideVersion := ""
					effectiveVersion := modelVersion

					if override, ok :=
						overrideBySoftwaretypeID[dbSoftware.SoftwaretypeID]; ok {

						overrideVersion = override.SoftwareVersion
						effectiveVersion = override.SoftwareVersion
					}

					comparison := classes.InstanceSoftwareUpdateView{
						DeviceInstanceID: inst.DeviceInstanceID,
						SoftwaretypeID:   dbSoftware.SoftwaretypeID,

						SoftwareName: dbSoftware.SoftwareName,

						ModelVersion:     modelVersion,
						OverrideVersion:  overrideVersion,
						EffectiveVersion: effectiveVersion,

						DBVersion: effectiveVersion,

						ShortenedSystemVersions: dbSoftware.ShortenedSystemVersions,
					}

					liveVersion, softwareFound :=
						liveSoftwareByName[dbSoftware.SoftwareName]

					if !softwareFound {
						comparison.FoundInLive = false
						comparison.DBLiveMatch = false
						comparison.StatusText = "not found in Live Report"
					} else {
						comparison.FoundInLive = true
						comparison.LiveVersion = liveVersion

						liveSoftware, err :=
							dbprovider.GetDBManager().
								GetSoftwareByTypeNameAndVersion(
									comparison.SoftwareName,
									comparison.LiveVersion,
								)

						if err != nil {
							ctx.StopWithError(iris.StatusInternalServerError, err)
							return
						}

						if liveSoftware != nil {
							comparison.LiveSoftwareKnown = true
							comparison.LiveSoftwareID = liveSoftware.Software_id()
						}

						comparison.DBLiveMatch =
							comparison.DBVersion == comparison.LiveVersion

						delete(liveSoftwareByName, dbSoftware.SoftwareName)

						if !comparison.DBLiveMatch {
							comparison.StatusText =
								"Live: " + comparison.LiveVersion
						}
					}

					// Gelbe Markierung: DB unterscheidet sich vom Live-Stand.
					comparison.DBOutdated =
						comparison.FoundInLive &&
							comparison.DBVersion != "" &&
							comparison.LiveVersion != "" &&
							comparison.DBVersion != comparison.LiveVersion

					// Zielversion der ausgewählten Aktualisierung.
					if bySoftware, ok := updateSw[dev.DeviceName]; ok {
						if targetVersion, found :=
							bySoftware[comparison.SoftwareName]; found {

							comparison.UpdateTargetVersion = targetVersion

							if comparison.FoundInLive {
								comparison.UpdateAvailable =
									targetVersion != comparison.LiveVersion
							} else {
								comparison.UpdateAvailable =
									targetVersion != comparison.DBVersion
							}
						}
					}

					inst.Software = append(
						inst.Software,
						comparison,
					)
				}

				// Zusätzliche Overrides, deren Softwaretyp nicht im offiziellen Modell vorkommt.
				for softwaretypeID, override := range overrideBySoftwaretypeID {

					if handledSoftwaretypeIDs[softwaretypeID] {
						continue
					}

					comparison := classes.InstanceSoftwareUpdateView{
						DeviceInstanceID: inst.DeviceInstanceID,
						SoftwaretypeID:   softwaretypeID,

						SoftwareName: override.SoftwareName,

						ModelVersion:     "",
						OverrideVersion:  override.SoftwareVersion,
						EffectiveVersion: override.SoftwareVersion,

						DBVersion: override.SoftwareVersion,
					}

					liveVersion, softwareFound :=
						liveSoftwareByName[override.SoftwareName]

					if !softwareFound {
						comparison.FoundInLive = false
						comparison.DBLiveMatch = false
						comparison.StatusText = "override not found in Live Report"
					} else {
						comparison.FoundInLive = true
						comparison.LiveVersion = liveVersion

						liveSoftware, err :=
							dbprovider.GetDBManager().
								GetSoftwareByTypeNameAndVersion(
									comparison.SoftwareName,
									comparison.LiveVersion,
								)

						if err != nil {
							ctx.StopWithError(iris.StatusInternalServerError, err)
							return
						}

						if liveSoftware != nil {
							comparison.LiveSoftwareKnown = true
							comparison.LiveSoftwareID = liveSoftware.Software_id()
						}

						comparison.DBLiveMatch =
							comparison.DBVersion == comparison.LiveVersion

						delete(liveSoftwareByName, override.SoftwareName)

						if !comparison.DBLiveMatch {
							comparison.StatusText =
								"Live: " + comparison.LiveVersion
						}
					}

					comparison.DBOutdated =
						comparison.FoundInLive &&
							comparison.DBVersion != "" &&
							comparison.LiveVersion != "" &&
							comparison.DBVersion != comparison.LiveVersion

					if bySoftware, ok := updateSw[dev.DeviceName]; ok {
						if targetVersion, found :=
							bySoftware[comparison.SoftwareName]; found {

							comparison.UpdateTargetVersion = targetVersion

							if comparison.FoundInLive {
								comparison.UpdateAvailable =
									targetVersion != comparison.LiveVersion
							} else {
								comparison.UpdateAvailable =
									targetVersion != comparison.DBVersion
							}
						}
					}

					inst.Software = append(
						inst.Software,
						comparison,
					)
				}

				// Zusätzliche Software, die nur im Live-Report vorhanden ist.
				for softwareName, liveVersion := range liveSoftwareByName {

					comparison := classes.InstanceSoftwareUpdateView{
						DeviceInstanceID: inst.DeviceInstanceID,
						SoftwareName: softwareName,

						// Diese Software ist weder Teil des offiziellen Modells
						// noch aktuell als Override der Instanz bekannt.
						ModelVersion:     "",
						OverrideVersion:  "",
						EffectiveVersion: "",

						DBVersion: "",

						LiveVersion: liveVersion,

						FoundInLive: true,
						DBLiveMatch: false,

						StatusText: "only in Live Report",

						DBOutdated: false,
					}

					liveSoftware, err :=
						dbprovider.GetDBManager().
							GetSoftwareByTypeNameAndVersion(
								comparison.SoftwareName,
								comparison.LiveVersion,
							)

					if err != nil {
						ctx.StopWithError(iris.StatusInternalServerError, err)
						return
					}

					if liveSoftware != nil {
						comparison.LiveSoftwareKnown = true
						comparison.LiveSoftwareID = liveSoftware.Software_id()
						comparison.SoftwaretypeID = liveSoftware.Softwaretype_id()
					}

					// Prüfen, ob für diese Software ein bekanntes Update-Ziel existiert.
					if bySoftware, ok := updateSw[dev.DeviceName]; ok {
						if targetVersion, found := bySoftware[softwareName]; found {

							comparison.UpdateTargetVersion = targetVersion
							comparison.UpdateAvailable =
								targetVersion != liveVersion
						}
					}

					inst.Software = append(
						inst.Software,
						comparison,
					)
				}
			}

			// Warntext erst nach dem Live-Vergleich erzeugen.
			warnings := make([]string, 0)

			for _, inst := range dev.Instances {
				if inst.StatusText == "" {
					continue
				}

				warnings = append(
					warnings,
					inst.Serialnumber+" ("+inst.StatusText+")",
				)
			}

			dev.InstanceWarningText = strings.Join(warnings, ", ")

		}
	}

	for bi := range blocks {
		block := &blocks[bi]

		block.InstanceGroups = make([]classes.DeviceInstanceDisplayGroup, 0)
		groupIndex := make(map[classes.DeviceInstanceDisplayGroupKey]int)

		for _, dev := range block.DevicesWithSW {
			for _, inst := range dev.Instances {
				softwareKey := buildInstanceSoftwareKey(inst.Software)

				key := classes.DeviceInstanceDisplayGroupKey{
					ProjectBOMID: block.ProjectBOMID,
					DeviceName:   dev.DeviceName,
					DBVersion:    inst.DBDeviceVersion,
					LiveVersion:  inst.LiveDeviceVersion,
					FoundInLive:  inst.FoundInLive,
					SoftwareKey:  softwareKey,
				}

				if index, found := groupIndex[key]; found {
					block.InstanceGroups[index].DeviceCount++
					block.InstanceGroups[index].Serialnumbers = append(
						block.InstanceGroups[index].Serialnumbers,
						inst.Serialnumber,
					)
					continue
				}

				groupIndex[key] = len(block.InstanceGroups)

				targetVersion := ""
				updateAvailable := false

				if version, found := updateDev[dev.DeviceName]; found {
					targetVersion = version

					if inst.FoundInLive {
						updateAvailable = version != inst.LiveDeviceVersion
					} else {
						updateAvailable = version != inst.DBDeviceVersion
					}
				}

				block.InstanceGroups = append(
					block.InstanceGroups,
					classes.DeviceInstanceDisplayGroup{
						ProjectBOMID:  block.ProjectBOMID,
						DeviceName:    dev.DeviceName,
						DBVersion:     inst.DBDeviceVersion,
						LiveVersion:   inst.LiveDeviceVersion,
						FoundInLive:   inst.FoundInLive,
						SoftwareKey:   softwareKey,
						DeviceCount:   1,
						UpdateTargetVersion: targetVersion,
						UpdateAvailable:     updateAvailable,
						Serialnumbers: []string{inst.Serialnumber},

						MostCommonSystemVersion: dev.MostCommonSystemVersion,
						ShortenedSystemVersions: dev.ShortenedSystemVersions,
						IsInvalidSystemVersion:  dev.IsInvalidSystemVersion,

						DBOutdated: inst.FoundInLive &&
							inst.DBDeviceVersion != "" &&
							inst.LiveDeviceVersion != "" &&
							inst.DBDeviceVersion != inst.LiveDeviceVersion,

						Software: append(
							[]classes.InstanceSoftwareUpdateView{},
							inst.Software...,
						),
					},
				)
			}
		}

		for i := range block.InstanceGroups {
			block.InstanceGroups[i].SerialnumberText =
				strings.Join(block.InstanceGroups[i].Serialnumbers, ", ")

		}
	}

	liveLine1 := ""
	liveLine2 := ""

	if liveState != nil {
		liveLine1 = "Live snapshot: " + liveState.CreatedAt
		liveLine2 = "received " + liveState.ReceivedAt
	}

	// Geräte ermitteln, die im Live-Report vorhanden sind,
	// aber zu keiner Geräteinstanz dieses Projekts gehören.
	liveOnlyDevices := make([]classes.LiveOnlyDeviceView, 0)

	if liveState != nil {
		knownProjectSerialnumbers := make(map[string]struct{})

		// Alle Seriennummern sammeln, die laut DB zu diesem Projekt gehören.
		for _, block := range blocks {
			for _, dev := range block.DevicesWithSW {
				for _, inst := range dev.Instances {
					if inst.Serialnumber == "" {
						continue
					}

					knownProjectSerialnumbers[inst.Serialnumber] = struct{}{}
				}
			}
		}

		// Alle Live-Geräte suchen, deren Seriennummer nicht im Projekt existiert.
		for serialnumber, liveDevice := range liveState.DeviceBySerialnumber {
			if _, exists := knownProjectSerialnumbers[serialnumber]; exists {
				continue
			}

			liveOnlyDevices = append(
				liveOnlyDevices,
				classes.LiveOnlyDeviceView{
					Serialnumber:  serialnumber,
					DeviceType:    liveDevice.DeviceType,
					DeviceVersion: liveDevice.DeviceVersion,
					Software: append(
						[]classes.LiveSoftwareState{},
						liveDevice.Software...,
					),
				},
			)
		}

		// Stabile Reihenfolge für die Anzeige.
		sort.Slice(
			liveOnlyDevices,
			func(i, j int) bool {
				return liveOnlyDevices[i].Serialnumber <
					liveOnlyDevices[j].Serialnumber
			},
		)
	}

	ctx.ViewData("liveOnlyDevices", liveOnlyDevices)


	ctx.ViewData("liveLine1", liveLine1)
	ctx.ViewData("liveLine2", liveLine2)

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

func buildInstanceSoftwareKey(
	software []classes.InstanceSoftwareUpdateView,
) string {
	if len(software) == 0 {
		return ""
	}

	parts := make([]string, 0, len(software))

	for _, sw := range software {
		part := sw.SoftwareName +
			"|db=" + sw.DBVersion +
			"|live=" + sw.LiveVersion +
			"|found=" + strconv.FormatBool(sw.FoundInLive)

		parts = append(parts, part)
	}

	sort.Strings(parts)

	return strings.Join(parts, "||")
}

func SMSProjectUpdateUploadLiveReport(ctx iris.Context) {
	projectID, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Invalid project id")
		return
	}

	file, _, err := ctx.FormFile("report")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Missing file 'report'")
		return
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("Failed to read file")
		return
	}

	// hash + size
	sum := sha256.Sum256(b)
	sha := hex.EncodeToString(sum[:])
	size := len(b)

	// parse JSON
	var lr classes.LiveReportV1
	if err := json.Unmarshal(b, &lr); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Invalid JSON")
		return
	}

	if lr.SchemaVersion == "" {
		lr.SchemaVersion = "1"
	}
	if lr.CreatedAt == "" {
		lr.CreatedAt = time.Now().UTC().Format(time.RFC3339)
	}

	// createdAt parse
	createdAt, parseErr := time.Parse(time.RFC3339, lr.CreatedAt)
	if parseErr != nil {
		// akzeptiere auch "YYYY-MM-DD HH:MM:SS"
		if t2, err2 := time.Parse("2006-01-02 15:04:05", lr.CreatedAt); err2 == nil {
			createdAt = t2
		} else {
			createdAt = time.Now().UTC()
		}
	}

	// Determine mapping project_id (priority: update_center_id)
	var updateCenterID *int = lr.UpdateCenterID

	mappedProjectID := 0
	if updateCenterID != nil && *updateCenterID > 0 {
		pid, err := dbprovider.GetDBManager().GetProjectIDByUpdateCenterID(*updateCenterID)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString("Invalid update_center_id (not found)")
			return
		}
		mappedProjectID = pid
	} else if lr.Project != nil && lr.Project.ProjectID != nil {
		mappedProjectID = *lr.Project.ProjectID
	}

	if mappedProjectID <= 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Report must include update_center_id or project.project_id")
		return
	}

	// Optional: URL projectID nicht erzwingen, aber du könntest hier warnen
	_ = projectID

	// Optional system_id (if present)
	var systemID *int
	if lr.System != nil && lr.System.SystemID != nil && *lr.System.SystemID > 0 {
		systemID = lr.System.SystemID
	}

	// report_name
	rn := strings.TrimSpace(ctx.PostValue("report_name"))
	var reportName *string
	if rn != "" {
		reportName = &rn
	}

	// received_by (optional)
	rb := strings.TrimSpace(ctx.PostValue("received_by"))
	var receivedBy *string
	if rb != "" {
		receivedBy = &rb
	}

	payload := string(b)
	format := "device_software_v1"

	// store
	sz := size
	reportID, err := dbprovider.GetDBManager().AddLiveReport(
		mappedProjectID,
		updateCenterID,
		nil,      // projectBOM_id (optional later)
		systemID, // system_id (optional)
		"upload",
		reportName,
		createdAt,
		receivedBy,
		lr.SchemaVersion,
		format,
		payload,
		&sha,
		&sz,
		nil,
	)
	if err != nil {
		log.Println("[LIVE-REPORT] insert failed:", err)
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("DB error saving report")
		return
	}

	// validate + store report items
	err = dbprovider.GetDBManager().ValidateAndStoreLiveReportItems(mappedProjectID, reportID, lr)
	if err != nil {
		log.Println("[LIVE-REPORT] validation failed:", err)
		// Report bleibt gespeichert, auch wenn Validation fehlschlägt
	}

	// back
	ctx.Redirect("/sms_projectUpdates/show/" + strconv.Itoa(mappedProjectID))
}

func SMSProjectUpdateApplyLivePlaceholder(ctx iris.Context) {
	projectID, _ := ctx.Params().GetInt("project_id")

	deviceType := ctx.PostValue("device_type")
	liveVersion := ctx.PostValue("live_device_version")

	// später: hier mappen device_type + version -> device_id und deviceInstances updaten
	ctx.ViewData("error", "TODO: Apply live version to DB. device_type="+deviceType+" live_version="+liveVersion)
	ctx.Redirect("/sms_projectUpdates/show/" + strconv.Itoa(projectID))
}