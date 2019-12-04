/**
 * Author:    Admiral Helmut
 * Created:   13.11.2019
 *
 * (C)
 **/

package classes

import "time"

type TestResult struct {
	testResult_id int `db:"testResult_id"`
	moduleName string `db:"moduleName"`
	result string `db:"result"`
	Created time.Time `db:"created"`
	firmware_id int `db:"firmware_id"`
	msg string

}

func (f *TestResult) SetCreated(created time.Time) {
	f.Created = created
}

func (t *TestResult) Msg() string {
	return t.msg
}

func (t *TestResult) SetMsg(msg string) {
	t.msg = msg
}

func (t *TestResult) Firmware_id() int {
	return t.firmware_id
}

func (t *TestResult) SetFirmware_id(firmware_id int) {
	t.firmware_id = firmware_id
}

func (t *TestResult) Result() string {
	return t.result
}

func (t *TestResult) SetResult(result string) {
	t.result = result
}

func (t *TestResult) ModuleName() string {
	return t.moduleName
}

func (t *TestResult) SetModuleName(moduleName string) {
	t.moduleName = moduleName
}

func (t *TestResult) TestResult_id() int {
	return t.testResult_id
}

func (t *TestResult) SetTestResult_id(testResult_id int) {
	t.testResult_id = testResult_id
}

func NewTestResult(testResult_id int, moduleName string, result string, created time.Time, firmware_id int) *TestResult {
	return &TestResult{testResult_id: testResult_id, moduleName: moduleName, result: result, Created: created, firmware_id: firmware_id}
}

