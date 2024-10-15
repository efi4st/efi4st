/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_DeviceType struct {
	devicetype_id int `db:"devicetype_id"`
	deviceType string `db:"deviceType"`
}

func (s *Sms_DeviceType) Devicetype_id() int {
	return s.devicetype_id
}

func (s *Sms_DeviceType) SetDevicetype_id(devicetype_id int) {
	s.devicetype_id = devicetype_id
}

func (s *Sms_DeviceType) DeviceType() string {
	return s.deviceType
}

func (s *Sms_DeviceType) SetDeviceType(deviceType string) {
	s.deviceType = deviceType
}

func NewSms_DeviceType(devicetype_id int, deviceType string) *Sms_DeviceType {
	return &Sms_DeviceType{devicetype_id: devicetype_id, deviceType: deviceType}
}
