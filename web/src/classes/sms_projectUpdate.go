/**
 * Author:    Admiral Helmut
 * Created:   14.02.2025
 *
 * (C)
 **/

package classes

type DeviceInstanceDisplayGroup struct {
	ProjectBOMID     int
	DeviceName       string
	DBVersion        string
	LiveVersion      string
	FoundInLive      bool
	SoftwareKey      string
	DeviceCount      int
	Serialnumbers    []string
	SerialnumberText string

	MostCommonSystemVersion string
	ShortenedSystemVersions string
	IsInvalidSystemVersion  bool

	DBOutdated          bool
	UpdateAvailable     bool
	UpdateTargetVersion string

	Software []InstanceSoftwareUpdateView
}

type DeviceInstanceDisplayGroupKey struct {
	ProjectBOMID int
	DeviceName   string
	DBVersion    string
	LiveVersion  string
	FoundInLive  bool
	SoftwareKey string
}

type SystemUpdateBlock struct {
	ProjectBOMID   int
	SystemID       int
	SystemTypeName string
	SystemVersion  string

	IsClean bool

	DevicesWithSW    []DeviceUpdateView
	InstanceGroups   []DeviceInstanceDisplayGroup
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
	DBOutdated bool
	UpdateAvailable bool
	UpdateTargetVersion string // optional, wenn du es anzeigen willst (können wir nutzen)
	Instances []DeviceInstanceUpdateView
	CollapseID string
	SerialnumberText string
	InstanceWarningText string
}


type SoftwareUpdateView struct {
	SoftwareName    string
	SoftwareVersion string
	UpdateVersion   string
	DBOutdated bool
	UpdateAvailable bool
	UpdateTargetVersion string
	ShortenedSystemVersions string
}

type DeviceInstanceUpdateView struct {
	DeviceInstanceID int
	Serialnumber      string

	DBDeviceVersion   string
	LiveDeviceVersion string

	FoundInLive bool
	DBLiveMatch bool
	StatusText string

	Software []InstanceSoftwareUpdateView
}