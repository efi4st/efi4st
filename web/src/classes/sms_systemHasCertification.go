/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_SystemHasCertification struct {
	system_id int `db:"system_id"`
	certification_id int `db:"certification_id"`
	additionalInfo string `db:"additionalInfo"`
	certification_name string `db:"certification_name"`
	system_name string `db:"system_name"`
	system_version string `db:"system_version"`
}

func (s *Sms_SystemHasCertification) System_id() int {
	return s.system_id
}

func (s *Sms_SystemHasCertification) SetSystem_id(system_id int) {
	s.system_id = system_id
}

func (s *Sms_SystemHasCertification) Certification_id() int {
	return s.certification_id
}

func (s *Sms_SystemHasCertification) SetCertification_id(certification_id int) {
	s.certification_id = certification_id
}

func (s *Sms_SystemHasCertification) AdditionalInfo() string {
	return s.additionalInfo
}

func (s *Sms_SystemHasCertification) SetAdditionalInfo(additionalInfo string) {
	s.additionalInfo = additionalInfo
}

func (s *Sms_SystemHasCertification) Certification_name() string {
	return s.certification_name
}

func (s *Sms_SystemHasCertification) SetCertification_name(certification_name string) {
	s.certification_name = certification_name
}

func (s *Sms_SystemHasCertification) System_name() string {
	return s.system_name
}

func (s *Sms_SystemHasCertification) SetSystem_name(system_name string) {
	s.system_name = system_name
}

func (s *Sms_SystemHasCertification) System_version() string {
	return s.system_version
}

func (s *Sms_SystemHasCertification) SetSystem_version(system_version string) {
	s.system_version = system_version
}

func NewSms_SystemHasCertification(system_id int, certification_id int, additionalInfo string, certification_name string, system_name string, system_version string) *Sms_SystemHasCertification {
	return &Sms_SystemHasCertification{system_id: system_id, certification_id: certification_id, additionalInfo: additionalInfo, certification_name: certification_name, system_name: system_name, system_version: system_version}
}

