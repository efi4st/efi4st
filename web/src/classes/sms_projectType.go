/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes


type Sms_ProjectType struct {
	projecttype_id int `db:"projecttype_id"`
	projectType string `db:"projectType"`
}

func (s *Sms_ProjectType) Projecttype_id() int {
	return s.projecttype_id
}

func (s *Sms_ProjectType) SetProjecttype_id(projecttype_id int) {
	s.projecttype_id = projecttype_id
}

func (s *Sms_ProjectType) ProjectType() string {
	return s.projectType
}

func (s *Sms_ProjectType) SetProjectType(projectType string) {
	s.projectType = projectType
}

func NewSms_ProjectType(projecttype_id int, projectType string) *Sms_ProjectType {
	return &Sms_ProjectType{projecttype_id: projecttype_id, projectType: projectType}
}
