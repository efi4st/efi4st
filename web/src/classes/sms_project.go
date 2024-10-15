/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_Project struct {
	project_id int `db:"project_id"`
	name string `db:"name"`
	customer string `db:"customer"`
	projecttype string `db:"projecttype_id"`
	reference string `db:"reference"`
	date string `db:"date"`
	active bool `db:"active"`
}

func (s *Sms_Project) Project_id() int {
	return s.project_id
}

func (s *Sms_Project) SetProject_id(project_id int) {
	s.project_id = project_id
}

func (s *Sms_Project) Name() string {
	return s.name
}

func (s *Sms_Project) SetName(name string) {
	s.name = name
}

func (s *Sms_Project) Customer() string {
	return s.customer
}

func (s *Sms_Project) SetCustomer(customer string) {
	s.customer = customer
}

func (s *Sms_Project) Projecttype() string {
	return s.projecttype
}

func (s *Sms_Project) SetProjecttype_id(projecttype string) {
	s.projecttype = projecttype
}

func (s *Sms_Project) Reference() string {
	return s.reference
}

func (s *Sms_Project) SetReference(reference string) {
	s.reference = reference
}

func (s *Sms_Project) Date() string {
	return s.date
}

func (s *Sms_Project) SetDate(date string) {
	s.date = date
}

func (s *Sms_Project) Active() bool {
	return s.active
}

func (s *Sms_Project) SetActive(active bool) {
	s.active = active
}

func NewSms_ProjectFromDB(project_id int, name string, customer string, projecttype string, reference string, date string, active bool) *Sms_Project {
	return &Sms_Project{project_id: project_id, name: name, customer: customer, projecttype: projecttype, reference: reference, date: date, active: active}
}

func NewSms_Project(name string, customer string, projecttype string, reference string, date string, active bool) *Sms_Project {
	return &Sms_Project{name: name, customer: customer, projecttype: projecttype, reference: reference, date: date, active: active}
}

