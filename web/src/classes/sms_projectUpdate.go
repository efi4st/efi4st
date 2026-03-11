/**
 * Author:    Admiral Helmut
 * Created:   14.02.2025
 *
 * (C)
 **/

package classes

type SystemUpdateBlock struct {
	ProjectBOMID   int
	SystemID       int
	SystemTypeName string
	SystemVersion  string

	IsClean bool

	DevicesWithSW    []DeviceUpdateView
	AvailableUpdates []Sms_UpdateDetails
}


type DeviceUpdateView struct {
	DeviceName              string
	DeviceVersion           string
	UpdateVersion           string
	DeviceCount             int
	SystemVersions           []string
	IsInvalidSystemVersion  bool
	MostCommonSystemVersion string
	ShortenedSystemVersions string
	SoftwareList            []SoftwareUpdateView
	LiveDeviceVersion string
	DBOutdated bool
	UpdateAvailable bool
	UpdateTargetVersion string // optional, wenn du es anzeigen willst (können wir nutzen)
}


type SoftwareUpdateView struct {
	SoftwareName    string
	SoftwareVersion string
	UpdateVersion   string
	LiveSoftwareVersion string
	DBOutdated bool
	UpdateAvailable bool
	UpdateTargetVersion string
	ShortenedSystemVersions string
}

