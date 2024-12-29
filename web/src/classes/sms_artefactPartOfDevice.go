/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_ArtefactPartOfDevice struct {
	device_id int `db:"device_id"`
	artefact_id int `db:"artefact_id"`
	additionalInfo string `db:"additionalInfo"`
	name string `db:"type"`
	version string `db:"version"`
}

func (s *Sms_ArtefactPartOfDevice) Device_id() int {
	return s.device_id
}

func (s *Sms_ArtefactPartOfDevice) SetDevice_id(device_id int) {
	s.device_id = device_id
}

func (s *Sms_ArtefactPartOfDevice) Artefact_id() int {
	return s.artefact_id
}

func (s *Sms_ArtefactPartOfDevice) SetArtefact_id(artefact_id int) {
	s.artefact_id = artefact_id
}

func (s *Sms_ArtefactPartOfDevice) AdditionalInfo() string {
	return s.additionalInfo
}

func (s *Sms_ArtefactPartOfDevice) SetAdditionalInfo(additionalInfo string) {
	s.additionalInfo = additionalInfo
}

func (s *Sms_ArtefactPartOfDevice) Name() string {
	return s.name
}

func (s *Sms_ArtefactPartOfDevice) SetName(name string) {
	s.name = name
}

func (s *Sms_ArtefactPartOfDevice) Version() string {
	return s.version
}

func (s *Sms_ArtefactPartOfDevice) SetVersion(version string) {
	s.version = version
}

func NewSms_ArtefactPartOfDevice(device_id int, artefact_id int, additionalInfo string, name string, version string) *Sms_ArtefactPartOfDevice {
	return &Sms_ArtefactPartOfDevice{device_id: device_id, artefact_id: artefact_id, additionalInfo: additionalInfo, name: name, version: version}
}
