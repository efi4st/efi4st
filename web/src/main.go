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

	// GET: http://localhost:8144/smsprojects
	app.Get("/sms_projects", routes.SMSProjects)

	// GET: http://localhost:8144/smsprojects/createSMSProject
	app.Get("/sms_projects/createSMSProject", routes.CreateSMSProject)

	// POST: http://localhost:8144/smsprojects/createSMSProject
	app.Post("/sms_projects/addSMSProject", routes.AddSMSProject)

	// GET: http://localhost:8144/smsprojects/show/1
	app.Get("/sms_projects/show/{id:string}", routes.ShowSMSProject)

	// GET: http://localhost:8144/smsprojects/remove/1
	app.Get("/sms_projects/remove/{id:string}", routes.RemoveSMSProject)


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

	// GET: http://localhost:8144/sms_issueAffectedDevice/createSMSIssueAffectedDevice/1
	app.Get("/sms_issueAffectedDevice/createSMSIssueAffectedDevice/{id:string}", routes.CreateSMSIssueAffectedDevice)

	// POST: http://localhost:8144/sms_issueAffectedDevice/addSMSIssueAffectedDevice
	app.Post("/sms_issueAffectedDevice/addSMSIssueAffectedDevice", routes.AddSMSIssueAffectedDevice)

	// GET: http://localhost:8144/sms_issueAffectedDevice/remove/1
	app.Get("/sms_issueAffectedDevice/remove/{id:string}", routes.RemoveSMSIssueAffectedDevice)

	// POST: http://localhost:8144/sms_solutions/addSMSSolution
	app.Post("/sms_solutions/addSMSSolution", routes.AddSMSSolution)

	// GET: http://localhost:8144/sms_solutions/show/1
	app.Get("/sms_solutions/show/{id:string}", routes.ShowSMSSolution)

	// GET: http://localhost:8144/sms_solutions/createSMSSolution/1
	app.Get("/sms_solutions/createSMSSolution/{id:string}", routes.CreateSMSSolution)

	// Application started. Press CTRL+C to shut down.
	app.Run(utils.Addr)
}
