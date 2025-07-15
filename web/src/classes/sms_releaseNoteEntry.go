/**
 * Author:    Admiral Helmut
 * Created:   13.10.2024
 *
 * (C)
 **/

package classes


type Sms_ReleaseNoteEntry struct {
	ElementType           string
	ElementID             int
	Name                  string
	ReleaseNote           string
	ReleaseDate           string
	IntroducedInVersion   string
}

func NewSms_ReleaseNoteEntry(elementType string, elementID int, name string, releaseNote string, releaseDate string, introducedInVersion string) *Sms_ReleaseNoteEntry {
	return &Sms_ReleaseNoteEntry{
		ElementType:         elementType,
		ElementID:           elementID,
		Name:                name,
		ReleaseNote:         releaseNote,
		ReleaseDate:         releaseDate,
		IntroducedInVersion: introducedInVersion,
	}
}