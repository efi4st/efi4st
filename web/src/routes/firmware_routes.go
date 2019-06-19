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

	/*projects := dbprovider.GetDBManager().GetProjects()

	ctx.ViewData("error", "")

	if len(projects) < 1 {
		ctx.ViewData("error", "Error: No projects available. Add one!")
	}

	ctx.ViewData("projectList", projects)
	ctx.View("projects.html")*/
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
	if err != nil {
		fmt.Println("Error getting file Size!")
		return
	}

	dbprovider.GetDBManager().AddFirmware(fname, int(fi.Size()),i)

	io.Copy(out, file)
}




