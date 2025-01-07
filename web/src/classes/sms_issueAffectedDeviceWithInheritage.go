/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_IssueAffectedDeviceWithInheritage struct {
	DeviceID       int     `db:"device_id"`
	IssueID        int     `db:"issue_id"`
	AdditionalInfo string `db:"additionalInfo"`
	Confirmed      bool    `db:"confirmed"`
	DeviceType     string  `db:"device_type"`
	DeviceVersion  string  `db:"device_version"`
	Inherit        bool    `db:"inherit"` // true, wenn durch Software oder Komponente betroffen
}

func NewSms_IssueAffectedDeviceWithInheritage(deviceID int, issueID int, additionalInfo string, confirmed bool, deviceType string, deviceVersion string, inherit bool) *Sms_IssueAffectedDeviceWithInheritage {
	return &Sms_IssueAffectedDeviceWithInheritage{DeviceID: deviceID, IssueID: issueID, AdditionalInfo: additionalInfo, Confirmed: confirmed, DeviceType: deviceType, DeviceVersion: deviceVersion, Inherit: inherit}
}