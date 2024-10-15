/**
 * Author:    Admiral Helmut
 * Created:   13.10.2024
 *
 * (C)
 **/

package classes

type Sms_UpdateHistory struct {
	updateHistory_id int `db:"updateHistory_id"`
	deviceInstance_id int `db:"deviceInstance_id"`
	deviceInstance_name string `db:"deviceInstance_name"`
	user string `db:"user"`
	updateType string `db:"updateType"`
	date string `db:"date"`
	description string `db:"description"`
}

func (s *Sms_UpdateHistory) UpdateHistory_id() int {
	return s.updateHistory_id
}

func (s *Sms_UpdateHistory) SetUpdateHistory_id(updateHistory_id int) {
	s.updateHistory_id = updateHistory_id
}

func (s *Sms_UpdateHistory) DeviceInstance_id() int {
	return s.deviceInstance_id
}

func (s *Sms_UpdateHistory) SetDeviceInstance_id(deviceInstance_id int) {
	s.deviceInstance_id = deviceInstance_id
}

func (s *Sms_UpdateHistory) DeviceInstance_name() string {
	return s.deviceInstance_name
}

func (s *Sms_UpdateHistory) SetDeviceInstance_name(deviceInstance_name string) {
	s.deviceInstance_name = deviceInstance_name
}

func (s *Sms_UpdateHistory) User() string {
	return s.user
}

func (s *Sms_UpdateHistory) SetUser(user string) {
	s.user = user
}

func (s *Sms_UpdateHistory) UpdateType() string {
	return s.updateType
}

func (s *Sms_UpdateHistory) SetUpdateType(updateType string) {
	s.updateType = updateType
}

func (s *Sms_UpdateHistory) Date() string {
	return s.date
}

func (s *Sms_UpdateHistory) SetDate(date string) {
	s.date = date
}

func (s *Sms_UpdateHistory) Description() string {
	return s.description
}

func (s *Sms_UpdateHistory) SetDescription(description string) {
	s.description = description
}

func NewSms_UpdateHistory(deviceInstance_id int, deviceInstance_name string, user string, updateType string, date string, description string) *Sms_UpdateHistory {
	return &Sms_UpdateHistory{deviceInstance_id: deviceInstance_id, deviceInstance_name: deviceInstance_name, user: user, updateType: updateType, date: date, description: description}
}

func NewSms_UpdateHistoryFromDB(updateHistory_id int, deviceInstance_id int, deviceInstance_name string, user string, updateType string, date string, description string) *Sms_UpdateHistory {
	return &Sms_UpdateHistory{updateHistory_id: updateHistory_id, deviceInstance_id: deviceInstance_id, deviceInstance_name: deviceInstance_name, user: user, updateType: updateType, date: date, description: description}
}

