/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

import "fmt"

type Sms_ArtefactPartOfSystem struct {
	System_id      int
	Artefact_id    int
	AdditionalInfo string
	ArtefactType   string // z. B. Aufbauanleitung, Image etc.
	Version        string // Artefakt-Version oder System-Version, je nach Query
}

func NewSms_ArtefactPartOfSystem(system_id int, artefact_id int, additionalInfo string, artefactType string, version string) *Sms_ArtefactPartOfSystem {
	return &Sms_ArtefactPartOfSystem{
		System_id:      system_id,
		Artefact_id:    artefact_id,
		AdditionalInfo: additionalInfo,
		ArtefactType:   artefactType,
		Version:        version,
	}
}

func (s *Sms_ArtefactPartOfSystem) String() string {
	return fmt.Sprintf("SystemID: %d, ArtefactID: %d, Info: %s, Type: %s, Version: %s",
		s.System_id, s.Artefact_id, s.AdditionalInfo, s.ArtefactType, s.Version)
}