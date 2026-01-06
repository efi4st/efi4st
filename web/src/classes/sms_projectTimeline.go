/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package classes

type Sms_ProjectTimelineItem struct {
	ProjectID int

	Source   string // doc|updateExecution|updateHistory
	SourceID int

	OccurredAt string // oder time.Time, je nach deinem Scan
	CreatedAt  string
	Actor      string

	EntryType string
	Title     string
	Body      string

	AccessGroup *string

	DeviceInstanceID *int
	UpdateCenterID   *int
	UpdateID         *int
	PackageID        *int
	FromSystemID      *int
	ToSystemID        *int
	ExecStatus        *string
	ExecExitCode      *int

	DeviceInstanceSerialnumber *string
	DeviceType                 *string
	DeviceVersion              *string

	FromSystemType    *string
	FromSystemVersion *string
	ToSystemType      *string
	ToSystemVersion   *string

	// UI convenience
	Assets []Sms_ProjectDocAsset
}

type Sms_ProjectDocAsset struct {
	AssetID int
	EntryID int

	Kind    string // image|file
	Storage string // file|db

	Mime            string
	OriginalFile    *string
	StoredFilename  *string
	FilePath        *string
	FileSize        *int
	Sha256          *string
	Width           *int
	Height          *int
	CreatedAt       *string
	CreatedBy       *string
}