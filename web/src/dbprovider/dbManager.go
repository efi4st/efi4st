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
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type Manager interface {
	GetProjects() []classes.Project
	AddProject(project *classes.Project) error
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

func (mgr *manager) AddProject(project *classes.Project) (err error) {
	// TODO
	//mgr.db..Create(project)
	//if errs := mgr.db.GetErrors(); len(errs) > 0 {
	//	err = errs[0]
	//}
	return
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
		var project = classes.NewProject(dbId, dbName, dbUploads, dbDate)
		projects=append(projects, *project)
		if err != nil {
			log.Fatal(err)
		}
	}

	return projects
}