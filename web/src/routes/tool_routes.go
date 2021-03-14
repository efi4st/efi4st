/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (
	"github.com/efi4st/efi4st/dbprovider"
	"fmt"
	"github.com/kataras/iris/v12"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
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

func EmulationRun(ctx iris.Context) {

	var (
		moduleName = ctx.Params().Get("moduleName")
		firmwareId = ctx.Params().Get("firmwareId")
		firmwareName = ctx.Params().Get("firmwareName")

	)

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := filepath.FromSlash(dir +"/../../modules/shell/")
	fmt.Println("Starting "+ moduleName +" on project "+ firmwareId+":"+path+moduleName)

	var cmd = exec.Command("/bin/sh", path+moduleName, firmwareId, firmwareName)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	fmt.Println(string(out))

	ctx.Redirect("/firmware/show/"+firmwareId)
}

func ModuleOnAppRun(ctx iris.Context) {

var (
moduleName = ctx.Params().Get("moduleName")
firmwareId = ctx.Params().Get("firmwareId")
relevantAppId = ctx.Params().Get("relevantAppId")
)

dir, err := os.Getwd()
if err != nil {
log.Fatal(err)
}

i, err := strconv.Atoi(relevantAppId)

relApp := dbprovider.GetDBManager().GetRelevantAppInfo(i)

path := filepath.FromSlash(dir +"/../../modules/shell/")
fmt.Println("Starting "+ moduleName +" on project "+ firmwareId+":"+path+moduleName)

var cmd = exec.Command("/bin/sh", path+moduleName, firmwareId, "../../working/filesystem" + relApp.Path())
out, err := cmd.Output()

if err != nil {
println(err.Error())
return
}

fmt.Println(string(out))

ctx.Redirect("/relevantApps/show/"+relevantAppId)
}