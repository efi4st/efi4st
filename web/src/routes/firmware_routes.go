/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (
	"../utils"
	"fmt"
	"github.com/kataras/iris"
	"io"
	"os"
	_ "github.com/go-sql-driver/mysql"
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

	ctx.View("firmwareUpload.html")
}

// GET
func UploadFirmware(ctx iris.Context) {

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
	fmt.Println(utils.UploadsDir)
	out, err := os.OpenFile(utils.UploadsDir+fname,
		os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Application().Logger().Warnf("Error while preparing the new file: %v", err.Error())
		return
	}
	defer out.Close()

	io.Copy(out, file)
}




