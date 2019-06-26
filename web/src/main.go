package main

import (
	"./dbUtils"
	"./routes"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"log"
	"./utils"
)

func main(){
	fmt.Printf("### Starting efi4st WEBUI...\n")
	dbInit()
	irisMain()
}

func dbInit()() {

	db, err := sqlx.Connect("mysql", "root:@(localhost:3307)/efi4st")
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
	app.StaticWeb("/static", "./static")

	// GET: http://localhost:8144
	app.Get("/", routes.Index)

	// GET: http://localhost:8144/documentation
	app.Get("/documentation", routes.Documentation)

	// GET: http://localhost:8144/modules/run/xxx/xxx
	app.Get("/modules/run/{moduleName:string}/{project:string}/", routes.ModuleRun)

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

	// GET: http://localhost:8144/firmware/upload/xxx
	app.Get("/firmware/show/upload/{project_id:string}", routes.ShowFirmwareUpload)

	// POST: http://localhost:8144/firmware/upload/xxx
	app.Post("/firmware/upload/{project_id:string}", iris.LimitRequestBodySize(10<<20), routes.UploadFirmware)

	// GET: http://localhost:8144/firmware/show/1
	app.Get("/firmware/show/{id:string}", routes.ShowFirmware)

	// GET: http://localhost:8144/firmware/remove/1
	app.Get("/firmware/remove/{id:string}", routes.RemoveFirmware)

	// Application started. Press CTRL+C to shut down.
	app.Run(utils.Addr)
}