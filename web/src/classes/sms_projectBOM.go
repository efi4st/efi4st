/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_ProjectBOM struct {
	projectBOM_id int `db:"projectBOM_id"`
	project_id int `db:"project_id"`
	system_id int `db:"system_id"`
	orderNumber string `db:"orderNumber"`
	additionalInfo string `db:"additionalInfo"`
	name string
	tmp string
}

func (s *Sms_ProjectBOM) ProjectBOM_id() int {
	return s.projectBOM_id
}

func (s *Sms_ProjectBOM) SetProjectBOM_id(projectBOM_id int) {
	s.projectBOM_id = projectBOM_id
}

func (s *Sms_ProjectBOM) Project_id() int {
	return s.project_id
}

func (s *Sms_ProjectBOM) SetProject_id(project_id int) {
	s.project_id = project_id
}

func (s *Sms_ProjectBOM) System_id() int {
	return s.system_id
}

func (s *Sms_ProjectBOM) SetSystem_id(system_id int) {
	s.system_id = system_id
}

func (s *Sms_ProjectBOM) OrderNumber() string {
	return s.orderNumber
}

func (s *Sms_ProjectBOM) SetOrderNumber(orderNumber string) {
	s.orderNumber = orderNumber
}

func (s *Sms_ProjectBOM) AdditionalInfo() string {
	return s.additionalInfo
}

func (s *Sms_ProjectBOM) SetAdditionalInfo(additionalInfo string) {
	s.additionalInfo = additionalInfo
}

func (s *Sms_ProjectBOM) Name() string {
	return s.name
}

func (s *Sms_ProjectBOM) SetName(name string) {
	s.name = name
}

func (s *Sms_ProjectBOM) Tmp() string {
	return s.tmp
}

func (s *Sms_ProjectBOM) SetTmp(tmp string) {
	s.tmp = tmp
}

func NewSms_ProjectBOM(project_id int, system_id int, orderNumber string, additionalInfo string, name string, tmp string) *Sms_ProjectBOM {
	return &Sms_ProjectBOM{project_id: project_id, system_id: system_id, orderNumber: orderNumber, additionalInfo: additionalInfo, name: name, tmp: tmp}
}

func NewSms_ProjectBOMFromDB(projectBOM_id int, project_id int, system_id int, orderNumber string, additionalInfo string, name string, tmp string) *Sms_ProjectBOM {
	return &Sms_ProjectBOM{projectBOM_id: projectBOM_id, project_id: project_id, system_id: system_id, orderNumber: orderNumber, additionalInfo: additionalInfo, name: name, tmp: tmp}
}

