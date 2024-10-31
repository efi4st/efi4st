/**
 * Author:    Admiral Helmut
 * Created:   13.10.2024
 *
 * (C)
 **/

package classes

type Sms_ReleaseNote struct {
	releasenote_id int `db:"releasenote_id"`
	device_id int `db:"device_id"`
	releaseNoteType string `db:"type"`
	date string `db:"date"`
	details string `db:"details"`
}

func (s *Sms_ReleaseNote) Releasenote_id() int {
	return s.releasenote_id
}

func (s *Sms_ReleaseNote) SetReleasenote_id(releasenote_id int) {
	s.releasenote_id = releasenote_id
}

func (s *Sms_ReleaseNote) Device_id() int {
	return s.device_id
}

func (s *Sms_ReleaseNote) SetDevice_id(device_id int) {
	s.device_id = device_id
}

func (s *Sms_ReleaseNote) ReleaseNoteType() string {
	return s.releaseNoteType
}

func (s *Sms_ReleaseNote) SetReleaseNoteType(releaseNoteType string) {
	s.releaseNoteType = releaseNoteType
}

func (s *Sms_ReleaseNote) Date() string {
	return s.date
}

func (s *Sms_ReleaseNote) SetDate(date string) {
	s.date = date
}

func (s *Sms_ReleaseNote) Details() string {
	return s.details
}

func (s *Sms_ReleaseNote) SetDetails(details string) {
	s.details = details
}

func NewSms_ReleaseNote(device_id int, releaseNoteType string, date string, details string) *Sms_ReleaseNote {
	return &Sms_ReleaseNote{device_id: device_id, releaseNoteType: releaseNoteType, date: date, details: details}
}

func NewSms_ReleaseNoteFromDB(releasenote_id int, device_id int, releaseNoteType string, date string, details string) *Sms_ReleaseNote {
	return &Sms_ReleaseNote{releasenote_id: releasenote_id, device_id: device_id, releaseNoteType: releaseNoteType, date: date, details: details}
}

