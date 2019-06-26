/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

import "time"

type Project struct {
	project_id int `db:"project_id"`
	name string `db:"name"`
	uploads int `db:"uploads"`
	date time.Time `db:"date"`
}

func NewProject(name string, uploads int, date time.Time) *Project {
	return &Project{name: name, uploads: uploads, date: date}
}

func NewProjectFromDB(project_id int, name string, uploads int, date time.Time) *Project {
	return &Project{project_id: project_id, name: name, uploads: uploads, date: date}
}

func (p *Project) Date() time.Time {
	return p.date
}

func (p *Project) SetDate(date time.Time) {
	p.date = date
}

func (p *Project) Uploads() int {
	return p.uploads
}

func (p *Project) SetUploads(uploads int) {
	p.uploads = uploads
}

func (p *Project) Project_id() int {
	return p.project_id
}

func (p *Project) SetProject_id(project_id int) {
	p.project_id = project_id
}

func (p *Project) Name() string {
	return p.name
}

func (p *Project) SetName(name string) {
	p.name = name
}
