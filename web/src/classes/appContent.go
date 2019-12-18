/**
 * Author:    Admiral Helmut
 * Created:   13.11.2019
 *
 * (C)
 **/

package classes

type AppContent struct {
	appContent_id int `db:"appContent_id"`
	contentPathList string `db:"contentPathList"`
	binwalkOutput string `db:"binwalkOutput"`
	relevantApps_path string `db:"relevantApps_path"`
}

func NewAppContent(appContent_id int, contentPathList string, binwalkOutput string, relevantApps_path string) *AppContent {
	return &AppContent{appContent_id: appContent_id, contentPathList: contentPathList, binwalkOutput: binwalkOutput, relevantApps_path: relevantApps_path}
}

func (a *AppContent) BinwalkOutput() string {
	return a.binwalkOutput
}

func (a *AppContent) SetBinwalkOutput(binwalkOutput string) {
	a.binwalkOutput = binwalkOutput
}

func (a *AppContent) RelevantApps_path() string {
	return a.relevantApps_path
}

func (a *AppContent) SetRelevantApps_path(relevantApps_path string) {
	a.relevantApps_path = relevantApps_path
}

func (a *AppContent) ContentPathList() string {
	return a.contentPathList
}

func (a *AppContent) SetContentPathList(contentPathList string) {
	a.contentPathList = contentPathList
}

func (a *AppContent) AppContent_id() int {
	return a.appContent_id
}

func (a *AppContent) SetAppContent_id(appContent_id int) {
	a.appContent_id = appContent_id
}

