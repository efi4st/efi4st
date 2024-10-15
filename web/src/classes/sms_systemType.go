/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_SystemType struct {
	systemtype_id int `db:"systemtype_id"`
	systemType string `db:"type"`
}

func (s *Sms_SystemType) Systemtype_id() int {
	return s.systemtype_id
}

func (s *Sms_SystemType) SetSystemtype_id(systemtype_id int) {
	s.systemtype_id = systemtype_id
}

func (s *Sms_SystemType) SystemType() string {
	return s.systemType
}

func (s *Sms_SystemType) SetSystemType(systemType string) {
	s.systemType = systemType
}

func NewSms_SystemType(systemtype_id int, systemType string) *Sms_SystemType {
	return &Sms_SystemType{systemtype_id: systemtype_id, systemType: systemType}
}

