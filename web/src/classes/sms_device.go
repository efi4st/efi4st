/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_Device struct {
	device_id int `db:"device_id"`
	devicetype_id string `db:"devicetype_id"`
	version string `db:"version"`
	date string `db:"date"`
}

func (s *Sms_Device) Device_id() int {
	return s.device_id
}

func (s *Sms_Device) SetDevice_id(device_id int) {
	s.device_id = device_id
}

func (s *Sms_Device) Devicetype_id() string {
	return s.devicetype_id
}

func (s *Sms_Device) SetDevicetype_id(devicetype_id string) {
	s.devicetype_id = devicetype_id
}

func (s *Sms_Device) Version() string {
	return s.version
}

func (s *Sms_Device) SetVersion(version string) {
	s.version = version
}

func (s *Sms_Device) Date() string {
	return s.date
}

func (s *Sms_Device) SetDate(date string) {
	s.date = date
}

func NewSms_Device(devicetype_id string, version string, date string) *Sms_Device {
	return &Sms_Device{devicetype_id: devicetype_id, version: version, date: date}
}

func NewSms_DeviceFromDB(device_id int, devicetype_id string, version string, date string) *Sms_Device {
	return &Sms_Device{device_id: device_id, devicetype_id: devicetype_id, version: version, date: date}
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
	SoftwareID       int
	SoftwareName     string
	SoftwareVersion  string
	UpdateVersion    string  // Neues Feld für Software-Update-Version
}