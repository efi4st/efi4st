package utils

import "sort"
import "github.com/efi4st/efi4st/classes"

func MostCommonString(items []string) string {
	if len(items) == 0 {
		return ""
	}

	counts := map[string]int{}

	for _, s := range items {
		if s == "" {
			continue
		}

		counts[s]++
	}

	if len(counts) == 0 {
		return ""
	}

	type kv struct {
		key   string
		count int
	}

	arr := make([]kv, 0, len(counts))

	for key, count := range counts {
		arr = append(arr, kv{
			key:   key,
			count: count,
		})
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i].count != arr[j].count {
			return arr[i].count > arr[j].count
		}

		// Deterministischer Tie-Breaker.
		return arr[i].key < arr[j].key
	})

	return arr[0].key
}

func BuildLiveProjectStateFromReport(
	lr classes.LiveReportV1,
	createdAt string,
	receivedAt string,
) *classes.LiveProjectState {
	state := &classes.LiveProjectState{
		CreatedAt:  createdAt,
		ReceivedAt: receivedAt,

		DeviceVersionByType:   map[string]string{},
		SoftwareVersionByType: map[string]map[string]string{},

		DeviceBySerialnumber: map[string]classes.LiveDeviceState{},
	}

	// Aggregation für die bisherige View.
	devVers := map[string][]string{}
	swVers := map[string]map[string][]string{}

	for _, d := range lr.Devices {
		// Konkrete Instanz erhalten.
		liveDevice := classes.LiveDeviceState{
			Serialnumber:  d.Serialnumber,
			DeviceType:    d.DeviceType,
			DeviceVersion: d.DeviceVersion,
			Software:      make([]classes.LiveSoftwareState, 0, len(d.Software)),
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

		if d.Serialnumber != "" {
			state.DeviceBySerialnumber[d.Serialnumber] = liveDevice
		}

		// Bisherige Aggregation nach Gerätetyp.
		if d.DeviceType != "" && d.DeviceVersion != "" {
			devVers[d.DeviceType] = append(
				devVers[d.DeviceType],
				d.DeviceVersion,
			)
		}

		if d.DeviceType == "" {
			continue
		}

		if _, ok := swVers[d.DeviceType]; !ok {
			swVers[d.DeviceType] = map[string][]string{}
		}

		for _, sw := range d.Software {
			if sw.Name == "" || sw.Version == "" {
				continue
			}

			swVers[d.DeviceType][sw.Name] = append(
				swVers[d.DeviceType][sw.Name],
				sw.Version,
			)
		}
	}

	for deviceType, versions := range devVers {
		state.DeviceVersionByType[deviceType] = MostCommonString(versions)
	}

	for deviceType, softwareVersions := range swVers {
		if _, ok := state.SoftwareVersionByType[deviceType]; !ok {
			state.SoftwareVersionByType[deviceType] = map[string]string{}
		}

		for softwareName, versions := range softwareVersions {
			state.SoftwareVersionByType[deviceType][softwareName] =
				MostCommonString(versions)
		}
	}

	return state
}