/**
 * Author:    Admiral Helmut
 * Created:   13.10.2024
 *
 * (C)
 **/

package classes

type Sms_ManufacturingOrder struct {
	manufacturingOrder_id int `db:"manufacturingOrder_id"`
	system_id int `db:"system_id"`
	packageReference string `db:"packageReference"`
	start string `db:"start"`
	end string `db:"end"`
	description string `db:"description"`
}

func (s *Sms_ManufacturingOrder) ManufacturingOrder_id() int {
	return s.manufacturingOrder_id
}

func (s *Sms_ManufacturingOrder) SetManufacturingOrder_id(manufacturingOrder_id int) {
	s.manufacturingOrder_id = manufacturingOrder_id
}

func (s *Sms_ManufacturingOrder) System_id() int {
	return s.system_id
}

func (s *Sms_ManufacturingOrder) SetSystem_id(system_id int) {
	s.system_id = system_id
}

func (s *Sms_ManufacturingOrder) PackageReference() string {
	return s.packageReference
}

func (s *Sms_ManufacturingOrder) SetPackageReference(packageReference string) {
	s.packageReference = packageReference
}

func (s *Sms_ManufacturingOrder) Start() string {
	return s.start
}

func (s *Sms_ManufacturingOrder) SetStart(start string) {
	s.start = start
}

func (s *Sms_ManufacturingOrder) End() string {
	return s.end
}

func (s *Sms_ManufacturingOrder) SetEnd(end string) {
	s.end = end
}

func (s *Sms_ManufacturingOrder) Description() string {
	return s.description
}

func (s *Sms_ManufacturingOrder) SetDescription(description string) {
	s.description = description
}

func NewSms_ManufacturingOrder(system_id int, packageReference string, start string, end string, description string) *Sms_ManufacturingOrder {
	return &Sms_ManufacturingOrder{system_id: system_id, packageReference: packageReference, start: start, end: end, description: description}
}

func NewSms_ManufacturingOrderFromDB(manufacturingOrder_id int, system_id int, packageReference string, start string, end string, description string) *Sms_ManufacturingOrder {
	return &Sms_ManufacturingOrder{manufacturingOrder_id: manufacturingOrder_id, system_id: system_id, packageReference: packageReference, start: start, end: end, description: description}
}

