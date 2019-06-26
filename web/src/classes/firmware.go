/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

import "time"

type Firmware struct {
	firmware_id int `db:"firmware_id"`
	name string `db:"name"`
	version string `db:"version"`
	binwalkOutput string `db:"binwalkOutput"`
	sizeInBytes int `db:"sizeInBytes"`
	project_id int `db:"project_id"`
	Created time.Time `db:"created"`
	msg string
}

func (f *Firmware) Msg() string {
	return f.msg
}

func (f *Firmware) SetMsg(msg string) {
	f.msg = msg
}

func NewFirmware(firmware_id int, name string, version string, binwalkOutput string, sizeInBytes int, project_id int, created time.Time) *Firmware {
	return &Firmware{firmware_id: firmware_id, name: name, version: version, binwalkOutput: binwalkOutput, sizeInBytes: sizeInBytes, project_id: project_id, Created: created}
}

func (f *Firmware) Firmware_id() int {
	return f.firmware_id
}

func (f *Firmware) SetFirmware_id(firmware_id int) {
	f.firmware_id = firmware_id
}

func (f *Firmware) SizeInBytes() int {
	return f.sizeInBytes
}

func (f *Firmware) SetSizeInBytes(sizeInBytes int) {
	f.sizeInBytes = sizeInBytes
}

func (f *Firmware) BinwalkOutput() string {
	return f.binwalkOutput
}

func (f *Firmware) SetBinwalkOutput(binwalkOutput string) {
	f.binwalkOutput = binwalkOutput
}

func (f *Firmware) Name() string {
	return f.name
}

func (f *Firmware) SetName(name string) {
	f.name = name
}

func (f *Firmware) Version() string {
	return f.version
}

func (f *Firmware) SetVersion(version string) {
	f.version = version
}

func (f *Firmware) SetCreated(created time.Time) {
	f.Created = created
}
