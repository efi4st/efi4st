/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_Solution struct {
	solution_id int `db:"solution_id"`
	issue_id int `db:"issue_id"`
	devicetype_id int `db:"devicetype_id"`
	date string `db:"date"`
	name string `db:"name"`
	description string `db:"description"`
	reference string `db:"reference"`
	deviceTypeJoined string `db:"type"`
}

func (s *Sms_Solution) Solution_id() int {
	return s.solution_id
}

func (s *Sms_Solution) SetSolution_id(solution_id int) {
	s.solution_id = solution_id
}

func (s *Sms_Solution) Issue_id() int {
	return s.issue_id
}

func (s *Sms_Solution) SetIssue_id(issue_id int) {
	s.issue_id = issue_id
}

func (s *Sms_Solution) Devicetype_id() int {
	return s.devicetype_id
}

func (s *Sms_Solution) SetDevicetype_id(devicetype_id int) {
	s.devicetype_id = devicetype_id
}

func (s *Sms_Solution) Date() string {
	return s.date
}

func (s *Sms_Solution) SetDate(date string) {
	s.date = date
}

func (s *Sms_Solution) Name() string {
	return s.name
}

func (s *Sms_Solution) SetName(name string) {
	s.name = name
}

func (s *Sms_Solution) Description() string {
	return s.description
}

func (s *Sms_Solution) SetDescription(description string) {
	s.description = description
}

func (s *Sms_Solution) Reference() string {
	return s.reference
}

func (s *Sms_Solution) SetReference(reference string) {
	s.reference = reference
}

func (s *Sms_Solution) DeviceTypeJoined() string {
	return s.deviceTypeJoined
}

func (s *Sms_Solution) SetDeviceTypeJoined(deviceTypeJoined string) {
	s.deviceTypeJoined = deviceTypeJoined
}

func NewSms_Solution(issue_id int, devicetype_id int, date string, name string, description string, reference string, deviceTypeJoined string) *Sms_Solution {
	return &Sms_Solution{issue_id: issue_id, devicetype_id: devicetype_id, date: date, name: name, description: description, reference: reference, deviceTypeJoined: deviceTypeJoined}
}

func NewSms_SolutionFromDB(solution_id int, issue_id int, devicetype_id int, date string, name string, description string, reference string, deviceTypeJoined string) *Sms_Solution {
	return &Sms_Solution{solution_id: solution_id, issue_id: issue_id, devicetype_id: devicetype_id, date: date, name: name, description: description, reference: reference, deviceTypeJoined: deviceTypeJoined}
}

