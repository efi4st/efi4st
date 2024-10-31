/**
 * Author:    Admiral Helmut
 * Created:   29.10.2024
 *
 * (C)
 **/

package classes

type Sms_Component struct {
	component_id int `db:"component_id"`
	name string `db:"name"`
	componentType string `db:"componentType"`
	version string `db:"version"`
	date string `db:"date"`
	license string `db:"license"`
	thirdParty bool `db:"thirdParty"`
	releaseNote string `db:"releaseNote"`
}

func (s *Sms_Component) Component_id() int {
	return s.component_id
}

func (s *Sms_Component) SetComponent_id(component_id int) {
	s.component_id = component_id
}

func (s *Sms_Component) Name() string {
	return s.name
}

func (s *Sms_Component) SetName(name string) {
	s.name = name
}

func (s *Sms_Component) ComponentType() string {
	return s.componentType
}

func (s *Sms_Component) SetComponentType(componentType string) {
	s.componentType = componentType
}

func (s *Sms_Component) Version() string {
	return s.version
}

func (s *Sms_Component) SetVersion(version string) {
	s.version = version
}

func (s *Sms_Component) Date() string {
	return s.date
}

func (s *Sms_Component) SetDate(date string) {
	s.date = date
}

func (s *Sms_Component) License() string {
	return s.license
}

func (s *Sms_Component) SetLicense(license string) {
	s.license = license
}

func (s *Sms_Component) ThirdParty() bool {
	return s.thirdParty
}

func (s *Sms_Component) SetThirdParty(thirdParty bool) {
	s.thirdParty = thirdParty
}

func (s *Sms_Component) ReleaseNote() string {
	return s.releaseNote
}

func (s *Sms_Component) SetReleaseNote(releaseNote string) {
	s.releaseNote = releaseNote
}

func NewSms_Component(name string, componentType string, version string, date string, license string, thirdParty bool, releaseNote string) *Sms_Component {
	return &Sms_Component{name: name, componentType: componentType, version: version, date: date, license: license, thirdParty: thirdParty, releaseNote: releaseNote}
}

func NewSms_ComponentFromDB(component_id int, name string, componentType string, version string, date string, license string, thirdParty bool, releaseNote string) *Sms_Component {
	return &Sms_Component{component_id: component_id, name: name, componentType: componentType, version: version, date: date, license: license, thirdParty: thirdParty, releaseNote: releaseNote}
}
