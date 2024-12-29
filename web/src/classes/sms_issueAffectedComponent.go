/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_IssueAffectedComponent struct {
	component_id        int    `db:"component_id"`
	issue_id            int    `db:"issue_id"`
	additionalInfo      string `db:"additionalInfo"`
	confirmed           bool   `db:"confirmed"`
	component_name      string `db:"component_name"`
	component_version   string `db:"component_version"`
}

func (s *Sms_IssueAffectedComponent) Component_id() int {
	return s.component_id
}

func (s *Sms_IssueAffectedComponent) SetComponent_id(component_id int) {
	s.component_id = component_id
}

func (s *Sms_IssueAffectedComponent) Issue_id() int {
	return s.issue_id
}

func (s *Sms_IssueAffectedComponent) SetIssue_id(issue_id int) {
	s.issue_id = issue_id
}

func (s *Sms_IssueAffectedComponent) AdditionalInfo() string {
	return s.additionalInfo
}

func (s *Sms_IssueAffectedComponent) SetAdditionalInfo(additionalInfo string) {
	s.additionalInfo = additionalInfo
}

func (s *Sms_IssueAffectedComponent) Confirmed() bool {
	return s.confirmed
}

func (s *Sms_IssueAffectedComponent) SetConfirmed(confirmed bool) {
	s.confirmed = confirmed
}

func (s *Sms_IssueAffectedComponent) Component_name() string {
	return s.component_name
}

func (s *Sms_IssueAffectedComponent) SetComponent_name(component_name string) {
	s.component_name = component_name
}

func (s *Sms_IssueAffectedComponent) Component_version() string {
	return s.component_version
}

func (s *Sms_IssueAffectedComponent) SetComponent_version(component_version string) {
	s.component_version = component_version
}

func NewSms_IssueAffectedComponent(component_id int, issue_id int, additionalInfo string, confirmed bool, component_name string, component_version string) *Sms_IssueAffectedComponent {
	return &Sms_IssueAffectedComponent{component_id: component_id, issue_id: issue_id, additionalInfo: additionalInfo, confirmed: confirmed, component_name: component_name, component_version: component_version}
}

