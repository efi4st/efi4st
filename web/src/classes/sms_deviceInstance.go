/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_DeviceInstance struct {
	deviceInstance_id int `db:"deviceInstance_id"`
	project_id int `db:"project_id"`
	device_id int `db:"device_id"`
	serialnumber string `db:"serialnumber"`
	provisioner string `db:"provisioner"`
	configuration string `db:"configuration"`
	date string `db:"date"`
	projectName string `db:"name"`
	deviceType string `db:"type"`
	deviceVersion string `db:"version"`

}

func (s *Sms_DeviceInstance) DeviceInstance_id() int {
	return s.deviceInstance_id
}

func (s *Sms_DeviceInstance) SetDeviceInstance_id(deviceInstance_id int) {
	s.deviceInstance_id = deviceInstance_id
}

func (s *Sms_DeviceInstance) Project_id() int {
	return s.project_id
}

func (s *Sms_DeviceInstance) SetProject_id(project_id int) {
	s.project_id = project_id
}

func (s *Sms_DeviceInstance) Device_id() int {
	return s.device_id
}

func (s *Sms_DeviceInstance) SetDevice_id(device_id int) {
	s.device_id = device_id
}

func (s *Sms_DeviceInstance) Serialnumber() string {
	return s.serialnumber
}

func (s *Sms_DeviceInstance) SetSerialnumber(serialnumber string) {
	s.serialnumber = serialnumber
}

func (s *Sms_DeviceInstance) Provisioner() string {
	return s.provisioner
}

func (s *Sms_DeviceInstance) SetProvisioner(provisioner string) {
	s.provisioner = provisioner
}

func (s *Sms_DeviceInstance) Configuration() string {
	return s.configuration
}

func (s *Sms_DeviceInstance) SetConfiguration(configuration string) {
	s.configuration = configuration
}

func (s *Sms_DeviceInstance) Date() string {
	return s.date
}

func (s *Sms_DeviceInstance) SetDate(date string) {
	s.date = date
}

func (s *Sms_DeviceInstance) ProjectName() string {
	return s.projectName
}

func (s *Sms_DeviceInstance) SetProjectName(projectName string) {
	s.projectName = projectName
}

func (s *Sms_DeviceInstance) DeviceType() string {
	return s.deviceType
}

func (s *Sms_DeviceInstance) SetDeviceType(deviceType string) {
	s.deviceType = deviceType
}

func (s *Sms_DeviceInstance) DeviceVersion() string {
	return s.deviceVersion
}

func (s *Sms_DeviceInstance) SetDeviceVersion(deviceVersion string) {
	s.deviceVersion = deviceVersion
}

func NewSms_DeviceInstance(project_id int, device_id int, serialnumber string, provisioner string, configuration string, date string, projectName string, deviceType string, deviceVersion string) *Sms_DeviceInstance {
	return &Sms_DeviceInstance{project_id: project_id, device_id: device_id, serialnumber: serialnumber, provisioner: provisioner, configuration: configuration, date: date, projectName: projectName, deviceType: deviceType, deviceVersion: deviceVersion}
}

func NewSms_DeviceInstanceFromDB(deviceInstance_id int, project_id int, device_id int, serialnumber string, provisioner string, configuration string, date string, projectName string, deviceType string, deviceVersion string) *Sms_DeviceInstance {
	return &Sms_DeviceInstance{deviceInstance_id: deviceInstance_id, project_id: project_id, device_id: device_id, serialnumber: serialnumber, provisioner: provisioner, configuration: configuration, date: date, projectName: projectName, deviceType: deviceType, deviceVersion: deviceVersion}
}

