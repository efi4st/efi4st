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
