/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package main

import (
	"fmt"
	"github.com/efi4st/efi4st/dbUtils"
	"github.com/efi4st/efi4st/routes"
	"github.com/efi4st/efi4st/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"log"
)

func main(){
	fmt.Printf("### Starting efi4st WEBUI...\n")
	dbInit()
	irisMain()
}

func dbInit()() {

	db, err := sqlx.Connect("mysql", "efi4db:efi4db@tcp(127.0.0.1:3306)/efi4st")

	if err != nil {
		log.Fatalln(err)
	}
	dbUtils.CreateDB(db)

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
	app.HandleDir("/static", "./static")

	// GET: http://localhost:8144
	app.Get("/", routes.Index)

	// GET: http://localhost:8144/documentation
	app.Get("/documentation", routes.Documentation)

	// GET: http://localhost:8144/modules/run/xxx/xxx
	app.Get("/modules/run/{moduleName:string}/{firmwareId:string}", routes.ModuleRun)

	// GET: http://localhost:8144/modules/run/xxx/xxx
	app.Get("/modules/runEmulation/{moduleName:string}/{firmwareId:string}/{firmwareName:string}", routes.EmulationRun)

	// GET: http://localhost:8144/modules/run/xxx/xxx/xxx
	app.Get("/modules/run/{moduleName:string}/{firmwareId:string}/{relevantAppId:string}", routes.ModuleOnAppRun)

	// GET: http://localhost:8144/projects
	app.Get("/projects", routes.Projects)

	// GET: http://localhost:8144/projects/createProject
	app.Get("/projects/createProject", routes.CreateProject)

	// POST: http://localhost:8144/projects/createProject
	app.Post("/projects/addProject", routes.AddProject)

	// GET: http://localhost:8144/projects/show/1
	app.Get("/projects/show/{id:string}", routes.ShowProject)

	// GET: http://localhost:8144/projects/remove/1
	app.Get("/projects/remove/{id:string}", routes.RemoveProject)

	// GET: http://localhost:8144/firmwares
	app.Get("/firmwares", routes.Firmwares)

	// GET: http://localhost:8144/firmware/show/upload/xxx
	app.Get("/firmware/show/upload/{project_id:string}", routes.ShowFirmwareUpload)

	// POST: http://localhost:8144/firmware/upload/xxx
	app.Post("/firmware/upload/{project_id:string}", iris.LimitRequestBodySize(10<<50), routes.UploadFirmware)

	// GET: http://localhost:8144/firmware/show/1
	app.Get("/firmware/show/{id:string}", routes.ShowFirmware)

	// GET: http://localhost:8144/firmware/remove/1
	app.Get("/firmware/remove/{id:string}", routes.RemoveFirmware)

	// GET: http://localhost:8144/relevantApps
	app.Get("/relevantApps", routes.RelevantApps)

	// GET: http://localhost:8144/relevantApps/show/1
	app.Get("/relevantApps/show/{id:string}", routes.ShowRelevantApp)

	// GET: http://localhost:8144/relevantApps/show/1
	app.Get("/relevantApps/showEmu/{id:string}", routes.ShowRelevantAppEmu)

	// GET: http://localhost:8144/relevantApps/download/1
	app.Get("/relevantApps/download/{id:string}", routes.DownloadRelevantApp)

	// GET: http://localhost:8144/relevantApps/remove/1
	app.Get("/relevantApps/remove/{id:string}", routes.RemoveRelevantApp)

	// GET: http://localhost:8144/testResults
	app.Get("/testResults", routes.TestResults)

	// GET: http://localhost:8144/testResults/show/1
	app.Get("/testResults/show/{id:string}", routes.ShowTestResult)

	// GET: http://localhost:8144/testResults/remove/1
	app.Get("/testResults/remove/{id:string}", routes.RemoveTestResult)

	// POST: http://localhost:8144/testResults/addResultSet/xxx
	app.Post("/testResults/addResultSet/{project_id:string}", iris.LimitRequestBodySize(10<<20), routes.AddResultSet)

	// POST: http://localhost:8144/testResults/addRelevantApp/xxx
	app.Post("/testResults/addRelevantApp/{project_id:string}", iris.LimitRequestBodySize(10<<20), routes.AddRelevantApp)

	/**
	 * Security Management System
	 * Created:   29.09.2024
	 *
	 * (C)
	 **/

	// GET: http://localhost:8144/sms_projects
	app.Get("/sms_projects", routes.SMSProjects)

	// GET: http://localhost:8144/sms_projects/createSMSProject
	app.Get("/sms_projects/createSMSProject", routes.CreateSMSProject)

	// POST: http://localhost:8144/sms_projects/createSMSProject
	app.Post("/sms_projects/addSMSProject", routes.AddSMSProject)

	// GET: http://localhost:8144/sms_projects/show/1
	app.Get("/sms_projects/show/{id:string}", routes.ShowSMSProject)

	// GET: http://localhost:8144/sms_projects/remove/1
	app.Get("/sms_projects/remove/{id:string}", routes.RemoveSMSProject)

	// GET: http://localhost:8144/sms_projects/getiplist/1
	app.Get("/sms_projects/getiplist/{project_id:string}", routes.SMSProjectIPs)

	// GET: http://localhost:8144/sms_projects/downloadiplist/1
	app.Get("/sms_projects/downloadiplist/{project_id:string}", routes.SMSExportProjectIPsCSV)

	// GET: http://localhost:8144/sms_projects/downloadiplistCustomer/1
	app.Get("/sms_projects/downloadiplistCustomer/{project_id:string}", routes.SMSExportProjectIPsCSVCustomer)

	app.Get("/sms_projects/checklist/{id:int}/{check_type:string}", routes.SMSProjectCheckList)

	// GET: http://localhost:8144/smssystems
	app.Get("/sms_systems", routes.SMSSystems)

	// GET: http://localhost:8144/smssystems/createSMSSystem
	app.Get("/sms_systems/createSMSSystem", routes.CreateSMSSystem)

	// POST: http://localhost:8144/smssystems/createSMSSystem
	app.Post("/sms_systems/addSMSSystem", routes.AddSMSSystem)

	// GET: http://localhost:8144/smssystems/show/1
	app.Get("/sms_systems/show/{id:string}", routes.ShowSMSSystem)

	// GET: http://localhost:8144/smssystems/remove/1
	app.Get("/sms_systems/remove/{id:string}", routes.RemoveSMSSystem)

	// GET: http://localhost:8144/sms_systems/downloadSystemStructureJSON/1
	app.Get("/sms_systems/downloadSystemStructureJSON/{system_id:string}", routes.DownloadSystemTreeJSON)

	// GET: http://localhost:8144/sms_systems/releasenotes/1
	app.Get("/sms_systems/releasenotes/{id}", routes.ShowSMSReleaseNotesForSystem)

	// GET: http://localhost:8144/smsdevices
	app.Get("/sms_devices", routes.SMSDevice)

	// GET: http://localhost:8144/smsdevices/createSMSDevice
	app.Get("/sms_devices/createSMSDevice", routes.CreateSMSDevice)

	// POST: http://localhost:8144/smsdevices/createSMSDevice
	app.Post("/sms_devices/addSMSDevice", routes.AddSMSDevice)

	// GET: http://localhost:8144/smsdevices/show/1
	app.Get("/sms_devices/show/{id:string}", routes.ShowSMSDevice)

	// GET: http://localhost:8144/smsdevices/remove/1
	app.Get("/sms_devices/remove/{id:string}", routes.RemoveSMSDevice)

	// GET: http://localhost:8144/smsdeviceInstances
	app.Get("/sms_deviceInstances", routes.SMSDeviceInstance)

	// GET: http://localhost:8144/smsdeviceInstances/createSMSDeviceInstance
	app.Get("/sms_deviceInstances/createSMSDeviceInstance", routes.CreateSMSDeviceInstance)

	// GET: http://localhost:8144/smsdeviceInstances/createSMSDeviceInstanceForProject/1
	app.Get("/sms_deviceInstances/createSMSDeviceInstanceForProject/{id:string}", routes.CreateSMSDeviceInstanceForProject)

	// POST: http://localhost:8144/smsdeviceInstances/addSMSDeviceInstance
	app.Post("/sms_deviceInstances/addSMSDeviceInstance", routes.AddSMSDeviceInstance)

	// GET: http://localhost:8144/smsdeviceInstances/show/1
	app.Get("/sms_deviceInstances/show/{id:string}", routes.ShowSMSDeviceInstance)

	app.Post("/sms_deviceInstances/upgrade/{id:string}", routes.UpgradeSMSDeviceInstance)

	// GET: http://localhost:8144/smsdeviceInstances/remove/1
	app.Get("/sms_deviceInstances/remove/{id:string}", routes.RemoveSMSDeviceInstance)

	// GET: http://localhost:8144/smsupdateHistory/createSMSUpdateHistory/1
	app.Get("/smsupdateHistory/createSMSUpdateHistory/{id:string}", routes.CreateSMSUpdateHistory)

	// POST: http://localhost:8144/smsupdateHistory/addSMSUpdateHistory
	app.Post("/smsupdateHistory/addSMSUpdateHistory", routes.AddSMSUpdateHistory)

	// GET: http://localhost:8144/sms_updateHistory/show/1
	app.Get("/sms_updateHistory/show/{id:string}", routes.ShowSMSUpdateHistory)

	// GET: http://localhost:8144/sms_issues
	app.Get("/sms_issues", routes.SMSIssues)

	// GET: http://localhost:8144/sms_issues/createSMSIssue
	app.Get("/sms_issues/createSMSIssue", routes.CreateSMSIssue)

	// POST: http://localhost:8144/sms_issues/addSMSIssue
	app.Post("/sms_issues/addSMSIssue", routes.AddSMSIssue)

	// GET: http://localhost:8144/sms_issues/show/1
	app.Get("/sms_issues/show/{id:string}", routes.ShowSMSIssue)

	// GET: http://localhost:8144/sms_issues/remove/1
	app.Get("/sms_issues/remove/{id:string}", routes.RemoveSMSIssue)

	// GET: http://localhost:8144/sms_issues/show/1
	app.Get("/sms_issues/serviceletter/{id:string}", routes.SMSIssueServiceLetter)

	// GET: http://localhost:8144/sms_issueAffectedDevice/createSMSIssueAffectedDevice/1
	app.Get("/sms_issueAffectedDevice/createSMSIssueAffectedDevice/{id:string}", routes.CreateSMSIssueAffectedDevice)

	// POST: http://localhost:8144/sms_issueAffectedDevice/addSMSIssueAffectedDevice
	app.Post("/sms_issueAffectedDevice/addSMSIssueAffectedDevice", routes.AddSMSIssueAffectedDevice)

	// GET: http://localhost:8144/sms_issueAffectedDevice/remove/1
	app.Get("/sms_issueAffectedDevice/remove/{issueId:string}/{deviceId:string}", routes.RemoveSMSIssueAffectedDevice)

	// POST: http://localhost:8144/sms_solutions/addSMSSolution
	app.Post("/sms_solutions/addSMSSolution", routes.AddSMSSolution)

	// GET: http://localhost:8144/sms_solutions/show/1
	app.Get("/sms_solutions/show/{id:string}", routes.ShowSMSSolution)

	// GET: http://localhost:8144/sms_solutions/createSMSSolution/1
	app.Get("/sms_solutions/createSMSSolution/{id:string}", routes.CreateSMSSolution)

	// GET: http://localhost:8144/sms_artefacts
	app.Get("/sms_artefacts", routes.SMSArtefact)

	// GET: http://localhost:8144/sms_artefacts/createSMSArtefact
	app.Get("/sms_artefacts/createSMSArtefact", routes.CreateSMSArtefact)

	// POST: http://localhost:8144/sms_artefacts/addSMSArtefact
	app.Post("/sms_artefacts/addSMSArtefact", routes.AddSMSArtefact)

	// GET: http://localhost:8144/sms_artefacts/show/1
	app.Get("/sms_artefacts/show/{id:string}", routes.ShowSMSArtefact)

	// GET: http://localhost:8144/sms_artefacts/remove/1
	app.Get("/sms_artefacts/remove/{id:string}", routes.RemoveSMSArtefact)

	// GET: http://localhost:8144/sms_releaseNote/createSMSReleaseNote/1
	app.Get("/sms_releaseNote/createSMSReleaseNote/{id:string}", routes.CreateSMSReleaseNote)

	// POST: http://localhost:8144/sms_releaseNote/addSMSReleaseNote
	app.Post("/sms_releaseNote/addSMSReleaseNote", routes.AddSMSReleaseNote)

	// GET: http://localhost:8144/sms_releaseNote/show/1
	app.Get("/sms_releaseNote/show/{id:string}", routes.ShowSMSReleaseNote)

	// GET: http://localhost:8144/sms_softwares
	app.Get("/sms_softwares", routes.SMSSoftware)

	// GET: http://localhost:8144/sms_softwares/createSMSSoftware
	app.Get("/sms_softwares/createSMSSoftware", routes.CreateSMSSoftware)

	// POST: http://localhost:8144/sms_softwares/addSMSSoftware
	app.Post("/sms_softwares/addSMSSoftware", routes.AddSMSSoftware)

	// GET: http://localhost:8144/sms_softwares/show/1
	app.Get("/sms_softwares/show/{id:string}", routes.ShowSMSSoftware)

	// GET: http://localhost:8144/sms_softwares/remove/1
	app.Get("/sms_softwares/remove/{id:string}", routes.RemoveSMSSoftware)

	// GET: http://localhost:8144/sms_components
	app.Get("/sms_components", routes.SMSComponent)

	// GET: http://localhost:8144/sms_components/createSMSComponent
	app.Get("/sms_components/createSMSComponent", routes.CreateSMSComponent)

	// POST: http://localhost:8144/sms_components/addSMSComponent
	app.Post("/sms_components/addSMSComponent", routes.AddSMSComponent)

	// GET: http://localhost:8144/sms_components/show/1
	app.Get("/sms_components/show/{id:string}", routes.ShowSMSComponent)

	// GET: http://localhost:8144/sms_components/remove/1
	app.Get("/sms_components/remove/{id:string}", routes.RemoveSMSComponent)

	// GET: http://localhost:8144/sms_componentPartOfSoftware/createSMSComponentPartOfSoftware/1
	app.Get("/sms_componentPartOfSoftware/createSMSComponentPartOfSoftware/{id:string}", routes.CreateSMSComponentPartOfSoftware)

	// POST: http://localhost:8144/sms_componentPartOfSoftware/addSMSComponentPartOfSoftware
	app.Post("/sms_componentPartOfSoftware/addSMSComponentPartOfSoftware", routes.AddSMSComponentPartOfSoftware)

	// GET: http://localhost:8144/sms_componentPartOfSoftware/importSCAreportToComponentPartOfSoftwareView/1
	app.Get("/sms_componentPartOfSoftware/importSCAreportToComponentPartOfSoftwareView/{id:string}", routes.UploadViewSMSComponentPartOfSoftwareSCAReport)

	// POST: http://localhost:8144/sms_componentPartOfSoftware/1/upload-sbom
	app.Post("/sms_componentPartOfSoftware/{id:string}/upload-sbom", routes.UploadSMSComponentPartOfSoftwareSCAReport)

	// GET: http://localhost:8144/sms_componentPartOfSoftware/remove/1
	app.Get("/sms_componentPartOfSoftware/remove/{id:string}", routes.RemoveSMSComponentPartOfSoftware)

	// GET: http://localhost:8144/sms_softwarePartOfDevice/createSMSSoftwarePartOfDevice/1
	app.Get("/sms_softwarePartOfDevice/createSMSSoftwarePartOfDevice/{id:string}", routes.CreateSMSSoftwarePartOfDevice)

	// POST: http://localhost:8144/sms_softwarePartOfDevice/addSMSSoftwarePartOfDevice
	app.Post("/sms_softwarePartOfDevice/addSMSSoftwarePartOfDevice", routes.AddSMSSoftwarePartOfDevice)

	// GET: http://localhost:8144/sms_softwarePartOfDevice/remove/1
	app.Get("/sms_softwarePartOfDevice/remove/{id:string}", routes.RemoveSMSSoftwarePartOfDevice)

	// GET: http://localhost:8144/sms_devicePartOfSystem/createSMSDevicePartOfSystem/1
	app.Get("/sms_devicePartOfSystem/createSMSDevicePartOfSystem/{id:string}", routes.CreateSMSDevicePartOfSystem)

	// POST: http://localhost:8144/sms_devicePartOfSystem/addSMSDevicePartOfSystem
	app.Post("/sms_devicePartOfSystem/addSMSDevicePartOfSystem", routes.AddSMSDevicePartOfSystem)

	// GET: http://localhost:8144/sms_devicePartOfSystem/remove/1
	app.Get("/sms_devicePartOfSystem/remove/{id:string}", routes.RemoveSMSDevicePartOfSystem)

	// GET: http://localhost:8144/sms_projectBOM/createSMSProjectBOMForProject/1
	app.Get("/sms_projectBOM/createSMSProjectBOMForProject/{id:string}", routes.CreateSMSProjectBOMForProject)

	// POST: http://localhost:8144/sms_projectBOM/addSMSProjectBOM
	app.Post("/sms_projectBOM/addSMSProjectBOM", routes.AddSMSProjectBOM)

	// GET: http://localhost:8144/sms_projectBOM/remove/1
	app.Get("/sms_projectBOM/remove/{id:string}", routes.RemoveSMSProjectBOM)

	// GET: http://localhost:8144/sms_issueAffectedSoftware/createSMSIssueAffectedSoftware/1
	app.Get("/sms_issueAffectedSoftware/createSMSIssueAffectedSoftware/{id:string}", routes.CreateSMSIssueAffectedSoftware)

	// POST: http://localhost:8144/sms_issueAffectedSoftware/addSMSIssueAffectedSoftware
	app.Post("/sms_issueAffectedSoftware/addSMSIssueAffectedSoftware", routes.AddSMSIssueAffectedSoftware)

	// GET: http://localhost:8144/sms_issueAffectedSoftware/remove/1/1
	app.Get("/sms_issueAffectedSoftware/remove/{issue_id:string}/{software_id:string}", routes.RemoveSMSIssueAffectedSoftware)

	// GET: http://localhost:8144/sms_artefactPartOfDevice/createSMSArtefactPartOfDevice/1
	app.Get("/sms_artefactPartOfDevice/createSMSArtefactPartOfDevice/{id:string}", routes.CreateSMSArtefactPartOfDevice)

	// POST: http://localhost:8144/sms_artefactPartOfDevice/addSMSArtefactPartOfDevice
	app.Post("/sms_artefactPartOfDevice/addSMSArtefactPartOfDevice", routes.AddSMSArtefactPartOfDevice)

	// GET: http://localhost:8144/sms_artefactPartOfDevice/remove/1
	app.Get("/sms_artefactPartOfDevice/remove/{id:string}", routes.RemoveSMSArtefactPartOfDevice)

	// GET: http://localhost:8144/sms_manufacturingOrder/createSMSManufacturingOrder/1
	app.Get("/sms_manufacturingOrder/createSMSManufacturingOrder/{id:string}", routes.CreateSMSManufacturingOrder)

	// POST: http://localhost:8144/sms_manufacturingOrder/addSMSManufacturingOrder
	app.Post("/sms_manufacturingOrder/addSMSManufacturingOrder", routes.AddSMSManufacturingOrder)

	// GET: http://localhost:8144/sms_manufacturingOrder/show/1
	app.Get("/sms_manufacturingOrder/show/{id:string}", routes.ShowSMSManufacturingOrder)

	// GET: http://localhost:8144/sms_certifications
	app.Get("/sms_certifications", routes.SMSCertification)

	// GET: http://localhost:8144/sms_certifications/createSMSCertification
	app.Get("/sms_certifications/createSMSCertification", routes.CreateSMSCertification)

	// POST: http://localhost:8144/sms_certifications/addSMSCertification
	app.Post("/sms_certifications/addSMSCertification", routes.AddSMSCertification)

	// GET: http://localhost:8144/sms_certifications/show/1
	app.Get("/sms_certifications/show/{id:string}", routes.ShowSMSCertification)

	// GET: http://localhost:8144/sms_certifications/remove/1
	app.Get("/sms_certifications/remove/{id:string}", routes.RemoveSMSCertification)

	// GET: http://localhost:8144/sms_systemHasCertification/createSMSSystemHasCertification/1
	app.Get("/sms_systemHasCertification/createSMSSystemHasCertification/{id:string}", routes.CreateSMSSystemHasCertification)

	// POST: http://localhost:8144/sms_systemHasCertification/addSMSSystemHasCertification
	app.Post("/sms_systemHasCertification/addSMSSystemHasCertification", routes.AddSMSSystemHasCertification)

	// GET: http://localhost:8144/sms_systemHasCertification/removeSMSSystemHasCertification/1/2
	app.Get("/sms_systemHasCertification/removeSMSSystemHasCertification/{systemId:string}/{certificationId:string}", routes.RemoveSMSSystemHasCertification)

	// GET: http://localhost:8144/sms_issueAffectedComponent/createSMSIssueAffectedComponent/1
	app.Get("/sms_issueAffectedComponent/createSMSIssueAffectedComponent/{id:string}", routes.CreateSMSIssueAffectedComponent)

	// POST: http://localhost:8144/sms_issueAffectedComponent/addSMSIssueAffectedComponent
	app.Post("/sms_issueAffectedComponent/addSMSIssueAffectedComponent", routes.AddSMSIssueAffectedComponent)

	// GET: http://localhost:8144/sms_issueAffectedComponent/remove/1/1
	app.Get("/sms_issueAffectedComponent/remove/{issueId:string}/{componentId:string}", routes.RemoveSMSIssueAffectedComponent)

	// GET: http://localhost:8144/sms_issueAffectedArtefact/createSMSIssueAffectedArtefact/1
	app.Get("/sms_issueAffectedArtefact/createSMSIssueAffectedArtefact/{id:string}", routes.CreateSMSIssueAffectedArtefact)

	// POST: http://localhost:8144/sms_issueAffectedArtefact/addSMSIssueAffectedArtefact
	app.Post("/sms_issueAffectedArtefact/addSMSIssueAffectedArtefact", routes.AddSMSIssueAffectedArtefact)

	// GET: http://localhost:8144/sms_issueAffectedArtefact/remove/1/1
	app.Get("/sms_issueAffectedArtefact/remove/{issueId:string}/{artefactId:string}", routes.RemoveSMSIssueAffectedArtefact)

	// GET: http://localhost:8144/sms_securityReports
	app.Get("/sms_securityReports", routes.SMSSecurityReports)

	// GET: http://localhost:8144/sms_securityReports/createSMSSecurityReport
	app.Get("/sms_securityReports/createSMSSecurityReport", routes.CreateSMSSecurityReport)

	// POST: http://localhost:8144/sms_securityReports/addSMSSecurityReport
	app.Post("/sms_securityReports/addSMSSecurityReport", routes.AddSMSSecurityReport)

	// GET: http://localhost:8144/sms_securityReports/show/1
	app.Get("/sms_securityReports/show/{id:string}", routes.ShowSMSSecurityReport)

	// GET: http://localhost:8144/sms_securityReports/remove/1
	app.Get("/sms_securityReports/remove/{id:string}", routes.RemoveSMSSecurityReport)

	// POST: Datei-Upload für einen Security Report
	app.Post("/sms_securityReports/upload", routes.UploadSecurityReportFile)

	// GET: Anzeige des Security Reports als Datei
	app.Get("/sms_securityReports/view/{report_id:string}", routes.ViewSecurityReport)

	// GET: Download des Report Files, aber Anzeige direkt im Browser
	app.Get("/uploads/reports/{file:path}", routes.GetSecurityReportFile)

	// GET: http://localhost:8144/sms_securityReportLink/createSMSSecurityReportLink/1/Device
	app.Get("/sms_securityReportLink/createSMSSecurityReportLink/{linkedObjectId:string}/{linkedObjectType:string}", routes.CreateSMSSecurityReportLink)

	// POST: http://localhost:8144/sms_securityReportLink/addReportLink
	app.Post("/sms_securityReportLink/addReportLink", routes.AddSMSSecurityReportLink)

	// GET: http://localhost:8144/sms_securityReportLink/remove/1/1
	app.Get("/sms_securityReportLink/remove/{reportId:string}/{linkedObjectId:string}/{linkedObjectType:string}", routes.RemoveSMSSecurityReportLink)

	// GET: http://localhost:8144/sms_securityReportLink/linked/1/sms_device
	app.Get("/sms_securityReportLink/linked/{linkedObjectId:int}/{linkedObjectType:string}", routes.RedirectToLinkedObject)

	// GET: http://localhost:8144/sms_projectSettings
	app.Get("/sms_projectSettings", routes.SMSProjectSettings)

	// GET: http://localhost:8144/sms_projectSettings/createSMSProjectSettings
	app.Get("/sms_projectSettings/createSMSProjectSettings", routes.CreateSMSProjectSettings)

	// POST: http://localhost:8144/sms_projectSettings/addSMSProjectSettings
	app.Post("/sms_projectSettings/addSMSProjectSettings", routes.AddSMSProjectSetting)

	// GET: http://localhost:8144/sms_projectSettings/remove/1
	app.Get("/sms_projectSettings/remove/{id:string}", routes.RemoveSMSProjectSettings)

	// GET: http://localhost:8144/sms_projectSettingsLink/createSMSProjectSettingsLink/1
	app.Get("/sms_projectSettingsLink/createSMSProjectSettingsLink/{id:string}", routes.CreateSMSProjectSettingsLink)

	// POST: http://localhost:8144/sms_projectSettingsLink/addSMSProjectSettingsLink
	app.Post("/sms_projectSettingsLink/addSMSProjectSettingsLink", routes.AddSMSProjectSettingsLink)

	// GET: http://localhost:8144/sms_projectSettingsLink/remove/1/1
	app.Get("/sms_projectSettingsLink/remove/{project_id:string}/{setting_id:string}", routes.RemoveSMSProjectSettingsLink)

	// GET: http://localhost:8144/sms_deviceIPDefinitions
	app.Get("/sms_deviceIPDefinitions", routes.SMSDeviceIPDefinitions)

	// GET: http://localhost:8144/sms_deviceIPDefinitions/createSMSDeviceIPDefinitions
	app.Get("/sms_deviceIPDefinitions/createSMSDeviceIPDefinitions", routes.CreateSMSDeviceIPDefinitions)

	// POST: http://localhost:8144/sms_deviceIPDefinitions/addSMSDeviceIPDefinitions
	app.Post("/sms_deviceIPDefinitions/addSMSDeviceIPDefinitions", routes.AddSMSDeviceIPDefinition)

	// GET: http://localhost:8144/sms_deviceIPDefinitions/remove/1
	app.Get("/sms_deviceIPDefinitions/remove/{id:string}", routes.RemoveSMSDeviceIPDefinition)

	// GET: Übersicht aller Device Check Definitions
	app.Get("/sms_deviceCheckDefinitions", routes.SMSDeviceCheckDefinitions)

	// GET: Seite zum Erstellen einer neuen Device Check Definition
	app.Get("/sms_deviceCheckDefinitions/createSMSDeviceCheckDefinition", routes.CreateSMSDeviceCheckDefinitions)

	// POST: Neue Device Check Definition hinzufügen
	app.Post("/sms_deviceCheckDefinitions/addSMSDeviceCheckDefinition", routes.AddSMSDeviceCheckDefinition)

	// GET: Device Check Definition entfernen
	app.Get("/sms_deviceCheckDefinitions/remove/{id:string}", routes.RemoveSMSDeviceCheckDefinition)

	app.Get("/sms_deviceCheckDefinitions/show/{id:int}", routes.SMSDeviceCheckDetails)

	app.Post("/sms_deviceCheckDefinitions/update/{id:int}", routes.SMSUpdateCheckDefinition)

	app.Get("/sms_deviceCheckDefinitions/edit/{id:int}", routes.SMSEditProjectCheck)

	// GET: http://localhost:8144/sms_statistics
	app.Get("/sms_statistics", routes.ShowStatistics)

	app.Get("/sms_projectUpdates/show/{id:int}", routes.SMSprojectUpdate)

	// GET: http://localhost:8144/sms_updates
	app.Get("/sms_updates", routes.SMSUpdates)

	// GET: http://localhost:8144/sms_updates/create
	app.Get("/sms_updates/create", routes.CreateSMSUpdate)

	// POST: http://localhost:8144/sms_updates/add
	app.Post("/sms_updates/add", routes.AddSMSUpdate)

	// POST: http://localhost:8144/sms_updates/add
	app.Post("/sms_updates/update/{id:string}", routes.SMSUpdateEditPost)

	// GET: http://localhost:8144/sms_updates
	app.Get("/sms_updates/edit/{id:string}", routes.EditSMSUpdateForm)

	// GET: http://localhost:8144/sms_updates/show/1
	app.Get("/sms_updates/show/{id:string}", routes.ShowSMSUpdateDetails)

	// GET: http://localhost:8144/sms_updates/remove/1
	app.Get("/sms_updates/remove/{id:string}", routes.RemoveSMSUpdate)

	// GET: http://localhost:8144/sms_update_packages
	app.Get("/sms_update_packages", routes.SMSUpdatePackages)

	// GET: http://localhost:8144/sms_update_packages/create
	app.Get("/sms_update_packages/create", routes.CreateSMSUpdatePackage)

	// POST: http://localhost:8144/sms_update_packages/add
	app.Post("/sms_update_packages/add", routes.AddSMSUpdatePackage)

	// GET: http://localhost:8144/sms_update_packages/show/1
	app.Get("/sms_update_packages/show/{id:string}", routes.ShowSMSUpdatePackage)

	// GET: http://localhost:8144/sms_update_packages/remove/1
	app.Get("/sms_update_packages/remove/{id:string}", routes.RemoveSMSUpdatePackage)

	// GET: http://localhost:8144/sms_update_centers
	app.Get("/sms_update_centers", routes.SMSUpdateCenters)

	// GET: http://localhost:8144/sms_update_centers/create
	app.Get("/sms_update_centers/create", routes.CreateSMSUpdateCenter)

	// POST: http://localhost:8144/sms_update_centers/add
	app.Post("/sms_update_centers/add", routes.AddSMSUpdateCenter)

	// GET: http://localhost:8144/sms_update_centers/show/{id:string}
	app.Get("/sms_update_centers/show/{id:string}", routes.ShowSMSUpdateCenter)

	// GET: http://localhost:8144/sms_update_centers/remove/{id:string}
	app.Get("/sms_update_centers/remove/{id:string}", routes.RemoveSMSUpdateCenter)

	// GET: http://localhost:8144/sms_update_centers/edit/{id:string}
	app.Get("/sms_update_centers/edit/{id:string}", routes.EditSMSUpdateCenter)

	// POST: http://localhost:8144/sms_update_centers/update
	app.Post("/sms_update_centers/update/{id:string}", routes.UpdateSMSUpdateCenter)

	// POST: http://localhost:8144/sms_update_centers/update_last_contact/{id:string}
	app.Post("/sms_update_centers/ping/{id:string}", routes.PingSMSUpdateCenter)

	// GET: Artefakt-Zuordnung zu einer Device-Instanz erstellen
	app.Get("/sms_artefactPartOfDeviceInstance/create/{id:string}", routes.CreateSMSArtefactPartOfDeviceInstance)

	// POST: Artefakt(e) einer Device-Instanz zuordnen
	app.Post("/sms_artefactPartOfDeviceInstance/add", routes.AddSMSArtefactPartOfDeviceInstance)

	// GET: Artefakt-Zuordnung von einer Device-Instanz entfernen
	app.Get("/sms_artefactPartOfDeviceInstance/remove/{deviceInstanceId:string}/{artefactId:string}", routes.RemoveSMSArtefactPartOfDeviceInstance)

	// GET: Zeigt Formular zur Verknüpfung eines Artefakts mit einem System
	// Beispiel: http://localhost:8144/sms_artefactPartOfSystem/createSMSArtefactPartOfSystem/1
	app.Get("/sms_artefactPartOfSystem/createSMSArtefactPartOfSystem/{id:string}", routes.CreateSMSArtefactPartOfSystem)

	// POST: Fügt eine Artefakt-System-Verknüpfung hinzu
	// Beispiel: http://localhost:8144/sms_artefactPartOfSystem/addSMSArtefactPartOfSystem
	app.Post("/sms_artefactPartOfSystem/addSMSArtefactPartOfSystem", routes.AddSMSArtefactPartOfSystem)

	// GET: Entfernt eine Artefakt-System-Verknüpfung (benötigt Query-Parameter)
	// Beispiel: http://localhost:8144/sms_artefactPartOfSystem/remove?system_id=1&artefact_id=4
	app.Get("/sms_artefactPartOfSystem/remove", routes.RemoveSMSArtefactPartOfSystem)

	app.Get("/sms_projectstatus/create/{project_id}", routes.CreateSMSProjectStatusLog)
	app.Post("/sms_projectstatus/add", routes.AddSMSProjectStatusLog)
	app.Get("/sms_projectstatus/show/{id}", routes.ShowSMSProjectStatusLog)

	app.Get("/sms_project/{project_id:int}/export-structure", routes.SMSExportProjectStructureCSV)
	app.Get("/sms_project/exportprojectstructureyaml/{project_id:int}", routes.SMSExportProjectStructureYAML)

	// element search
	//http://localhost:8144/sms_elementsearch-ui
	app.Get("/sms_elementsearch-ui", routes.ShowElementSearchPage)
	app.Get("/elementsearch", routes.SearchElementsAPI)

	// Application started. Press CTRL+C to shut down.
	app.Run(utils.Addr)
}