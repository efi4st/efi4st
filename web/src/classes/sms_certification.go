/**
 * Author:    Admiral Helmut
 * Created:   29.10.2024
 *
 * (C)
 **/

package classes

type Sms_Certification struct {
	certification_id int `db:"certification_id"`
	name string `db:"name"`
	date string `db:"date"`
	description string `db:"description"`
}

func (s *Sms_Certification) Certification_id() int {
	return s.certification_id
}

func (s *Sms_Certification) SetCertification_id(certification_id int) {
	s.certification_id = certification_id
}

func (s *Sms_Certification) Name() string {
	return s.name
}

func (s *Sms_Certification) SetName(name string) {
	s.name = name
}

func (s *Sms_Certification) Date() string {
	return s.date
}

func (s *Sms_Certification) SetDate(date string) {
	s.date = date
}

func (s *Sms_Certification) Description() string {
	return s.description
}

func (s *Sms_Certification) SetDescription(description string) {
	s.description = description
}

func NewSms_Certification(name string, date string, description string) *Sms_Certification {
	return &Sms_Certification{name: name, date: date, description: description}
}

func NewSms_CertificationFromDB(certification_id int, name string, date string, description string) *Sms_Certification {
	return &Sms_Certification{certification_id: certification_id, name: name, date: date, description: description}
}

