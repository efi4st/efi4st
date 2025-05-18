/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_ArtefactPartOfDeviceInstance struct {
	DeviceInstanceID int    `db:"deviceInstance_id"`
	ArtefactID       int    `db:"artefact_id"`
	AdditionalInfo   string `db:"additionalInfo"`
	ArtefactType     string `db:"artefactType"`
	Version          string `db:"version"`
}

type Sms_ArtefactPartOfDeviceInstanceDetailed struct {
	DeviceInstanceID  int
	ArtefactID        int
	AdditionalInfo    string
	DeviceType        string
	DeviceVersion     string
	SerialNumber      string
	ArtefactType      string
	ArtefactName      string
	ArtefactVersion   string
	OverridesModel    bool // ⬅️ NEU!
}

func NewSms_ArtefactPartOfDeviceInstance(deviceInstanceID int, artefactID int, additionalInfo string, artefactType string, version string) *Sms_ArtefactPartOfDeviceInstance {
	return &Sms_ArtefactPartOfDeviceInstance{
		DeviceInstanceID: deviceInstanceID,
		ArtefactID:       artefactID,
		AdditionalInfo:   additionalInfo,
		ArtefactType:     artefactType,
		Version:          version,
	}
}


func NewSms_ArtefactPartOfDeviceInstanceDetailed(
	deviceInstanceID int,
	artefactID int,
	additionalInfo string,
	deviceType string,
	deviceVersion string,
	serialNumber string,
	artefactType string,
	artefactName string,
	artefactVersion string,
	overridesModel bool, // ⬅️ NEU
) *Sms_ArtefactPartOfDeviceInstanceDetailed {
	return &Sms_ArtefactPartOfDeviceInstanceDetailed{
		DeviceInstanceID: deviceInstanceID,
		ArtefactID:       artefactID,
		AdditionalInfo:   additionalInfo,
		DeviceType:       deviceType,
		DeviceVersion:    deviceVersion,
		SerialNumber:     serialNumber,
		ArtefactType:     artefactType,
		ArtefactName:     artefactName,
		ArtefactVersion:  artefactVersion,
		OverridesModel:   overridesModel,
	}
}