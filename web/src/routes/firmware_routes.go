/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (
	"../utils"
	"../dbprovider"
	"fmt"
	"github.com/kataras/iris"
	"io"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func Firmwares(ctx iris.Context) {

	firmwares := dbprovider.GetDBManager().GetFirmwares()

	ctx.ViewData("error", "")

	if len(firmwares) < 1 {
		ctx.ViewData("error", "Error: No firmwares available. Add one!")
	}

	ctx.ViewData("firmwareList", firmwares)
	ctx.View("firmwares.html")
}

func ShowFirmware(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing firmware Id!")
	}

	firmware := dbprovider.GetDBManager().GetFirmwareInfo(i)

	ctx.ViewData("firmware", firmware)
	ctx.View("showFirmware.html")
}

// GET
func ShowFirmwareUpload(ctx iris.Context) {

	id := ctx.Params().Get("project_id")
	i, err := strconv.Atoi(id)
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing project Id!")
	}
	project := dbprovider.GetDBManager().GetProjectInfo(i)

	ctx.ViewData("project", project)
	ctx.View("firmwareUpload.html")
}

// GET
func ShowFirmwareApps(ctx iris.Context) {

	id := ctx.Params().Get("firmware_id")
	i, err := strconv.Atoi(id)
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing project Id!")
	}
	apps := dbprovider.GetDBManager().GetFirmwareApps(i)

	ctx.ViewData("apps", apps)
	ctx.View("showFirmwareApplications.html")
}

// POST
func UploadFirmware(ctx iris.Context) {

	id := ctx.Params().Get("project_id")

	i, err := strconv.Atoi(id)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing project Id!")
	}

	// Get the file from the dropzone request
	file, info, err := ctx.FormFile("file")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Application().Logger().Warnf("Error while uploading: %v", err.Error())
		return
	}

	defer file.Close()
	fname := info.Filename

	// Create a file with the same name
	// assuming that you have a folder named 'uploads'
	out, err := os.OpenFile(utils.UploadsDir+fname,
		os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Application().Logger().Warnf("Error while preparing the new file: %v", err.Error())
		return
	}
	defer out.Close()

	fi, err := out.Stat()
	fileSize:=int(fi.Size())
	if err != nil {
		fmt.Println("Error getting file Size!")
		return
	}

	dbprovider.GetDBManager().AddFirmware(fname, fileSize,i)

	proj := dbprovider.GetDBManager().GetProjectInfo(i)
	dbprovider.GetDBManager().UpdateProjectsUploads(i, proj.Uploads()+1)

	io.Copy(out, file)
}

func RemoveFirmware(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveFirmware(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing project!")
	}

	firmwares := dbprovider.GetDBManager().GetFirmwares()

	ctx.ViewData("firmwareList", firmwares)
	ctx.View("firmwares.html")
}




