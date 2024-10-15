/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_AffectedDeviceInstancesAndProjects struct {
	deviceInstance_id int `db:"deviceInstance_id"`
	devicetype string `db:"type"`
	project_id int `db:"project_id"`
	version string `db:"version"`
}

func (s *Sms_AffectedDeviceInstancesAndProjects) DeviceInstance_id() int {
	return s.deviceInstance_id
}

func (s *Sms_AffectedDeviceInstancesAndProjects) SetDeviceInstance_id(deviceInstance_id int) {
	s.deviceInstance_id = deviceInstance_id
}

func (s *Sms_AffectedDeviceInstancesAndProjects) Devicetype() string {
	return s.devicetype
}

func (s *Sms_AffectedDeviceInstancesAndProjects) SetDevicetype(devicetype string) {
	s.devicetype = devicetype
}

func (s *Sms_AffectedDeviceInstancesAndProjects) Project_id() int {
	return s.project_id
}

func (s *Sms_AffectedDeviceInstancesAndProjects) SetProject_id(project_id int) {
	s.project_id = project_id
}

func (s *Sms_AffectedDeviceInstancesAndProjects) Version() string {
	return s.version
}

func (s *Sms_AffectedDeviceInstancesAndProjects) SetVersion(version string) {
	s.version = version
}

func NewSms_AffectedDeviceInstancesAndProjects(deviceInstance_id int, devicetype string, project_id int, version string) *Sms_AffectedDeviceInstancesAndProjects {
	return &Sms_AffectedDeviceInstancesAndProjects{deviceInstance_id: deviceInstance_id, devicetype: devicetype, project_id: project_id, version: version}
}
