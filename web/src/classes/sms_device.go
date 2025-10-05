/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_Device struct {
	Device_id int `db:"device_id"`
	Devicetype_id string `db:"devicetype_id"`
	Version string `db:"version"`
	Date string `db:"date"`
}

func NewSms_Device(Devicetype_id string, Version string, Date string) *Sms_Device {
	return &Sms_Device{Devicetype_id: Devicetype_id, Version: Version, Date: Date}
}

func NewSms_DeviceFromDB(Device_id int, Devicetype_id string, Version string, Date string) *Sms_Device {
	return &Sms_Device{Device_id: Device_id, Devicetype_id: Devicetype_id, Version: Version, Date: Date}
}

// DeviceSoftwareInfo speichert Informationen zu einem Gerät und der zugehörigen Software.
// DeviceSoftwareInfo - Information über das Gerät und seine Software
type DeviceSoftwareInfo struct {
	DeviceID                int
	DeviceName              string
	DeviceVersion           string
	DeviceCount             int
	SoftwareList            []SoftwareInfo
	SystemVersions          []string
	MostCommonSystemVersion string
	IsInvalidSystemVersion  bool
	ShortenedSystemVersions string
	UpdateVersion           string  // Neues Feld für die Update-Version
}

// SoftwareInfo - Information über die Software eines Geräts
type SoftwareInfo struct {
	SoftwareID              int
	SoftwareName            string
	SoftwareVersion         string
	UpdateVersion           string
	ShortenedSystemVersions string // 👈 HINZUGEFÜGT
}

// Minimaler DTO für die Device-Auswahlliste
type Sms_DeviceCatalogMinimal struct {
	DeviceID      int
	DeviceType    string
	DeviceVersion string
}