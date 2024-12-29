/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_IssueAffectedArtefact struct {
	artefact_id       int    `db:"artefact_id"`
	issue_id          int    `db:"issue_id"`
	additionalInfo    string `db:"additionalInfo"`
	confirmed         bool   `db:"confirmed"`
	artefact_name     string `db:"artefact_name"`
	artefact_version  string `db:"artefact_version"`
}

func (s *Sms_IssueAffectedArtefact) Artefact_id() int {
	return s.artefact_id
}

func (s *Sms_IssueAffectedArtefact) SetArtefact_id(artefact_id int) {
	s.artefact_id = artefact_id
}

func (s *Sms_IssueAffectedArtefact) Issue_id() int {
	return s.issue_id
}

func (s *Sms_IssueAffectedArtefact) SetIssue_id(issue_id int) {
	s.issue_id = issue_id
}

func (s *Sms_IssueAffectedArtefact) AdditionalInfo() string {
	return s.additionalInfo
}

func (s *Sms_IssueAffectedArtefact) SetAdditionalInfo(additionalInfo string) {
	s.additionalInfo = additionalInfo
}

func (s *Sms_IssueAffectedArtefact) Confirmed() bool {
	return s.confirmed
}

func (s *Sms_IssueAffectedArtefact) SetConfirmed(confirmed bool) {
	s.confirmed = confirmed
}

func (s *Sms_IssueAffectedArtefact) Artefact_name() string {
	return s.artefact_name
}

func (s *Sms_IssueAffectedArtefact) SetArtefact_name(artefact_name string) {
	s.artefact_name = artefact_name
}

func (s *Sms_IssueAffectedArtefact) Artefact_version() string {
	return s.artefact_version
}

func (s *Sms_IssueAffectedArtefact) SetArtefact_version(artefact_version string) {
	s.artefact_version = artefact_version
}

func NewSms_IssueAffectedArtefact(artefact_id int, issue_id int, additionalInfo string, confirmed bool, artefact_name string, artefact_version string) *Sms_IssueAffectedArtefact {
	return &Sms_IssueAffectedArtefact{artefact_id: artefact_id, issue_id: issue_id, additionalInfo: additionalInfo, confirmed: confirmed, artefact_name: artefact_name, artefact_version: artefact_version}
}
