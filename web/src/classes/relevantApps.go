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
	moduleDefault bool `db:"moduleDefault"`
	moduleInitSystem bool `db:"moduleInitSystem"`
	moduleFileContent bool `db:"moduleFileContent"`
	moduleBash bool `db:"moduleBash"`
	moduleCronJob bool `db:"moduleCronJob"`
	moduleProcesses bool `db:"moduleProcesses"`
	moduleInterfaces bool `db:"moduleInterfaces"`
	moduleSystemControls bool `db:"moduleSystemControls"`
	moduleFileSystem bool `db:"moduleFileSystem"`
	modulePortscanner bool `db:"modulePortscanner"`
	moduleProtocolls bool `db:"moduleProtocolls"`
	moduleNetInterfaces bool `db:"moduleNetInterfaces"`
	moduleFileSystemInterfaces bool `db:"moduleFileSystemInterfaces"`
	moduleFileHandles bool `db:"moduleFileHandles"`
	firmware_id int `db:"firmware_id"`
	msg string

}

func (r *RelevantApps) ModuleFileHandles() bool {
	return r.moduleFileHandles
}

func (r *RelevantApps) SetModuleFileHandles(moduleFileHandles bool) {
	r.moduleFileHandles = moduleFileHandles
}

func (r *RelevantApps) ModuleFileSystemInterfaces() bool {
	return r.moduleFileSystemInterfaces
}

func (r *RelevantApps) SetModuleFileSystemInterfaces(moduleFileSystemInterfaces bool) {
	r.moduleFileSystemInterfaces = moduleFileSystemInterfaces
}

func (r *RelevantApps) ModuleNetInterfaces() bool {
	return r.moduleNetInterfaces
}

func (r *RelevantApps) SetModuleNetInterfaces(moduleNetInterfaces bool) {
	r.moduleNetInterfaces = moduleNetInterfaces
}

func (r *RelevantApps) ModuleProtocolls() bool {
	return r.moduleProtocolls
}

func (r *RelevantApps) SetModuleProtocolls(moduleProtocolls bool) {
	r.moduleProtocolls = moduleProtocolls
}

func (r *RelevantApps) ModulePortscanner() bool {
	return r.modulePortscanner
}

func (r *RelevantApps) SetModulePortscanner(modulePortscanner bool) {
	r.modulePortscanner = modulePortscanner
}

func (r *RelevantApps) ModuleFileSystem() bool {
	return r.moduleFileSystem
}

func (r *RelevantApps) SetModuleFileSystem(moduleFileSystem bool) {
	r.moduleFileSystem = moduleFileSystem
}

func (r *RelevantApps) ModuleSystemControls() bool {
	return r.moduleSystemControls
}

func (r *RelevantApps) SetModuleSystemControls(moduleSystemControls bool) {
	r.moduleSystemControls = moduleSystemControls
}

func (r *RelevantApps) ModuleInterfaces() bool {
	return r.moduleInterfaces
}

func (r *RelevantApps) SetModuleInterfaces(moduleInterfaces bool) {
	r.moduleInterfaces = moduleInterfaces
}

func (r *RelevantApps) ModuleDefault() bool {
	return r.moduleDefault
}

func (r *RelevantApps) SetModuleDefault(moduleDefault bool) {
	r.moduleDefault = moduleDefault
}

func (r *RelevantApps) ModuleInitSystem() bool {
	return r.moduleInitSystem
}

func (r *RelevantApps) SetModuleInitSystem(moduleInitSystem bool) {
	r.moduleInitSystem = moduleInitSystem
}

func (r *RelevantApps) ModuleFileContent() bool {
	return r.moduleFileContent
}

func (r *RelevantApps) SetModuleFileContent(moduleFileContent bool) {
	r.moduleFileContent = moduleFileContent
}

func (r *RelevantApps) ModuleBash() bool {
	return r.moduleBash
}

func (r *RelevantApps) SetModuleBash(moduleBash bool) {
	r.moduleBash = moduleBash
}

func (r *RelevantApps) ModuleCronJob() bool {
	return r.moduleCronJob
}

func (r *RelevantApps) SetModuleCronJob(moduleCronJob bool) {
	r.moduleCronJob = moduleCronJob
}

func (r *RelevantApps) ModuleProcesses() bool {
	return r.moduleProcesses
}

func (r *RelevantApps) SetModuleProcesses(moduleProcesses bool) {
	r.moduleProcesses = moduleProcesses
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

