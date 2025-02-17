/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (
	"fmt"
	"github.com/efi4st/efi4st/dbprovider"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func SMSSoftware(ctx iris.Context) {

	softwares := dbprovider.GetDBManager().GetSMSSoftware()
	ctx.ViewData("error", "")

	if len(softwares) < 1 {
		ctx.ViewData("error", "Error: No softwares available. Add one!")
	}
	ctx.ViewData("softwareList", softwares)
	ctx.View("sms_softwares.html")
}

// GET
func CreateSMSSoftware(ctx iris.Context) {

	softwareTypes := dbprovider.GetDBManager().GetSMSSoftwareTypes()

	ctx.ViewData("typeList", softwareTypes)
	ctx.View("sms_createSoftware.html")
}

// POST
func AddSMSSoftware(ctx iris.Context) {

	softwaretypeId := ctx.PostValue("SoftwaretypeId")
	version := ctx.PostValue("Version")
	license := ctx.PostValue("License")
	thirdParty := ctx.PostValue("ThirdParty")
	thirdPartyBool, err := strconv.ParseBool(thirdParty)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing thirdParty Bool!")
	}
	releaseNote := ctx.PostValue("ReleaseNote")

	i, err := strconv.Atoi(softwaretypeId)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing softwaretypeId!")
	}

	err = dbprovider.GetDBManager().AddSMSSoftware(i, version, license, thirdPartyBool, releaseNote)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Not able to add software!")
	}
	softwares := dbprovider.GetDBManager().GetSMSSoftware()
	ctx.ViewData("softwareList", softwares)
	ctx.View("sms_softwares.html")
}

func directoryExists(path string) bool {
	// Versuche, die Informationen über den Pfad zu erhalten
	info, err := os.Stat(path)
	if err != nil {
		// Wenn ein Fehler auftritt, prüfen wir, ob es sich um einen "Existenzfehler" handelt
		if os.IsNotExist(err) {
			return false // Das Verzeichnis existiert nicht
		}
		// Falls es sich um einen anderen Fehler handelt, geben wir diesen zurück
		fmt.Println("Error checking directory:", err)
		return false
	}
	// Prüfe, ob es sich um ein Verzeichnis handelt
	return info.IsDir()
}

func ShowSMSSoftware(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)

	//ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing software Id!")
	}

	// Verzeichnis, in dem die SBOMs gespeichert werden
	safeAppID := filepath.Base(id)
	uploadDir := "./uploads/" + safeAppID

	var sbomFiles []string

	if directoryExists(uploadDir) {
		files, err := ioutil.ReadDir(uploadDir)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString("Error reading SBOM files: " + err.Error())
			return
		}

		for _, file := range files {
			if !file.IsDir() && (strings.HasSuffix(file.Name(), ".json") || strings.HasSuffix(file.Name(), ".xml")) {
				sbomFiles = append(sbomFiles, file.Name())
			}
		}
	} else {
		fmt.Println("Directory does not exist!")
	}

	software := dbprovider.GetDBManager().GetSMSSoftwareInfo(i)
	componentsUnderSoftware := dbprovider.GetDBManager().GetSMSComponentPartOfSoftwareForSoftware(i)
	devicesParentsOfSoftware := dbprovider.GetDBManager().GetSMSSoftwarePartOfDeviceForSoftware(i)
	issuesForThisSoftware := dbprovider.GetDBManager().GetSMSIssuesForSoftware(i)
	reportsForThisSoftware, err := dbprovider.GetDBManager().GetReportsForLinkedObject(i,"sms_software")

	ctx.ViewData("devicesParentsOfSoftware", devicesParentsOfSoftware)
	ctx.ViewData("componentsUnderSoftware", componentsUnderSoftware)
	ctx.ViewData("reportsForThisSoftware", reportsForThisSoftware)
	ctx.ViewData("software", software)
	ctx.ViewData("SBOMFiles", sbomFiles)
	ctx.ViewData("issuesForThisSoftware", issuesForThisSoftware)
	ctx.View("sms_showSoftware.html")
}

// GET
func RemoveSMSSoftware(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveSMSSoftware(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing software!")
	}

	softwares := dbprovider.GetDBManager().GetSMSSoftware()

	ctx.ViewData("softwareList", softwares)
	ctx.View("sms_softwares.html")
}


