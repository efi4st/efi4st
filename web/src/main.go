package main

import (
	"github.com/efi4st/efi4st/dbUtils"
	"github.com/efi4st/efi4st/routes"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"log"
	"github.com/efi4st/efi4st/utils"
)

func main(){
	fmt.Printf("### Starting efi4st WEBUI...\n")
	dbInit()
	irisMain()
}

func dbInit()() {

	db, err := sqlx.Connect("mysql", "efi4db:efi4db@tcp(127.0.0.1:3306)/efi4st")

	if err != nil {
		log.Fatalln(err)
	}
	dbUtils.CreateDB(db)

	db.Close()
}

func irisMain()() {

	fmt.Println("### Started WEBUI!!! Now ready to use..")
	app := iris.New()

	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	// Register templates and embed them into layout
	app.RegisterView(iris.Django("./templates", ".html"))

	// Serve static content like css, js, images
	app.HandleDir("/static", "./static")

	// GET: http://localhost:8144
	app.Get("/", routes.Index)

	// GET: http://localhost:8144/documentation
	app.Get("/documentation", routes.Documentation)

	// GET: http://localhost:8144/modules/run/xxx/xxx
	app.Get("/modules/run/{moduleName:string}/{firmwareId:string}", routes.ModuleRun)

	// GET: http://localhost:8144/modules/run/xxx/xxx
	app.Get("/modules/runEmulation/{moduleName:string}/{firmwareId:string}/{firmwareName:string}", routes.EmulationRun)

	// GET: http://localhost:8144/modules/run/xxx/xxx/xxx
	app.Get("/modules/run/{moduleName:string}/{firmwareId:string}/{relevantAppId:string}", routes.ModuleOnAppRun)

	// GET: http://localhost:8144/projects
	app.Get("/projects", routes.Projects)

	// GET: http://localhost:8144/projects/createProject
	app.Get("/projects/createProject", routes.CreateProject)

	// POST: http://localhost:8144/projects/createProject
	app.Post("/projects/addProject", routes.AddProject)

	// GET: http://localhost:8144/projects/show/1
	app.Get("/projects/show/{id:string}", routes.ShowProject)

	// GET: http://localhost:8144/projects/remove/1
	app.Get("/projects/remove/{id:string}", routes.RemoveProject)

	// GET: http://localhost:8144/firmwares
	app.Get("/firmwares", routes.Firmwares)

	// GET: http://localhost:8144/firmware/show/upload/xxx
	app.Get("/firmware/show/upload/{project_id:string}", routes.ShowFirmwareUpload)

	// POST: http://localhost:8144/firmware/upload/xxx
	app.Post("/firmware/upload/{project_id:string}", iris.LimitRequestBodySize(10<<50), routes.UploadFirmware)

	// GET: http://localhost:8144/firmware/show/1
	app.Get("/firmware/show/{id:string}", routes.ShowFirmware)

	// GET: http://localhost:8144/firmware/remove/1
	app.Get("/firmware/remove/{id:string}", routes.RemoveFirmware)

	// GET: http://localhost:8144/relevantApps
	app.Get("/relevantApps", routes.RelevantApps)

	// GET: http://localhost:8144/relevantApps/show/1
	app.Get("/relevantApps/show/{id:string}", routes.ShowRelevantApp)

	// GET: http://localhost:8144/relevantApps/show/1
	app.Get("/relevantApps/showEmu/{id:string}", routes.ShowRelevantAppEmu)

	// GET: http://localhost:8144/relevantApps/download/1
	app.Get("/relevantApps/download/{id:string}", routes.DownloadRelevantApp)

	// GET: http://localhost:8144/relevantApps/remove/1
	app.Get("/relevantApps/remove/{id:string}", routes.RemoveRelevantApp)

	// GET: http://localhost:8144/testResults
	app.Get("/testResults", routes.TestResults)

	// GET: http://localhost:8144/testResults/show/1
	app.Get("/testResults/show/{id:string}", routes.ShowTestResult)

	// GET: http://localhost:8144/testResults/remove/1
	app.Get("/testResults/remove/{id:string}", routes.RemoveTestResult)

	// POST: http://localhost:8144/testResults/addResultSet/xxx
	app.Post("/testResults/addResultSet/{project_id:string}", iris.LimitRequestBodySize(10<<20), routes.AddResultSet)

	// POST: http://localhost:8144/testResults/addRelevantApp/xxx
	app.Post("/testResults/addRelevantApp/{project_id:string}", iris.LimitRequestBodySize(10<<20), routes.AddRelevantApp)

	// Application started. Press CTRL+C to shut down.
	app.Run(utils.Addr)
}
