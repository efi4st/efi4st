/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Firmware struct {
	firmware_id int `db:"firmware_id"`
	name string `db:"name"`
	version string `db:"version"`
	binwalkOutput string `db:"binwalkOutput"`
	sizeInBytes int `db:"sizeInBytes"`
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
