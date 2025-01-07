/**
 * Author:    Admiral Helmut
 * Created:   28.12.2024
 *
 * (C)
 **/

package classes

type Sms_SecurityReportLink struct {
	ReportID         int    `db:"report_id"`          // ID des Reports
	LinkedObjectID   int    `db:"linked_object_id"`   // ID des verknüpften Objekts
	LinkedObjectType string `db:"linked_object_type"` // Typ des verknüpften Objekts ('sms_device', 'sms_application', 'sms_system')
}