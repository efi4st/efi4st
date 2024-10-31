/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_ArtefactType struct {
	artefacttype_id int `db:"artefacttype_id"`
	artefactType string `db:"artefactType"`
}

func (s *Sms_ArtefactType) Artefacttype_id() int {
	return s.artefacttype_id
}

func (s *Sms_ArtefactType) SetArtefacttype_id(artefacttype_id int) {
	s.artefacttype_id = artefacttype_id
}

func (s *Sms_ArtefactType) ArtefactType() string {
	return s.artefactType
}

func (s *Sms_ArtefactType) SetArtefactType(artefactType string) {
	s.artefactType = artefactType
}

func NewSms_ArtefactTypeFromDB(artefacttype_id int, artefactType string) *Sms_ArtefactType {
	return &Sms_ArtefactType{artefacttype_id: artefacttype_id, artefactType: artefactType}
}
