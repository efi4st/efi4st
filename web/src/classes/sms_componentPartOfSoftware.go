/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_ComponentPartOfSoftware struct {
	software_id int `db:"software_id"`
	component_id int `db:"component_id"`
	additionalInfo string `db:"additionalInfo"`
	name string `db:"name"`
	version string `db:"version"`
}

func (s *Sms_ComponentPartOfSoftware) Software_id() int {
	return s.software_id
}

func (s *Sms_ComponentPartOfSoftware) SetSoftware_id(software_id int) {
	s.software_id = software_id
}

func (s *Sms_ComponentPartOfSoftware) Component_id() int {
	return s.component_id
}

func (s *Sms_ComponentPartOfSoftware) SetComponent_id(component_id int) {
	s.component_id = component_id
}

func (s *Sms_ComponentPartOfSoftware) AdditionalInfo() string {
	return s.additionalInfo
}

func (s *Sms_ComponentPartOfSoftware) SetAdditionalInfo(additionalInfo string) {
	s.additionalInfo = additionalInfo
}

func (s *Sms_ComponentPartOfSoftware) Name() string {
	return s.name
}

func (s *Sms_ComponentPartOfSoftware) SetName(name string) {
	s.name = name
}

func (s *Sms_ComponentPartOfSoftware) Version() string {
	return s.version
}

func (s *Sms_ComponentPartOfSoftware) SetVersion(version string) {
	s.version = version
}

func NewSms_ComponentPartOfSoftwareFromDB(software_id int, component_id int, additionalInfo string, name string, version string) *Sms_ComponentPartOfSoftware {
	return &Sms_ComponentPartOfSoftware{software_id: software_id, component_id: component_id, additionalInfo: additionalInfo, name: name, version: version}
}
