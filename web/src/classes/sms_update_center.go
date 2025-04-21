/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

import "time"

type Sms_UpdateCenter struct {
	ID           int        `json:"id"`
	ProjectID    int        `json:"project_id"`
	UpdaterID    int        `json:"updater_id"`
	UpdaterType  string     `json:"updater_type"`
	Version      string     `json:"version"`
	Environment  string     `json:"environment"`
	Status       string     `json:"status"`
	Description  string     `json:"description"`
	Note         string     `json:"note"`
	Owner        string     `json:"owner"`
	LastContact  *time.Time `json:"last_contact"`
	CreatedAt    time.Time  `json:"created_at"`
}