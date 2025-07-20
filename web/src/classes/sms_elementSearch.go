/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_ElementSearch struct {
	EntityType string `json:"entity_type"`
	EntityID   int    `json:"entity_id"`
	Name       string `json:"name"`
	Version    string `json:"version"`
	Type       string `json:"type"`
	Systems    string `json:"systems"`
}

func NewSms_ElementSearch(entityType string, entityID int, name, version, typeStr, systems string) *Sms_ElementSearch {
	return &Sms_ElementSearch{
		EntityType: entityType,
		EntityID:   entityID,
		Name:       name,
		Version:    version,
		Type:       typeStr,
		Systems:    systems,
	}
}