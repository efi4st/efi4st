/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package routes

import (
	"encoding/json"
	"fmt"
	"github.com/efi4st/efi4st/classes"
	"github.com/efi4st/efi4st/dbprovider"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// GET
func CreateSMSComponentPartOfSoftware(ctx iris.Context) {

	id := ctx.Params().Get("id")
	i, err := strconv.Atoi(id)
	components := dbprovider.GetDBManager().GetSMSComponent()
	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing software id!")
	}

	ctx.ViewData("softwareId", i)
	ctx.ViewData("componentList", components)
	ctx.View("sms_createComponentPartOfSoftware.html")
}

// POST
func AddSMSComponentPartOfSoftware(ctx iris.Context) {

	component_ids, err := ctx.PostValueMany("Component_ids")
	software_id := ctx.PostValue("Software_id")
	additionalInfo := ctx.PostValue("AdditionalInfo")

	iI, err := strconv.Atoi(software_id)
	if err != nil {
		ctx.ViewData("error", "Error: Error parsing software_id!")
	}

	for index, componentId := range strings.Split(component_ids,","){
		iD, err := strconv.Atoi(componentId)
		if err != nil {
			ctx.ViewData("error", "Error: Error parsing componentId!")
		}
		err = dbprovider.GetDBManager().AddSMSComponentPartOfSoftware(iI, iD, additionalInfo)
		fmt.Printf(strconv.Itoa(index))
		ctx.ViewData("error", "")
		if err !=nil {
			ctx.ViewData("error", "Error: Not able to add component software link!")
		}
	}

	ctx.Params().Set("id", software_id)
	ShowSMSSoftware(ctx)
}

// GET
func RemoveSMSComponentPartOfSoftware(ctx iris.Context) {

	id := ctx.Params().Get("id")

	i, err := strconv.Atoi(id)
	err = dbprovider.GetDBManager().RemoveSMSComponentPartOfSoftware(i)

	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error removing component software link!")
	}

	SMSProjects(ctx)
}

// GET
func UploadViewSMSComponentPartOfSoftwareSCAReport(ctx iris.Context) {

	id := ctx.Params().Get("id")
	i, err := strconv.Atoi(id)
	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing software id!")
	}

	ctx.ViewData("softwareId", i)
	ctx.View("sms_uploadSCAReportForSoftware.html")
}

// POST
func UploadSMSComponentPartOfSoftwareSCAReport(ctx iris.Context) {

	ctx.SetMaxRequestBodySize(1 * iris.MB)

	appID := ctx.Params().Get("id")
	softwareID, err := strconv.Atoi(appID)
	ctx.ViewData("error", "")
	if err !=nil {
		ctx.ViewData("error", "Error: Error parsing software id!")
	}

	file, fileHeader, err := ctx.FormFile("sbomFile")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("Error uploading file: " + err.Error())
		return
	}
	defer file.Close()

	// Zielverzeichnis definieren
	sanitizedAppID := filepath.Base(appID)
	uploadDir := "./uploads/" + sanitizedAppID
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("Error creating directory: " + err.Error())
		return
	}

	// Dateipfad erstellen
	filePath := uploadDir + "/" + fileHeader.Filename

	// Datei auf dem Server speichern
	outFile, err := os.Create(filePath)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("Error saving file: " + err.Error())
		return
	}
	defer outFile.Close()

	if _, err := io.Copy(outFile, file); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("Error writing file: " + err.Error())
		return
	}

	// Datei verarbeiten (z.B. speichern oder analysieren)
	sbomData, err := io.ReadAll(file)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("Error reading file: " + err.Error())
		return
	}
	fmt.Println(len(sbomData))
	// Beispielpfad zur SBOM-Datei
	sbomFile := filePath

	components, err := parseSBOMFile(sbomFile)
	if err != nil {
		fmt.Printf("Error processing SBOM: %v\n", err)
		return
	}

	// Ausgabe der verarbeiteten Komponenten
	for _, comp := range components {
		fmt.Printf("Component: %+v\n", comp)
	}

	// SBOM in die Datenbank oder einen anderen Speicher integrieren
	dbprovider.GetDBManager().ProcessComponents(components, softwareID)

	ctx.ViewData("error", "SBOM uploaded successfully for Application ID: " + appID)

	ctx.Params().Set("id", appID)
	ShowSMSSoftware(ctx)

}

// CycloneDXComponent Definition (entsprechend CycloneDX JSON)
type CycloneDXComponent struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Version string `json:"version"`
	Licenses []struct {
		License struct{
			ID string `json:"id"`
		} `json:"license"`
	} `json:"licenses"`
}

type CycloneDXSBOM struct {
	Components []CycloneDXComponent `json:"components"`
}

func parseSBOMFile(filePath string) ([]classes.Sms_Component, error) {

	dt := time.Now()

	// Datei öffnen
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open SBOM file: %w", err)
	}
	defer file.Close()

	// SBOM einlesen
	var sbom CycloneDXSBOM
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&sbom); err != nil {
		return nil, fmt.Errorf("failed to decode SBOM JSON: %w", err)
	}

	// Komponenten verarbeiten
	var components []classes.Sms_Component
	for _, comp := range sbom.Components {
		license := ""
		if len(comp.Licenses) > 0 {
			license = comp.Licenses[0].License.ID
		}

		var component = classes.NewSms_Component(comp.Name, comp.Type, comp.Version, dt.String(), license, true, "")

		components = append(components, *component)
	}

	return components, nil
}

