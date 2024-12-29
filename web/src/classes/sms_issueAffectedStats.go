/**
 * Author:    Admiral Helmut
 * Created:   28.12.2024
 *
 * (C)
 **/

package classes

// AffectedStats represents the statistics of affected devices and projects.
type Sms_IssueAffectedStats struct {
	AffectedDeviceInstances        		int `db:"affected_device_instances"`        // Anzahl betroffener Device-Instanzen
	AffectedDevicesWithoutInstances 	int `db:"affected_devices_without_instances"` // Anzahl betroffener Geräte ohne Instanzen
	AffectedProjects               		int `db:"affected_projects"`               // Anzahl betroffener Projekte
	DistinctDeviceVersionCombinations 	int `db:"distinct_device_version_combinations"` // Anzahl einzigartiger Gerät+Version-Kombinationen
}

func NewSms_IssueAffectedStats(affectedDeviceInstances int, affectedDevicesWithoutInstances int, affectedProjects int, distinctDeviceVersionCombinations int) *Sms_IssueAffectedStats {
	return &Sms_IssueAffectedStats{AffectedDeviceInstances: affectedDeviceInstances, AffectedDevicesWithoutInstances: affectedDevicesWithoutInstances, AffectedProjects: affectedProjects, DistinctDeviceVersionCombinations: distinctDeviceVersionCombinations}
}
