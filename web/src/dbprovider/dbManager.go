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
}

type manager struct {
	db *sqlx.DB
}

var dbMgr Manager
func GetDBManager() Manager { return dbMgr }

func init() {
	db, err := sqlx.Connect("mysql", "root:@(localhost:3306)/efi4st?parseTime=true")
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
			dbCall string			)

	for rows.Next() {
		err := rows.Scan(&dbAnalysisTool_id, &dbName, &dbCall)
		var analysisTool = classes.NewAnalysisTool(dbAnalysisTool_id, dbName, dbCall)
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
			dbCall string			)

	row := stmt.QueryRow(id)
	row.Scan(&dbAnalysisTool_id, &dbName, &dbCall)

	var analysisTool = classes.NewAnalysisTool(dbAnalysisTool_id, dbName, dbCall)

	return analysisTool
}

func (mgr *manager) AddAnalysisTool(analysisToolName string,  callPattern string) (err error) {

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_newAnalysisTool)
	if err != nil{
		fmt.Print(err)
	}
	rows, err := stmt.Query(analysisToolName, callPattern)

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
