/**
 * Author:    Admiral Helmut
 * Created:   13.10.2024
 *
 * (C)
 **/

package classes

type Sms_UpdateHistory struct {
	UpdateHistory_id int `db:"updateHistory_id"`
	DeviceInstance_id int `db:"deviceInstance_id"`
	DeviceInstance_name string `db:"deviceInstance_name"`
	User string `db:"user"`
	UpdateType string `db:"updateType"`
	Date string `db:"date"`
	Description string `db:"description"`
}

func NewSms_UpdateHistory(DeviceInstance_id int, DeviceInstance_name string, User string, UpdateType string, Date string, Description string) *Sms_UpdateHistory {
	return &Sms_UpdateHistory{DeviceInstance_id: DeviceInstance_id, DeviceInstance_name: DeviceInstance_name, User: User, UpdateType: UpdateType, Date: Date, Description: Description}
}

func NewSms_UpdateHistoryFromDB(UpdateHistory_id int, DeviceInstance_id int, DeviceInstance_name string, User string, UpdateType string, Date string, Description string) *Sms_UpdateHistory {
	return &Sms_UpdateHistory{UpdateHistory_id: UpdateHistory_id, DeviceInstance_id: DeviceInstance_id, DeviceInstance_name: DeviceInstance_name, User: User, UpdateType: UpdateType, Date: Date, Description: Description}
}

