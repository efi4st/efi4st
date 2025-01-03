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
	RemoveSMSIssueAffectedDevice(device_id int, issue_id int) error
	GetSMSAffectedDeviceInstancesAndProjects(issue_id int) []classes.Sms_AffectedDeviceInstancesAndProjects
	GetIssueAffectedStats(issue_id int) (*classes.Sms_IssueAffectedStats, error)
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
	AddSMSDevicePartOfSystem(system_id int, device_id int, additionalInfo string) error
	GetSMSDevicePartOfSystemForSystem(system_id int) []classes.Sms_DevicePartOfSystem
	GetSMSDevicePartOfSystemForDevice(device_id int) []classes.Sms_DevicePartOfSystem
	RemoveSMSDevicePartOfSystem(id int) error
	AddSMSProjectBOM(project_id int, system_id int, orderNumber string, additionalInfo string) error
	GetSMSProjectBOMForProject(project_id int) []classes.Sms_ProjectBOM
	GetSMSProjectBOMForSystem(system_id int) []classes.Sms_ProjectBOM
	RemoveSMSProjectBOM(id int) error
	AddSMSIssueAffectedSoftware(software_id int, issue_id int, additionalInfo string, confirmed bool) error
	GetSMSIssueAffectedSoftwareForIssueID(issue_id int) []classes.Sms_IssueAffectedSoftware
	GetSMSIssuesForSoftware(software_id int) []classes.Sms_IssueAffectedSoftware
	RemoveSMSIssueAffectedSoftware(software_id int, issue_id int) (err error)
	AddSMSArtefactPartOfDevice(device_id int, artefact_id int, additionalInfo string) error
	GetSMSArtefactPartOfDeviceForDevice(device_id int) []classes.Sms_ArtefactPartOfDevice
	GetSMSArtefactPartOfDeviceForArtefact(artefact_id int) []classes.Sms_ArtefactPartOfDevice
	RemoveSMSArtefactPartOfDevice(id int) error
	GetSMSManufactoringOrderForSystem(id int) []classes.Sms_ManufacturingOrder
	AddSMSManufacturingOrder(system_id int, packageReference string, description string) error
	GetSMSManufacturingOrderInfo(id int) *classes.Sms_ManufacturingOrder
	GetSMSSystemTreeForSystem(id int) *classes.Sms_Tree_System
	// Certification
	AddSMSCertification(name string, description string) error
	GetSMSCertification() []classes.Sms_Certification
	GetSMSCertificationInfo(id int) *classes.Sms_Certification
	RemoveSMSCertification(id int) error
	ComponentExists(name string, componentType string, version string) (bool, int, error)
	ProcessComponents(components []classes.Sms_Component, softwareID int) error
	// SystemHasCertification
	AddSystemHasCertification(system_id int, certification_id int, additionalInfo string) error
	GetCertificationsForSystem(systemID int) (certifications []classes.Sms_SystemHasCertification, err error)
	GetSystemsForCertification(certificationID int) (systems []classes.Sms_SystemHasCertification, err error)
	RemoveSystemHasCertification(systemID int, certificationID int) error
	//IssueAffectedComponent
	AddSMSIssueAffectedComponent(component_id int, issue_id int, additionalInfo string, confirmed bool) error
	GetSMSIssueAffectedComponentsForIssueID(issue_id int) (issueAffectedComponents []classes.Sms_IssueAffectedComponent, err error)
	GetSMSIssuesForComponent(component_id int) (issueAffectedComponents []classes.Sms_IssueAffectedComponent, err error)
	RemoveSMSIssueAffectedComponent(component_id int, issue_id int) error
	//IssueAffectedArtefact
	AddSMSIssueAffectedArtefact(artefact_id int, issue_id int, additionalInfo string, confirmed bool) error
	GetSMSIssueAffectedArtefactsForIssueID(issue_id int) (issueAffectedArtefacts []classes.Sms_IssueAffectedArtefact, err error)
	GetSMSIssuesForArtefact(artefact_id int) (issueAffectedArtefacts []classes.Sms_IssueAffectedArtefact, err error)
	RemoveSMSIssueAffectedArtefact(artefact_id int, issue_id int) error
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

func (mgr *manager) GetSMSSystemTreeForSystem(id int) (*classes.Sms_Tree_System){
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_DevicePartOfSystemForSystem)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(id)

	var ( 	db1System_id int
		db1Device_id int
		db1AdditionalInfo string
		db1Name string
		db1Version string
	)

	var deviceList []classes.Sms_Tree_Device
	for rows.Next() {
		err := rows.Scan(&db1System_id, &db1Device_id, &db1AdditionalInfo, &db1Name, &db1Version)

		stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_SoftwarePartOfDeviceForDevice)
		if err != nil{
			fmt.Print(err)
		}
		rows, err := stmt.Query(db1Device_id)

		var ( 	db2Device_id int
			db2Software_id int
			db2AdditionalInfo string
			db2Name string
			db2Version string
		)

		var applicationList []classes.Sms_Tree_Application
		for rows.Next() {
			err := rows.Scan(&db2Device_id, &db2Software_id, &db2AdditionalInfo, &db2Name, &db2Version)

			stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_ComponentPartOfSoftwareForSoftware)
			if err != nil{
				fmt.Print(err)
			}
			rows, err := stmt.Query(db2Software_id)

			var ( 	db3Software_id int
				db3Component_id int
				db3AdditionalInfo string
				db3Name string
				db3Version string
			)
			var componentList []classes.Sms_Tree_Component
			for rows.Next() {
				err := rows.Scan(&db3Software_id, &db3Component_id, &db3AdditionalInfo, &db3Name, &db3Version)

				var treeComp = classes.NewSms_Tree_Component(db3Name, db3Version)
				componentList=append(componentList, *treeComp)
				if err != nil {
					log.Fatal(err)
				}
			}
			var treeSoft = classes.NewSms_Tree_Application(db2Name, db2Version, componentList)
			applicationList=append(applicationList, *treeSoft)
			if err != nil {
				log.Fatal(err)
			}
		}
		var treeDev = classes.NewSms_Tree_Device(db1Name, db1Version, applicationList)
		deviceList=append(deviceList, *treeDev)
		if err != nil {
			log.Fatal(err)
		}
	}

	var systemTree = classes.NewSms_Tree_System("System:", strconv.Itoa(id), deviceList)
	return systemTree
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
	// Prepare the statement
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_affectedDeviceInstancesAndProjects)
	if err != nil {
		log.Fatalf("Error preparing statement: %v", err)
	}
	defer stmt.Close() // Ensure the statement is closed

	// Execute the query
	rows, err := stmt.Query(issue_id, issue_id, issue_id, issue_id, issue_id, issue_id, issue_id, issue_id)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	defer rows.Close() // Ensure rows are closed

	// Variables for scanning
	var (
		dbDeviceInstance_id sql.NullInt32
		dbDevicetype        sql.NullString
		dbProject_id        sql.NullInt32
		dbVersion           sql.NullString
	)

	// Iterate over rows
	for rows.Next() {
		err := rows.Scan(&dbDeviceInstance_id, &dbDevicetype, &dbProject_id, &dbVersion)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}

		// Check for NULL values in important fields
		if !dbDeviceInstance_id.Valid || !dbDevicetype.Valid || !dbProject_id.Valid || !dbVersion.Valid {
			log.Printf("Skipping row with NULL values: DeviceInstanceID=%v, DeviceType=%v, ProjectID=%v, Version=%v",
				dbDeviceInstance_id, dbDevicetype, dbProject_id, dbVersion)
			continue
		}

		// Create the object only if all fields are valid
		affectedDeviceInstancesAndProject := classes.NewSms_AffectedDeviceInstancesAndProjects(
			int(dbDeviceInstance_id.Int32),
			dbDevicetype.String,
			int(dbProject_id.Int32),
			dbVersion.String,
		)
		affectedDevicInstancessAndProjects = append(affectedDevicInstancessAndProjects, *affectedDeviceInstancesAndProject)
	}

	// Check for errors after iteration
	if err = rows.Err(); err != nil {
		log.Fatalf("Error during row iteration: %v", err)
	}

	return affectedDevicInstancessAndProjects
}

func (mgr *manager) GetIssueAffectedStats(issue_id int) (*classes.Sms_IssueAffectedStats, error) {
	// Prepare the statement
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_statisticsForaffectedDeviceInstancesAndProjects)
	if err != nil {
		return nil, fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close() // Ensure the statement is closed

	// Execute the query
	row := stmt.QueryRow(issue_id, issue_id, issue_id, issue_id, issue_id, issue_id, issue_id, issue_id)

	// Variables for scanning
	var (
		affectedDeviceInstances         	sql.NullInt32
		affectedDevicesWithoutInstances 	sql.NullInt32
		affectedProjects                	sql.NullInt32
		distinctDeviceVersionCombinations   sql.NullInt32
	)

	// Scan the single row of results
	err = row.Scan(
		&affectedDeviceInstances,
		&affectedDevicesWithoutInstances,
		&affectedProjects,
		&distinctDeviceVersionCombinations,
	)
	if err != nil {
		return nil, fmt.Errorf("error scanning row: %w", err)
	}

	// Create and return the Sms_IssueAffectedStats object
	stats := &classes.Sms_IssueAffectedStats{
		AffectedDeviceInstances:        	0, // default values
		AffectedDevicesWithoutInstances: 	0,
		AffectedProjects:                	0,
		DistinctDeviceVersionCombinations:  0,
	}

	// Populate fields only if they are valid
	if affectedDeviceInstances.Valid {
		stats.AffectedDeviceInstances = int(affectedDeviceInstances.Int32)
	}
	if affectedDevicesWithoutInstances.Valid {
		stats.AffectedDevicesWithoutInstances = int(affectedDevicesWithoutInstances.Int32)
	}
	if affectedProjects.Valid {
		stats.AffectedProjects = int(affectedProjects.Int32)
	}
	if distinctDeviceVersionCombinations.Valid {
		stats.DistinctDeviceVersionCombinations = int(distinctDeviceVersionCombinations.Int32)
	}

	return stats, nil
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


func (mgr *manager) RemoveSMSIssueAffectedDevice(device_id int, issue_id int) (err error) {
	// Bereite die SQL-Anweisung vor
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_sms_IssueAffectedDevice)
	if err != nil {
		return fmt.Errorf("failed to prepare DELETE statement: %w", err)
	}
	defer stmt.Close()

	// Führe die SQL-Abfrage mit den richtigen Parametern aus
	_, err = stmt.Exec(device_id, issue_id)
	if err != nil {
		return fmt.Errorf("failed to execute DELETE statement: %w", err)
	}

	return nil
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

func (mgr *manager) ComponentExists(name string, componentType string, version string) (bool, int, error) {
	var componentID int
	stmt, err := mgr.db.Prepare(dbUtils.Check_sms_component)
	if err != nil{
		fmt.Print(err)
	}
	err = stmt.QueryRow(name, componentType, version).Scan(&componentID)

	if err == sql.ErrNoRows {
		fmt.Print("Component not found in Database!")
		// Komponente nicht gefunden
		return false, 0, nil
	} else if err != nil {
		// Ein anderer Fehler ist aufgetreten
		return false, 0, err
	}

	// Komponente gefunden, Rückgabe der ID
	return true, componentID, nil
}

// übergibt Liste der eingelesenen Subkomponenten
// checkt ob diese schon in der DB enthalten sind
// wenn nein -> eintragen
func (mgr *manager) ProcessComponents(components []classes.Sms_Component,  softwareID int) error {
	for _, comp := range components {

		exists, id, err := GetDBManager().ComponentExists(comp.Name(), comp.ComponentType(), comp.Version())
		if err != nil {
			return fmt.Errorf("error checking component existence: %w", err)
		}

		if exists {
			log.Printf("Component already exists in database with ID %d: %s", id, comp.Name)
			GetDBManager().AddSMSComponentPartOfSoftware(softwareID, id, "inserted by SBOM upload")
			continue
		}

		// Komponente hinzufügen
		dt := time.Now()
		stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newComponent)
		if err != nil{
			fmt.Print(err)
		}

		rows, err := stmt.Query(comp.Name(), comp.ComponentType(), comp.Version(), dt, comp.License(), comp.ThirdParty(), comp.ReleaseNote())

		insertSuccess := true
		if rows == nil{
			fmt.Println("rows should be null, Add Component")
			insertSuccess = false
		}

		if err != nil {
			return fmt.Errorf("error inserting component: %w", err)
			insertSuccess = false
		}

		exists, id, err = GetDBManager().ComponentExists(comp.Name(), comp.ComponentType(), comp.Version())
		if err != nil {
			return fmt.Errorf("error checking component existence: %w", err)
		}

		if insertSuccess == true && exists{
			GetDBManager().AddSMSComponentPartOfSoftware(softwareID, id, "inserted by SBOM upload")
			log.Printf("Inserted new component: %s", comp.Name())
		}
	}
	return nil
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

/////////////////////////////////////////
////	SMS DevicePartOfSystem
////////////////////////////////////////
func (mgr *manager) AddSMSDevicePartOfSystem(system_id int, device_id int, additionalInfo string) (err error) {

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newDevicePartOfSystem)
	if err != nil{
		fmt.Print(err)
	}

	rows, err := stmt.Query(system_id, device_id, additionalInfo)

	if rows == nil{
		fmt.Println("rows should be null AddSMSDevicePartOfSystem -> insert query")
	}

	return err
}

func (mgr *manager) GetSMSDevicePartOfSystemForSystem(system_id int) (devicesPartOfSystem []classes.Sms_DevicePartOfSystem) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_DevicePartOfSystemForSystem)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(system_id)

	var ( 	dbSystem_id int
		dbDevice_id int
		dbAdditionalInfo string
		dbName string
		dbVersion string
	)

	for rows.Next() {
		err := rows.Scan(&dbSystem_id, &dbDevice_id, &dbAdditionalInfo, &dbName, &dbVersion)

		var device = classes.NewSms_DevicePartOfSystem(dbSystem_id, dbDevice_id, dbAdditionalInfo, dbName, dbVersion)
		devicesPartOfSystem=append(devicesPartOfSystem, *device)
		if err != nil {
			log.Fatal(err)
		}
	}

	return devicesPartOfSystem
}


func (mgr *manager) GetSMSDevicePartOfSystemForDevice(device_id int) (systemsParentOfDevice []classes.Sms_DevicePartOfSystem) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_DevicePartOfSystemForDevice)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(device_id)

	var ( 	dbSystem_id int
		dbDevice_id int
		dbAdditionalInfo string
		dbName string
		dbVersion string
	)

	for rows.Next() {
		err := rows.Scan(&dbSystem_id, &dbDevice_id, &dbAdditionalInfo, &dbName, &dbVersion)
		var system = classes.NewSms_DevicePartOfSystem(dbSystem_id, dbDevice_id, dbAdditionalInfo, dbName, dbVersion)
		systemsParentOfDevice=append(systemsParentOfDevice, *system)
		if err != nil {
			log.Fatal(err)
		}
	}

	return systemsParentOfDevice
}


func (mgr *manager) RemoveSMSDevicePartOfSystem(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_sms_DevicePartOfSystem)

	stmt.QueryRow(id)

	return err
}


/////////////////////////////////////////
////	SMS DeviceInstance
////////////////////////////////////////
func (mgr *manager) AddSMSProjectBOM(project_id int, system_id int, orderNumber string, additionalInfo string) (err error) {

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newProjectBOM)
	if err != nil{
		fmt.Print(err)
	}

	rows, err := stmt.Query(project_id, system_id, orderNumber, additionalInfo)

	if rows == nil{
		fmt.Println("rows should be null -> AddSMSprojectBOM")
	}

	return err
}

func (mgr *manager) GetSMSProjectBOMForProject(project_id int) (soldSystemsPartOfProject []classes.Sms_ProjectBOM) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_ProjectBOMForProject)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(project_id)

	var ( 	dbProjectBOM_id int
		dbProject_id int
		dbSystem_id int
		dbOrderNumber string
		dbAdditionalInfo string
		dbName string
		dbTmp string
	)

	for rows.Next() {
		err := rows.Scan(&dbProjectBOM_id, &dbProject_id, &dbSystem_id, &dbOrderNumber, &dbAdditionalInfo, &dbName, &dbTmp)

		var system = classes.NewSms_ProjectBOMFromDB(dbProjectBOM_id, dbProject_id, dbSystem_id, dbOrderNumber, dbAdditionalInfo, dbName, dbTmp)
		soldSystemsPartOfProject=append(soldSystemsPartOfProject, *system)
		if err != nil {
			log.Fatal(err)
		}
	}

	return soldSystemsPartOfProject
}


func (mgr *manager) GetSMSProjectBOMForSystem(system_id int) (projectsUsingSystem []classes.Sms_ProjectBOM) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_ProjectBOMForSystem)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(system_id)

	var ( 	dbProjectBOM_id int
		dbProject_id int
		dbSystem_id int
		dbOrderNumber string
		dbAdditionalInfo string
		dbName string
		dbTmp string
	)

	for rows.Next() {
		err := rows.Scan(&dbProjectBOM_id, &dbProject_id, &dbSystem_id, &dbOrderNumber, &dbAdditionalInfo, &dbName, &dbTmp)
		var project = classes.NewSms_ProjectBOMFromDB(dbProjectBOM_id, dbProject_id, dbSystem_id, dbOrderNumber, dbAdditionalInfo, dbName, dbTmp)
		projectsUsingSystem=append(projectsUsingSystem, *project)
		if err != nil {
			log.Fatal(err)
		}
	}

	return projectsUsingSystem
}

func (mgr *manager) RemoveSMSProjectBOM(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_sms_ProjectBOM)

	stmt.QueryRow(id)

	return err
}

/////////////////////////////////////////
////	SMS IssueAffectedSoftware
////////////////////////////////////////
func (mgr *manager) AddSMSIssueAffectedSoftware(software_id int, issue_id int, additionalInfo string, confirmed bool) (err error) {

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newIssueAffectedSoftware)
	if err != nil{
		fmt.Print(err)
	}

	rows, err := stmt.Query(software_id, issue_id, additionalInfo, confirmed)

	if rows == nil{
		fmt.Println("rows should be null AddSMSIssueAffectedSoftware")
	}

	return err
}

func (mgr *manager) GetSMSIssueAffectedSoftwareForIssueID(issue_id int) (issueAffectedSoftwares []classes.Sms_IssueAffectedSoftware) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_IssueAffectedSoftwaresForIssueID)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(issue_id)

	var ( 	dbSoftware_id int
		dbIssue_id int
		dbAdditionalInfo string
		dbConfirmed bool
		dbTmp string
		dbTmp2 string
	)

	for rows.Next() {
		err := rows.Scan(&dbSoftware_id, &dbIssue_id, &dbAdditionalInfo, &dbConfirmed, &dbTmp, &dbTmp2)

		var affectedSoftware = classes.NewSms_IssueAffectedSoftware(dbSoftware_id, dbIssue_id, dbAdditionalInfo, dbConfirmed, dbTmp, dbTmp2)
		issueAffectedSoftwares=append(issueAffectedSoftwares, *affectedSoftware)
		if err != nil {
			log.Fatal(err)
		}
	}

	return issueAffectedSoftwares
}


func (mgr *manager) GetSMSIssuesForSoftware(software_id int) (issueAffectedSoftwares []classes.Sms_IssueAffectedSoftware) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_IssuesForSoftware)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(software_id)

	var ( 	dbSoftware_id int
		dbIssue_id int
		dbAdditionalInfo string
		dbConfirmed bool
		dbTmp string
		dbTmp2 string
	)

	for rows.Next() {
		err := rows.Scan(&dbSoftware_id, &dbIssue_id, &dbAdditionalInfo, &dbConfirmed, &dbTmp)
		dbTmp2 = ""
		var issueList = classes.NewSms_IssueAffectedSoftware(dbSoftware_id, dbIssue_id, dbAdditionalInfo, dbConfirmed, dbTmp, dbTmp2)
		issueAffectedSoftwares=append(issueAffectedSoftwares, *issueList)
		if err != nil {
			log.Fatal(err)
		}
	}

	return issueAffectedSoftwares
}


func (mgr *manager) RemoveSMSIssueAffectedSoftware(software_id int, issue_id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_sms_IssueAffectedSoftware)
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return err
	}
	defer stmt.Close() // Sicherstellen, dass die Ressource freigegeben wird

	_, err = stmt.Exec(software_id, issue_id)
	if err != nil {
		fmt.Println("Error executing DELETE statement:", err)
	}

	return err
}

/////////////////////////////////////////
////	SMS ArtefactPartOfDevice
////////////////////////////////////////
func (mgr *manager) AddSMSArtefactPartOfDevice(device_id int, artefact_id int, additionalInfo string) (err error) {

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newArtefactPartOfDevice)
	if err != nil{
		fmt.Print(err)
	}

	rows, err := stmt.Query(device_id, artefact_id, additionalInfo)

	if rows == nil{
		fmt.Println("rows should be null AddSMSArtefactPartOfDevice -> insert query")
	}

	return err
}

func (mgr *manager) GetSMSArtefactPartOfDeviceForDevice(device_id int) (artefactsPartOfDevice []classes.Sms_ArtefactPartOfDevice) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_ArtefactPartOfDeviceForDevice)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(device_id)

	var ( 	dbDevice_id int
		dbArtefact_id int
		dbAdditionalInfo string
		dbName string
		dbVersion string
	)

	for rows.Next() {
		err := rows.Scan(&dbDevice_id, &dbArtefact_id, &dbAdditionalInfo, &dbName, &dbVersion)

		var artefact = classes.NewSms_ArtefactPartOfDevice(dbDevice_id, dbArtefact_id, dbAdditionalInfo, dbName, dbVersion)
		artefactsPartOfDevice=append(artefactsPartOfDevice, *artefact)
		if err != nil {
			log.Fatal(err)
		}
	}

	return artefactsPartOfDevice
}


func (mgr *manager) GetSMSArtefactPartOfDeviceForArtefact(artefact_id int) (devicesParentOfArtefact []classes.Sms_ArtefactPartOfDevice) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_ArtefactPartOfDeviceForArtefact)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(artefact_id)

	var ( 	dbDevice_id int
		dbArtefact_id int
		dbAdditionalInfo string
		dbName string
		dbVersion string
	)

	for rows.Next() {
		err := rows.Scan(&dbDevice_id, &dbArtefact_id, &dbAdditionalInfo, &dbName, &dbVersion)
		var device = classes.NewSms_ArtefactPartOfDevice(dbDevice_id, dbArtefact_id, dbAdditionalInfo, dbName, dbVersion)
		devicesParentOfArtefact=append(devicesParentOfArtefact, *device)
		if err != nil {
			log.Fatal(err)
		}
	}

	return devicesParentOfArtefact
}


func (mgr *manager) RemoveSMSArtefactPartOfDevice(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_sms_ArtefactPartOfDevice)

	stmt.QueryRow(id)

	return err
}


/////////////////////////////////////////
////	SMS ManufactoringOrder
////////////////////////////////////////
func (mgr *manager) GetSMSManufactoringOrderForSystem(id int) (manufactoringOrders []classes.Sms_ManufacturingOrder) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_ManufacturingOrdersForSystem)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(id)

	var ( 	dbManufacturingOrder_id int
		dbSystem_id int
		dbPackageReference string
		dbStart time.Time
		dbEnd sql.NullTime
		dbDescription string
	)

	for rows.Next() {
		err := rows.Scan(&dbManufacturingOrder_id,&dbSystem_id,&dbPackageReference,&dbStart,&dbEnd,&dbDescription)

		var manufactoringOrder *classes.Sms_ManufacturingOrder
		if dbEnd.Valid == true{
			manufactoringOrder = classes.NewSms_ManufacturingOrderFromDB(dbManufacturingOrder_id, dbSystem_id, dbPackageReference, dbStart.String(), dbEnd.Time.String(), dbDescription)
		} else{
			manufactoringOrder = classes.NewSms_ManufacturingOrderFromDB(dbManufacturingOrder_id, dbSystem_id, dbPackageReference, dbStart.String(), "", dbDescription)
		}
		manufactoringOrders=append(manufactoringOrders, *manufactoringOrder)
		if err != nil {
			log.Fatal(err)
		}
	}

	return manufactoringOrders
}

func (mgr *manager) AddSMSManufacturingOrder(system_id int, packageReference string, description string) (err error) {
	dt := time.Now()

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newManufacturingOrder)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(system_id, packageReference, dt, description)

	if rows == nil{
		fmt.Println("rows should be null, AddSMSManufacturingOrder")
	}

	return err
}

func (mgr *manager) GetSMSManufacturingOrderInfo(id int) (*classes.Sms_ManufacturingOrder) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_ManufacturingOrderInfo)
	if err != nil{
		fmt.Print(err)
	}

	var ( 	dbManufacturingOrder_id int
		dbSystem_id int
		dbPackageReference string
		dbStart time.Time
		dbEnd sql.NullTime
		dbDescription string
	)

	row := stmt.QueryRow(id)
	row.Scan(&dbManufacturingOrder_id,&dbSystem_id,&dbPackageReference,&dbStart,&dbEnd,&dbDescription)

	var manufactoringOrder *classes.Sms_ManufacturingOrder
	if dbEnd.Valid == true{
		manufactoringOrder = classes.NewSms_ManufacturingOrderFromDB(dbManufacturingOrder_id, dbSystem_id, dbPackageReference, dbStart.String(), dbEnd.Time.String(), dbDescription)
	} else{
		manufactoringOrder = classes.NewSms_ManufacturingOrderFromDB(dbManufacturingOrder_id, dbSystem_id, dbPackageReference, dbStart.String(), "", dbDescription)
	}

	return manufactoringOrder
}


/////////////////////////////////////////
////	SMS Certification
////////////////////////////////////////
func (mgr *manager) AddSMSCertification(name string, description string) (err error) {

	dt := time.Now()
	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newCertification)
	if err != nil{
		fmt.Print(err)
	}

	rows, err := stmt.Query(name, dt, description)

	if rows == nil{
		fmt.Println("rows should be null, Add Certification")
	}

	return err
}

func (mgr *manager) GetSMSCertification() (certifications []classes.Sms_Certification) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_certification)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query()

	var ( 	dbCertification_id int
		dbName string
		dbDate time.Time
		dbDescription string)

	for rows.Next() {
		err := rows.Scan(&dbCertification_id, &dbName, &dbDate, &dbDescription)

		var certification = classes.NewSms_CertificationFromDB(dbCertification_id, dbName, dbDate.String(), dbDescription)
		certifications=append(certifications, *certification)
		if err != nil {
			log.Fatal(err)
		}
	}

	return certifications
}

func (mgr *manager) GetSMSCertificationInfo(id int) (*classes.Sms_Certification) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_certificationInfo)
	if err != nil{
		fmt.Print(err)
	}

	var ( 	dbCertification_id int
		dbName string
		dbDate time.Time
		dbDescription string)

	row := stmt.QueryRow(id)
	row.Scan(&dbCertification_id, &dbName, &dbDate, &dbDescription)

	var certification = classes.NewSms_CertificationFromDB(dbCertification_id, dbName, dbDate.String(), dbDescription)

	return certification
}

func (mgr *manager) RemoveSMSCertification(id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_sms_certification)

	stmt.QueryRow(id)

	return err
}

/////////////////////////////////////////
////	SMS SystemHasCertification
////////////////////////////////////////
func (mgr *manager) AddSystemHasCertification(system_id int, certification_id int, additionalInfo string) (err error) {
	// Vorbereiten des Statements
	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_systemHasCertification)
	if err != nil {
		fmt.Printf("Error preparing statement: %v\n", err)
		return err
	}
	defer stmt.Close() // Sicherstellen, dass das Statement geschlossen wird

	// Ausführen des Statements
	_, err = stmt.Exec(system_id, certification_id, additionalInfo)
	if err != nil {
		fmt.Printf("Error executing statement: %v\n", err)
		return err
	}

	return nil
}

func (mgr *manager) GetCertificationsForSystem(systemID int) (certifications []classes.Sms_SystemHasCertification, err error) {
	// Bereite das SELECT-Statement vor
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_systemHasCertificationForSystem)
	if err != nil {
		fmt.Printf("Error preparing statement: %v\n", err)
		return nil, err
	}
	defer stmt.Close() // Schließe das Statement nach der Nutzung

	// Führe die Abfrage aus
	rows, err := stmt.Query(systemID)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close() // Schließe die Rows nach der Nutzung

	// Variablen zum Speichern der abgerufenen Daten
	var (
		dbSystemID          int
		dbCertificationID   int
		dbAdditionalInfo    string
		dbCertificationName string
	)

	// Iteriere durch die Ergebnisse und baue die Liste der Zertifizierungen
	for rows.Next() {
		err := rows.Scan(&dbSystemID, &dbCertificationID, &dbAdditionalInfo, &dbCertificationName)
		if err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			return nil, err
		}

		// Erstelle eine Instanz der Datenklasse
		certification := classes.NewSms_SystemHasCertification(dbSystemID, dbCertificationID, dbAdditionalInfo, dbCertificationName, "", "")

		// Füge die Zertifizierung zur Liste hinzu
		certifications = append(certifications, *certification)
	}

	// Prüfe auf Fehler beim Durchlaufen der Zeilen
	if err := rows.Err(); err != nil {
		fmt.Printf("Error iterating rows: %v\n", err)
		return nil, err
	}

	return certifications, nil
}

func (mgr *manager) GetSystemsForCertification(certificationID int) (systems []classes.Sms_SystemHasCertification, err error) {
	// Bereite das SELECT-Statement vor
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_systemHasCertificationForCertification)
	if err != nil {
		fmt.Printf("Error preparing statement: %v\n", err)
		return nil, err
	}
	defer stmt.Close() // Schließe das Statement nach der Nutzung

	// Führe die Abfrage aus
	rows, err := stmt.Query(certificationID)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return nil, err
	}
	defer rows.Close() // Schließe die Rows nach der Nutzung

	// Variablen zum Speichern der abgerufenen Daten
	var (
		dbSystemID       int
		dbCertificationID int
		dbAdditionalInfo string
		dbSystemName     string
		dbSystemVersion  string
	)

	// Iteriere durch die Ergebnisse und baue die Liste der Systeme
	for rows.Next() {
		err := rows.Scan(&dbSystemID, &dbCertificationID, &dbAdditionalInfo, &dbSystemName, &dbSystemVersion)
		if err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			return nil, err
		}

		// Erstelle eine Instanz der Datenklasse
		system := classes.NewSms_SystemHasCertification(dbSystemID, dbCertificationID, dbAdditionalInfo, "", dbSystemName, dbSystemVersion)

		// Füge das System zur Liste hinzu
		systems = append(systems, *system)
	}

	// Prüfe auf Fehler beim Durchlaufen der Zeilen
	if err := rows.Err(); err != nil {
		fmt.Printf("Error iterating rows: %v\n", err)
		return nil, err
	}

	return systems, nil
}

func (mgr *manager) RemoveSystemHasCertification(systemID int, certificationID int) (err error) {
	// Bereite das DELETE-Statement vor
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_sms_systemHasCertification)
	if err != nil {
		fmt.Printf("Error preparing statement: %v\n", err)
		return err
	}
	defer stmt.Close() // Schließe das Statement nach der Ausführung

	// Führt die Abfrage aus und übergibt die Parameter
	_, err = stmt.Exec(systemID, certificationID)
	if err != nil {
		fmt.Printf("Error executing statement: %v\n", err)
		return err
	}

	return nil
}


/////////////////////////////////////////
////	SMS IssueAffectedComponent
////////////////////////////////////////
func (mgr *manager) AddSMSIssueAffectedComponent(component_id int, issue_id int, additionalInfo string, confirmed bool) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newIssueAffectedComponent)
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return err
	}
	defer stmt.Close() // Sicherstellen, dass das Statement geschlossen wird

	_, err = stmt.Exec(component_id, issue_id, additionalInfo, confirmed)
	if err != nil {
		fmt.Println("Error executing statement:", err)
	}

	return err
}

func (mgr *manager) GetSMSIssueAffectedComponentsForIssueID(issue_id int) (issueAffectedComponents []classes.Sms_IssueAffectedComponent, err error) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_IssueAffectedComponentsForIssueID)
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(issue_id)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dbComponent_id int
		var dbIssue_id int
		var dbAdditionalInfo string
		var dbConfirmed bool
		var dbComponent_name string
		var dbComponent_version string

		err := rows.Scan(&dbComponent_id, &dbIssue_id, &dbAdditionalInfo, &dbConfirmed, &dbComponent_name, &dbComponent_version)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		affectedComponent := classes.NewSms_IssueAffectedComponent(dbComponent_id, dbIssue_id, dbAdditionalInfo, dbConfirmed, dbComponent_name, dbComponent_version)
		issueAffectedComponents = append(issueAffectedComponents, *affectedComponent)
	}

	return issueAffectedComponents, nil
}


func (mgr *manager) GetSMSIssuesForComponent(component_id int) (issueAffectedComponents []classes.Sms_IssueAffectedComponent, err error) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_IssuesForComponent)
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(component_id)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dbComponent_id int
		var dbIssue_id int
		var dbAdditionalInfo string
		var dbConfirmed bool
		var dbIssue_name string

		err := rows.Scan(&dbComponent_id, &dbIssue_id, &dbAdditionalInfo, &dbConfirmed, &dbIssue_name)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		issueList := classes.NewSms_IssueAffectedComponent(dbComponent_id, dbIssue_id, dbAdditionalInfo, dbConfirmed, dbIssue_name, "")
		issueAffectedComponents = append(issueAffectedComponents, *issueList)
	}

	return issueAffectedComponents, nil
}

func (mgr *manager) RemoveSMSIssueAffectedComponent(component_id int, issue_id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_sms_IssueAffectedComponent)
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(component_id, issue_id)
	if err != nil {
		fmt.Println("Error executing DELETE statement:", err)
	}

	return err
}

/////////////////////////////////////////
////	SMS IssueAffectedArtefact
////////////////////////////////////////
func (mgr *manager) AddSMSIssueAffectedArtefact(artefact_id int, issue_id int, additionalInfo string, confirmed bool) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.INSERT_sms_newIssueAffectedArtefact)
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return err
	}
	defer stmt.Close() // Sicherstellen, dass das Statement geschlossen wird

	_, err = stmt.Exec(artefact_id, issue_id, additionalInfo, confirmed)
	if err != nil {
		fmt.Println("Error executing statement:", err)
	}

	return err
}

func (mgr *manager) GetSMSIssueAffectedArtefactsForIssueID(issue_id int) (issueAffectedArtefacts []classes.Sms_IssueAffectedArtefact, err error) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_IssueAffectedArtefactsForIssueID)
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(issue_id)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dbArtefact_id int
		var dbIssue_id int
		var dbAdditionalInfo string
		var dbConfirmed bool
		var dbArtefact_name string
		var dbArtefact_version string

		err := rows.Scan(&dbArtefact_id, &dbIssue_id, &dbAdditionalInfo, &dbConfirmed, &dbArtefact_name, &dbArtefact_version)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		affectedArtefact := classes.NewSms_IssueAffectedArtefact(dbArtefact_id, dbIssue_id, dbAdditionalInfo, dbConfirmed, dbArtefact_name, dbArtefact_version)
		issueAffectedArtefacts = append(issueAffectedArtefacts, *affectedArtefact)
	}

	return issueAffectedArtefacts, nil
}

func (mgr *manager) GetSMSIssuesForArtefact(artefact_id int) (issueAffectedArtefacts []classes.Sms_IssueAffectedArtefact, err error) {
	stmt, err := mgr.db.Prepare(dbUtils.SELECT_sms_IssuesForArtefact)
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(artefact_id)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dbArtefact_id int
		var dbIssue_id int
		var dbAdditionalInfo string
		var dbConfirmed bool
		var dbArtefact_name string

		err := rows.Scan(&dbArtefact_id, &dbIssue_id, &dbAdditionalInfo, &dbConfirmed, &dbArtefact_name)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		issueList := classes.NewSms_IssueAffectedArtefact(dbArtefact_id, dbIssue_id, dbAdditionalInfo, dbConfirmed, dbArtefact_name, "")
		issueAffectedArtefacts = append(issueAffectedArtefacts, *issueList)
	}

	return issueAffectedArtefacts, nil
}

func (mgr *manager) RemoveSMSIssueAffectedArtefact(artefact_id int, issue_id int) (err error) {
	stmt, err := mgr.db.Prepare(dbUtils.DELETE_sms_IssueAffectedArtefact)
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(artefact_id, issue_id)
	if err != nil {
		fmt.Println("Error executing DELETE statement:", err)
	}

	return err
}