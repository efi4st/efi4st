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
		firmwareId = ctx.Params().Get("firmwareId")
	)

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := filepath.FromSlash(dir +"/../../modules/shell/")
	fmt.Println("Starting "+ moduleName +" on project "+ firmwareId+":"+path+moduleName)

	var cmd = exec.Command("/bin/sh", path+moduleName, firmwareId)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	fmt.Println(string(out))

	ctx.Redirect("/firmware/show/"+firmwareId)
}