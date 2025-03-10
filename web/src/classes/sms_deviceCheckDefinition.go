/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_DeviceCheckDefinition struct {
	ID                 int     `db:"id"`
	DeviceTypeID       int     `db:"device_type_id"`
	DeviceTypeName     string  `db:"device_type_name"`
	ApplicableVersions string  `db:"applicable_versions"`
	TestName           string  `db:"test_name"`
	TestDescription    string  `db:"test_description"`
	Explanation        string  `db:"explanation"`
	ExpectedResult     string  `db:"expected_result"`
	FilterCondition    *string `db:"filter_condition"`
	CheckType          string  `db:"check_type"` // Neues Feld für die Check-Typen
}

// Konstruktor für eine neue Check-Definition (ohne ID, für Inserts)
func NewSms_DeviceCheckDefinition(deviceTypeID int, deviceTypeName, applicableVersions, testName, testDescription, explanation, expectedResult, filterCondition, checkType string) *Sms_DeviceCheckDefinition {
	return &Sms_DeviceCheckDefinition{
		DeviceTypeID:       deviceTypeID,
		DeviceTypeName:     deviceTypeName,
		ApplicableVersions: applicableVersions,
		TestName:           testName,
		TestDescription:    testDescription,
		Explanation:        explanation,
		ExpectedResult:     expectedResult,
		FilterCondition:    &filterCondition,
		CheckType:          checkType,
	}
}

// Konstruktor für eine Check-Definition aus der Datenbank (mit ID, für Updates oder Reads)
func NewSms_DeviceCheckDefinitionFromDB(ID, deviceTypeID int, deviceTypeName, applicableVersions, testName, testDescription, explanation, expectedResult, filterCondition, checkType string) *Sms_DeviceCheckDefinition {
	return &Sms_DeviceCheckDefinition{
		ID:                 ID,
		DeviceTypeID:       deviceTypeID,
		DeviceTypeName:     deviceTypeName,
		ApplicableVersions: applicableVersions,
		TestName:           testName,
		TestDescription:    testDescription,
		Explanation:        explanation,
		ExpectedResult:     expectedResult,
		FilterCondition:    &filterCondition,
		CheckType:          checkType,
	}
}

type ProjectDeviceCheck struct {
	DeviceInstanceID  int    `db:"deviceInstance_id"`
	DeviceID          int    `db:"device_id"`
	DeviceTypeID      int    `db:"device_type_id"`
	TestName          string `db:"test_name"`
	TestDescription   string `db:"test_description"`
	Explanation       string `db:"explanation"`
	ExpectedResult    string `db:"expected_result"`
	CheckType         string `db:"check_type"` // Neues Feld für die Check-Typen
}

type ResultProjectCheck struct {
	TestName           string  `db:"test_name"`
	TestDescription    string  `db:"test_description"`
	ApplicableVersions string  `db:"applicable_versions"`
	Explanation        string  `db:"explanation"`
	ExpectedResult     string  `db:"expected_result"`
	FilterCondition    *string `db:"filter_condition"`
	DeviceType         string  `db:"device_type"`
	InstanceCount      int     `db:"instance_count"`
	Versions           string  `db:"versions"`
	CheckType          string  `db:"check_type"` // Neues Feld für die Check-Typen
}