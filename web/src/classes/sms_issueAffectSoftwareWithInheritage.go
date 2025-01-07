/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_IssueAffectedSoftwareWithInheritage struct {
	SoftwareID     int     `db:"software_id"`
	IssueID        int     `db:"issue_id"`
	AdditionalInfo string `db:"additionalInfo"`
	Confirmed      bool    `db:"confirmed"`
	TypeName       string  `db:"type_name"`
	Version        string  `db:"version"`
	Inherit        bool    `db:"inherit"` // true, wenn durch eine Komponente betroffen
}

func NewSms_IssueAffectedSoftwareWithInheritage(softwareID int, issueID int, additionalInfo string, confirmed bool, typeName string, version string, inherit bool) *Sms_IssueAffectedSoftwareWithInheritage {
	return &Sms_IssueAffectedSoftwareWithInheritage{SoftwareID: softwareID, IssueID: issueID, AdditionalInfo: additionalInfo, Confirmed: confirmed, TypeName: typeName, Version: version, Inherit: inherit}
}

