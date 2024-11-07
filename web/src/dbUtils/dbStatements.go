/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package dbUtils

// Project
var SELECT_projects = `SELECT * FROM projects`
var SELECT_projectInfo = `SELECT project_id, name, uploads, date FROM projects WHERE project_id = ?;`
var INSERT_newProject = `INSERT INTO projects (name, uploads, date) VALUES (?,?,?);`
var DELETE_project = `DELETE FROM projects WHERE project_id = ?;`
var UPDATE_projectUploads = `UPDATE projects SET uploads = ? WHERE project_id = ?;`

// Firmware
var SELECT_firmware = `SELECT firmware.firmware_id, firmware.name, firmware.version, firmware.binwalkOutput, firmware.sizeInBytes, firmware.project_id, firmware.created, projects.name FROM firmware JOIN projects ON firmware.project_id = projects.project_id;`
var SELECT_firmwareInfo = `SELECT firmware_id, name, version, binwalkOutput, sizeInBytes, project_id, created FROM firmware WHERE firmware_id = ?;`
var SELECT_firmwareForProject = `SELECT * FROM firmware WHERE project_id = ?`
var INSERT_newFirmware = `INSERT INTO firmware (name, version, sizeInBytes, project_id, created) VALUES (?,?,?,?,?);`
var DELETE_firmware = `DELETE FROM firmware WHERE firmware_id = ?;`

// relevantApps
var SELECT_relevantApps = `SELECT relevantApps.relevantApps_id, relevantApps.name, relevantApps.path, relevantApps.extPort, relevantApps.extProtocoll, relevantApps.intInterface, relevantApps.firmware_id, firmware.name FROM relevantApps JOIN firmware ON relevantApps.firmware_id = firmware.firmware_id;`
var SELECT_relevantAppInfo = `SELECT * FROM relevantApps WHERE relevantApps_id = ?;`
var SELECT_relevantAppsForFirmware = `SELECT * FROM relevantApps WHERE firmware_id = ?`
var INSERT_newrelevantApps = `INSERT INTO relevantApps (name, path, extPort, extProtocoll, intInterface, firmware_id) VALUES (?,?,?,?,?,?);`
var DELETE_relevantApps = `DELETE FROM relevantApps WHERE relevantApps_id = ?;`
var SELECT_relevantAppByPath = `SELECT relevantApps_id FROM relevantApps WHERE relevantApps.path = ? AND relevantApps.firmware_id = ?;`
var SELECT_relevantAppByName = `SELECT relevantApps_id FROM relevantApps WHERE relevantApps.name = ? AND relevantApps.firmware_id = ?;`
var UPDATE_relevantAppmoduleDefault = `UPDATE relevantApps SET relevantApps.moduleDefault = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATE_relevantAppmoduleInitSystem = `UPDATE relevantApps SET relevantApps.moduleInitSystem = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATE_relevantAppmoduleFileContent = `UPDATE relevantApps SET relevantApps.moduleFileContent = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATE_relevantAppmoduleBash = `UPDATE relevantApps SET relevantApps.moduleBash = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATE_relevantAppmoduleCronJob = `UPDATE relevantApps SET relevantApps.moduleCronJob = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATE_relevantAppmoduleProcesses = `UPDATE relevantApps SET relevantApps.moduleProcesses = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATE_relevantAppmoduleInterfaces = `UPDATE relevantApps SET relevantApps.moduleInterfaces = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATE_relevantAppmoduleSystemControls = `UPDATE relevantApps SET relevantApps.moduleSystemControls = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATE_relevantAppmoduleFileSystem = `UPDATE relevantApps SET relevantApps.moduleFileSystem = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATE_relevantAppmodulePortscanner = `UPDATE relevantApps SET relevantApps.modulePortscanner = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATE_relevantAppmoduleProtocolls = `UPDATE relevantApps SET relevantApps.moduleProtocolls = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATE_relevantAppmoduleNetInterfaces = `UPDATE relevantApps SET relevantApps.moduleNetInterfaces = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATE_relevantAppmoduleFileSystemInterfaces = `UPDATE relevantApps SET relevantApps.moduleFileSystemInterfaces = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATE_relevantAppmoduleFileHandles = `UPDATE relevantApps SET relevantApps.moduleFileHandles = ? WHERE relevantApps.relevantApps_id = ?;`

var UPDATEWITHINTERFACE_relevantAppmoduleDefault = `UPDATE relevantApps SET relevantApps.moduleDefault = ?, relevantApps.extPort = ?, relevantApps.extProtocoll = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATEWITHINTERFACE_relevantAppmoduleInitSystem = `UPDATE relevantApps SET relevantApps.moduleInitSystem = ?, relevantApps.extPort = ?, relevantApps.extProtocoll = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATEWITHINTERFACE_relevantAppmoduleFileContent = `UPDATE relevantApps SET relevantApps.moduleFileContent = ?, relevantApps.extPort = ?, relevantApps.extProtocoll = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATEWITHINTERFACE_relevantAppmoduleBash = `UPDATE relevantApps SET relevantApps.moduleBash = ?, relevantApps.extPort = ?, relevantApps.extProtocoll = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATEWITHINTERFACE_relevantAppmoduleCronJob = `UPDATE relevantApps SET relevantApps.moduleCronJob = ?, relevantApps.extPort = ?, relevantApps.extProtocoll = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATEWITHINTERFACE_relevantAppmoduleProcesses = `UPDATE relevantApps SET relevantApps.moduleProcesses = ?, relevantApps.extPort = ?, relevantApps.extProtocoll = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATEWITHINTERFACE_relevantAppmoduleInterfaces = `UPDATE relevantApps SET relevantApps.moduleInterfaces = ?, relevantApps.extPort = ?, relevantApps.extProtocoll = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATEWITHINTERFACE_relevantAppmoduleSystemControls = `UPDATE relevantApps SET relevantApps.moduleSystemControls = ?, relevantApps.extPort = ?, relevantApps.extProtocoll = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATEWITHINTERFACE_relevantAppmoduleFileSystem = `UPDATE relevantApps SET relevantApps.moduleFileSystem = ?, relevantApps.extPort = ?, relevantApps.extProtocoll = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATEWITHINTERFACE_relevantAppmodulePortscanner = `UPDATE relevantApps SET relevantApps.modulePortscanner = ?, relevantApps.extPort = ?, relevantApps.extProtocoll = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATEWITHINTERFACE_relevantAppmoduleProtocolls = `UPDATE relevantApps SET relevantApps.moduleProtocolls = ?, relevantApps.extPort = ?, relevantApps.extProtocoll = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATEWITHINTERFACE_relevantAppmoduleNetInterfaces = `UPDATE relevantApps SET relevantApps.moduleNetInterfaces = ?, relevantApps.extPort = ?, relevantApps.extProtocoll = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATEWITHINTERFACE_relevantAppmoduleFileSystemInterfaces = `UPDATE relevantApps SET relevantApps.moduleFileSystemInterfaces = ?, relevantApps.extPort = ?, relevantApps.extProtocoll = ? WHERE relevantApps.relevantApps_id = ?;`
var UPDATEWITHINTERFACE_relevantAppmoduleFileHandles = `UPDATE relevantApps SET relevantApps.moduleFileHandles = ?, relevantApps.extPort = ?, relevantApps.extProtocoll = ? WHERE relevantApps.relevantApps_id = ?;`

// Results
var SELECT_results = `SELECT testResult.testResult_id, testResult.moduleName, testResult.result, testResult.created, testResult.firmware_id, firmware.name FROM testResult JOIN firmware ON testResult.firmware_id = firmware.firmware_id;`
var SELECT_resultInfo = `SELECT testResult.testResult_id, testResult.moduleName, testResult.result, testResult.created, testResult.firmware_id, firmware.name FROM testResult JOIN firmware ON testResult.firmware_id = firmware.firmware_id WHERE testResult.testResult_id = ?;`
var SELECT_resultsForFirmware = `SELECT testResult.testResult_id, testResult.moduleName, testResult.created, testResult.firmware_id FROM testResult WHERE firmware_id = ?`
var INSERT_newresults = `INSERT INTO testResult (moduleName, result, created, firmware_id) VALUES (?,?,?,?);`
var DELETE_result = `DELETE FROM testResult WHERE testResult_id = ?;`

// AppContent
var SELECT_appContent = `SELECT * FROM appContent WHERE appContent.appContent_id = ?;`
var SELECT_appContentForRelevantApp = `SELECT appContent.appContent_id, appContent.contentPathList, appContent.binwalkOutput, appContent.readelfOutput, appContent.lddOutput, appContent.straceOutput, appContent.relevantApps_path FROM appContent WHERE relevantApps_path = ?`
var INSERT_newappContent = `INSERT INTO appContent (contentPathList, binwalkOutput, readelfOutput, lddOutput, straceOutput, relevantApps_path) VALUES (?,?,?,?,?,?);`
var DELETE_appContent = `DELETE FROM appContent WHERE appContent_id = ?;`
var DELETE_appContentByRelevantAppPath = `DELETE FROM appContent WHERE relevantApps_path = ?;`
var SELECT_appContentByPath = `SELECT * FROM appContent WHERE appContent.relevantApps_path = ?;`
var UPDATE_appContentbinwalk = `UPDATE appContent SET appContent.binwalkOutput = ? WHERE appContent.appContent_id = ?;`
var UPDATE_appContentreadelf = `UPDATE appContent SET appContent.readelfOutput = ? WHERE appContent.appContent_id = ?;`
var UPDATE_appContentldd = `UPDATE appContent SET appContent.lddOutput = ? WHERE appContent.appContent_id = ?;`
var UPDATE_appContentstrace = `UPDATE appContent SET appContent.straceOutput = ? WHERE appContent.appContent_id = ?;`

// BinaryAnalysis
var SELECT_binaryAnalysis = `SELECT * FROM binaryAnalysis WHERE binaryAnalysis.binaryAnalysis_id = ?;`
var SELECT_binaryAnalysisForRelevantApp = `SELECT binaryAnalysis.binaryAnalysis_id, binaryAnalysis.toolOutput, binaryAnalysis.analysisTool_id, analysisTool.name, binaryAnalysis.relevantApps_id FROM binaryAnalysis JOIN analysisTool ON binaryAnalysis.analysisTool_id = analysisTool.analysisTool_id WHERE binaryAnalysis.relevantApps_id = ?`
var SELECT_binaryAnalysisForRelevantAppAndTool = `SELECT binaryAnalysis.binaryAnalysis_id, binaryAnalysis.toolOutput, binaryAnalysis.analysisTool_id, analysisTool.name, binaryAnalysis.relevantApps_id FROM binaryAnalysis JOIN analysisTool ON binaryAnalysis.analysisTool_id = analysisTool.analysisTool_id WHERE binaryAnalysis.relevantApps_id = ? AND binaryAnalysis.analysisTool_id = ?`
var INSERT_newbinaryAnalysis = `INSERT INTO binaryAnalysis (toolOutput, analysisTool_id, relevantApps_id) VALUES (?,?,?);`
var DELETE_binaryAnalysis = `DELETE FROM binaryAnalysis WHERE binaryAnalysis_id = ?;`
var DELETE_binaryAnalysisByRelevantApp = `DELETE FROM binaryAnalysis WHERE relevantApps_id = ?;`
var UPDATE_binaryAnalysis = `UPDATE binaryAnalysis SET binaryAnalysis.toolOutput = ? WHERE binaryAnalysis.binaryAnalysis_id = ?;`

// AnalysisTool
var SELECT_analysisTool = `SELECT analysisTool.analysisTool_id, analysisTool.name, analysisTool.executionString from analysisTool;`
var SELECT_analysisToolInfo = `SELECT analysisTool_id, name, executionString FROM analysisTool WHERE analysisTool_id = ?;`
var INSERT_newAnalysisTool = `INSERT INTO analysisTool (name, executionString) VALUES (?,?);`
var DELETE_analysisTool = `DELETE FROM analysisTool WHERE analysisTool_id = ?;`


/**
 * Security Management System
 * Created:   29.09.2024
 *
 * (C)
 **/


// SMS Project
var SELECT_sms_projects = `SELECT sms_project.project_id, sms_project.name, sms_project.customer, sms_projecttype.type, sms_project.reference, sms_project.date, sms_project.active FROM sms_project LEFT JOIN sms_projecttype ON sms_project.projecttype_id = sms_projecttype.projecttype_id `
var SELECT_sms_projectInfo = `SELECT sms_project.project_id, sms_project.name, sms_project.customer, sms_projecttype.type, sms_project.reference, sms_project.date, sms_project.active FROM sms_project LEFT JOIN sms_projecttype ON sms_project.projecttype_id = sms_projecttype.projecttype_id WHERE project_id = ?;`
var INSERT_sms_newProject = `INSERT INTO sms_project (name, customer, projecttype_id, reference, date, active) VALUES (?,?,?,?,?,?);`
var DELETE_sms_project = `DELETE FROM sms_project WHERE project_id = ?;`
var UPDATE_sms_projectActive = `UPDATE sms_project SET active = ? WHERE project_id = ?;`
var SELECT_sms_projectTypes = `SELECT sms_projecttype.projecttype_id, sms_projecttype.type FROM sms_projecttype;`

// SMS System
var SELECT_sms_systems = `SELECT sms_system.system_id, sms_system.version, sms_system.date, sms_systemtype.type FROM sms_system LEFT JOIN sms_systemtype ON sms_system.systemtype_id = sms_systemtype.systemtype_id `
var SELECT_sms_systemInfo = `SELECT sms_system.system_id, sms_system.version, sms_system.date, sms_systemtype.type FROM sms_system LEFT JOIN sms_systemtype ON sms_system.systemtype_id = sms_systemtype.systemtype_id WHERE system_id = ?;`
var INSERT_sms_newSystem = `INSERT INTO sms_system (systemtype_id, version, date) VALUES (?,?,?);`
var DELETE_sms_system = `DELETE FROM sms_system WHERE system_id = ?;`
var SELECT_sms_systemTypes = `SELECT sms_systemtype.systemtype_id, sms_systemtype.type FROM sms_systemtype;`

// SMS Device
var SELECT_sms_devices = `SELECT sms_device.device_id, sms_device.version, sms_device.date, sms_devicetype.type FROM sms_device LEFT JOIN sms_devicetype ON sms_device.devicetype_id = sms_devicetype.devicetype_id `
var SELECT_sms_deviceInfo = `SELECT sms_device.device_id, sms_device.version, sms_device.date, sms_devicetype.type FROM sms_device LEFT JOIN sms_devicetype ON sms_device.devicetype_id = sms_devicetype.devicetype_id WHERE device_id = ?;`
var INSERT_sms_newDevice = `INSERT INTO sms_device (devicetype_id, version, date) VALUES (?,?,?);`
var DELETE_sms_device = `DELETE FROM sms_device WHERE device_id = ?;`
var SELECT_sms_deviceTypes = `SELECT sms_devicetype.devicetype_id, sms_devicetype.type FROM sms_devicetype;`

// SMS DeviceInstance
var SELECT_sms_deviceInstances = `SELECT sms_deviceInstance.deviceInstance_id, sms_deviceInstance.project_id, sms_deviceInstance.device_id, sms_deviceInstance.serialnumber, sms_deviceInstance.provisioner, sms_deviceInstance.configuration, sms_deviceInstance.date, sms_project.name, sms_device.devicetype_id, sms_device.version, sms_devicetype.type FROM sms_deviceInstance LEFT JOIN sms_project ON sms_deviceInstance.project_id = sms_project.project_id LEFT JOIN sms_device ON sms_deviceInstance.device_id = sms_device.device_id LEFT JOIN sms_devicetype ON sms_device.devicetype_id = sms_devicetype.devicetype_id ; `
var SELECT_sms_deviceInstanceInfo = `SELECT sms_deviceInstance.deviceInstance_id, sms_deviceInstance.project_id, sms_deviceInstance.device_id, sms_deviceInstance.serialnumber, sms_deviceInstance.provisioner, sms_deviceInstance.configuration, sms_deviceInstance.date, sms_project.name, sms_device.devicetype_id, sms_device.version, sms_devicetype.type FROM sms_deviceInstance LEFT JOIN sms_project ON sms_deviceInstance.project_id = sms_project.project_id LEFT JOIN sms_device ON sms_deviceInstance.device_id = sms_device.device_id LEFT JOIN sms_devicetype ON sms_device.devicetype_id = sms_devicetype.devicetype_id WHERE deviceInstance_id = ?;`
var INSERT_sms_newDeviceInstance = `INSERT INTO sms_deviceInstance (project_id, device_id, serialnumber, provisioner, configuration, date) VALUES (?,?,?,?,?,?);`
var DELETE_sms_deviceInstance = `DELETE FROM sms_deviceInstance WHERE deviceInstance_id = ?;`
var SELECT_sms_deviceInstancesForProject = `SELECT sms_deviceInstance.deviceInstance_id, sms_deviceInstance.project_id, sms_deviceInstance.device_id, sms_deviceInstance.serialnumber, sms_deviceInstance.provisioner, sms_deviceInstance.configuration, sms_deviceInstance.date, sms_project.name, sms_device.devicetype_id, sms_device.version, sms_devicetype.type FROM sms_deviceInstance LEFT JOIN sms_project ON sms_deviceInstance.project_id = sms_project.project_id LEFT JOIN sms_device ON sms_deviceInstance.device_id = sms_device.device_id LEFT JOIN sms_devicetype ON sms_device.devicetype_id = sms_devicetype.devicetype_id WHERE sms_deviceInstance.project_id = ?; `

// SMS_UpdateHistory
var SELECT_sms_updatehistoriesForDevice = `SELECT sms_updateHistory.updateHistory_id, sms_updateHistory.deviceInstance_id, sms_updateHistory.user, sms_updateHistory.updateType, sms_updateHistory.date, sms_updateHistory.description FROM sms_updateHistory WHERE deviceInstance_id = ? `
var INSERT_sms_newUpdateHistory = `INSERT INTO sms_updateHistory (deviceInstance_id, user, updateType, date, description) VALUES (?,?,?,?,?);`
var SELECT_sms_UpdateHistoryInfo = `SELECT sms_updateHistory.updateHistory_id, sms_updateHistory.deviceInstance_id, sms_updateHistory.user, sms_updateHistory.updateType, sms_updateHistory.date, sms_updateHistory.description FROM sms_updateHistory WHERE updateHistory_id = ?`

// SMS_Issue
var SELECT_sms_issues = `SELECT sms_issue.issue_id, sms_issue.name, sms_issue.date, sms_issue.issueType, sms_issue.reference, sms_issue.criticality, sms_issue.cve, sms_issue.description FROM sms_issue; `
var SELECT_sms_issueInfo = `SELECT sms_issue.issue_id, sms_issue.name, sms_issue.date, sms_issue.issueType, sms_issue.reference, sms_issue.criticality, sms_issue.cve, sms_issue.description FROM sms_issue WHERE issue_id = ?;`
var INSERT_sms_newIssue = `INSERT INTO sms_issue (name, date, issueType, reference, criticality, cve, description) VALUES (?,?,?,?,?,?,?);`
var DELETE_sms_issue = `DELETE FROM sms_issue WHERE issue_id = ?;`

// SMS IssueAffectedDevice
var INSERT_sms_newIssueAffectedDevice = `INSERT INTO sms_issueAffectedDevice (device_id, issue_id, additionalInfo, confirmed) VALUES (?,?,?,?);`
var DELETE_sms_IssueAffectedDevice = `DELETE FROM sms_issueAffectedDevice WHERE device_id = ? AND issue_id = ?;`
var SELECT_sms_IssueAffectedDevicesForIssueID = `SELECT sms_issueAffectedDevice.device_id, sms_issueAffectedDevice.issue_id, sms_issueAffectedDevice.additionalInfo, sms_issueAffectedDevice.confirmed, sms_devicetype.type, sms_device.version FROM sms_issueAffectedDevice LEFT JOIN sms_device ON sms_issueAffectedDevice.device_id = sms_device.device_id LEFT JOIN sms_devicetype ON sms_device.devicetype_id = sms_devicetype.devicetype_id WHERE sms_issueAffectedDevice.issue_id = ?; `
var SELECT_sms_IssuesForDevice = `SELECT sms_issueAffectedDevice.device_id, sms_issueAffectedDevice.issue_id, sms_issueAffectedDevice.additionalInfo, sms_issueAffectedDevice.confirmed, sms_issue.name FROM sms_issueAffectedDevice LEFT JOIN sms_issue ON sms_issueAffectedDevice.issue_id = sms_issue.issue_id WHERE sms_issueAffectedDevice.device_id = ?; `
var SELECT_sms_affectedDeviceInstancesAndProjects = `SELECT deviceInstance_id, type, project_id, version FROM sms_deviceInstance sdi LEFT JOIN sms_device sd1 ON sdi.device_id = sd1.device_id LEFT JOIN sms_devicetype sd2 ON sd1.devicetype_id = sd2.devicetype_id LEFT JOIN sms_issueAffectedDevice siad ON sdi.device_id = siad.device_id WHERE siad.issue_id = ? UNION ALL Select deviceInstance_id, type, project_id, version FROM sms_deviceInstance sdi LEFT JOIN sms_device sd1 ON sdi.device_id = sd1.device_id LEFT JOIN sms_devicetype sd2 ON sd1.devicetype_id = sd2.devicetype_id RIGHT JOIN sms_issueAffectedDevice siad ON sdi.device_id = siad.device_id WHERE sdi.device_id IS NULL AND siad.issue_id = ?; `

// SMS Solution
var INSERT_sms_newSolution = `INSERT INTO sms_solution (issue_id, devicetype_id, date, name, description, reference) VALUES (?,?,?,?,?,?);`
var DELETE_sms_Solution = `DELETE FROM sms_solution WHERE solution_id = ?;`
var SELECT_sms_solutionsForIssue = `SELECT sms_solution.solution_id, sms_solution.issue_id, sms_solution.devicetype_id, sms_solution.date, sms_solution.name, sms_solution.description, sms_solution.reference, sms_devicetype.type FROM sms_solution LEFT JOIN sms_devicetype ON sms_solution.devicetype_id = sms_devicetype.devicetype_id WHERE issue_id = ?;`
var SELECT_sms_solutionInfo = `SELECT sms_solution.solution_id, sms_solution.issue_id, sms_solution.devicetype_id, sms_solution.date, sms_solution.name, sms_solution.description, sms_solution.reference, sms_devicetype.type FROM sms_solution LEFT JOIN sms_devicetype ON sms_solution.devicetype_id = sms_devicetype.devicetype_id WHERE solution_id = ?;`

// SMS Artefact
var SELECT_sms_artefact = `SELECT sms_artefact.artefact_id, sms_artefact.artefacttype_id, sms_artefact.name, sms_artefact.version, sms_artefacttype.artefactType FROM sms_artefact LEFT JOIN sms_artefacttype ON sms_artefact.artefacttype_id = sms_artefacttype.artefacttype_id `
var SELECT_sms_artefactInfo = `SELECT sms_artefact.artefact_id, sms_artefact.artefacttype_id, sms_artefact.name, sms_artefact.version, sms_artefacttype.artefactType FROM sms_artefact LEFT JOIN sms_artefacttype ON sms_artefact.artefacttype_id = sms_artefacttype.artefacttype_id WHERE artefact_id = ?;`
var INSERT_sms_newArtefact = `INSERT INTO sms_artefact (artefacttype_id, name, version) VALUES (?,?,?);`
var DELETE_sms_artefact = `DELETE FROM sms_artefact WHERE artefact_id = ?;`
var SELECT_sms_artefactTypes = `SELECT sms_artefacttype.artefacttype_id, sms_artefacttype.artefactType FROM sms_artefacttype;`

// SMS_UpdateHistory
var SELECT_sms_releaseNoteForDevice = `SELECT sms_releasenote.releasenote_id, sms_releasenote.device_id, sms_releasenote.type, sms_releasenote.date, sms_releasenote.details FROM sms_releasenote WHERE device_id = ? `
var INSERT_sms_newReleaseNote = `INSERT INTO sms_releasenote (device_id, type, date, details) VALUES (?,?,?,?);`
var SELECT_sms_ReleaseNoteInfo = `SELECT sms_releasenote.releasenote_id, sms_releasenote.device_id, sms_releasenote.type, sms_releasenote.date, sms_releasenote.details FROM sms_releasenote WHERE releasenote_id = ?`

// SMS Software
var SELECT_sms_softwares = `SELECT sms_software.software_id, sms_software.softwaretype_id, sms_software.version, sms_software.date, sms_softwaretype.typeName, sms_software.license, sms_software.thirdParty, sms_software.releaseNote FROM sms_software LEFT JOIN sms_softwaretype ON sms_software.softwaretype_id = sms_softwaretype.softwaretype_id `
var SELECT_sms_softwareInfo = `SELECT sms_software.software_id, sms_software.softwaretype_id, sms_software.version, sms_software.date, sms_softwaretype.typeName, sms_software.license, sms_software.thirdParty, sms_software.releaseNote FROM sms_software LEFT JOIN sms_softwaretype ON sms_software.softwaretype_id = sms_softwaretype.softwaretype_id WHERE software_id = ?;`
var INSERT_sms_newSoftware = `INSERT INTO sms_software (softwaretype_id, version, date, license, thirdParty, releaseNote) VALUES (?,?,?,?,?,?);`
var DELETE_sms_software = `DELETE FROM sms_software WHERE software_id = ?;`
var SELECT_sms_softwareTypes = `SELECT sms_softwaretype.softwaretype_id, sms_softwaretype.typeName FROM sms_softwaretype;`

// SMS Component
var SELECT_sms_components = `SELECT sms_component.component_id, sms_component.name, sms_component.componentType, sms_component.version, sms_component.date, sms_component.license, sms_component.thirdParty, sms_component.releaseNote FROM sms_component;`
var SELECT_sms_componentInfo = `SELECT sms_component.component_id, sms_component.name, sms_component.componentType, sms_component.version, sms_component.date, sms_component.license, sms_component.thirdParty, sms_component.releaseNote FROM sms_component WHERE component_id = ?;`
var INSERT_sms_newComponent = `INSERT INTO sms_component (name, componentType, version, date, license, thirdParty, releaseNote) VALUES (?,?,?,?,?,?,?);`
var DELETE_sms_component = `DELETE FROM sms_component WHERE component_id = ?;`

// SMS ComponentPartOfSoftware
var INSERT_sms_newComponentPartOfSoftware = `INSERT INTO sms_componentPartOfSoftware (software_id, component_id, additionalInfo) VALUES (?,?,?);`
var DELETE_sms_ComponentPartOfSoftware = `DELETE FROM sms_componentPartOfSoftware WHERE software_id = ? AND component_id = ?;`
var SELECT_sms_ComponentPartOfSoftwareForSoftware = `SELECT sms_componentPartOfSoftware.software_id, sms_componentPartOfSoftware.component_id, sms_componentPartOfSoftware.additionalInfo, sms_component.name, sms_component.version FROM sms_componentPartOfSoftware LEFT JOIN sms_component ON sms_componentPartOfSoftware.component_id = sms_component.component_id WHERE sms_componentPartOfSoftware.software_id = ?; `
var SELECT_sms_ComponentPartOfSoftwareForComponent = `SELECT sms_componentPartOfSoftware.software_id, sms_componentPartOfSoftware.component_id, sms_componentPartOfSoftware.additionalInfo, sms_softwaretype.typeName, sms_software.version FROM sms_componentPartOfSoftware LEFT JOIN sms_software ON sms_componentPartOfSoftware.software_id = sms_software.software_id LEFT JOIN sms_softwaretype ON sms_software.softwaretype_id = sms_softwaretype.softwaretype_id WHERE sms_componentPartOfSoftware.component_id = ?; `

// SMS SoftwarePartOfDevice
var INSERT_sms_newSoftwarePartOfDevice = `INSERT INTO sms_softwarePartOfDevice (device_id, software_id, additionalInfo) VALUES (?,?,?);`
var DELETE_sms_SoftwarePartOfDevice = `DELETE FROM sms_softwarePartOfDevice WHERE device_id = ? AND software_id = ?;`
var SELECT_sms_SoftwarePartOfDeviceForDevice = `SELECT sms_softwarePartOfDevice.device_id, sms_softwarePartOfDevice.software_id, sms_softwarePartOfDevice.additionalInfo, sms_softwaretype.typeName, sms_software.version FROM sms_softwarePartOfDevice LEFT JOIN sms_software ON sms_softwarePartOfDevice.software_id = sms_software.software_id LEFT JOIN sms_softwaretype ON sms_software.softwaretype_id = sms_softwaretype.softwaretype_id WHERE sms_softwarePartOfDevice.device_id = ?; `
var SELECT_sms_SoftwarePartOfDeviceForSoftware = `SELECT sms_softwarePartOfDevice.device_id, sms_softwarePartOfDevice.software_id, sms_softwarePartOfDevice.additionalInfo, sms_devicetype.type, sms_device.version FROM sms_softwarePartOfDevice LEFT JOIN sms_device ON sms_softwarePartOfDevice.device_id = sms_device.device_id LEFT JOIN sms_devicetype ON sms_device.devicetype_id = sms_devicetype.devicetype_id WHERE sms_softwarePartOfDevice.software_id = ?; `

// SMS DevicePartOfSystem
var INSERT_sms_newDevicePartOfSystem = `INSERT INTO sms_devicePartOfSystem (system_id, device_id, additionalInfo) VALUES (?,?,?);`
var DELETE_sms_DevicePartOfSystem = `DELETE FROM sms_devicePartOfSystem WHERE system_id = ? AND device_id = ?;`
var SELECT_sms_DevicePartOfSystemForSystem = `SELECT sms_devicePartOfSystem.system_id, sms_devicePartOfSystem.device_id, sms_devicePartOfSystem.additionalInfo, sms_devicetype.type, sms_device.version FROM sms_devicePartOfSystem LEFT JOIN sms_device ON sms_devicePartOfSystem.device_id = sms_device.device_id LEFT JOIN sms_devicetype ON sms_device.devicetype_id = sms_devicetype.devicetype_id WHERE sms_devicePartOfSystem.system_id = ?; `
var SELECT_sms_DevicePartOfSystemForDevice = `SELECT sms_devicePartOfSystem.system_id, sms_devicePartOfSystem.device_id, sms_devicePartOfSystem.additionalInfo, sms_systemtype.type, sms_system.version FROM sms_devicePartOfSystem LEFT JOIN sms_system ON sms_devicePartOfSystem.system_id = sms_system.system_id LEFT JOIN sms_systemtype ON sms_system.systemtype_id = sms_systemtype.systemtype_id WHERE sms_devicePartOfSystem.device_id = ?; `

// SMS projectBOM
var INSERT_sms_newProjectBOM = `INSERT INTO sms_projectBOM (project_id, system_id, orderNumber, additionalInfo) VALUES (?,?,?,?);`
var DELETE_sms_ProjectBOM = `DELETE FROM sms_projectBOM WHERE projectBOM_id = ?;`
var SELECT_sms_ProjectBOMForProject = `SELECT sms_projectBOM.projectBOM_id, sms_projectBOM.project_id, sms_projectBOM.system_id, sms_projectBOM.orderNumber, sms_projectBOM.additionalInfo, sms_systemtype.type, sms_system.version FROM sms_projectBOM LEFT JOIN sms_system ON sms_projectBOM.system_id = sms_system.system_id LEFT JOIN sms_systemtype ON sms_system.systemtype_id = sms_systemtype.systemtype_id WHERE sms_projectBOM.project_id = ?; `
var SELECT_sms_ProjectBOMForSystem = `SELECT sms_projectBOM.projectBOM_id, sms_projectBOM.project_id, sms_projectBOM.system_id, sms_projectBOM.orderNumber, sms_projectBOM.additionalInfo, sms_project.name, sms_project.customer FROM sms_projectBOM LEFT JOIN sms_project ON sms_projectBOM.project_id = sms_project.project_id WHERE sms_projectBOM.system_id = ?; `
