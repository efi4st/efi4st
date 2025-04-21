/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_Project struct {
	Project_id int `db:"project_id"`
	Name string `db:"name"`
	Customer string `db:"customer"`
	Projecttype string `db:"projecttype_id"`
	Reference string `db:"reference"`
	Date string `db:"date"`
	Active bool `db:"active"`
}


func NewSms_ProjectFromDB(Project_id int, Name string, Customer string, Projecttype string, Reference string, Date string, Active bool) *Sms_Project {
	return &Sms_Project{Project_id: Project_id, Name: Name, Customer: Customer, Projecttype: Projecttype, Reference: Reference, Date: Date, Active: Active}
}

func NewSms_Project(Name string, Customer string, Projecttype string, Reference string, Date string, Active bool) *Sms_Project {
	return &Sms_Project{Name: Name, Customer: Customer, Projecttype: Projecttype, Reference: Reference, Date: Date, Active: Active}
}

