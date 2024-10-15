/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_Issue struct {
	issue_id int `db:"issue_id"`
	name string `db:"name"`
	date string `db:"date"`
	issueType string `db:"issueType"`
	reference string `db:"reference"`
	criticality int `db:"criticality"`
	cve string `db:"cve"`
	description string `db:"description"`
}

func NewSms_Issue(name string, date string, issueType string, reference string, criticality int, cve string, description string) *Sms_Issue {
	return &Sms_Issue{name: name, date: date, issueType: issueType, reference: reference, criticality: criticality, cve: cve, description: description}
}

func NewSms_IssueFromDB(issue_id int, name string, date string, issueType string, reference string, criticality int, cve string, description string) *Sms_Issue {
	return &Sms_Issue{issue_id: issue_id, name: name, date: date, issueType: issueType, reference: reference, criticality: criticality, cve: cve, description: description}
}

