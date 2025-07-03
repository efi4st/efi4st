/**
 * Author:    Admiral Helmut
 * Created:   29.10.2024
 *
 * (C)
 **/

package classes

type Component struct {
	Name    string
	Version string
}

type Software struct {
	Type       string
	Version    string
	Components []Component
}

type ProjectDeviceStructure struct {
	DeviceType string
	DeviceVersion string
	SerialNumber  string
	Software      []Software
}