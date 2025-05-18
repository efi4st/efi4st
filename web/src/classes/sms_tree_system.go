/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_Tree_Component struct {
	Name string
	Version string
}

func NewSms_Tree_Component(name string, version string) *Sms_Tree_Component {
	return &Sms_Tree_Component{Name: name, Version: version}
}

type Sms_Tree_Application struct {
	Name       string
	Version    string
	Components []Sms_Tree_Component
}

func NewSms_Tree_Application(name string, version string, components []Sms_Tree_Component) *Sms_Tree_Application {
	return &Sms_Tree_Application{Name: name, Version: version, Components: components}
}

type Sms_Tree_Device struct {
	Name        string
	Version     string
	Applications []Sms_Tree_Application
	Artefacts    []Sms_Tree_Artefact
}


func NewSms_Tree_Device(name string, version string, applications []Sms_Tree_Application, artefacts []Sms_Tree_Artefact) *Sms_Tree_Device {
	return &Sms_Tree_Device{Name: name, Version: version, Applications: applications, Artefacts: artefacts}
}

type Sms_Tree_System struct {
	Name     string
	Version  string
	Devices  []Sms_Tree_Device
	Artefacts []Sms_Tree_Artefact
}

func NewSms_Tree_System(name string, version string, devices []Sms_Tree_Device, artefacts []Sms_Tree_Artefact) *Sms_Tree_System {
	return &Sms_Tree_System{Name: name, Version: version, Devices: devices, Artefacts: artefacts}
}

type Sms_Tree_Artefact struct {
	Name    string
	Version string
}

func NewSms_Tree_Artefact(name string, version string) *Sms_Tree_Artefact {
	return &Sms_Tree_Artefact{
		Name: name,
		Version: version,
	}
}