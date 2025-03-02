/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

// Sms_AffectedProjects speichert betroffene Projekte eines Issues
type Sms_AffectedProjects struct {
	ProjectID int32  `json:"project_id" db:"project_id"`
	Name      string `json:"name" db:"name"`
	Customer  string `json:"customer" db:"customer"`
}