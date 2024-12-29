/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_IssueAffectedSoftware struct {
	software_id int `db:"software_id"`
	issue_id int `db:"issue_id"`
	additionalInfo string `db:"additionalInfo"`
	confirmed bool `db:"confirmed"`
	tmp string
	tmp2 string
}

func (s *Sms_IssueAffectedSoftware) Software_id() int {
	return s.software_id
}

func (s *Sms_IssueAffectedSoftware) SetSoftware_id(software_id int) {
	s.software_id = software_id
}

func (s *Sms_IssueAffectedSoftware) Issue_id() int {
	return s.issue_id
}

func (s *Sms_IssueAffectedSoftware) SetIssue_id(issue_id int) {
	s.issue_id = issue_id
}

func (s *Sms_IssueAffectedSoftware) AdditionalInfo() string {
	return s.additionalInfo
}

func (s *Sms_IssueAffectedSoftware) SetAdditionalInfo(additionalInfo string) {
	s.additionalInfo = additionalInfo
}

func (s *Sms_IssueAffectedSoftware) Confirmed() bool {
	return s.confirmed
}

func (s *Sms_IssueAffectedSoftware) SetConfirmed(confirmed bool) {
	s.confirmed = confirmed
}

func (s *Sms_IssueAffectedSoftware) Tmp() string {
	return s.tmp
}

func (s *Sms_IssueAffectedSoftware) SetTmp(tmp string) {
	s.tmp = tmp
}

func (s *Sms_IssueAffectedSoftware) Tmp2() string {
	return s.tmp2
}

func (s *Sms_IssueAffectedSoftware) SetTmp2(tmp2 string) {
	s.tmp2 = tmp2
}

func NewSms_IssueAffectedSoftware(software_id int, issue_id int, additionalInfo string, confirmed bool, tmp string, tmp2 string) *Sms_IssueAffectedSoftware {
	return &Sms_IssueAffectedSoftware{software_id: software_id, issue_id: issue_id, additionalInfo: additionalInfo, confirmed: confirmed, tmp: tmp, tmp2: tmp2}
}


