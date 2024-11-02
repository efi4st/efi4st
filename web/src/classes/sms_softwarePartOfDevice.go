/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_SoftwarePartOfDevice struct {
	device_id int `db:"device_id"`
	software_id int `db:"software_id"`
	additionalInfo string `db:"additionalInfo"`
	name string `db:"name"`
	version string `db:"version"`
}

func (s *Sms_SoftwarePartOfDevice) Device_id() int {
	return s.device_id
}

func (s *Sms_SoftwarePartOfDevice) SetDevice_id(device_id int) {
	s.device_id = device_id
}

func (s *Sms_SoftwarePartOfDevice) Software_id() int {
	return s.software_id
}

func (s *Sms_SoftwarePartOfDevice) SetSoftware_id(software_id int) {
	s.software_id = software_id
}

func (s *Sms_SoftwarePartOfDevice) AdditionalInfo() string {
	return s.additionalInfo
}

func (s *Sms_SoftwarePartOfDevice) SetAdditionalInfo(additionalInfo string) {
	s.additionalInfo = additionalInfo
}

func (s *Sms_SoftwarePartOfDevice) Name() string {
	return s.name
}

func (s *Sms_SoftwarePartOfDevice) SetName(name string) {
	s.name = name
}

func (s *Sms_SoftwarePartOfDevice) Version() string {
	return s.version
}

func (s *Sms_SoftwarePartOfDevice) SetVersion(version string) {
	s.version = version
}

func NewSms_SoftwarePartOfDevice(device_id int, software_id int, additionalInfo string, name string, version string) *Sms_SoftwarePartOfDevice {
	return &Sms_SoftwarePartOfDevice{device_id: device_id, software_id: software_id, additionalInfo: additionalInfo, name: name, version: version}
}
