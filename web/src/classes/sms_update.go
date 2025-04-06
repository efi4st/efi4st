/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

import "time"

type Sms_Update struct {
	ID                    int    `db:"update_id"`
	UpdateType            string `db:"update_type"`
	IsApproved            bool   `db:"is_approved"`
	CreatedAt             string `db:"created_at"`
	FromSystemType        string `db:"from_system_type"`         // Systemname des From-Systems
	FromSystemVersion     string `db:"from_system_version"`      // Version des From-Systems
	ToSystemType          string `db:"to_system_type"`           // Systemname des To-Systems
	ToSystemVersion       string `db:"to_system_version"`        // Version des To-Systems
	MandatorySystemType   string `db:"mandatory_system_type"`    // Systemname des Mandatory-Systems
	MandatorySystemVersion string `db:"mandatory_system_version"` // Version des Mandatory-Systems
}

type Sms_UpdateDetails struct {
	ID                     int
	FromSystemID           int
	ToSystemID             int
	MandatorySystemID      int
	UpdateType             string
	AdditionalInfo         string
	IsApproved             bool
	IssueLink              string
	ProjectName            string
	CreatedAt              time.Time
	FromSystemTypeID       int     // systemtype_id von From-System
	FromSystemType         string  // Name des From-Systems
	FromSystemVersion      string  // Version des From-Systems
	ToSystemTypeID         int     // systemtype_id von To-System
	ToSystemType           string  // Name des To-Systems
	ToSystemVersion        string  // Version des To-Systems
	MandatorySystemTypeID  int     // systemtype_id von Mandatory-System
	MandatorySystemType    string  // Name des Mandatory-Systems
	MandatorySystemVersion string  // Version des Mandatory-Systems
}

