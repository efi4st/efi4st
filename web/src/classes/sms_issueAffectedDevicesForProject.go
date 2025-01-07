/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_IssueWithAffectedDevices struct {
	IssueID        int    `db:"issue_id"`
	IssueName      string `db:"issue_name"`
	Criticality    string `db:"criticality"`
	AffectedDevices []struct {
		DeviceID      int    `db:"device_id"`
		DeviceName    string `db:"device_name"`
		DeviceVersion string `db:"device_version"`
		Inherit       bool   `db:"inherit"`
	}
}

func NewSms_IssueWithAffectedDevices(issueID int, issueName string, criticality string, affectedDevices []struct {
	DeviceID      int    `db:"device_id"`
	DeviceName    string `db:"device_name"`
	DeviceVersion string `db:"device_version"`
	Inherit       bool   `db:"inherit"`
}) *Sms_IssueWithAffectedDevices {
	return &Sms_IssueWithAffectedDevices{IssueID: issueID, IssueName: issueName, Criticality: criticality, AffectedDevices: affectedDevices}
}

