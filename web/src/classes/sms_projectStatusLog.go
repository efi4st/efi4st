/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_ProjectStatusLog struct {
	StatusID     int    `db:"status_id"`     // Primärschlüssel für den Status-Eintrag
	ProjectID    int    `db:"project_id"`    // Fremdschlüssel zum Projekt
	Status       string `db:"status"`        // Statusbezeichnung (z. B. "bestellt", "in_konstruktion", etc.)
	Note         string `db:"note"`          // Zusätzliche Beschreibung oder Kommentar
	AccessGroup  string `db:"access_group"`  // Zugriffskontrollgruppe
	CreatedAt    string `db:"created_at"`    // Zeitstempel des Statuswechsels
}

func NewSms_ProjectStatusLog(statusID int, projectID int, status string, note string, accessGroup string, createdAt string) *Sms_ProjectStatusLog {
	return &Sms_ProjectStatusLog{
		StatusID:    statusID,
		ProjectID:   projectID,
		Status:      status,
		Note:        note,
		AccessGroup: accessGroup,
		CreatedAt:   createdAt,
	}
}