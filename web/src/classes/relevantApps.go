/**
 * Author:    Admiral Helmut
 * Created:   13.11.2019
 *
 * (C)
 **/

package classes

type RelevantApps struct {
	relevantApps_id int `db:"relevantApps_id"`
	name string `db:"name"`
	path string `db:"path"`
	extPort int `db:"extPort"`
	extProtocoll string `db:"extProtocoll"`
	intInterface string `db:"intInterface"`
	firmware_id int `db:"firmware_id"`
	msg string

}

func (r *RelevantApps) Msg() string {
	return r.msg
}

func (r *RelevantApps) SetMsg(msg string) {
	r.msg = msg
}

func (r *RelevantApps) RelevantApps_id() int {
	return r.relevantApps_id
}

func (r *RelevantApps) SetRelevantApps_id(relevantApps_id int) {
	r.relevantApps_id = relevantApps_id
}

func (r *RelevantApps) Name() string {
	return r.name
}

func (r *RelevantApps) SetName(name string) {
	r.name = name
}

func (r *RelevantApps) Path() string {
	return r.path
}

func (r *RelevantApps) SetPath(path string) {
	r.path = path
}

func (r *RelevantApps) ExtPort() int {
	return r.extPort
}

func (r *RelevantApps) SetExtPort(extPort int) {
	r.extPort = extPort
}

func (r *RelevantApps) ExtProtocoll() string {
	return r.extProtocoll
}

func (r *RelevantApps) SetExtProtocoll(extProtocoll string) {
	r.extProtocoll = extProtocoll
}

func (r *RelevantApps) IntInterface() string {
	return r.intInterface
}

func (r *RelevantApps) SetIntInterface(intInterface string) {
	r.intInterface = intInterface
}

func (r *RelevantApps) Firmware_id() int {
	return r.firmware_id
}

func (r *RelevantApps) SetFirmware_id(firmware_id int) {
	r.firmware_id = firmware_id
}

func NewRelevantApps(relevantApps_id int, name string, path string, extPort int, extProtocoll string, intInterface string, firmware_id int) *RelevantApps {
	return &RelevantApps{relevantApps_id: relevantApps_id, name: name, path: path, extPort: extPort, extProtocoll: extProtocoll, intInterface: intInterface, firmware_id: firmware_id}
}

