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
	// Add other methods
}

type manager struct {
	db *sqlx.DB
}

var dbMgr Manager
func GetDBManager() Manager { return dbMgr }

func init() {
	db, err := sqlx.Connect("mysql", "root:@(localhost:3307)/efi4st")
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
	dt.Format("01-02-2006")

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_newProject)
	if err != nil{
		fmt.Print(err)
	}

	rows, err := stmt.Query(projectName, 0, dt.Format("01-02-2006"))

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
			dbDate string	)

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
			dbDate string	)

	row := stmt.QueryRow(id)
	row.Scan(&dbId, &dbName, &dbUploads, &dbDate)

	var project = classes.NewProjectFromDB(dbId, dbName, dbUploads, dbDate)

	return project
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
			dbproject_id int	)

	for rows.Next() {
		err := rows.Scan(&dbfirmware_id, &dbname, &dbversion, &dbbinwalkOutput, &dbsizeInBytes, &dbproject_id)
		var firmware = classes.NewFirmware(dbfirmware_id, dbname, dbversion, dbbinwalkOutput.String, dbsizeInBytes, dbproject_id)
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
func (mgr *manager) AddFirmware(firmwareName string, size int, proj_id int) (err error) {
	dt := time.Now()
	dt.Format("01-02-2006")

	stmt, err := mgr.db.Prepare(dbUtils.INSERT_newFirmware)
	if err != nil{
		fmt.Print(err)
	}

	rows, err := stmt.Query(firmwareName, "", size, proj_id)

	if rows == nil{
		fmt.Println("rows should be null")
	}

	return err
}