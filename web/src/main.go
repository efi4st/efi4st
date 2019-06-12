package main

import (
	_ "github.com/go-sql-driver/mysql"
	"./routes"
	"./utils"
	"log"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"fmt"
	"github.com/jmoiron/sqlx"
)

// Serve using a host:port form.
var addr = iris.Addr("0.0.0.0:8144")

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
	utils.CreateDB(db)

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

	// GET: http://localhost:3000/modules/run/xxx/xxx
	app.Get("/modules/run/{moduleName:string}{project:string}/", routes.ModuleRun)

	// Application started. Press CTRL+C to shut down.
	app.Run(addr)
}