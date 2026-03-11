package utils

import "sort"
import "github.com/efi4st/efi4st/classes"

func MostCommonString(items []string) string {
	if len(items) == 0 { return "" }
	counts := map[string]int{}
	for _, s := range items { counts[s]++ }
	type kv struct{ k string; v int }
	var arr []kv
	for k, v := range counts { arr = append(arr, kv{k,v}) }
	sort.Slice(arr, func(i,j int) bool { return arr[i].v > arr[j].v })
	return arr[0].k
}

func BuildLiveProjectStateFromReport(lr classes.LiveReportV1, createdAt, receivedAt string) *classes.LiveProjectState {
	state := &classes.LiveProjectState{
		CreatedAt: createdAt,
		ReceivedAt: receivedAt,
		DeviceVersionByType: map[string]string{},
		SoftwareVersionByType: map[string]map[string]string{},
	}

	// Collect versions per device_type
	devVers := map[string][]string{}
	swVers := map[string]map[string][]string{} // device_type -> sw_name -> []versions

	for _, d := range lr.Devices {
		if d.DeviceType != "" && d.DeviceVersion != "" {
			devVers[d.DeviceType] = append(devVers[d.DeviceType], d.DeviceVersion)
		}
		if _, ok := swVers[d.DeviceType]; !ok {
			swVers[d.DeviceType] = map[string][]string{}
		}
		for _, sw := range d.Software {
			if sw.Name == "" || sw.Version == "" { continue }
			swVers[d.DeviceType][sw.Name] = append(swVers[d.DeviceType][sw.Name], sw.Version)
		}
	}

	for dt, versions := range devVers {
		state.DeviceVersionByType[dt] = MostCommonString(versions)
	}
	for dt, bySw := range swVers {
		if _, ok := state.SoftwareVersionByType[dt]; !ok {
			state.SoftwareVersionByType[dt] = map[string]string{}
		}
		for swName, versions := range bySw {
			state.SoftwareVersionByType[dt][swName] = MostCommonString(versions)
		}
	}

	return state
}