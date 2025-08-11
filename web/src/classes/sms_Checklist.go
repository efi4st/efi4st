/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_ChecklistTemplate struct {
	ChecklistTemplateID int
	Name                string
	Description         string
}

type Sms_ChecklistTemplateItem struct {
	ChecklistTemplateItemID int
	ChecklistTemplateID     int
	CheckDefinitionID       *int   // optional
	ArtefactTypeID          *int   // optional
	TargetScope             string // e.g., "project", "device", "instance", "system"
	ExpectedValue           string
	Optional                bool
}

type Sms_ChecklistInstance struct {
	ChecklistInstanceID  int
	ChecklistTemplateID  int
	ProjectID            *int
	DeviceID             *int
	GeneratedAt          string
	GeneratedBy          string
	Status               string
	TemplateName         string
}

type Sms_ChecklistItemInstance struct {
	ChecklistItemInstanceID int
	ChecklistInstanceID     int
	ChecklistTemplateItemID int
	TargetObjectID          int
	TargetObjectType        string
	IsOK                    *bool
	ActualValue             string
	Comment                 string
	IsOKStr 				string // Hilfsfeld für das Template
	ExpectedValue 			string
	// CheckDefinition (optional)
	CheckDefinitionID   *int
	CheckDefName        string
	CheckDefDescription string
	CheckDefExplanation string
	CheckDefExpected    string
	DeviceTypeID        *int
	DeviceTypeName      string
	ApplicableVersions  string

	// Gematchte Device Instances im Projekt
	MatchingDevices []MatchingDevice

	// 🆕 Device-Kontext:
	DeviceContextTypeName string // z.B. "Router"
	DeviceContextVersion  string // z.B. "1.2.3"
	AppliesToThisDevice   *bool  // nil = nicht geprüft, true/false = Ergebnis
	AppliesToThisDeviceStr string // "true" | "false" | "none"
}

type MatchingDevice struct {
	DeviceInstanceID int
	Serialnumber     string
	DeviceVersion    string
	DeviceTypeName   string
}

// classes – falls noch nicht vorhanden
type DeviceBasic struct {
	DeviceID     int
	DeviceTypeID int
	DeviceType   string
	Version      string
}