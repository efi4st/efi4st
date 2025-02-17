/**
 * Author:    Admiral Helmut
 * Created:   14.02.2025
 *
 * (C)
 **/

package classes

type ProjectSetting struct {
	SettingID    int    `db:"setting_id"`    // Eindeutige ID der Einstellung
	KeyName      string `db:"key_name"`       // Schlüsselname der Einstellung
	ValueType    string `db:"value_type"`     // Datentyp des Wertes (z. B. 'string', 'int', 'boolean', 'json')
	DefaultValue string `db:"default_value"`  // Standardwert der Einstellung
}

func NewProjectSetting(keyName string, valueType string, defaultValue string) *ProjectSetting {
	return &ProjectSetting{KeyName: keyName, ValueType: valueType, DefaultValue: defaultValue}
}

func NewProjectSettingFromDB(settingID int, keyName string, valueType string, defaultValue string) *ProjectSetting {
	return &ProjectSetting{SettingID: settingID, KeyName: keyName, ValueType: valueType, DefaultValue: defaultValue}
}

