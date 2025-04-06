/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_System struct {
	system_id int `db:"system_id"`
	systemtype_id string `db:"systemtype_id"`
	version string `db:"version"`
	date string `db:"date"`
}

func NewSms_SystemFromDB(system_id int, systemtype_id string, version string, date string) *Sms_System {
	return &Sms_System{system_id: system_id, systemtype_id: systemtype_id, version: version, date: date}
}

func NewSms_System(systemtype_id string, version string, date string) *Sms_System {
	return &Sms_System{systemtype_id: systemtype_id, version: version, date: date}
}

func (s *Sms_System) System_id() int {
	return s.system_id
}

func (s *Sms_System) SetSystem_id(system_id int) {
	s.system_id = system_id
}

func (s *Sms_System) Systemtype_id() string {
	return s.systemtype_id
}

func (s *Sms_System) SetSystemtype_id(systemtype_id string) {
	s.systemtype_id = systemtype_id
}

func (s *Sms_System) Version() string {
	return s.version
}

func (s *Sms_System) SetVersion(version string) {
	s.version = version
}

func (s *Sms_System) Date() string {
	return s.date
}

func (s *Sms_System) SetDate(date string) {
	s.date = date
}

type Sms_System_Query struct {
	SystemID     int    `db:"system_id"`
	SystemTypeID string `db:"systemtype_id"`
	SystemType   string `db:"system_type"` // Den systemType aus der sms_systemtype Tabelle
	Version      string `db:"version"`
	Date         string `db:"date"`
}