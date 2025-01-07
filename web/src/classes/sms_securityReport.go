/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_SecurityReport struct {
	ReportID         int    `db:"report_id"`         // Eindeutige ID des Reports
	ReportName       string `db:"report_name"`       // Name des Reports
	ScannerName      string `db:"scanner_name"`      // Name des verwendeten Scanners
	ScannerVersion   string `db:"scanner_version"`   // Version des Scanners
	CreationDate     string `db:"creation_date"`     // Datum der Erstellung des Reports
	UploadDate       string `db:"upload_date"`       // Datum des Uploads
	UploadedBy       string `db:"uploaded_by"`       // Wer hat den Report hochgeladen
	ScanScope        string `db:"scan_scope"`        // Umfang des Scans
	VulnerabilityCount int   `db:"vulnerability_count"` // Anzahl der gefundenen Schwachstellen
	ComponentCount   int    `db:"component_count"`   // Anzahl der gefundenen Subkomponenten
}

func NewSms_SecurityReport(reportID int, reportName string, scannerName string, scannerVersion string, creationDate string, uploadDate string, uploadedBy string, scanScope string, vulnerabilityCount int, componentCount int) *Sms_SecurityReport {
	return &Sms_SecurityReport{ReportID: reportID, ReportName: reportName, ScannerName: scannerName, ScannerVersion: scannerVersion, CreationDate: creationDate, UploadDate: uploadDate, UploadedBy: uploadedBy, ScanScope: scanScope, VulnerabilityCount: vulnerabilityCount, ComponentCount: componentCount}
}