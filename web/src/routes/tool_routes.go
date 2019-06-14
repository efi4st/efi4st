/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (
	"fmt"
	"github.com/kataras/iris"
	"log"
	"os"
	"os/exec"
	"path/filepath"

)

func ModuleRun(ctx iris.Context) {

	var (
		moduleName = ctx.Params().Get("moduleName")
		project = ctx.Params().Get("project")
	)

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}






	path := filepath.FromSlash(dir +"../../modules/python/")
	fmt.Println("Starting "+ moduleName +" on project "+ project+":")

	var cmd = exec.Command("python", path+moduleName, "-c", "-m", "-r", "report", dir)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	fmt.Println(string(out))

	ctx.Redirect("/")
}