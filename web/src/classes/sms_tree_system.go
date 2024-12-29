/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_Tree_Component struct {
	name string
	version string
}

func (s *Sms_Tree_Component) Name() string {
	return s.name
}

func (s *Sms_Tree_Component) SetName(name string) {
	s.name = name
}

func (s *Sms_Tree_Component) Version() string {
	return s.version
}

func (s *Sms_Tree_Component) SetVersion(version string) {
	s.version = version
}

func NewSms_Tree_Component(name string, version string) *Sms_Tree_Component {
	return &Sms_Tree_Component{name: name, version: version}
}

type Sms_Tree_Application struct {
	name string
	version string
	components []Sms_Tree_Component
}

func (s *Sms_Tree_Application) Name() string {
	return s.name
}

func (s *Sms_Tree_Application) SetName(name string) {
	s.name = name
}

func (s *Sms_Tree_Application) Version() string {
	return s.version
}

func (s *Sms_Tree_Application) SetVersion(version string) {
	s.version = version
}

func (s *Sms_Tree_Application) Components() []Sms_Tree_Component {
	return s.components
}

func (s *Sms_Tree_Application) SetComponents(components []Sms_Tree_Component) {
	s.components = components
}

func NewSms_Tree_Application(name string, version string, components []Sms_Tree_Component) *Sms_Tree_Application {
	return &Sms_Tree_Application{name: name, version: version, components: components}
}

type Sms_Tree_Device struct {
	name string
	version string
	applications []Sms_Tree_Application
}

func (s *Sms_Tree_Device) Name() string {
	return s.name
}

func (s *Sms_Tree_Device) SetName(name string) {
	s.name = name
}

func (s *Sms_Tree_Device) Version() string {
	return s.version
}

func (s *Sms_Tree_Device) SetVersion(version string) {
	s.version = version
}

func (s *Sms_Tree_Device) Applications() []Sms_Tree_Application {
	return s.applications
}

func (s *Sms_Tree_Device) SetApplications(applications []Sms_Tree_Application) {
	s.applications = applications
}

func NewSms_Tree_Device(name string, version string, applications []Sms_Tree_Application) *Sms_Tree_Device {
	return &Sms_Tree_Device{name: name, version: version, applications: applications}
}

type Sms_Tree_System struct {
	name string
	version string
	devices []Sms_Tree_Device
}

func (s *Sms_Tree_System) Name() string {
	return s.name
}

func (s *Sms_Tree_System) SetName(name string) {
	s.name = name
}

func (s *Sms_Tree_System) Version() string {
	return s.version
}

func (s *Sms_Tree_System) SetVersion(version string) {
	s.version = version
}

func (s *Sms_Tree_System) Devices() []Sms_Tree_Device {
	return s.devices
}

func (s *Sms_Tree_System) SetDevices(devices []Sms_Tree_Device) {
	s.devices = devices
}

func NewSms_Tree_System(name string, version string, devices []Sms_Tree_Device) *Sms_Tree_System {
	return &Sms_Tree_System{name: name, version: version, devices: devices}
}
