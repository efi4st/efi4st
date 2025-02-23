/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_DeviceIPDefinition struct {
	ID               int    `db:"id"`
	DeviceTypeID     int    `db:"devicetype_id"`
	DeviceTypeName   string `db:"device_type_name"`
	ApplicableVersions string `db:"applicable_versions"`
	IPAddress        string `db:"ip_address"`
	VLANID           int    `db:"vlan_id"`
	Description      string `db:"description"`
	FilterCondition  string `db:"filter_condition"`
}

func NewSms_DeviceIPDefinition(deviceTypeID int, deviceTypeName string, applicableVersions string, IPAddress string, VLANID int, description string, filterCondition string) *Sms_DeviceIPDefinition {
	return &Sms_DeviceIPDefinition{DeviceTypeID: deviceTypeID, DeviceTypeName: deviceTypeName, ApplicableVersions: applicableVersions, IPAddress: IPAddress, VLANID: VLANID, Description: description, FilterCondition: filterCondition}
}

func NewSms_DeviceIPDefinitionFromDB(ID int, deviceTypeID int, deviceTypeName string, applicableVersions string, IPAddress string, VLANID int, description string, filterCondition string) *Sms_DeviceIPDefinition {
	return &Sms_DeviceIPDefinition{ID: ID, DeviceTypeID: deviceTypeID, DeviceTypeName: deviceTypeName, ApplicableVersions: applicableVersions, IPAddress: IPAddress, VLANID: VLANID, Description: description, FilterCondition: filterCondition}
}

type ProjectDeviceIP struct {
	DeviceInstanceID int    `db:"deviceInstance_id"`
	DeviceID         int    `db:"device_id"`
	DeviceTypeID     int    `db:"devicetype_id"`
	IPAddress        string `db:"ip_address"`
	VLANID           int    `db:"vlan_id"`
	Description      string `db:"description"`
}

type ResultProjectIP struct {
	IPAddress          string  `db:"ip_address"`
	ApplicableVersions string  `db:"applicable_versions"`
	VLANID            *int    `db:"vlan_id"`
	Description       *string `db:"description"`
	FilterCondition   *string `db:"filter_condition"`
	DeviceType        string  `db:"device_type"`
	InstanceCount     int     `db:"instance_count"`
	Versions          string  `db:"versions"`
}