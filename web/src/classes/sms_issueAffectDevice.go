/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_IssueAffectedDevice struct {
	device_id int `db:"device_id"`
	issue_id int `db:"issue_id"`
	additionalInfo string `db:"additionalInfo"`
	confirmed bool `db:"confirmed"`
	tmp string
	tmp2 string
}

func (s *Sms_IssueAffectedDevice) Device_id() int {
	return s.device_id
}

func (s *Sms_IssueAffectedDevice) SetDevice_id(device_id int) {
	s.device_id = device_id
}

func (s *Sms_IssueAffectedDevice) Issue_id() int {
	return s.issue_id
}

func (s *Sms_IssueAffectedDevice) SetIssue_id(issue_id int) {
	s.issue_id = issue_id
}

func (s *Sms_IssueAffectedDevice) AdditionalInfo() string {
	return s.additionalInfo
}

func (s *Sms_IssueAffectedDevice) SetAdditionalInfo(additionalInfo string) {
	s.additionalInfo = additionalInfo
}

func (s *Sms_IssueAffectedDevice) Confirmed() bool {
	return s.confirmed
}

func (s *Sms_IssueAffectedDevice) SetConfirmed(confirmed bool) {
	s.confirmed = confirmed
}

func (s *Sms_IssueAffectedDevice) Tmp() string {
	return s.tmp
}

func (s *Sms_IssueAffectedDevice) SetTmp(tmp string) {
	s.tmp = tmp
}

func (s *Sms_IssueAffectedDevice) Tmp2() string {
	return s.tmp2
}

func (s *Sms_IssueAffectedDevice) SetTmp2(tmp2 string) {
	s.tmp2 = tmp2
}

func NewSms_IssueAffectedDevice(device_id int, issue_id int, additionalInfo string, confirmed bool, tmp string, tmp2 string) *Sms_IssueAffectedDevice {
	return &Sms_IssueAffectedDevice{device_id: device_id, issue_id: issue_id, additionalInfo: additionalInfo, confirmed: confirmed, tmp: tmp, tmp2: tmp2}
}

