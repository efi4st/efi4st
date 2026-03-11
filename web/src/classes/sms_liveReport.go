/**
 * Author:    Admiral Helmut
 * Created:   29.10.2024
 *
 * (C)
 **/

package classes



type LiveReportV1 struct {
	SchemaVersion  string `json:"schema_version"`
	CreatedAt      string `json:"created_at"`
	UpdateCenterID *int   `json:"update_center_id,omitempty"`

	Project *struct {
		ProjectID   *int    `json:"project_id,omitempty"`
		PlantNumber *string `json:"plant_number,omitempty"`
	} `json:"project,omitempty"`

	System *struct {
		SystemID      *int    `json:"system_id,omitempty"`
		SystemType    *string `json:"system_type,omitempty"`
		SystemVersion *string `json:"system_version,omitempty"`
	} `json:"system,omitempty"`

	Devices []struct {
		Serialnumber  string `json:"serialnumber"`
		DeviceType    string `json:"device_type"`
		DeviceVersion string `json:"device_version"`
		Software      []struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"software"`
	} `json:"devices"`
}

type Sms_LiveReportRow struct {
	ReportID       int
	CreatedAt      string
	ReceivedAt     string
	SchemaVersion  string
	ReportFormat   string
	PayloadJSON    string
}

// Aggregiert pro DeviceType, wie deine Update-Tabelle es darstellt.
type LiveProjectState struct {
	CreatedAt  string
	ReceivedAt string

	// device_type -> most common device_version
	DeviceVersionByType map[string]string

	// device_type -> software_name -> most common version
	SoftwareVersionByType map[string]map[string]string
}


type Sms_LiveReportItem struct {
	ItemID int
	ReportID int
	ProjectID int

	DeviceInstanceID *int
	Serialnumber string

	LiveDeviceType string
	LiveDeviceVersion string

	MatchStatus string
	MatchedDeviceID *int
	MismatchSummary *string

	CreatedAt string
}