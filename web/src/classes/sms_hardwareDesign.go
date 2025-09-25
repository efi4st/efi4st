/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_HardwareDesign struct {
	HardwareDesignID     int
	Name                 string
	Version              string
	Date                 string
	Description          string
	Image                []byte
	Author               string
	IsApproved           bool
	RevisionNote         string
	DocumentNumber       string

	// aus Mapping-Tabelle (sms_hardwaredesignPartOfSystem):
	AdditionalInfo       string
	IsDefault            bool
	CompatibilityStatus  string

	// View-only
	ImageBase64          string
}

type Sms_HardwareDesignVariant struct {
	HardwareDesignVariantID int
	HardwareDesignID        int
	Code                    string
	Name                    string
	Description             string
	Spec                    string // JSON als String (DB-Spalte ist JSON, MySQL validiert)
	IsActive                bool
	CreatedAt               string
}