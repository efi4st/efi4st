package main

import (
	"./routes"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	"fmt"
)

// Serve using a host:port form.
var addr = iris.Addr("0.0.0.0:8144")

func main(){
	fmt.Printf("### Starting efi4st WEBUI...\n")
	irisMain()
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

	// GET: http://localhost:3000/tools/run/xxx
	app.Get("/tools/run/{toolname:string}", routes.ToolRun)

	// Application started. Press CTRL+C to shut down.
	app.Run(addr)
}