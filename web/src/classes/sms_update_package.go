/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_UpdatePackage struct {
	ID                 int     `db:"package_id"`
	UpdateID           int     `db:"update_id"`
	DeviceTypeID       int     `db:"device_type_id"`
	PackageIdentifier  string  `db:"package_identifier"`
	PackageVersion     string  `db:"package_version"`
	PackageName        string  `db:"package_name"`
	PackageDescription *string `db:"package_description"`
	UpdatePackageFile  string  `db:"update_package_file"`
	Creator            string  `db:"creator"`
	IsTested           bool    `db:"is_tested"`
	CreatedAt          string  `db:"created_at"`
}