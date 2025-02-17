/**
 * Author:    Admiral Helmut
 * Created:   14.02.2025
 *
 * (C)
 **/

package classes

type ProjectSettingsLink struct {
	ProjectID int    `db:"project_id"`  // ID des Projekts
	SettingID int    `db:"setting_id"`  // ID der Einstellung
	Value     string `db:"value"`       // Gespeicherter Wert der Einstellung
}

func NewProjectSettingsLink(projectID int, settingID int, value string) *ProjectSettingsLink {
	return &ProjectSettingsLink{ProjectID: projectID, SettingID: settingID, Value: value}
}