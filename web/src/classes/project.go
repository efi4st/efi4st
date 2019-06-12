package classes

type Project struct {
	project_id int `db:"project_id"`
	name string `db:"name"`
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