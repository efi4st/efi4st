/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_Project struct {
	Project_id                    int     `db:"project_id"`
	Name                           string  `db:"name"`
	Customer                       string  `db:"customer"`
	Projecttype                    string  `db:"projecttype_id"`
	Reference                      string  `db:"reference"`
	Date                           string  `db:"date"`
	Active                         bool    `db:"active"`
	PlantNumber                    *string `db:"plant_number"`
	ProjectReference               *string `db:"project_reference"`
	IMOPlantPowerPlantFactory      *string `db:"imo_plant_powerplant_factory"`
	PlantType                      *string `db:"plant_type"`
	Note                           *string `db:"note"`
}


func NewSms_ProjectFromDB(
	Project_id int,
	Name string,
	Customer string,
	Projecttype string,
	Reference string,
	Date string,
	Active bool,
	PlantNumber *string,
	ProjectReference *string,
	IMOPlantPowerPlantFactory *string,
	PlantType *string,
	Note *string,
) *Sms_Project {
	return &Sms_Project{
		Project_id: Project_id,
		Name: Name,
		Customer: Customer,
		Projecttype: Projecttype,
		Reference: Reference,
		Date: Date,
		Active: Active,
		PlantNumber: PlantNumber,
		ProjectReference: ProjectReference,
		IMOPlantPowerPlantFactory: IMOPlantPowerPlantFactory,
		PlantType: PlantType,
		Note: Note,
	}
}

