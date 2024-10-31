/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_SoftwareType struct {
	softwaretype_id int `db:"softwaretype_id"`
	typeName string `db:"typeName"`
}

func (s *Sms_SoftwareType) Softwaretype_id() int {
	return s.softwaretype_id
}

func (s *Sms_SoftwareType) SetSoftwaretype_id(softwaretype_id int) {
	s.softwaretype_id = softwaretype_id
}

func (s *Sms_SoftwareType) TypeName() string {
	return s.typeName
}

func (s *Sms_SoftwareType) SetTypeName(typeName string) {
	s.typeName = typeName
}

func NewSms_SoftwareType(typeName string) *Sms_SoftwareType {
	return &Sms_SoftwareType{typeName: typeName}
}

func NewSms_SoftwareTypeFromDB(softwaretype_id int, typeName string) *Sms_SoftwareType {
	return &Sms_SoftwareType{softwaretype_id: softwaretype_id, typeName: typeName}
}
