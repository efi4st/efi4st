/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_ProjectBOM struct {
	ProjectBOMID           int
	ProjectID              int
	SystemID               int
	HardwareDesignID       int
	HardwareDesignVariantID int
	OrderNumber            string
	AdditionalInfo         string
	AssignedAt             string
}

type Sms_ProjectBOMView struct {
	ProjectBOMID int
	ProjectID    int
	SystemID     int

	OrderNumber    string
	AdditionalInfo string
	AssignedAt     string

	// System
	SystemType    string
	SystemVersion string

	// Hardware-Design
	HardwareDesignID       int
	HardwareDesignName     string
	HardwareDesignVersion  string

	// Variante
	VariantID   int
	VariantCode string
	VariantName string
	VariantSpec string

	DeviceList  []Sms_DeviceInstancePBOMView
	DeviceCount int
}