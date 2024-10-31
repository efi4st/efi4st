/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_Artefact struct {
	artefact_id int `db:"artefact_id"`
	artefactype_id int `db:"artefactype_id"`
	name string `db:"name"`
	version string `db:"version"`
	artefactype_join string `db:"artefactType"`
}

func (s *Sms_Artefact) Artefact_id() int {
	return s.artefact_id
}

func (s *Sms_Artefact) SetArtefact_id(artefact_id int) {
	s.artefact_id = artefact_id
}

func (s *Sms_Artefact) Artefactype_id() int {
	return s.artefactype_id
}

func (s *Sms_Artefact) SetArtefactype_id(artefactype_id int) {
	s.artefactype_id = artefactype_id
}

func (s *Sms_Artefact) Name() string {
	return s.name
}

func (s *Sms_Artefact) SetName(name string) {
	s.name = name
}

func (s *Sms_Artefact) Version() string {
	return s.version
}

func (s *Sms_Artefact) SetVersion(version string) {
	s.version = version
}

func (s *Sms_Artefact) Artefactype_join() string {
	return s.artefactype_join
}

func (s *Sms_Artefact) SetArtefactype_join(artefactype_join string) {
	s.artefactype_join = artefactype_join
}

func NewSms_Artefact(artefactype_id int, name string, version string, artefactype_join string) *Sms_Artefact {
	return &Sms_Artefact{artefactype_id: artefactype_id, name: name, version: version, artefactype_join: artefactype_join}
}

func NewSms_ArtefactFromDB(artefact_id int, artefactype_id int, name string, version string, artefactype_join string) *Sms_Artefact {
	return &Sms_Artefact{artefact_id: artefact_id, artefactype_id: artefactype_id, name: name, version: version, artefactype_join: artefactype_join}
}
