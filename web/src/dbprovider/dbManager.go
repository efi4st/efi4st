/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package dbprovider

import (
	"../classes"
	"../dbUtils"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
	"time"
	_ "github.com/go-sql-driver/mysql"
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
	UpdateRelevantApp(column string, relevantApp_id string) error
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
		}else if(dbmoduleInitSystem.Bool){
			count++
		}else if(dbmoduleFileContent.Bool){
			count++
		} else if(dbmoduleBash.Bool){
			count++
		}else if(dbmoduleCronJob.Bool){
			count++
		} else if(dbmoduleProcesses.Bool){
			count++
		} else if(dbmoduleInterfaces.Bool){
			count++
		} else if(dbmoduleSystemControls.Bool){
			count++
		} else if(dbmoduleFileSystem.Bool){
			count++
		} else if(dbmodulePortscanner.Bool){
			count++
		} else if(dbmoduleProtocolls.Bool){
			count++
		} else if(dbmoduleNetInterfaces.Bool){
			count++
		} else if(dbmoduleFileSystemInterfaces.Bool){
			count++
		} else if(dbmoduleFileHandles.Bool){
			count++
		}

		relevantApp.SetMsg(strconv.Itoa(count))
		relevantApps=append(relevantApps, *relevantApp)
		if err != nil {
			log.Fatal(err)
		}
	}

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

