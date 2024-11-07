/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_DevicePartOfSystem struct {
	system_id int `db:"system_id"`
	device_id int `db:"device_id"`
	additionalInfo string `db:"additionalInfo"`
	name string `db:"type"`
	version string `db:"version"`
}

func (s *Sms_DevicePartOfSystem) System_id() int {
	return s.system_id
}

func (s *Sms_DevicePartOfSystem) SetSystem_id(system_id int) {
	s.system_id = system_id
}

func (s *Sms_DevicePartOfSystem) Device_id() int {
	return s.device_id
}

func (s *Sms_DevicePartOfSystem) SetDevice_id(device_id int) {
	s.device_id = device_id
}

func (s *Sms_DevicePartOfSystem) AdditionalInfo() string {
	return s.additionalInfo
}

func (s *Sms_DevicePartOfSystem) SetAdditionalInfo(additionalInfo string) {
	s.additionalInfo = additionalInfo
}

func (s *Sms_DevicePartOfSystem) Name() string {
	return s.name
}

func (s *Sms_DevicePartOfSystem) SetName(name string) {
	s.name = name
}

func (s *Sms_DevicePartOfSystem) Version() string {
	return s.version
}

func (s *Sms_DevicePartOfSystem) SetVersion(version string) {
	s.version = version
}

func NewSms_DevicePartOfSystem(system_id int, device_id int, additionalInfo string, name string, version string) *Sms_DevicePartOfSystem {
	return &Sms_DevicePartOfSystem{system_id: system_id, device_id: device_id, additionalInfo: additionalInfo, name: name, version: version}
}
