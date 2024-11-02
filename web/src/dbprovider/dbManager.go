/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package dbprovider

import (
	"database/sql"
	"fmt"
	"github.com/efi4st/efi4st/classes"
	"github.com/efi4st/efi4st/dbUtils"
	"github.com/efi4st/efi4st/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"sort"
	"strconv"
	"time"
)

type Manager interface {
	GetProjects() []classes.Project
	AddProject(projectName string) error
	GetProjectInfo(id int) *classes.Project
	RemoveProject(id int) error
	AddFirmware(firmwareName string, size int, proj_id int) (err error)
	GetFirmwareListForProject(id int) []classes.Firmware
	UpdateProjectsUploads(id int, count int) error
	GetFirmwares() []classes.Firmware
	RemoveFirmware(id int) error
	GetFirmwareInfo(id int) *classes.Firmware
	GetRelevantApps() []classes.RelevantApps
	AddRelevantApp(relevantAppName string, path string, extPort int, extProtocoll string, intInterface string, firmware_id int)(err error)
	RemoveRelevantApp(id int) error
	GetRelevantAppInfo(id int) *classes.RelevantApps
	GetAppListForFirmware(id int) []classes.RelevantApps
	GetTestResults() []classes.TestResult
	AddTestResult(moduleName string, path string, created time.Time, firmware_id int) (err error)
	RemoveTestResult(id int) error
	GetTestResultInfo(id int) *classes.TestResult
	GetResultListForFirmware(id int) []classes.TestResult
	GetRelevantAppByPath(path string, firmwareId int) int
	GetRelevantAppByName(name string, firmwareId int) int
	UpdateRelevantApp(column string, relevantApp_id string) error
	UpdateRelevantAppForInterface(column string, relevantApp_id string, port int, protocol string) error
	GetAppContent() *classes.AppContent
	AddAppContent(contentPathList string, binwalkOutput string, readelfOutput string, lddOutput string, straceOutput string, relevantApps_path string) error
	RemoveAppContent(id int) error
	GetAppContentForRelevantApp(id int) *classes.AppContent
	RemoveAppContentByRelevantAppPath(path string) error
	GetAppContentForRelevantAppByPath(path string) *classes.AppContent
	UpdateAppContent(id int, module string, content string) error
	GetAnalysisTools() []classes.AnalysisTool
	GetAnalysisToolInfo(id int) *classes.AnalysisTool
	AddAnalysisTool(analysisToolName string,  callPattern string) (err error)
	RemoveAnalysisTool(id int) error
	GetBinaryAnalysis(id int) *classes.BinaryAnalysis
	GetBinaryAnalysisForRelevantApp(id int) []classes.BinaryAnalysis
	GetBinaryAnalysisForRelevantAppAndTool(id int, toolId int) []classes.BinaryAnalysis
	AddBinaryAnalysis(toolOutput string, analysisTool_id int, relevantApps_id int)  error
	RemoveBinaryAnalysis(id int) error
	RemoveBinaryAnalysisByRelevantApp(id int) error
	UpdateBinaryAnalysis(id int, output string) error
	GetSMSProjects() []classes.Sms_Project
	AddSMSProject(projectName string, customer string, projecttypeId int, reference string) (err error)
	GetSMSProjectInfo(id int) *classes.Sms_Project
	UpdateSMSProjectsActive(id int, active bool) error
	RemoveSMSProject(id int) error
	AddSMSSystem(systemtypeId int, version string, date string) error
	GetSMSSystems() []classes.Sms_System
	GetSMSSystemInfo(id int) *classes.Sms_System
	RemoveSMSSystem(id int) error
	AddSMSDevice(devicetypeId int, version string, date string) error
	GetSMSDevice() []classes.Sms_Device
	GetSMSDeviceInfo(id int) *classes.Sms_Device
	RemoveSMSDevice(id int) error
	AddSMSDeviceInstance(project_id int, device_id int, serialnumber string, provisioner string, configuration string) error
	GetSMSDeviceInstances() []classes.Sms_DeviceInstance
	GetSMSDeviceInstanceInfo(id int) *classes.Sms_DeviceInstance
	RemoveSMSDeviceInstances(id int) error
	GetDeviceInstanceListForProject(id int) []classes.Sms_DeviceInstance
	GetSMSUpdateHistoryForDevice(id int) []classes.Sms_UpdateHistory
	AddSMSUpdateHistory(deviceInstance_id int, user string, updateType string, description string) error
	GetSMSUdateHistoryInfo(id int) *classes.Sms_UpdateHistory
	AddSMSIssue(name string, issueType string, reference string, criticality int, cve string, description string) error
	GetSMSIssues() []classes.Sms_Issue
	GetSMSIssueInfo(id int) *classes.Sms_Issue
	RemoveSMSIssue(id int) error
	AddSMSIssueAffectedDevice(device_id int, issue_id int, additionalInfo string, confirmed bool) error
	GetSMSIssueAffectedDevicesForIssueID(issue_id int) []classes.Sms_IssueAffectedDevice
	GetSMSIssuesForDevice(device_id int) []classes.Sms_IssueAffectedDevice
	RemoveSMSIssueAffectedDevice(id int) error
	GetSMSAffectedDeviceInstancesAndProjects(issue_id int) []classes.Sms_AffectedDeviceInstancesAndProjects
	GetSMSProjectTypes() []classes.Sms_ProjectType
	GetSMSSystemTypes() []classes.Sms_SystemType
	GetSMSDeviceTypes() []classes.Sms_DeviceType
	GetSMSSolutionsForIssue(issue_id int) []classes.Sms_Solution
	RemoveSMSSolution(id int) error
	AddSMSSolution(issue_id int, devicetype_id int, name string, description string, reference string) error
	GetSMSSolutionInfo(id int) *classes.Sms_Solution
	AddSMSArtefact(artefactype_id int, name string, version string) error
	GetSMSArtefact() []classes.Sms_Artefact
	GetSMSArtefactTypes() []classes.Sms_ArtefactType
	GetSMSArtefactInfo(id int) *classes.Sms_Artefact
	RemoveSMSArtefact(id int) error
	GetSMSReleaseNoteForDevice(id int) []classes.Sms_ReleaseNote
	AddSMSReleaseNote(device_id int, releaseNoteType string, details string) error
	GetSMSReleaseNoteInfo(id int) *classes.Sms_ReleaseNote
	AddSMSSoftware(softwaretype_id int, version string, license string, thirdParty bool, releaseNote string) error
	GetSMSSoftware() []classes.Sms_Software
	GetSMSSoftwareTypes() []classes.Sms_SoftwareType
	GetSMSSoftwareInfo(id int) *classes.Sms_Software
	RemoveSMSSoftware(id int) error
	AddSMSComponent(name string, componentType string, version string, license string, thirdParty bool, releaseNote string) error
	GetSMSComponent() []classes.Sms_Component
	GetSMSComponentInfo(id int) *classes.Sms_Component
	RemoveSMSComponent(id int) error
	AddSMSComponentPartOfSoftware(software_id int, component_id int, additionalInfo string) error
	GetSMSComponentPartOfSoftwareForSoftware(software_id int) []classes.Sms_ComponentPartOfSoftware
	GetSMSComponentPartOfSoftwareForComponent(component_id int) []classes.Sms_ComponentPartOfSoftware
	RemoveSMSComponentPartOfSoftware(id int) error
	AddSMSSoftwarePartOfDevice(device_id int, software_id int, additionalInfo string) error
	GetSMSSoftwarePartOfDeviceForDevice(device_id int) []classes.Sms_SoftwarePartOfDevice
	GetSMSSoftwarePartOfDeviceForSoftware(software_id int) []classes.Sms_SoftwarePartOfDevice
	RemoveSMSSoftwarePartOfDevice(id int) error
}

type manager struct {
	db *sqlx.DB
}

var dbMgr Manager
func GetDBManager() Manager { return dbMgr }

func init() {
	db, err := sqlx.Connect("mysql", "efi4db:efi4db@tcp(127.0.0.1:3306)/efi4st?parseTime=true")
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	dbMgr = &manager{db: db}
}

/////////////////////////////////////////
////	Project
////////////////////////////////////////
func (mgr *manager) AddProject(projectName string) (err error) {
	dt := time.Now()

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_newProject)
	if err != nil{
		fmt.Print(err)
	}

	rows, err := stmt.Query(projectName, 0, dt)

	if rows == nil{
		fmt.Println("rows should be null")
	}

	return err
}

func (mgr *manager) GetProjects() (projects []classes.Project) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_projects)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query()

	var ( 	dbId int
		  	dbName string
			dbUploads int
			dbDate time.Time	)

	for rows.Next() {
		err := rows.Scan(&dbId, &dbName, &dbUploads, &dbDate)
		var project = classes.NewProjectFromDB(dbId, dbName, dbUploads, dbDate)
		projects=append(projects, *project)
		if err != nil {
			log.Fatal(err)
		}
	}

	return projects
}

func (mgr *manager) GetProjectInfo(id int) (*classes.Project) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_projectInfo)
	if err != nil{
		fmt.Print(err)
	}

	var ( 	dbId int
			dbName string
			dbUploads int
			dbDate time.Time	)

	row := stmt.QueryRow(id)
	row.Scan(&dbId, &dbName, &dbUploads, &dbDate)

	var project = classes.NewProjectFromDB(dbId, dbName, dbUploads, dbDate)

	return project
}

func (mgr *manager) UpdateProjectsUploads(id int, count int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.UPDATE_projectUploads)

	stmt.QueryRow(count, id)

	return err
}

func (mgr *manager) RemoveProject(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_project)

	stmt.QueryRow(id)

	return err
}

func (mgr *manager) GetFirmwareListForProject(id int) (firmwares []classes.Firmware) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_firmwareForProject)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(id)

	var ( 	dbfirmware_id int
			dbname string
			dbversion string
			dbbinwalkOutput sql.NullString
			dbsizeInBytes int
			dbproject_id int
			created time.Time)

	for rows.Next() {
		err := rows.Scan(&dbfirmware_id, &dbname, &dbversion, &dbbinwalkOutput, &dbsizeInBytes, &dbproject_id, &created)
		var firmware = classes.NewFirmware(dbfirmware_id, dbname, dbversion, dbbinwalkOutput.String, dbsizeInBytes, dbproject_id, created)
		firmwares=append(firmwares, *firmware)
		if err != nil {
			log.Fatal(err)
		}
	}

	return firmwares
}

/////////////////////////////////////////
////	Firmware
////////////////////////////////////////
func (mgr *manager) GetFirmwares() (firmwares []classes.Firmware){
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_firmware)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query()


	var (	dbFirmware_id int
		dbName string
		dbVersion string
		dbBinwalkOutput sql.NullString
		dbSizeInBytes int
		dbProject_id int
		dbCreated time.Time
		dbProjectName string			)

	for rows.Next() {
		err := rows.Scan(&dbFirmware_id, &dbName, &dbVersion, &dbBinwalkOutput, &dbSizeInBytes, &dbProject_id, &dbCreated, &dbProjectName)
		var firmware = classes.NewFirmware(dbFirmware_id, dbName, dbVersion, dbBinwalkOutput.String, dbSizeInBytes, dbProject_id, dbCreated)
		//Set ProjectName as Msg
		firmware.SetMsg(dbProjectName)
		firmwares=append(firmwares, *firmware)
		if err != nil {
			log.Fatal(err)
		}
	}

	return firmwares
}

func (mgr *manager) AddFirmware(firmwareName string, size int, proj_id int) (err error) {
	dt := time.Now()

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_newFirmware)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(firmwareName, "", size, proj_id, dt)

	if rows == nil{
		fmt.Print(err)
	}

	return err
}

func (mgr *manager) RemoveFirmware(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_firmware)

	stmt.QueryRow(id)

	return err
}

func (mgr *manager) GetFirmwareInfo(id int) (*classes.Firmware) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_firmwareInfo)
	if err != nil{
		fmt.Print(err)
	}

	var ( 	dbFirmware_id int
			dbName string
			dbVersion string
			dbBinwalkOutput sql.NullString
			dbSizeInBytes int
			dbProject_id int
			dbCreated time.Time	)

	row := stmt.QueryRow(id)
	row.Scan(&dbFirmware_id, &dbName, &dbVersion, &dbBinwalkOutput, &dbSizeInBytes, &dbProject_id, &dbCreated)

	var firmware = classes.NewFirmware(dbFirmware_id, dbName, dbVersion, dbBinwalkOutput.String, dbSizeInBytes, dbProject_id, dbCreated)

	return firmware
}


/////////////////////////////////////////
////	relevant Apps
////////////////////////////////////////
func (mgr *manager) GetRelevantApps() (relevantApps []classes.RelevantApps){
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_relevantApps)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query()


	var (	dbrelevantApps_id int
		dbName string
		dbPath sql.NullString
		dbExtPort int
		dbExtProtocoll sql.NullString
		dbIntInterface sql.NullString
		dbFirmware_id int
		dbFirmwareName string		)

	for rows.Next() {
		err := rows.Scan(&dbrelevantApps_id, &dbName, &dbPath, &dbExtPort, &dbExtProtocoll, &dbIntInterface, &dbFirmware_id, &dbFirmwareName)
		var relevantApp = classes.NewRelevantApps(dbrelevantApps_id, dbName, dbPath.String, dbExtPort, dbExtProtocoll.String, dbIntInterface.String, dbFirmware_id)

		//Set FirmwareName as Msg
		relevantApp.SetMsg(dbFirmwareName)
		relevantApps=append(relevantApps, *relevantApp)
		if err != nil {
			log.Fatal(err)
		}
	}

	return relevantApps
}

func (mgr *manager) AddRelevantApp(relevantAppName string, path string, extPort int, extProtocoll string, intInterface string, firmware_id int) (err error) {

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_newrelevantApps)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(relevantAppName, path, extPort, extProtocoll, intInterface, firmware_id)

	if rows == nil{
		fmt.Print(err)
	}
	rows.Close()

	return err
}

func (mgr *manager) UpdateRelevantApp(column string, relevantApp_id string) (err error) {

	stmt, err := mgr.db.Prepare(dbUtils.UPDATE_relevantAppmoduleDefault)

	switch column {
	case "moduleDefault":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATE_relevantAppmoduleDefault)
	case "moduleInitSystem":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATE_relevantAppmoduleInitSystem)
	case "moduleFileContent":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATE_relevantAppmoduleFileContent)
	case "moduleBash":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATE_relevantAppmoduleBash)
	case "moduleCronJob":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATE_relevantAppmoduleCronJob)
	case "moduleProcesses":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATE_relevantAppmoduleProcesses)
	case "moduleInterfaces":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATE_relevantAppmoduleInterfaces)
	case "moduleSystemControls":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATE_relevantAppmoduleSystemControls)
	case "moduleFileSystem":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATE_relevantAppmoduleFileSystem)
	case "modulePortscanner":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATE_relevantAppmodulePortscanner)
	case "moduleProtocolls":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATE_relevantAppmoduleProtocolls)
	case "moduleNetInterfaces":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATE_relevantAppmoduleNetInterfaces)
	case "moduleFileSystemInterfaces":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATE_relevantAppmoduleFileSystemInterfaces)
	case "moduleFileHandles":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATE_relevantAppmoduleFileHandles)

	default:
		fmt.Printf("Error updating relevant app! Unknown column!");
	}

	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(1, relevantApp_id)

	if rows == nil{
		fmt.Print(err)
	}
	rows.Close()

	return err
}

func (mgr *manager) UpdateRelevantAppForInterface(column string, relevantApp_id string, port int, protocol string) (err error) {

	stmt, err := mgr.db.Prepare(dbUtils.UPDATEWITHINTERFACE_relevantAppmoduleDefault)

	switch column {
	case "moduleDefault":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATEWITHINTERFACE_relevantAppmoduleDefault)
	case "moduleInitSystem":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATEWITHINTERFACE_relevantAppmoduleInitSystem)
	case "moduleFileContent":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATEWITHINTERFACE_relevantAppmoduleFileContent)
	case "moduleBash":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATEWITHINTERFACE_relevantAppmoduleBash)
	case "moduleCronJob":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATEWITHINTERFACE_relevantAppmoduleCronJob)
	case "moduleProcesses":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATEWITHINTERFACE_relevantAppmoduleProcesses)
	case "moduleInterfaces":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATEWITHINTERFACE_relevantAppmoduleInterfaces)
	case "moduleSystemControls":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATEWITHINTERFACE_relevantAppmoduleSystemControls)
	case "moduleFileSystem":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATEWITHINTERFACE_relevantAppmoduleFileSystem)
	case "modulePortscanner":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATEWITHINTERFACE_relevantAppmodulePortscanner)
	case "moduleProtocolls":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATEWITHINTERFACE_relevantAppmoduleProtocolls)
	case "moduleNetInterfaces":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATEWITHINTERFACE_relevantAppmoduleNetInterfaces)
	case "moduleFileSystemInterfaces":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATEWITHINTERFACE_relevantAppmoduleFileSystemInterfaces)
	case "moduleFileHandles":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATEWITHINTERFACE_relevantAppmoduleFileHandles)

	default:
		fmt.Printf("Error updating relevant app! Unknown column!");
	}

	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(1, port, protocol, relevantApp_id)
	fmt.Printf("fvfvfvfv")
	fmt.Printf(relevantApp_id)
	fmt.Printf(strconv.Itoa(port))
	fmt.Printf(protocol)


	if rows == nil{
		fmt.Print(err)
	}
	rows.Close()

	return err
}

func (mgr *manager) RemoveRelevantApp(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_relevantApps)

	stmt.QueryRow(id)

	return err
}

func (mgr *manager) GetRelevantAppInfo(id int) (*classes.RelevantApps) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_relevantAppInfo)
	if err != nil{
		fmt.Print(err)
	}

	var (	dbrelevantApps_id int
		dbName string
		dbPath sql.NullString
		dbExtPort int
		dbExtProtocoll sql.NullString
		dbIntInterface sql.NullString
		dbmoduleDefault sql.NullBool
		dbmoduleInitSystem sql.NullBool
		dbmoduleFileContent sql.NullBool
		dbmoduleBash sql.NullBool
		dbmoduleCronJob sql.NullBool
		dbmoduleProcesses sql.NullBool
		dbmoduleInterfaces sql.NullBool
		dbmoduleSystemControls sql.NullBool
		dbmoduleFileSystem sql.NullBool
		dbmodulePortscanner sql.NullBool
		dbmoduleProtocolls sql.NullBool
		dbmoduleNetInterfaces sql.NullBool
		dbmoduleFileSystemInterfaces sql.NullBool
		dbmoduleFileHandles sql.NullBool
		dbFirmware_id int			)

	row := stmt.QueryRow(id)
	row.Scan(&dbrelevantApps_id, &dbName, &dbPath, &dbExtPort, &dbExtProtocoll, &dbIntInterface, &dbmoduleDefault, &dbmoduleInitSystem, &dbmoduleFileContent, &dbmoduleBash, &dbmoduleCronJob, &dbmoduleProcesses, &dbmoduleInterfaces, &dbmoduleSystemControls, &dbmoduleFileSystem, &dbmodulePortscanner, &dbmoduleProtocolls, &dbmoduleNetInterfaces, &dbmoduleFileSystemInterfaces, &dbmoduleFileHandles, &dbFirmware_id)

	var relevantApp = classes.NewRelevantApps(dbrelevantApps_id, dbName, dbPath.String, dbExtPort, dbExtProtocoll.String, dbIntInterface.String, dbFirmware_id)
	relevantApp.SetModuleDefault(dbmoduleDefault.Bool)
	relevantApp.SetModuleInitSystem(dbmoduleInitSystem.Bool)
	relevantApp.SetModuleFileContent(dbmoduleFileContent.Bool)
	relevantApp.SetModuleBash(dbmoduleBash.Bool)
	relevantApp.SetModuleCronJob(dbmoduleCronJob.Bool)
	relevantApp.SetModuleProcesses(dbmoduleProcesses.Bool)
	relevantApp.SetModuleInterfaces(dbmoduleInterfaces.Bool)
	relevantApp.SetModuleSystemControls(dbmoduleSystemControls.Bool)
	relevantApp.SetModuleFileSystem(dbmoduleFileSystem.Bool)
	relevantApp.SetModulePortscanner(dbmodulePortscanner.Bool)
	relevantApp.SetModuleProtocolls(dbmoduleProtocolls.Bool)
	relevantApp.SetModuleNetInterfaces(dbmoduleNetInterfaces.Bool)
	relevantApp.SetModuleFileSystemInterfaces(dbmoduleFileSystemInterfaces.Bool)
	relevantApp.SetModuleFileHandles(dbmoduleFileHandles.Bool)

	return relevantApp
}

func (mgr *manager) GetRelevantAppByPath(path string, firmwareId int) (appId int) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_relevantAppByPath)
	if err != nil{
		fmt.Print(err)
	}

	var (	dbrelevantApps_id int
									)
	row := stmt.QueryRow(path, firmwareId)
	row.Scan(&dbrelevantApps_id)
	return dbrelevantApps_id
}

func (mgr *manager) GetRelevantAppByName(name string, firmwareId int) (appId int) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_relevantAppByName)
	if err != nil{
		fmt.Print(err)
	}

	var (	dbrelevantApps_id int
	)
	row := stmt.QueryRow(name, firmwareId)
	row.Scan(&dbrelevantApps_id)
	return dbrelevantApps_id
}

func (mgr *manager) GetAppListForFirmware(id int) (relevantApps []classes.RelevantApps) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_relevantAppsForFirmware)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(id)

	var (	dbrelevantApps_id int
		dbName string
		dbPath sql.NullString
		dbExtPort int
		dbExtProtocoll sql.NullString
		dbIntInterface sql.NullString
		dbmoduleDefault sql.NullBool
		dbmoduleInitSystem sql.NullBool
		dbmoduleFileContent sql.NullBool
		dbmoduleBash sql.NullBool
		dbmoduleCronJob sql.NullBool
		dbmoduleProcesses sql.NullBool
		dbmoduleInterfaces sql.NullBool
		dbmoduleSystemControls sql.NullBool
		dbmoduleFileSystem sql.NullBool
		dbmodulePortscanner sql.NullBool
		dbmoduleProtocolls sql.NullBool
		dbmoduleNetInterfaces sql.NullBool
		dbmoduleFileSystemInterfaces sql.NullBool
		dbmoduleFileHandles sql.NullBool
		dbFirmware_id int			)

	for rows.Next() {
		err := rows.Scan(&dbrelevantApps_id, &dbName, &dbPath, &dbExtPort, &dbExtProtocoll, &dbIntInterface, &dbmoduleDefault, &dbmoduleInitSystem, &dbmoduleFileContent, &dbmoduleBash, &dbmoduleCronJob, &dbmoduleProcesses, &dbmoduleInterfaces, &dbmoduleSystemControls, &dbmoduleFileSystem, &dbmodulePortscanner, &dbmoduleProtocolls, &dbmoduleNetInterfaces, &dbmoduleFileSystemInterfaces, &dbmoduleFileHandles, &dbFirmware_id)
		var relevantApp = classes.NewRelevantApps(dbrelevantApps_id, dbName, dbPath.String, dbExtPort, dbExtProtocoll.String, dbIntInterface.String, dbFirmware_id)
		count:= 0
		if(dbmoduleDefault.Bool){
			count++
		}
		if(dbmoduleInitSystem.Bool){
			count++
		}
		if(dbmoduleFileContent.Bool){
			count++
		}
		if(dbmoduleBash.Bool){
			count++
		}
		if(dbmoduleCronJob.Bool){
			count++
		}
		if(dbmoduleProcesses.Bool){
			count++
		}
		if(dbmoduleInterfaces.Bool){
			count++
		}
		if(dbmoduleSystemControls.Bool){
			count++
		}
		if(dbmoduleFileSystem.Bool){
			count++
		}
		if(dbmodulePortscanner.Bool){
			count++
		}
		if(dbmoduleProtocolls.Bool){
			count++
		}
		if(dbmoduleNetInterfaces.Bool){
			count++
		}
		if(dbmoduleFileSystemInterfaces.Bool){
			count++
		}
		if(dbmoduleFileHandles.Bool){
			count++
		}

		relevantApp.SetMsg(strconv.Itoa(count))
		relevantApps=append(relevantApps, *relevantApp)
		if err != nil {
			log.Fatal(err)
		}
	}

	sort.Sort(utils.RelevantAppsByScoreSorter(relevantApps))

	return relevantApps
}

func (mgr *manager) GetResultListForFirmware(id int) (firmwareResults []classes.TestResult) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_resultsForFirmware)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(id)
	
	var (	dbTestResult_id int
			dbModuleName string
			dbCreated time.Time
			dbFirmware_id int	)

	for rows.Next() {
		err := rows.Scan(&dbTestResult_id, &dbModuleName, &dbCreated, &dbFirmware_id)
		var firmwareResult = classes.NewTestResult(dbTestResult_id, dbModuleName, "", dbCreated, dbFirmware_id)

		firmwareResults=append(firmwareResults, *firmwareResult)
		if err != nil {
			log.Fatal(err)
		}
	}

	return firmwareResults
}

/////////////////////////////////////////
////	Test Results
////////////////////////////////////////
func (mgr *manager) GetTestResults() (testResults []classes.TestResult){
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_results)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query()


	var (	dbTestResult_id int
			dbModuleName string
			dbResult sql.NullString
			dbCreated time.Time
			dbFirmware_id int
			dbFirmwareName string	)

	for rows.Next() {
		err := rows.Scan(&dbTestResult_id, &dbModuleName, &dbResult, &dbCreated, &dbFirmware_id, &dbFirmwareName)
		var testResult = classes.NewTestResult(dbTestResult_id, dbModuleName, dbResult.String, dbCreated, dbFirmware_id)

		//Set FirmwareName as Msg
		testResult.SetMsg(dbFirmwareName)
		testResults=append(testResults, *testResult)
		if err != nil {
			log.Fatal(err)
		}
	}

	return testResults
}

func (mgr *manager) AddTestResult(moduleName string, result string, created time.Time, firmware_id int) (err error) {

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_newresults)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(moduleName , result , created, firmware_id)

	if rows == nil{
		fmt.Print(err)
	}

	return err
}

func (mgr *manager) RemoveTestResult(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_result)

	stmt.QueryRow(id)

	return err
}

func (mgr *manager) GetTestResultInfo(id int) (*classes.TestResult) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_resultInfo)
	if err != nil{
		fmt.Print(err)
	}

	var (	dbTestResult_id int
			dbModuleName string
			dbResult sql.NullString
			dbCreated time.Time
			dbFirmware_id int
			dbFirmwareName string)

	row := stmt.QueryRow(id)
	row.Scan(&dbTestResult_id, &dbModuleName, &dbResult, &dbCreated, &dbFirmware_id, &dbFirmwareName)
	var testResult = classes.NewTestResult(dbTestResult_id, dbModuleName, dbResult.String, dbCreated, dbFirmware_id)

	testResult.SetMsg(dbFirmwareName)

	return testResult
}


/////////////////////////////////////////
////	App Content
////////////////////////////////////////
func (mgr *manager) GetAppContent() (appcontent *classes.AppContent){
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_appContent)
	if err != nil{
		fmt.Print(err)
	}

	var (	dbAppContent_id int
		dbContentPathList sql.NullString
		dbBinwalkOutput sql.NullString
		dbReadelfOutput sql.NullString
		dbLDDOutput sql.NullString
		dbStraceOutput sql.NullString
		dbRelevantApps_id string			)

	row := stmt.QueryRow()
	row.Scan(&dbAppContent_id, &dbContentPathList, &dbBinwalkOutput, &dbReadelfOutput, &dbLDDOutput, &dbStraceOutput, &dbRelevantApps_id)

	appcontent = classes.NewAppContent(dbAppContent_id, dbContentPathList.String, dbBinwalkOutput.String, dbReadelfOutput.String, dbLDDOutput.String, dbStraceOutput.String, dbRelevantApps_id)

	return appcontent
}

func (mgr *manager) AddAppContent(contentPathList string, binwalkOutput string, readelfOutput string, lddOutput string, straceOutput string, relevantApps_path string) (err error) {

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_newappContent)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(contentPathList, binwalkOutput, readelfOutput, lddOutput, straceOutput, relevantApps_path)

	if rows == nil{
		fmt.Print(err)
	}

	return err
}

func (mgr *manager) UpdateAppContent(id int, module string, content string) (err error) {

	stmt, err := mgr.db.Prepare(dbUtils.UPDATE_appContentbinwalk)

	switch module {
	case "binwalk":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATE_appContentbinwalk)
	case "readelf":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATE_appContentreadelf)
	case "ldd":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATE_appContentldd)
	case "strace":
		stmt, err = mgr.db.Prepare(dbUtils.UPDATE_appContentstrace)

	default:

	}
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(content, id)

	if rows == nil{
		fmt.Print(err)
	}

	return err
}

func (mgr *manager) RemoveAppContent(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_appContent)

	stmt.QueryRow(id)

	return err
}

func (mgr *manager) RemoveAppContentByRelevantAppPath(path string) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_appContentByRelevantAppPath)

	stmt.QueryRow(path)

	return err
}

func (mgr *manager) GetAppContentForRelevantApp(id int) (appContentInfo *classes.AppContent) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_appContentForRelevantApp)
	if err != nil{
		fmt.Print(err)
	}

	var (	dbAppContent_id int
		dbContentPathList sql.NullString
		dbBinwalkOutput sql.NullString
		dbReadelfOutput sql.NullString
		dbLDDOutput sql.NullString
		dbStraceOutput sql.NullString
		dbRelevantApps_id string			)

	row := stmt.QueryRow(id)

	err2 := row.Scan(&dbAppContent_id, &dbContentPathList, &dbBinwalkOutput, &dbReadelfOutput, &dbLDDOutput, &dbStraceOutput, &dbRelevantApps_id)
	if err2 != nil {
		return nil
	}

	appContentInfo = classes.NewAppContent(dbAppContent_id, dbContentPathList.String, dbBinwalkOutput.String, dbReadelfOutput.String, dbLDDOutput.String, dbStraceOutput.String, dbRelevantApps_id)

	return appContentInfo
}

func (mgr *manager) GetAppContentForRelevantAppByPath(path string) (appContentInfo *classes.AppContent) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_appContentByPath)
	if err != nil{
		fmt.Print(err)
	}

	var (	dbAppContent_id int
		dbContentPathList sql.NullString
		dbBinwalkOutput sql.NullString
		dbReadelfOutput sql.NullString
		dbLDDOutput sql.NullString
		dbStraceOutput sql.NullString
		dbRelevantApps_id string			)

	row := stmt.QueryRow(path)
	err2 := row.Scan(&dbAppContent_id, &dbContentPathList, &dbBinwalkOutput, &dbReadelfOutput, &dbLDDOutput, &dbStraceOutput, &dbRelevantApps_id)
	if err2 != nil {
		return nil
	}

	appContentInfo = classes.NewAppContent(dbAppContent_id, dbContentPathList.String, dbBinwalkOutput.String, dbReadelfOutput.String, dbLDDOutput.String, dbStraceOutput.String, dbRelevantApps_id)

	return appContentInfo
}


/////////////////////////////////////////
////	AnalysisTool
////////////////////////////////////////
func (mgr *manager) GetAnalysisTools() (analysisTools []classes.AnalysisTool){
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_analysisTool)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query()


	var (	dbAnalysisTool_id int
			dbName string
			dbExecutionString string			)

	for rows.Next() {
		err := rows.Scan(&dbAnalysisTool_id, &dbName, &dbExecutionString)
		var analysisTool = classes.NewAnalysisTool(dbAnalysisTool_id, dbName, dbExecutionString)
		analysisTools=append(analysisTools, *analysisTool)
		if err != nil {
			log.Fatal(err)
		}
	}

	return analysisTools
}

func (mgr *manager) GetAnalysisToolInfo(id int) (*classes.AnalysisTool) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_analysisToolInfo)
	if err != nil{
		fmt.Print(err)
	}

	var (	dbAnalysisTool_id int
			dbName string
			dbExecutionString string			)

	row := stmt.QueryRow(id)
	row.Scan(&dbAnalysisTool_id, &dbName, &dbExecutionString)

	var analysisTool = classes.NewAnalysisTool(dbAnalysisTool_id, dbName, dbExecutionString)

	return analysisTool
}

func (mgr *manager) AddAnalysisTool(analysisToolName string,  executionStringPattern string) (err error) {

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_newAnalysisTool)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(analysisToolName, executionStringPattern)

	if rows == nil{
		fmt.Print(err)
	}

	return err
}

func (mgr *manager) RemoveAnalysisTool(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_analysisTool)

	stmt.QueryRow(id)

	return err
}

/////////////////////////////////////////
////	BinaryAnalysis
////////////////////////////////////////
func (mgr *manager) GetBinaryAnalysis(id int) (binaryAnalysis *classes.BinaryAnalysis){
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_binaryAnalysis)
	if err != nil{
		fmt.Print(err)
	}

	var (	dbBinaryAnalysis_id int
			dbToolOutput sql.NullString
			dbAnalysisTool_id int
			dbRelevantApps_id int			)

	row := stmt.QueryRow(id)
	row.Scan(&dbBinaryAnalysis_id, &dbToolOutput, &dbAnalysisTool_id, &dbRelevantApps_id)

	binaryAnalysis = classes.NewBinaryAnalysis(dbBinaryAnalysis_id, dbToolOutput.String, dbAnalysisTool_id, dbRelevantApps_id)

	return binaryAnalysis
}

func (mgr *manager) GetBinaryAnalysisForRelevantApp(id int) (binaryAnalysisList []classes.BinaryAnalysis) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_binaryAnalysisForRelevantApp)
	if err != nil{
		fmt.Print(err)
	}

	var (	dbBinaryAnalysis_id int
			dbToolOutput sql.NullString
			dbAnalysisTool_id int
			dbRelevantApps_id int			)

	rows, err := stmt.Query(id)

	for rows.Next() {
		err := rows.Scan(&dbBinaryAnalysis_id, &dbToolOutput, &dbAnalysisTool_id, &dbRelevantApps_id)
		var binaryAnalysis = classes.NewBinaryAnalysis(dbBinaryAnalysis_id, dbToolOutput.String, dbAnalysisTool_id, dbRelevantApps_id)
		binaryAnalysisList=append(binaryAnalysisList, *binaryAnalysis)
		if err != nil {
			log.Fatal(err)
		}
	}

	return binaryAnalysisList
}

func (mgr *manager) GetBinaryAnalysisForRelevantAppAndTool(id int, toolId int) (binaryAnalysisList []classes.BinaryAnalysis) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_binaryAnalysisForRelevantAppAndTool)
	if err != nil{
		fmt.Print(err)
	}

	var (	dbBinaryAnalysis_id int
			dbToolOutput sql.NullString
			dbAnalysisTool_id int
			dbRelevantApps_id int			)

	rows, err := stmt.Query(id, toolId)

	for rows.Next() {
		err := rows.Scan(&dbBinaryAnalysis_id, &dbToolOutput, &dbAnalysisTool_id, &dbRelevantApps_id)
		var binaryAnalysis = classes.NewBinaryAnalysis(dbBinaryAnalysis_id, dbToolOutput.String, dbAnalysisTool_id, dbRelevantApps_id)
		binaryAnalysisList=append(binaryAnalysisList, *binaryAnalysis)
		if err != nil {
			log.Fatal(err)
		}
	}

	return binaryAnalysisList
}

func (mgr *manager) AddBinaryAnalysis(toolOutput string, analysisTool_id int, relevantApps_id int) (err error) {

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_newbinaryAnalysis)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(toolOutput, analysisTool_id, relevantApps_id)

	if rows == nil{
		fmt.Print(err)
	}

	return err
}

func (mgr *manager) RemoveBinaryAnalysis(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_binaryAnalysis)

	stmt.QueryRow(id)

	return err
}

func (mgr *manager) RemoveBinaryAnalysisByRelevantApp(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_binaryAnalysisByRelevantApp)

	stmt.QueryRow(id)

	return err
}

func (mgr *manager) UpdateBinaryAnalysis(id int, output string) (err error) {

	stmt, err := mgr.db.Prepare(dbUtils.UPDATE_binaryAnalysis)

	if err != nil{
		fmt.Print(err)
	}

	rows, err := stmt.Query(output, id)

	if rows == nil{
		fmt.Print(err)
	}

	return err
}


/**
 * Security Management System
 * Created:   29.09.2024
 *
 * (C)
 **/

/////////////////////////////////////////
////	SMS Project
////////////////////////////////////////
func (mgr *manager) AddSMSProject(projectName string, customer string, projecttypeId int, reference string) (err error) {
	dt := time.Now()
	act := false

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newProject)
	if err != nil{
		fmt.Print(err)
	}

	rows, err := stmt.Query(projectName, customer, projecttypeId, reference, dt, act)

	if rows == nil{
		fmt.Println("rows should be null")
	}

	return err
}

func (mgr *manager) GetSMSProjects() (projects []classes.Sms_Project) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_projects)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query()

	var ( 	dbId int
		dbName string
		dbCustomer string
		dbProjectType string
		dbReference string
		dbDate time.Time
		dbActive bool)

	for rows.Next() {
		err := rows.Scan(&dbId, &dbName, &dbCustomer, &dbProjectType, &dbReference, &dbDate, &dbActive)

		var project = classes.NewSms_ProjectFromDB(dbId, dbName, dbCustomer, dbProjectType, dbReference, dbDate.String(), dbActive)
		projects=append(projects, *project)
		if err != nil {
			log.Fatal(err)
		}
	}

	return projects
}

func (mgr *manager) GetSMSProjectTypes() (projectTypes []classes.Sms_ProjectType) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_projectTypes)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query()

	var ( 	dbProjecttype_id int
		dbProjectType string)

	for rows.Next() {
		err := rows.Scan(&dbProjecttype_id, &dbProjectType)

		var projectType = classes.NewSms_ProjectType(dbProjecttype_id, dbProjectType)
		projectTypes=append(projectTypes, *projectType)
		if err != nil {
			log.Fatal(err)
		}
	}

	return projectTypes
}

func (mgr *manager) GetSMSProjectInfo(id int) (*classes.Sms_Project) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_projectInfo)
	if err != nil{
		fmt.Print(err)
	}

	var ( 	dbId int
		dbName string
		dbCustomer string
		dbProjectType string
		dbReference string
		dbDate time.Time
		dbActive bool)

	row := stmt.QueryRow(id)
	row.Scan(&dbId, &dbName, &dbCustomer, &dbProjectType, &dbReference, &dbDate, &dbActive)

	var project = classes.NewSms_ProjectFromDB(dbId, dbName, dbCustomer, dbProjectType, dbReference, dbDate.String(), dbActive)

	return project
}

func (mgr *manager) UpdateSMSProjectsActive(id int, active bool) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.UPDATE_sms_projectActive)

	stmt.QueryRow(active, id)

	return err
}

func (mgr *manager) RemoveSMSProject(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_sms_project)

	stmt.QueryRow(id)

	return err
}

func (mgr *manager) GetDeviceInstanceListForProject(id int) (deviceInstances []classes.Sms_DeviceInstance){

	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_deviceInstancesForProject)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(id)

	var (	dbId int
		dbProject_id int
		dbDevice_id int
		dbSerialnumber string
		dbProvisioner string
		dbConfiguration string
		dbDate time.Time
		dbprojectName string
		dbdevicetypeId int
		dbdeviceVersion string
		dbdeviceType string	)

	for rows.Next() {
		err := rows.Scan(&dbId, &dbProject_id, &dbDevice_id, &dbSerialnumber, &dbProvisioner, &dbConfiguration, &dbDate, &dbprojectName, &dbdevicetypeId, &dbdeviceVersion, &dbdeviceType)

		var deviceInstance = classes.NewSms_DeviceInstanceFromDB(dbId, dbProject_id, dbDevice_id, dbSerialnumber, dbProvisioner, dbConfiguration, dbDate.String(), dbprojectName, dbdeviceType, dbdeviceVersion)
		deviceInstances=append(deviceInstances, *deviceInstance)
		if err != nil {
			log.Fatal(err)
		}
	}

	return deviceInstances
}

/////////////////////////////////////////
////	SMS System
////////////////////////////////////////
func (mgr *manager) AddSMSSystem(systemtypeId int, version string, date string) (err error) {
	dt := time.Now()

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newSystem)
	if err != nil{
		fmt.Print(err)
	}

	rows, err := stmt.Query(systemtypeId, version, dt)

	if rows == nil{
		fmt.Println("rows should be null")
	}

	return err
}

func (mgr *manager) GetSMSSystems() (systems []classes.Sms_System) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_systems)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query()

	var ( 	dbId int
		dbSystemType string
		dbVersion string
		dbDate time.Time)

	for rows.Next() {
		err := rows.Scan(&dbId, &dbVersion, &dbDate, &dbSystemType)

		var system = classes.NewSms_SystemFromDB(dbId, dbSystemType, dbVersion, dbDate.String())
		systems=append(systems, *system)
		if err != nil {
			log.Fatal(err)
		}
	}

	return systems
}

func (mgr *manager) GetSMSSystemTypes() (systemTypes []classes.Sms_SystemType) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_systemTypes)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query()

	var ( 	dbId int
		dbSystemType string)

	for rows.Next() {
		err := rows.Scan(&dbId, &dbSystemType)

		//Misused project class to collect project types
		var systemType = classes.NewSms_SystemType(dbId, dbSystemType)
		systemTypes=append(systemTypes, *systemType)
		if err != nil {
			log.Fatal(err)
		}
	}

	return systemTypes
}

func (mgr *manager) GetSMSSystemInfo(id int) (*classes.Sms_System) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_systemInfo)
	if err != nil{
		fmt.Print(err)
	}

	var ( 	dbId int
		dbSystemType string
		dbVersion string
		dbDate time.Time)

	row := stmt.QueryRow(id)
	row.Scan(&dbId, &dbVersion, &dbDate, &dbSystemType)

	var system = classes.NewSms_SystemFromDB(dbId, dbSystemType, dbVersion, dbDate.String())

	return system
}

func (mgr *manager) RemoveSMSSystem(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_sms_system)

	stmt.QueryRow(id)

	return err
}


/////////////////////////////////////////
////	SMS Device
////////////////////////////////////////
func (mgr *manager) AddSMSDevice(devicetypeId int, version string, date string) (err error) {
	dt := time.Now()

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newDevice)
	if err != nil{
		fmt.Print(err)
	}

	rows, err := stmt.Query(devicetypeId, version, dt)

	if rows == nil{
		fmt.Println("rows should be null")
	}

	return err
}

func (mgr *manager) GetSMSDevice() (devices []classes.Sms_Device) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_devices)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query()

	var ( 	dbId int
		dbDeviceType string
		dbVersion string
		dbDate time.Time)

	for rows.Next() {
		err := rows.Scan(&dbId, &dbVersion, &dbDate, &dbDeviceType)

		var device = classes.NewSms_DeviceFromDB(dbId, dbDeviceType, dbVersion, dbDate.String())
		devices=append(devices, *device)
		if err != nil {
			log.Fatal(err)
		}
	}

	return devices
}

func (mgr *manager) GetSMSDeviceTypes() (deviceTypes []classes.Sms_DeviceType) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_deviceTypes)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query()

	var ( 	dbDevicetype_id int
		dbDeviceType string)

	for rows.Next() {
		err := rows.Scan(&dbDevicetype_id, &dbDeviceType)

		var deviceType = classes.NewSms_DeviceType(dbDevicetype_id, dbDeviceType)
		deviceTypes=append(deviceTypes, *deviceType)
		if err != nil {
			log.Fatal(err)
		}
	}

	return deviceTypes
}

func (mgr *manager) GetSMSDeviceInfo(id int) (*classes.Sms_Device) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_deviceInfo)
	if err != nil{
		fmt.Print(err)
	}

	var ( 	dbId int
		dbDeviceType string
		dbVersion string
		dbDate time.Time)

	row := stmt.QueryRow(id)
	row.Scan(&dbId, &dbVersion, &dbDate, &dbDeviceType)

	var device = classes.NewSms_DeviceFromDB(dbId, dbDeviceType, dbVersion, dbDate.String())

	return device
}

func (mgr *manager) RemoveSMSDevice(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_sms_device)

	stmt.QueryRow(id)

	return err
}

/////////////////////////////////////////
////	SMS DeviceInstance
////////////////////////////////////////
func (mgr *manager) AddSMSDeviceInstance(project_id int, device_id int, serialnumber string, provisioner string, configuration string) (err error) {
	dt := time.Now()

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newDeviceInstance)
	if err != nil{
		fmt.Print(err)
	}

	rows, err := stmt.Query(project_id, device_id, serialnumber, provisioner, configuration, dt)

	if rows == nil{
		fmt.Println("rows should be null")
	}

	return err
}

func (mgr *manager) GetSMSDeviceInstances() (deviceInstances []classes.Sms_DeviceInstance) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_deviceInstances)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query()

	var ( 	dbId int
		dbProject_id int
		dbDevice_id int
		dbSerialnumber string
		dbProvisioner string
		dbConfiguration string
		dbDate time.Time
		dbprojectName string
		dbdevicetypeId int
		dbdeviceVersion string
		dbdeviceType string
	)

	for rows.Next() {
		err := rows.Scan(&dbId, &dbProject_id, &dbDevice_id, &dbSerialnumber, &dbProvisioner, &dbConfiguration, &dbDate, &dbprojectName, &dbdevicetypeId, &dbdeviceVersion, &dbdeviceType)

		var deviceInstance = classes.NewSms_DeviceInstanceFromDB(dbId, dbProject_id, dbDevice_id, dbSerialnumber, dbProvisioner, dbConfiguration, dbDate.String(), dbprojectName, dbdeviceType, dbdeviceVersion)
		deviceInstances=append(deviceInstances, *deviceInstance)
		if err != nil {
			log.Fatal(err)
		}
	}

	return deviceInstances
}

func (mgr *manager) GetSMSDeviceInstanceInfo(id int) (*classes.Sms_DeviceInstance) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_deviceInstanceInfo)
	if err != nil{
		fmt.Print(err)
	}

	var ( 	dbId int
		dbProject_id int
		dbDevice_id int
		dbSerialnumber string
		dbProvisioner string
		dbConfiguration string
		dbDate time.Time
		dbprojectName string
		dbdevicetypeId int
		dbdeviceVersion string
		dbdeviceType string)

	row := stmt.QueryRow(id)
	row.Scan(&dbId, &dbProject_id, &dbDevice_id, &dbSerialnumber, &dbProvisioner, &dbConfiguration, &dbDate, &dbprojectName, &dbdevicetypeId, &dbdeviceVersion, &dbdeviceType)

	var deviceInstance = classes.NewSms_DeviceInstanceFromDB(dbId, dbProject_id, dbDevice_id, dbSerialnumber, dbProvisioner, dbConfiguration, dbDate.String(), dbprojectName, dbdeviceType, dbdeviceVersion)

	return deviceInstance
}

func (mgr *manager) RemoveSMSDeviceInstances(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_sms_deviceInstance)

	stmt.QueryRow(id)

	return err
}


/////////////////////////////////////////
////	SMS UpdateHistory
////////////////////////////////////////
func (mgr *manager) GetSMSUpdateHistoryForDevice(id int) (updateHistories []classes.Sms_UpdateHistory) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_updatehistoriesForDevice)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(id)

	var ( 	dbUpdateHistory_id int
		dbDeviceInstance_id int
		dbUser string
		dbUpdateType string
		dbDate time.Time
		dbDescription string
	)

	for rows.Next() {
		err := rows.Scan(&dbUpdateHistory_id,&dbDeviceInstance_id,&dbUser,&dbUpdateType,&dbDate,&dbDescription)

		var updateHistory = classes.NewSms_UpdateHistoryFromDB(dbUpdateHistory_id, dbDeviceInstance_id, "", dbUser, dbUpdateType, dbDate.String(), dbDescription)
		updateHistories=append(updateHistories, *updateHistory)
		if err != nil {
			log.Fatal(err)
		}
	}

	return updateHistories
}

func (mgr *manager) AddSMSUpdateHistory(deviceInstance_id int, user string, updateType string, description string) (err error) {
	dt := time.Now()

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newUpdateHistory)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(deviceInstance_id, user, updateType, dt, description)

	if rows == nil{
		fmt.Println("rows should be null")
	}

	return err
}

func (mgr *manager) GetSMSUdateHistoryInfo(id int) (*classes.Sms_UpdateHistory) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_UpdateHistoryInfo)
	if err != nil{
		fmt.Print(err)
	}

	var ( 	dbId int
		dbDeviceInstance_id int
		dbUser string
		dbUpdateType string
		dbDate time.Time
		dbDescription string
	)

	row := stmt.QueryRow(id)
	row.Scan(&dbId,&dbDeviceInstance_id,&dbUser,&dbUpdateType,&dbDate,&dbDescription)

	var updateHistory = classes.NewSms_UpdateHistoryFromDB(dbId, dbDeviceInstance_id, "", dbUser, dbUpdateType, dbDate.String(), dbDescription)

	return updateHistory
}

/////////////////////////////////////////
////	SMS Issue
////////////////////////////////////////
func (mgr *manager) AddSMSIssue(name string, issueType string, reference string, criticality int, cve string, description string) (err error) {
	dt := time.Now()

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newIssue)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(name, dt, issueType, reference, criticality, cve, description)

	if rows == nil{
		fmt.Println("rows should be null")
	}

	return err
}

func (mgr *manager) GetSMSIssues() (issues []classes.Sms_Issue) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_issues)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query()

	var ( 	dbId int
		dbName string
		dbDate time.Time
		dbIssueType string
		dbReference string
		dbCriticality int
		dbCve string
		dbDescription string
	)

	for rows.Next() {
		err := rows.Scan(&dbId, &dbName, &dbDate, &dbIssueType, &dbReference, &dbCriticality, &dbCve, &dbDescription)
		var issue = classes.NewSms_IssueFromDB(dbId, dbName, dbDate.String(), dbIssueType, dbReference, dbCriticality, dbCve, dbDescription)
		issues=append(issues, *issue)
		if err != nil {
			log.Fatal(err)
		}
	}

	return issues
}

func (mgr *manager) GetSMSIssueInfo(id int) (*classes.Sms_Issue) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_issueInfo)
	if err != nil{
		fmt.Print(err)
	}

	var ( 	dbId int
		dbName string
		dbDate time.Time
		dbIssueType string
		dbReference string
		dbCriticality int
		dbCve string
		dbDescription string
	)

	row := stmt.QueryRow(id)
	row.Scan(&dbId, &dbName, &dbDate, &dbIssueType, &dbReference, &dbCriticality, &dbCve, &dbDescription)

	var issue = classes.NewSms_IssueFromDB(dbId, dbName, dbDate.String(), dbIssueType, dbReference, dbCriticality, dbCve, dbDescription)

	return issue
}

func (mgr *manager) RemoveSMSIssue(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_sms_issue)

	stmt.QueryRow(id)

	return err
}

/////////////////////////////////////////
////	SMS IssueAffectedDevice
////////////////////////////////////////
func (mgr *manager) AddSMSIssueAffectedDevice(device_id int, issue_id int, additionalInfo string, confirmed bool) (err error) {

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newIssueAffectedDevice)
	if err != nil{
		fmt.Print(err)
	}

	rows, err := stmt.Query(device_id, issue_id, additionalInfo, confirmed)

	if rows == nil{
		fmt.Println("rows should be null")
	}

	return err
}

func (mgr *manager) GetSMSIssueAffectedDevicesForIssueID(issue_id int) (issueAffectedDevices []classes.Sms_IssueAffectedDevice) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_IssueAffectedDevicesForIssueID)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(issue_id)

	var ( 	dbDevice_id int
			dbIssue_id int
			dbAdditionalInfo string
			dbConfirmed bool
			dbTmp string
			dbTmp2 string
	)

	for rows.Next() {
		err := rows.Scan(&dbDevice_id, &dbIssue_id, &dbAdditionalInfo, &dbConfirmed, &dbTmp, &dbTmp2)

		var affectedDevice = classes.NewSms_IssueAffectedDevice(dbDevice_id, dbIssue_id, dbAdditionalInfo, dbConfirmed, dbTmp, dbTmp2)
		issueAffectedDevices=append(issueAffectedDevices, *affectedDevice)
		if err != nil {
			log.Fatal(err)
		}
	}

	return issueAffectedDevices
}

func (mgr *manager) GetSMSAffectedDeviceInstancesAndProjects(issue_id int) (affectedDevicInstancessAndProjects []classes.Sms_AffectedDeviceInstancesAndProjects) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_affectedDeviceInstancesAndProjects)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(issue_id, issue_id)

	var ( 	dbDeviceInstance_id int
		dbDevicetype string
		dbProject_id int
		dbVersion string
	)

	for rows.Next() {
		err := rows.Scan(&dbDeviceInstance_id, &dbDevicetype, &dbProject_id, &dbVersion)

		var affectedDeviceInstancesAndProject = classes.NewSms_AffectedDeviceInstancesAndProjects(dbDeviceInstance_id, dbDevicetype, dbProject_id, dbVersion)
		affectedDevicInstancessAndProjects=append(affectedDevicInstancessAndProjects, *affectedDeviceInstancesAndProject)
		if err != nil {
			log.Fatal(err)
		}
	}

	return affectedDevicInstancessAndProjects
}

func (mgr *manager) GetSMSIssuesForDevice(device_id int) (issueAffectedDevices []classes.Sms_IssueAffectedDevice) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_IssuesForDevice)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(device_id)

	var ( 	dbDevice_id int
		dbIssue_id int
		dbAdditionalInfo string
		dbConfirmed bool
		dbTmp string
		dbTmp2 string
	)

	for rows.Next() {
		err := rows.Scan(&dbDevice_id, &dbIssue_id, &dbAdditionalInfo, &dbConfirmed, &dbTmp)
		dbTmp2 = ""
		var issueList = classes.NewSms_IssueAffectedDevice(dbDevice_id, dbIssue_id, dbAdditionalInfo, dbConfirmed, dbTmp, dbTmp2)
		issueAffectedDevices=append(issueAffectedDevices, *issueList)
		if err != nil {
			log.Fatal(err)
		}
	}

	return issueAffectedDevices
}


func (mgr *manager) RemoveSMSIssueAffectedDevice(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_sms_IssueAffectedDevice)

	stmt.QueryRow(id)

	return err
}

/////////////////////////////////////////
////	SMS Solution
////////////////////////////////////////
func (mgr *manager) AddSMSSolution(issue_id int, devicetype_id int, name string, description string, reference string) (err error) {

	dt := time.Now()
	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newSolution)
	if err != nil{
		fmt.Print(err)
	}

	rows, err := stmt.Query(issue_id, devicetype_id, dt, name, description, reference)

	if rows == nil{
		fmt.Println("rows should be null")
	}

	return err
}

func (mgr *manager) RemoveSMSSolution(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_sms_Solution)

	stmt.QueryRow(id)

	return err
}

func (mgr *manager) GetSMSSolutionsForIssue(issue_id int) (solutions []classes.Sms_Solution) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_solutionsForIssue)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(issue_id)

	var ( 	dbSolution_id int
		dbIssue_id int
		dbDevicetype_id int
		dbDate time.Time
		dbName string
		dbDescription string
		dbReference string
		dbDeviceType string
	)

	for rows.Next() {
		err := rows.Scan(&dbSolution_id, &dbIssue_id, &dbDevicetype_id, &dbDate, &dbName, &dbDescription, &dbReference, &dbDeviceType)

		var solution = classes.NewSms_SolutionFromDB(dbSolution_id, dbIssue_id, dbDevicetype_id, dbDate.String(), dbName, dbDescription, dbReference, dbDeviceType)
		solutions=append(solutions, *solution)
		if err != nil {
			log.Fatal(err)
		}
	}

	return solutions
}

func (mgr *manager) GetSMSSolutionInfo(id int) (*classes.Sms_Solution) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_solutionInfo)
	if err != nil{
		fmt.Print(err)
	}

	var ( 	dbSolution_id int
		dbIssue_id int
		dbDevicetype_id int
		dbDate time.Time
		dbName string
		dbDescription string
		dbReference string
		dbDeviceType string
	)

	row := stmt.QueryRow(id)
	row.Scan(&dbSolution_id, &dbIssue_id, &dbDevicetype_id, &dbDate, &dbName, &dbDescription, &dbReference, &dbDeviceType)

	var solution = classes.NewSms_SolutionFromDB(dbSolution_id, dbIssue_id, dbDevicetype_id, dbDate.String(), dbName, dbDescription, dbReference, dbDeviceType)

	return solution
}

/////////////////////////////////////////
////	SMS Artefact
////////////////////////////////////////
func (mgr *manager) AddSMSArtefact(artefactype_id int, name string, version string) (err error) {

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newArtefact)
	if err != nil{
		fmt.Print(err)
	}

	rows, err := stmt.Query(artefactype_id, name, version)

	if rows == nil{
		fmt.Println("rows should be null")
	}

	return err
}

func (mgr *manager) GetSMSArtefact() (artefacts []classes.Sms_Artefact) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_artefact)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query()

	var ( 	dbArtefact_id int
		dbArtefactype_id int
		dbName string
		dbVersion string
		dbArtefactype_join string)

	for rows.Next() {
		err := rows.Scan(&dbArtefact_id, &dbArtefactype_id, &dbName, &dbVersion, &dbArtefactype_join)

		var artefact = classes.NewSms_ArtefactFromDB(dbArtefact_id, dbArtefactype_id, dbName, dbVersion, dbArtefactype_join)
		artefacts=append(artefacts, *artefact)
		if err != nil {
			log.Fatal(err)
		}
	}

	return artefacts
}

func (mgr *manager) GetSMSArtefactTypes() (artefactTypes []classes.Sms_ArtefactType) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_artefactTypes)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query()

	var ( 	dbArtefacttype_id int
		dbArtefactType string)

	for rows.Next() {
		err := rows.Scan(&dbArtefacttype_id, &dbArtefactType)

		var artefactType = classes.NewSms_ArtefactTypeFromDB(dbArtefacttype_id, dbArtefactType)
		artefactTypes=append(artefactTypes, *artefactType)
		if err != nil {
			log.Fatal(err)
		}
	}

	return artefactTypes
}

func (mgr *manager) GetSMSArtefactInfo(id int) (*classes.Sms_Artefact) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_artefactInfo)
	if err != nil{
		fmt.Print(err)
	}

	var ( 	dbArtefact_id int
		dbArtefactype_id int
		dbName string
		dbVersion string
		dbArtefactype_join string)

	row := stmt.QueryRow(id)
	row.Scan(&dbArtefact_id, &dbArtefactype_id, &dbName, &dbVersion, &dbArtefactype_join)

	var artefact = classes.NewSms_ArtefactFromDB(dbArtefact_id, dbArtefactype_id, dbName, dbVersion,dbArtefactype_join)

	return artefact
}

func (mgr *manager) RemoveSMSArtefact(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_sms_artefact)

	stmt.QueryRow(id)

	return err
}

/////////////////////////////////////////
////	SMS ReleaseNotes
////////////////////////////////////////
func (mgr *manager) GetSMSReleaseNoteForDevice(id int) (releaseNotes []classes.Sms_ReleaseNote) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_releaseNoteForDevice)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(id)

	var ( 	dbReleasenote_id int
		dbDevice_id int
		dbReleaseNoteType string
		dbDate time.Time
		dbDetails string
	)

	for rows.Next() {
		err := rows.Scan(&dbReleasenote_id,&dbDevice_id,&dbReleaseNoteType,&dbDate,&dbDetails)

		var releaseNote = classes.NewSms_ReleaseNoteFromDB(dbReleasenote_id, dbDevice_id, dbReleaseNoteType, dbDate.String(), dbDetails)
		releaseNotes=append(releaseNotes, *releaseNote)
		if err != nil {
			log.Fatal(err)
		}
	}

	return releaseNotes
}

func (mgr *manager) AddSMSReleaseNote(device_id int, releaseNoteType string, details string) (err error) {
	dt := time.Now()

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newReleaseNote)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(device_id, releaseNoteType, dt, details)

	if rows == nil{
		fmt.Println("rows should be null, releaseNoteInsert")
	}

	return err
}

func (mgr *manager) GetSMSReleaseNoteInfo(id int) (*classes.Sms_ReleaseNote) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_ReleaseNoteInfo)
	if err != nil{
		fmt.Print(err)
	}

	var ( 	dbReleasenote_id int
		dbDevice_id int
		dbReleaseNoteType string
		dbDate time.Time
		dbDetails string
	)

	row := stmt.QueryRow(id)
	row.Scan(&dbReleasenote_id,&dbDevice_id,&dbReleaseNoteType,&dbDate,&dbDetails)

	var releaseNote = classes.NewSms_ReleaseNoteFromDB(dbReleasenote_id, dbDevice_id, dbReleaseNoteType, dbDate.String(), dbDetails)

	return releaseNote
}

/////////////////////////////////////////
////	SMS Software
////////////////////////////////////////
func (mgr *manager) AddSMSSoftware(softwaretype_id int, version string, license string, thirdParty bool, releaseNote string) (err error) {
	dt := time.Now()

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newSoftware)
	if err != nil{
		fmt.Print(err)
	}

	rows, err := stmt.Query(softwaretype_id, version, dt, license, thirdParty, releaseNote)

	if rows == nil{
		fmt.Println("rows should be null, Add Software")
	}

	return err
}

func (mgr *manager) GetSMSSoftware() (softwares []classes.Sms_Software) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_softwares)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query()

	var ( 	dbSoftware_id int
		dbSoftwaretype_id int
		dbVersion string
		dbDate time.Time
		dbTypeName string
		dbLicense string
		dbThirdParty bool
		dbReleaseNote string)

	for rows.Next() {
		err := rows.Scan(&dbSoftware_id, &dbSoftwaretype_id, &dbVersion, &dbDate, &dbTypeName, &dbLicense, &dbThirdParty, &dbReleaseNote)

		var software = classes.NewSms_SoftwareFromDB(dbSoftware_id, dbSoftwaretype_id, dbVersion, dbDate.String(), dbLicense, dbThirdParty, dbReleaseNote, dbTypeName)
		softwares=append(softwares, *software)
		if err != nil {
			log.Fatal(err)
		}
	}

	return softwares
}

func (mgr *manager) GetSMSSoftwareTypes() (softwareTypes []classes.Sms_SoftwareType) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_softwareTypes)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query()

	var ( 	dbSoftwaretype_id int
		dbTypeName string)

	for rows.Next() {
		err := rows.Scan(&dbSoftwaretype_id, &dbTypeName)

		var softwareType = classes.NewSms_SoftwareTypeFromDB(dbSoftwaretype_id, dbTypeName)
		softwareTypes=append(softwareTypes, *softwareType)
		if err != nil {
			log.Fatal(err)
		}
	}

	return softwareTypes
}

func (mgr *manager) GetSMSSoftwareInfo(id int) (*classes.Sms_Software) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_softwareInfo)
	if err != nil{
		fmt.Print(err)
	}

	var ( 	dbSoftware_id int
		dbSoftwaretype_id int
		dbVersion string
		dbDate time.Time
		dbTypeName string
		dbLicense string
		dbThirdParty bool
		dbReleaseNote string)

	row := stmt.QueryRow(id)
	row.Scan(&dbSoftware_id, &dbSoftwaretype_id, &dbVersion, &dbDate, &dbTypeName, &dbLicense, &dbThirdParty, &dbReleaseNote)

	var software = classes.NewSms_SoftwareFromDB(dbSoftware_id, dbSoftwaretype_id, dbVersion, dbDate.String(), dbLicense, dbThirdParty, dbReleaseNote, dbTypeName)

	return software
}

func (mgr *manager) RemoveSMSSoftware(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_sms_software)

	stmt.QueryRow(id)

	return err
}

/////////////////////////////////////////
////	SMS Component
////////////////////////////////////////
func (mgr *manager) AddSMSComponent(name string, componentType string, version string, license string, thirdParty bool, releaseNote string) (err error) {

	dt := time.Now()
	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newComponent)
	if err != nil{
		fmt.Print(err)
	}

	rows, err := stmt.Query(name, componentType, version, dt, license, thirdParty, releaseNote)

	if rows == nil{
		fmt.Println("rows should be null, Add Component")
	}

	return err
}

func (mgr *manager) GetSMSComponent() (components []classes.Sms_Component) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_components)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query()

	var ( 	dbComponent_id int
		dbName string
		dbComponentType string
		dbVersion string
		dbDate time.Time
		dbLicense string
		dbThirdParty bool
		dbReleaseNote string)

	for rows.Next() {
		err := rows.Scan(&dbComponent_id, &dbName, &dbComponentType, &dbVersion, &dbDate, &dbLicense, &dbThirdParty, &dbReleaseNote)

		var component = classes.NewSms_ComponentFromDB(dbComponent_id, dbName, dbComponentType, dbVersion, dbDate.String(), dbLicense, dbThirdParty, dbReleaseNote)
		components=append(components, *component)
		if err != nil {
			log.Fatal(err)
		}
	}

	return components
}

func (mgr *manager) GetSMSComponentInfo(id int) (*classes.Sms_Component) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_componentInfo)
	if err != nil{
		fmt.Print(err)
	}

	var ( 	dbComponent_id int
		dbName string
		dbComponentType string
		dbVersion string
		dbDate time.Time
		dbLicense string
		dbThirdParty bool
		dbReleaseNote string)

	row := stmt.QueryRow(id)
	row.Scan(&dbComponent_id, &dbName, &dbComponentType, &dbVersion, &dbDate, &dbLicense, &dbThirdParty, &dbReleaseNote)

	var component = classes.NewSms_ComponentFromDB(dbComponent_id, dbName, dbComponentType, dbVersion, dbDate.String(), dbLicense, dbThirdParty, dbReleaseNote)

	return component
}

func (mgr *manager) RemoveSMSComponent(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_sms_component)

	stmt.QueryRow(id)

	return err
}

/////////////////////////////////////////
////	SMS ComponentPartOfSoftware
////////////////////////////////////////
func (mgr *manager) AddSMSComponentPartOfSoftware(software_id int, component_id int, additionalInfo string) (err error) {

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newComponentPartOfSoftware)
	if err != nil{
		fmt.Print(err)
	}

	rows, err := stmt.Query(software_id, component_id, additionalInfo)

	if rows == nil{
		fmt.Println("rows should be null AddSMSComponentPartOfSoftware")
	}

	return err
}

func (mgr *manager) GetSMSComponentPartOfSoftwareForSoftware(software_id int) (componentsPartOfSoftware []classes.Sms_ComponentPartOfSoftware) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_ComponentPartOfSoftwareForSoftware)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(software_id)

	var ( 	dbSoftware_id int
		dbComponent_id int
		dbAdditionalInfo string
		dbName string
		dbVersion string
	)

	for rows.Next() {
		err := rows.Scan(&dbSoftware_id, &dbComponent_id, &dbAdditionalInfo, &dbName, &dbVersion)

		var component = classes.NewSms_ComponentPartOfSoftwareFromDB(dbSoftware_id, dbComponent_id, dbAdditionalInfo, dbName, dbVersion)
		componentsPartOfSoftware=append(componentsPartOfSoftware, *component)
		if err != nil {
			log.Fatal(err)
		}
	}

	return componentsPartOfSoftware
}


func (mgr *manager) GetSMSComponentPartOfSoftwareForComponent(component_id int) (softwaresParentOfComponent []classes.Sms_ComponentPartOfSoftware) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_ComponentPartOfSoftwareForComponent)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(component_id)

	var ( 	dbSoftware_id int
		dbComponent_id int
		dbAdditionalInfo string
		dbName string
		dbVersion string
	)

	for rows.Next() {
		err := rows.Scan(&dbSoftware_id, &dbComponent_id, &dbAdditionalInfo, &dbName, &dbVersion)
		var software = classes.NewSms_ComponentPartOfSoftwareFromDB(dbSoftware_id, dbComponent_id, dbAdditionalInfo, dbName, dbVersion)
		softwaresParentOfComponent=append(softwaresParentOfComponent, *software)
		if err != nil {
			log.Fatal(err)
		}
	}

	return softwaresParentOfComponent
}


func (mgr *manager) RemoveSMSComponentPartOfSoftware(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_sms_ComponentPartOfSoftware)

	stmt.QueryRow(id)

	return err
}


/////////////////////////////////////////
////	SMS SoftwarePartOfDevice
////////////////////////////////////////
func (mgr *manager) AddSMSSoftwarePartOfDevice(device_id int, software_id int, additionalInfo string) (err error) {

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newSoftwarePartOfDevice)
	if err != nil{
		fmt.Print(err)
	}

	rows, err := stmt.Query(device_id, software_id, additionalInfo)

	if rows == nil{
		fmt.Println("rows should be null AddSMSSoftwarePartOfDevice -> insert query")
	}

	return err
}

func (mgr *manager) GetSMSSoftwarePartOfDeviceForDevice(device_id int) (sofwaresPartOfDevice []classes.Sms_SoftwarePartOfDevice) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_SoftwarePartOfDeviceForDevice)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(device_id)

	var ( 	dbDevice_id int
		dbSoftware_id int
		dbAdditionalInfo string
		dbName string
		dbVersion string
	)

	for rows.Next() {
		err := rows.Scan(&dbDevice_id, &dbSoftware_id, &dbAdditionalInfo, &dbName, &dbVersion)

		var software = classes.NewSms_SoftwarePartOfDevice(dbDevice_id, dbSoftware_id, dbAdditionalInfo, dbName, dbVersion)
		sofwaresPartOfDevice=append(sofwaresPartOfDevice, *software)
		if err != nil {
			log.Fatal(err)
		}
	}

	return sofwaresPartOfDevice
}


func (mgr *manager) GetSMSSoftwarePartOfDeviceForSoftware(software_id int) (devicesParentOfSoftware []classes.Sms_SoftwarePartOfDevice) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_SoftwarePartOfDeviceForSoftware)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(software_id)

	var ( 	dbDevice_id int
		dbSoftware_id int
		dbAdditionalInfo string
		dbName string
		dbVersion string
	)

	for rows.Next() {
		err := rows.Scan(&dbDevice_id, &dbSoftware_id, &dbAdditionalInfo, &dbName, &dbVersion)
		var device = classes.NewSms_SoftwarePartOfDevice(dbDevice_id, dbSoftware_id, dbAdditionalInfo, dbName, dbVersion)
		devicesParentOfSoftware=append(devicesParentOfSoftware, *device)
		if err != nil {
			log.Fatal(err)
		}
	}

	return devicesParentOfSoftware
}


func (mgr *manager) RemoveSMSSoftwarePartOfDevice(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_sms_SoftwarePartOfDevice)

	stmt.QueryRow(id)

	return err
}