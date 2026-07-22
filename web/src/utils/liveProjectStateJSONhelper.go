package utils

import "github.com/efi4st/efi4st/classes"


func BuildLiveProjectStateFromReport(
	lr classes.LiveReportV1,
	createdAt string,
	receivedAt string,
) *classes.LiveProjectState {
	state := &classes.LiveProjectState{
		CreatedAt:  createdAt,
		ReceivedAt: receivedAt,

		DeviceBySerialnumber: map[string]classes.LiveDeviceState{},
	}

	for _, d := range lr.Devices {
		liveDevice := classes.LiveDeviceState{
			Serialnumber:  d.Serialnumber,
			DeviceType:    d.DeviceType,
			DeviceVersion: d.DeviceVersion,
			Software:      make(
				[]classes.LiveSoftwareState,
				0,
				len(d.Software),
			),
		}

		for _, sw := range d.Software {
			liveDevice.Software = append(
				liveDevice.Software,
				classes.LiveSoftwareState{
					Name:    sw.Name,
					Version: sw.Version,
				},
			)
		}

		if d.Serialnumber == "" {
			continue
		}

		state.DeviceBySerialnumber[d.Serialnumber] = liveDevice
	}

	return state
}