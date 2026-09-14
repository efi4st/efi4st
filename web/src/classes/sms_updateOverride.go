/**
 * Author:    Admiral Helmut
 * Created:   13.10.2024
 *
 * (C)
 **/

package classes

type Sms_DeviceInstanceSoftwareOverride struct {
	DeviceInstanceID int     `db:"deviceInstance_id"`
	SoftwaretypeID   int     `db:"softwaretype_id"`
	SoftwareID       int     `db:"software_id"`

	OccurredAt string  `db:"occurred_at"`
	RecordedAt string  `db:"recorded_at"`
	RecordedBy *string `db:"recorded_by"`

	Source string `db:"source"`

	LiveReportID      *int `db:"live_report_id"`
	UpdateExecutionID *int `db:"update_execution_id"`

	Note *string `db:"note"`

	// Nur für den Lesezugriff per JOIN:
	SoftwareName    string `db:"software_name"`
	SoftwareVersion string `db:"software_version"`
}


type Sms_DeviceInstanceSoftwareOverrideHistory struct {
	HistoryID int `db:"history_id"`

	DeviceInstanceID int `db:"deviceInstance_id"`
	SoftwaretypeID   int `db:"softwaretype_id"`

	OldSoftwareID *int `db:"old_software_id"`
	NewSoftwareID *int `db:"new_software_id"`

	Action string `db:"action"`

	OccurredAt string  `db:"occurred_at"`
	RecordedAt string  `db:"recorded_at"`
	RecordedBy *string `db:"recorded_by"`

	Source string `db:"source"`

	LiveReportID      *int `db:"live_report_id"`
	UpdateExecutionID *int `db:"update_execution_id"`

	Note *string `db:"note"`
}