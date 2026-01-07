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

	DevicesWithSW    []DeviceUpdateView
	AvailableUpdates []Sms_UpdateDetails
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