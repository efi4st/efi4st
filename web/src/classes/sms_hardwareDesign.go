/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_HardwareDesign struct {
	HardwareDesignID int
	Name             string
	Version          string
	Date             string
	Description      string
	Image            []byte
	Author           string
	IsApproved       bool
	RevisionNote     string
	DocumentNumber   string
	AdditionalInfo   string // aus Mapping-Tabelle
	// View-only
	ImageBase64       string // ✅ zum Anzeigen im HTML
}