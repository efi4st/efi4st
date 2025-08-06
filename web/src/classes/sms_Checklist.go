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
}
