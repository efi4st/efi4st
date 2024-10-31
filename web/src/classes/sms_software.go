/**
 * Author:    Admiral Helmut
 * Created:   29.10.2024
 *
 * (C)
 **/

package classes

type Sms_Software struct {
	software_id int `db:"software_id"`
	softwaretype_id int `db:"softwaretype_id"`
	version string `db:"version"`
	date string `db:"date"`
	license string `db:"license"`
	thirdParty bool `db:"thirdParty"`
	releaseNote string `db:"releaseNote"`
	typeName string `db:"typeName"`
}

func (s *Sms_Software) Software_id() int {
	return s.software_id
}

func (s *Sms_Software) SetSoftware_id(software_id int) {
	s.software_id = software_id
}

func (s *Sms_Software) Softwaretype_id() int {
	return s.softwaretype_id
}

func (s *Sms_Software) SetSoftwaretype_id(softwaretype_id int) {
	s.softwaretype_id = softwaretype_id
}

func (s *Sms_Software) Version() string {
	return s.version
}

func (s *Sms_Software) SetVersion(version string) {
	s.version = version
}

func (s *Sms_Software) Date() string {
	return s.date
}

func (s *Sms_Software) SetDate(date string) {
	s.date = date
}

func (s *Sms_Software) License() string {
	return s.license
}

func (s *Sms_Software) SetLicense(license string) {
	s.license = license
}

func (s *Sms_Software) ThirdParty() bool {
	return s.thirdParty
}

func (s *Sms_Software) SetThirdParty(thirdParty bool) {
	s.thirdParty = thirdParty
}

func (s *Sms_Software) ReleaseNote() string {
	return s.releaseNote
}

func (s *Sms_Software) SetReleaseNote(releaseNote string) {
	s.releaseNote = releaseNote
}

func (s *Sms_Software) TypeName() string {
	return s.typeName
}

func (s *Sms_Software) SetTypeName(typeName string) {
	s.typeName = typeName
}

func NewSms_Software(softwaretype_id int, version string, date string, license string, thirdParty bool, releaseNote string, typeName string) *Sms_Software {
	return &Sms_Software{softwaretype_id: softwaretype_id, version: version, date: date, license: license, thirdParty: thirdParty, releaseNote: releaseNote, typeName: typeName}
}

func NewSms_SoftwareFromDB(software_id int, softwaretype_id int, version string, date string, license string, thirdParty bool, releaseNote string, typeName string) *Sms_Software {
	return &Sms_Software{software_id: software_id, softwaretype_id: softwaretype_id, version: version, date: date, license: license, thirdParty: thirdParty, releaseNote: releaseNote, typeName: typeName}
}
