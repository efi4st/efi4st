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
var SELECT_sms_issuesForProject = `SELECT DISTINCT di.device_id, dt.type AS device_name, d.version AS device_version, combined_issues.issue_id, combined_issues.issue_name, combined_issues.criticality, combined_issues.inherit FROM sms_deviceInstance di JOIN sms_device d ON di.device_id = d.device_id JOIN sms_devicetype dt ON d.devicetype_id = dt.devicetype_id JOIN ( SELECT did.device_id, did.issue_id, i.name AS issue_name, i.criticality, false AS inherit FROM sms_issueAffectedDevice did LEFT JOIN sms_issue i ON did.issue_id = i.issue_id UNION ALL SELECT spd.device_id, dis.issue_id, i.name AS issue_name, i.criticality, true AS inherit FROM sms_softwarePartOfDevice spd LEFT JOIN sms_issueAffectedSoftware dis ON spd.software_id = dis.software_id LEFT JOIN sms_issue i ON dis.issue_id = i.issue_id UNION ALL SELECT spd.device_id, dic.issue_id, i.name AS issue_name, i.criticality, true AS inherit FROM sms_softwarePartOfDevice spd LEFT JOIN sms_componentPartOfSoftware scps ON spd.software_id = scps.software_id LEFT JOIN sms_issueAffectedComponent dic ON scps.component_id = dic.component_id LEFT JOIN sms_issue i ON dic.issue_id = i.issue_id UNION ALL SELECT apd.device_id, dia.issue_id, i.name AS issue_name, i.criticality, true AS inherit FROM sms_artefactPartOfDevice apd LEFT JOIN sms_issueAffectedArtefact dia ON apd.artefact_id = dia.artefact_id LEFT JOIN sms_issue i ON dia.issue_id = i.issue_id ) AS combined_issues ON di.device_id = combined_issues.device_id WHERE di.project_id = ? AND combined_issues.issue_id IS NOT NULL;`

// SMS System
var SELECT_sms_systems = `SELECT sms_system.system_id, sms_system.version, sms_system.date, sms_systemtype.type FROM sms_system LEFT JOIN sms_systemtype ON sms_system.systemtype_id = sms_systemtype.systemtype_id `
var SELECT_sms_systemInfo = `SELECT sms_system.system_id, sms_system.version, sms_system.date, sms_systemtype.type FROM sms_system LEFT JOIN sms_systemtype ON sms_system.systemtype_id = sms_systemtype.systemtype_id WHERE system_id = ?;`
var INSERT_sms_newSystem = `INSERT INTO sms_system (systemtype_id, version, date) VALUES (?,?,?);`
var DELETE_sms_system = `DELETE FROM sms_system WHERE system_id = ?;`
var SELECT_sms_systemTypes = `SELECT sms_systemtype.systemtype_id, sms_systemtype.type FROM sms_systemtype;`
var SELECT_sms_getIssuesForWholeSystem = `SELECT DISTINCT device_issues.device_id, device_issues.issue_id, device_issues.additionalInfo, device_issues.confirmed, device_issues.issue_name, '' AS device_version, device_issues.inherit FROM ( SELECT dps.device_id, did.issue_id, did.additionalInfo, did.confirmed, i.name AS issue_name, false AS inherit FROM sms_devicePartOfSystem dps LEFT JOIN sms_issueAffectedDevice did ON dps.device_id = did.device_id LEFT JOIN sms_issue i ON did.issue_id = i.issue_id WHERE dps.system_id = ? UNION ALL SELECT dps.device_id, dis.issue_id, dis.additionalInfo, dis.confirmed, i.name AS issue_name, true AS inherit FROM sms_devicePartOfSystem dps LEFT JOIN sms_softwarePartOfDevice spd ON dps.device_id = spd.device_id LEFT JOIN sms_issueAffectedSoftware dis ON spd.software_id = dis.software_id LEFT JOIN sms_issue i ON dis.issue_id = i.issue_id WHERE dps.system_id = ? UNION ALL SELECT dps.device_id, dic.issue_id, dic.additionalInfo, dic.confirmed, i.name AS issue_name, true AS inherit FROM sms_devicePartOfSystem dps LEFT JOIN sms_softwarePartOfDevice spd ON dps.device_id = spd.device_id LEFT JOIN sms_componentPartOfSoftware scps ON spd.software_id = scps.software_id LEFT JOIN sms_issueAffectedComponent dic ON scps.component_id = dic.component_id LEFT JOIN sms_issue i ON dic.issue_id = i.issue_id WHERE dps.system_id = ? UNION ALL SELECT dps.device_id, dia.issue_id, dia.additionalInfo, dia.confirmed, i.name AS issue_name, true AS inherit FROM sms_devicePartOfSystem dps LEFT JOIN sms_artefactPartOfDevice apd ON dps.device_id = apd.device_id LEFT JOIN sms_issueAffectedArtefact dia ON apd.artefact_id = dia.artefact_id LEFT JOIN sms_issue i ON dia.issue_id = i.issue_id WHERE dps.system_id = ?) AS device_issues WHERE device_issues.issue_id IS NOT NULL;`
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
var SELECT_sms_issuesForDeviceInstance = `SELECT DISTINCT device_issues.device_id, device_issues.issue_id, COALESCE(device_issues.additionalInfo, '') AS additionalInfo, device_issues.confirmed, device_issues.issue_name AS device_type, '' AS device_version, device_issues.inherit FROM (SELECT di.device_id, did.issue_id, COALESCE(did.additionalInfo, '') AS additionalInfo, did.confirmed, i.name AS issue_name, false AS inherit FROM sms_deviceInstance di LEFT JOIN sms_issueAffectedDevice did ON di.device_id = did.device_id LEFT JOIN sms_issue i ON did.issue_id = i.issue_id WHERE di.deviceInstance_id = ? UNION ALL SELECT di.device_id, dis.issue_id, COALESCE(dis.additionalInfo, '') AS additionalInfo, dis.confirmed, i.name AS issue_name, true AS inherit FROM sms_deviceInstance di LEFT JOIN sms_softwarePartOfDevice spd ON di.device_id = spd.device_id LEFT JOIN sms_issueAffectedSoftware dis ON spd.software_id = dis.software_id LEFT JOIN sms_issue i ON dis.issue_id = i.issue_id WHERE di.deviceInstance_id = ? UNION ALL SELECT di.device_id, dic.issue_id, COALESCE(dic.additionalInfo, '') AS additionalInfo, dic.confirmed, i.name AS issue_name, true AS inherit FROM sms_deviceInstance di LEFT JOIN sms_softwarePartOfDevice spd ON di.device_id = spd.device_id LEFT JOIN sms_componentPartOfSoftware scps ON spd.software_id = scps.software_id LEFT JOIN sms_issueAffectedComponent dic ON scps.component_id = dic.component_id LEFT JOIN sms_issue i ON dic.issue_id = i.issue_id WHERE di.deviceInstance_id = ? UNION ALL SELECT di.device_id, dia.issue_id, COALESCE(dia.additionalInfo, '') AS additionalInfo, dia.confirmed, i.name AS issue_name, true AS inherit FROM sms_deviceInstance di LEFT JOIN sms_artefactPartOfDevice apd ON di.device_id = apd.device_id LEFT JOIN sms_issueAffectedArtefact dia ON apd.artefact_id = dia.artefact_id LEFT JOIN sms_issue i ON dia.issue_id = i.issue_id WHERE di.deviceInstance_id = ?) AS device_issues WHERE device_issues.issue_id IS NOT NULL;`

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
var SELECT_sms_IssueAffectedDevicesForIssueIDWithInheritage = `SELECT DISTINCT device_id, issue_id, additionalInfo, confirmed, device_type, device_version, inherit FROM ( SELECT sms_device.device_id, sms_issueAffectedDevice.issue_id, sms_issueAffectedDevice.additionalInfo, sms_issueAffectedDevice.confirmed, sms_devicetype.type AS device_type, sms_device.version AS device_version, false AS inherit FROM sms_device LEFT JOIN sms_issueAffectedDevice ON sms_device.device_id = sms_issueAffectedDevice.device_id LEFT JOIN sms_devicetype ON sms_device.devicetype_id = sms_devicetype.devicetype_id WHERE sms_issueAffectedDevice.issue_id = ? UNION ALL SELECT sms_device.device_id, sas.issue_id, COALESCE(sas.additionalInfo, '') AS additionalInfo, sas.confirmed, sms_devicetype.type AS device_type, sms_device.version AS device_version, true AS inherit FROM sms_device LEFT JOIN sms_softwarePartOfDevice ON sms_device.device_id = sms_softwarePartOfDevice.device_id LEFT JOIN sms_software ON sms_softwarePartOfDevice.software_id = sms_software.software_id LEFT JOIN sms_issueAffectedSoftware sas ON sms_software.software_id = sas.software_id LEFT JOIN sms_devicetype ON sms_device.devicetype_id = sms_devicetype.devicetype_id WHERE sas.issue_id = ? UNION ALL SELECT sms_device.device_id, siac.issue_id, COALESCE(siac.additionalInfo, '') AS additionalInfo, siac.confirmed, sms_devicetype.type AS device_type, sms_device.version AS device_version, true AS inherit FROM sms_device LEFT JOIN sms_softwarePartOfDevice ON sms_device.device_id = sms_softwarePartOfDevice.device_id LEFT JOIN sms_software ON sms_softwarePartOfDevice.software_id = sms_software.software_id LEFT JOIN sms_componentPartOfSoftware scps ON sms_software.software_id = scps.software_id LEFT JOIN sms_issueAffectedComponent siac ON scps.component_id = siac.component_id LEFT JOIN sms_devicetype ON sms_device.devicetype_id = sms_devicetype.devicetype_id WHERE siac.issue_id = ? UNION ALL SELECT sms_device.device_id, sia.issue_id, COALESCE(sia.additionalInfo, '') AS additionalInfo, sia.confirmed, sms_devicetype.type AS device_type, sms_device.version AS device_version, true AS inherit FROM sms_device LEFT JOIN sms_artefactPartOfDevice apd ON sms_device.device_id = apd.device_id LEFT JOIN sms_artefact a ON apd.artefact_id = a.artefact_id LEFT JOIN sms_issueAffectedArtefact sia ON a.artefact_id = sia.artefact_id LEFT JOIN sms_devicetype ON sms_device.devicetype_id = sms_devicetype.devicetype_id WHERE sia.issue_id = ? ) AS combined_issues;`
var SELECT_sms_IssuesForDevice = `SELECT DISTINCT device_id, issue_id, additionalInfo, confirmed, issue_name, inherit FROM ( SELECT did.device_id, did.issue_id, did.additionalInfo, did.confirmed, i.name AS issue_name, false AS inherit FROM sms_issueAffectedDevice did LEFT JOIN sms_issue i ON did.issue_id = i.issue_id WHERE did.device_id = ? UNION ALL SELECT spd.device_id, dis.issue_id, dis.additionalInfo, dis.confirmed, i.name AS issue_name, true AS inherit FROM sms_softwarePartOfDevice spd LEFT JOIN sms_issueAffectedSoftware dis ON spd.software_id = dis.software_id LEFT JOIN sms_issue i ON dis.issue_id = i.issue_id WHERE spd.device_id = ? UNION ALL SELECT spd.device_id, dic.issue_id, dic.additionalInfo, dic.confirmed, i.name AS issue_name, true AS inherit FROM sms_softwarePartOfDevice spd LEFT JOIN sms_componentPartOfSoftware scps ON spd.software_id = scps.software_id LEFT JOIN sms_issueAffectedComponent dic ON scps.component_id = dic.component_id LEFT JOIN sms_issue i ON dic.issue_id = i.issue_id WHERE spd.device_id = ? UNION ALL SELECT apd.device_id, dia.issue_id, dia.additionalInfo, dia.confirmed, i.name AS issue_name, true AS inherit FROM sms_artefactPartOfDevice apd LEFT JOIN sms_issueAffectedArtefact dia ON apd.artefact_id = dia.artefact_id LEFT JOIN sms_issue i ON dia.issue_id = i.issue_id WHERE apd.device_id = ? ) AS combined_issues; `
// Komplexe Abfrage für alle betroffenen Instanzen, (rekursiv auch über die betroffenen Software, componenten, Artefakte...., liefert auch die Devices ohne Instanzen, die werden aber erstmal rausgefiltert
// Teil 1: Geräte mit Instanzen, die direkt durch ein Issue betroffen sind
// Teil 2: Geräte ohne Instanzen, die direkt durch ein Issue betroffen sind
// Teil 3: Geräteinstanzen, die über ihre Software-Komponenten von einem Issue betroffen sind
// Teil 4: Geräte ohne Instanzen, die eine betroffene Software haben
// Teil 5: Geräteinstanzen, die über ihre Software-Komponenten durch betroffene Komponenten betroffen sind
// Teil 6: Geräte ohne Instanzen, die über betroffene Komponenten innerhalb ihrer Software betroffen sind
// Teil 7: Geräteinstanzen, die durch betroffene Artefakte innerhalb des Geräts betroffen sind
// Teil 8: Geräte ohne Instanzen, die durch betroffene Artefakte innerhalb des Geräts betroffen sind
var SELECT_sms_affectedDeviceInstancesAndProjects = `SELECT DISTINCT deviceInstance_id, type, project_id, version FROM (SELECT sdi.deviceInstance_id, sd2.type, sdi.project_id, version FROM sms_deviceInstance sdi LEFT JOIN sms_device sd1 ON sdi.device_id = sd1.device_id LEFT JOIN sms_devicetype sd2 ON sd1.devicetype_id = sd2.devicetype_id LEFT JOIN sms_issueAffectedDevice siad ON sdi.device_id = siad.device_id WHERE siad.issue_id = ?
UNION ALL
SELECT NULL AS deviceInstance_id, sd2.type, NULL AS project_id, COALESCE(sd1.version, 'Unknown') AS version FROM sms_issueAffectedDevice siad LEFT JOIN sms_device sd1 ON siad.device_id = sd1.device_id LEFT JOIN sms_devicetype sd2 ON sd1.devicetype_id = sd2.devicetype_id WHERE siad.device_id NOT IN (SELECT device_id FROM sms_deviceInstance) AND siad.issue_id = ?
UNION ALL
SELECT sdi.deviceInstance_id, sd2.type, sdi.project_id, version FROM sms_deviceInstance sdi LEFT JOIN sms_device sd1 ON sdi.device_id = sd1.device_id LEFT JOIN sms_softwarePartOfDevice ssod ON sd1.device_id = ssod.device_id LEFT JOIN sms_issueAffectedSoftware sias ON ssod.software_id = sias.software_id LEFT JOIN sms_devicetype sd2 ON sd1.devicetype_id = sd2.devicetype_id WHERE sias.issue_id = ?
UNION ALL
SELECT NULL AS deviceInstance_id, sd2.type, NULL AS project_id, COALESCE(sd1.version, 'Unknown') AS version FROM sms_device sd1 LEFT JOIN sms_softwarePartOfDevice ssod ON sd1.device_id = ssod.device_id LEFT JOIN sms_issueAffectedSoftware sias ON ssod.software_id = sias.software_id LEFT JOIN sms_devicetype sd2 ON sd1.devicetype_id = sd2.devicetype_id WHERE sias.issue_id = ? AND sd1.device_id NOT IN (SELECT device_id FROM sms_deviceInstance) AND sias.software_id IS NOT NULL
UNION ALL
SELECT sdi.deviceInstance_id, sd2.type, sdi.project_id, version FROM sms_deviceInstance sdi LEFT JOIN sms_device sd1 ON sdi.device_id = sd1.device_id LEFT JOIN sms_softwarePartOfDevice ssod ON sd1.device_id = ssod.device_id LEFT JOIN sms_componentPartOfSoftware scps ON ssod.software_id = scps.software_id LEFT JOIN sms_issueAffectedComponent siac ON scps.component_id = siac.component_id LEFT JOIN sms_devicetype sd2 ON sd1.devicetype_id = sd2.devicetype_id WHERE siac.issue_id = ?
UNION ALL
SELECT NULL AS deviceInstance_id, sd2.type, NULL AS project_id, COALESCE(sd1.version, 'Unknown') AS version FROM sms_device sd1 LEFT JOIN sms_softwarePartOfDevice ssod ON sd1.device_id = ssod.device_id LEFT JOIN sms_componentPartOfSoftware scps ON ssod.software_id = scps.software_id LEFT JOIN sms_issueAffectedComponent siac ON scps.component_id = siac.component_id LEFT JOIN sms_devicetype sd2 ON sd1.devicetype_id = sd2.devicetype_id WHERE siac.issue_id = ? AND sd1.device_id NOT IN (SELECT device_id FROM sms_deviceInstance)
UNION ALL
SELECT sdi.deviceInstance_id, sd2.type, sdi.project_id, version FROM sms_deviceInstance sdi LEFT JOIN sms_device sd1 ON sdi.device_id = sd1.device_id LEFT JOIN sms_artefactPartOfDevice sapd ON sd1.device_id = sapd.device_id LEFT JOIN sms_issueAffectedArtefact siaa ON sapd.artefact_id = siaa.artefact_id LEFT JOIN sms_devicetype sd2 ON sd1.devicetype_id = sd2.devicetype_id WHERE siaa.issue_id = ?
UNION ALL
SELECT NULL AS deviceInstance_id, sd2.type, NULL AS project_id, COALESCE(sd1.version, 'Unknown') AS version FROM sms_device sd1 LEFT JOIN sms_artefactPartOfDevice sapd ON sd1.device_id = sapd.device_id LEFT JOIN sms_issueAffectedArtefact siaa ON sapd.artefact_id = siaa.artefact_id LEFT JOIN sms_devicetype sd2 ON sd1.devicetype_id = sd2.devicetype_id WHERE siaa.issue_id = ? AND sd1.device_id NOT IN (SELECT device_id FROM sms_deviceInstance)
)AS combined_result;`
// Statistics for above Query
var SELECT_sms_statisticsForaffectedDeviceInstancesAndProjects = `WITH combined_result AS (SELECT DISTINCT deviceInstance_id, type, project_id, version FROM 
(SELECT sdi.deviceInstance_id, sd2.type, sdi.project_id, version FROM sms_deviceInstance sdi LEFT JOIN sms_device sd1 ON sdi.device_id = sd1.device_id LEFT JOIN sms_devicetype sd2 ON sd1.devicetype_id = sd2.devicetype_id LEFT JOIN sms_issueAffectedDevice siad ON sdi.device_id = siad.device_id WHERE siad.issue_id = ?
UNION ALL
SELECT NULL AS deviceInstance_id, sd2.type, NULL AS project_id, COALESCE(sd1.version, 'Unknown') AS version FROM sms_issueAffectedDevice siad LEFT JOIN sms_device sd1 ON siad.device_id = sd1.device_id LEFT JOIN sms_devicetype sd2 ON sd1.devicetype_id = sd2.devicetype_id WHERE siad.device_id NOT IN (SELECT device_id FROM sms_deviceInstance) AND siad.issue_id = ?
UNION ALL
SELECT sdi.deviceInstance_id, sd2.type, sdi.project_id, version FROM sms_deviceInstance sdi LEFT JOIN sms_device sd1 ON sdi.device_id = sd1.device_id LEFT JOIN sms_softwarePartOfDevice ssod ON sd1.device_id = ssod.device_id LEFT JOIN sms_issueAffectedSoftware sias ON ssod.software_id = sias.software_id LEFT JOIN sms_devicetype sd2 ON sd1.devicetype_id = sd2.devicetype_id WHERE sias.issue_id = ?
UNION ALL
SELECT NULL AS deviceInstance_id, sd2.type, NULL AS project_id, COALESCE(sd1.version, 'Unknown') AS version FROM sms_device sd1 LEFT JOIN sms_softwarePartOfDevice ssod ON sd1.device_id = ssod.device_id LEFT JOIN sms_issueAffectedSoftware sias ON ssod.software_id = sias.software_id LEFT JOIN sms_devicetype sd2 ON sd1.devicetype_id = sd2.devicetype_id WHERE sias.issue_id = ? AND sd1.device_id NOT IN (SELECT device_id FROM sms_deviceInstance) AND sias.software_id IS NOT NULL
UNION ALL
SELECT sdi.deviceInstance_id, sd2.type, sdi.project_id, version FROM sms_deviceInstance sdi LEFT JOIN sms_device sd1 ON sdi.device_id = sd1.device_id LEFT JOIN sms_softwarePartOfDevice ssod ON sd1.device_id = ssod.device_id LEFT JOIN sms_componentPartOfSoftware scps ON ssod.software_id = scps.software_id LEFT JOIN sms_issueAffectedComponent siac ON scps.component_id = siac.component_id LEFT JOIN sms_devicetype sd2 ON sd1.devicetype_id = sd2.devicetype_id WHERE siac.issue_id = ?
UNION ALL
SELECT NULL AS deviceInstance_id, sd2.type, NULL AS project_id, COALESCE(sd1.version, 'Unknown') AS version FROM sms_device sd1 LEFT JOIN sms_softwarePartOfDevice ssod ON sd1.device_id = ssod.device_id LEFT JOIN sms_componentPartOfSoftware scps ON ssod.software_id = scps.software_id LEFT JOIN sms_issueAffectedComponent siac ON scps.component_id = siac.component_id LEFT JOIN sms_devicetype sd2 ON sd1.devicetype_id = sd2.devicetype_id WHERE siac.issue_id = ? AND sd1.device_id NOT IN (SELECT device_id FROM sms_deviceInstance)
UNION ALL
SELECT sdi.deviceInstance_id, sd2.type, sdi.project_id, version FROM sms_deviceInstance sdi LEFT JOIN sms_device sd1 ON sdi.device_id = sd1.device_id LEFT JOIN sms_artefactPartOfDevice sapd ON sd1.device_id = sapd.device_id LEFT JOIN sms_issueAffectedArtefact siaa ON sapd.artefact_id = siaa.artefact_id LEFT JOIN sms_devicetype sd2 ON sd1.devicetype_id = sd2.devicetype_id WHERE siaa.issue_id = ?
UNION ALL
SELECT NULL AS deviceInstance_id, sd2.type, NULL AS project_id, COALESCE(sd1.version, 'Unknown') AS version FROM sms_device sd1 LEFT JOIN sms_artefactPartOfDevice sapd ON sd1.device_id = sapd.device_id LEFT JOIN sms_issueAffectedArtefact siaa ON sapd.artefact_id = siaa.artefact_id LEFT JOIN sms_devicetype sd2 ON sd1.devicetype_id = sd2.devicetype_id WHERE siaa.issue_id = ? AND sd1.device_id NOT IN (SELECT device_id FROM sms_deviceInstance)
)AS subquery
)
SELECT
    COUNT(DISTINCT deviceInstance_id) AS affected_device_instances,
    COUNT(DISTINCT CASE WHEN deviceInstance_id IS NULL THEN CONCAT(type, version) ELSE NULL END) AS affected_devices_without_instances,
    COUNT(DISTINCT project_id) AS affected_projects,
    COUNT(DISTINCT CONCAT(type, version)) AS distinct_device_version_combinations
FROM combined_result;`
var SELECT_sms_issueAffectedProjects = `SELECT DISTINCT p.project_id, p.name, p.customer FROM sms_project p JOIN ( SELECT DISTINCT project_id FROM (
SELECT sdi.project_id FROM sms_deviceInstance sdi LEFT JOIN sms_device sd1 ON sdi.device_id = sd1.device_id LEFT JOIN sms_issueAffectedDevice siad ON sdi.device_id = siad.device_id WHERE siad.issue_id = ?
UNION ALL
SELECT sdi.project_id FROM sms_deviceInstance sdi LEFT JOIN sms_device sd1 ON sdi.device_id = sd1.device_id LEFT JOIN sms_softwarePartOfDevice ssod ON sd1.device_id = ssod.device_id LEFT JOIN sms_issueAffectedSoftware sias ON ssod.software_id = sias.software_id WHERE sias.issue_id = ?
UNION ALL
SELECT sdi.project_id FROM sms_deviceInstance sdi LEFT JOIN sms_device sd1 ON sdi.device_id = sd1.device_id LEFT JOIN sms_softwarePartOfDevice ssod ON sd1.device_id = ssod.device_id LEFT JOIN sms_componentPartOfSoftware scps ON ssod.software_id = scps.software_id LEFT JOIN sms_issueAffectedComponent siac ON scps.component_id = siac.component_id WHERE siac.issue_id = ?
UNION ALL
SELECT sdi.project_id FROM sms_deviceInstance sdi LEFT JOIN sms_device sd1 ON sdi.device_id = sd1.device_id LEFT JOIN sms_artefactPartOfDevice sapd ON sd1.device_id = sapd.device_id LEFT JOIN sms_issueAffectedArtefact siaa ON sapd.artefact_id = siaa.artefact_id WHERE siaa.issue_id = ? ) AS affected_projects WHERE project_id IS NOT NULL ) AS filtered_projects ON p.project_id = filtered_projects.project_id;`

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
var Check_sms_component = `SELECT sms_component.component_id FROM sms_component WHERE sms_component.name = ? AND sms_component.componentType = ? AND sms_component.version = ? LIMIT 1;`

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

// SMS IssueAffectedSoftware
var INSERT_sms_newIssueAffectedSoftware = `INSERT INTO sms_issueAffectedSoftware (software_id, issue_id, additionalInfo, confirmed) VALUES (?,?,?,?);`
var DELETE_sms_IssueAffectedSoftware = `DELETE FROM sms_issueAffectedSoftware WHERE software_id = ? AND issue_id = ?;`
var SELECT_sms_IssueAffectedSoftwaresForIssueID = `SELECT sms_issueAffectedSoftware.software_id, sms_issueAffectedSoftware.issue_id, sms_issueAffectedSoftware.additionalInfo, sms_issueAffectedSoftware.confirmed, sms_softwaretype.typeName, sms_software.version FROM sms_issueAffectedSoftware LEFT JOIN sms_software ON sms_issueAffectedSoftware.software_id = sms_software.software_id LEFT JOIN sms_softwaretype ON sms_software.softwaretype_id = sms_softwaretype.softwaretype_id WHERE sms_issueAffectedSoftware.issue_id = ?; `
var SELECT_sms_IssueAffectedSoftwaresForIssueIDWithInheritage = `SELECT DISTINCT sms_issueAffectedSoftware.software_id, sms_issueAffectedSoftware.issue_id, sms_issueAffectedSoftware.additionalInfo, sms_issueAffectedSoftware.confirmed, sms_softwaretype.typeName, sms_software.version, CASE WHEN siac.component_id IS NOT NULL THEN true ELSE false END AS inherit FROM sms_issueAffectedSoftware LEFT JOIN sms_software ON sms_issueAffectedSoftware.software_id = sms_software.software_id LEFT JOIN sms_softwaretype ON sms_software.softwaretype_id = sms_softwaretype.softwaretype_id LEFT JOIN sms_componentPartOfSoftware scps ON sms_software.software_id = scps.software_id LEFT JOIN sms_issueAffectedComponent siac ON scps.component_id = siac.component_id AND siac.issue_id = ? WHERE sms_issueAffectedSoftware.issue_id = ? AND ( siac.component_id IS NULL OR siac.issue_id = ? );`
var SELECT_sms_IssuesForSoftware = `SELECT DISTINCT sms_issueAffectedSoftware.software_id, sms_issue.issue_id, COALESCE(sms_issueAffectedSoftware.additionalInfo, siac.additionalInfo) AS additionalInfo, COALESCE(sms_issueAffectedSoftware.confirmed, siac.confirmed) AS confirmed, sms_issue.name, CASE WHEN siac.component_id IS NOT NULL THEN true ELSE false END AS inherit FROM sms_issueAffectedSoftware LEFT JOIN sms_issue ON sms_issue.issue_id = sms_issueAffectedSoftware.issue_id LEFT JOIN sms_componentPartOfSoftware scps ON sms_issueAffectedSoftware.software_id = scps.software_id LEFT JOIN sms_issueAffectedComponent siac ON scps.component_id = siac.component_id AND siac.issue_id = sms_issueAffectedSoftware.issue_id WHERE sms_issueAffectedSoftware.software_id = ? OR siac.component_id IS NOT NULL; `

// SMS ArtefactPartOfDevice
var INSERT_sms_newArtefactPartOfDevice = `INSERT INTO sms_artefactPartOfDevice (device_id, artefact_id, additionalInfo) VALUES (?,?,?);`
var DELETE_sms_ArtefactPartOfDevice = `DELETE FROM sms_artefactPartOfDevice WHERE device_id = ? AND artefact_id = ?;`
var SELECT_sms_ArtefactPartOfDeviceForDevice = `SELECT sms_artefactPartOfDevice.device_id, sms_artefactPartOfDevice.artefact_id, sms_artefactPartOfDevice.additionalInfo, sms_artefacttype.artefactType, sms_artefact.version FROM sms_artefactPartOfDevice LEFT JOIN sms_artefact ON sms_artefactPartOfDevice.artefact_id = sms_artefact.artefact_id LEFT JOIN sms_artefacttype ON sms_artefact.artefacttype_id = sms_artefacttype.artefacttype_id WHERE sms_artefactPartOfDevice.device_id =  ?; `
var SELECT_sms_ArtefactPartOfDeviceForArtefact = `SELECT sms_artefactPartOfDevice.device_id, sms_artefactPartOfDevice.artefact_id, sms_artefactPartOfDevice.additionalInfo, sms_devicetype.type, sms_device.version FROM sms_artefactPartOfDevice LEFT JOIN sms_device ON sms_artefactPartOfDevice.device_id = sms_device.device_id LEFT JOIN sms_devicetype ON sms_device.devicetype_id = sms_devicetype.devicetype_id WHERE sms_artefactPartOfDevice.artefact_id = ?; `

// SMS_ManufacturingOrder
var SELECT_sms_ManufacturingOrdersForSystem = `SELECT sms_manufacturingOrder.manufacturingOrder_id, sms_manufacturingOrder.system_id, sms_manufacturingOrder.packageReference, sms_manufacturingOrder.start, sms_manufacturingOrder.end, sms_manufacturingOrder.description FROM sms_manufacturingOrder WHERE system_id = ? `
var INSERT_sms_newManufacturingOrder = `INSERT INTO sms_manufacturingOrder (system_id, packageReference, start, description) VALUES (?,?,?,?);`
var SELECT_sms_ManufacturingOrderInfo = `SELECT sms_manufacturingOrder.manufacturingOrder_id, sms_manufacturingOrder.system_id, sms_manufacturingOrder.packageReference, sms_manufacturingOrder.start, sms_manufacturingOrder.end, sms_manufacturingOrder.description FROM sms_manufacturingOrder WHERE manufacturingOrder_id = ?`

// SMS Certification
var SELECT_sms_certification = `SELECT sms_certification.certification_id, sms_certification.name, sms_certification.date, sms_certification.description FROM sms_certification;`
var SELECT_sms_certificationInfo = `SELECT sms_certification.certification_id, sms_certification.name, sms_certification.date, sms_certification.description FROM sms_certification WHERE certification_id = ?;`
var INSERT_sms_newCertification = `INSERT INTO sms_certification (name, date, description) VALUES (?,?,?);`
var DELETE_sms_certification = `DELETE FROM sms_certification WHERE certification_id = ?;`

// SMS_SystemHasCertification
var SELECT_sms_systemHasCertification = `SELECT sms_systemHasCertification.system_id, sms_systemHasCertification.certification_id, sms_systemHasCertification.additionalInfo FROM sms_systemHasCertification;`
var DELETE_sms_systemHasCertification = `DELETE FROM sms_systemHasCertification WHERE system_id = ? AND certification_id = ?;`
var INSERT_sms_systemHasCertification = `INSERT INTO sms_systemHasCertification (system_id, certification_id, additionalInfo) VALUES (?, ?, ?);`
var SELECT_sms_systemHasCertificationForSystem = `SELECT sms_systemHasCertification.system_id, sms_systemHasCertification.certification_id, sms_systemHasCertification.additionalInfo, sms_certification.name AS certification_name FROM sms_systemHasCertification LEFT JOIN sms_certification ON sms_systemHasCertification.certification_id = sms_certification.certification_id WHERE sms_systemHasCertification.system_id = ?;`
var SELECT_sms_systemHasCertificationForCertification = `SELECT sms_systemHasCertification.system_id, sms_systemHasCertification.certification_id, sms_systemHasCertification.additionalInfo, sms_systemtype.type AS system_name, sms_system.version AS system_version FROM sms_systemHasCertification LEFT JOIN sms_system ON sms_systemHasCertification.system_id = sms_system.system_id LEFT JOIN sms_systemtype ON sms_system.systemtype_id = sms_systemtype.systemtype_id WHERE sms_systemHasCertification.certification_id = ?;`

// SMS IssueAffectedComponents
var INSERT_sms_newIssueAffectedComponent = `INSERT INTO sms_issueAffectedComponent (component_id, issue_id, additionalInfo, confirmed) VALUES (?,?,?,?);`
var DELETE_sms_IssueAffectedComponent = `DELETE FROM sms_issueAffectedComponent WHERE component_id = ? AND issue_id = ?;`
var SELECT_sms_IssueAffectedComponentsForIssueID = `SELECT sms_issueAffectedComponent.component_id, sms_issueAffectedComponent.issue_id, sms_issueAffectedComponent.additionalInfo, sms_issueAffectedComponent.confirmed, sms_component.name AS component_name, sms_component.version AS component_version FROM sms_issueAffectedComponent LEFT JOIN sms_component ON sms_issueAffectedComponent.component_id = sms_component.component_id WHERE sms_issueAffectedComponent.issue_id = ?;`
var SELECT_sms_IssuesForComponent = ` SELECT sms_issueAffectedComponent.component_id, sms_issueAffectedComponent.issue_id, sms_issueAffectedComponent.additionalInfo, sms_issueAffectedComponent.confirmed, sms_issue.name FROM sms_issueAffectedComponent LEFT JOIN sms_issue ON sms_issueAffectedComponent.issue_id = sms_issue.issue_id WHERE sms_issueAffectedComponent.component_id = ?;`

// SMS IssueAffectedArtefacts
var INSERT_sms_newIssueAffectedArtefact = `INSERT INTO sms_issueAffectedArtefact (artefact_id, issue_id, additionalInfo, confirmed) VALUES (?,?,?,?);`
var DELETE_sms_IssueAffectedArtefact = `DELETE FROM sms_issueAffectedArtefact WHERE artefact_id = ? AND issue_id = ?;`
var SELECT_sms_IssueAffectedArtefactsForIssueID = `SELECT sms_issueAffectedArtefact.artefact_id, sms_issueAffectedArtefact.issue_id, sms_issueAffectedArtefact.additionalInfo, sms_issueAffectedArtefact.confirmed, sms_artefact.name AS artefact_name, sms_artefact.version AS artefact_version FROM sms_issueAffectedArtefact LEFT JOIN sms_artefact ON sms_issueAffectedArtefact.artefact_id = sms_artefact.artefact_id WHERE sms_issueAffectedArtefact.issue_id = ?;`
var SELECT_sms_IssuesForArtefact = `SELECT sms_issueAffectedArtefact.artefact_id, sms_issueAffectedArtefact.issue_id, sms_issueAffectedArtefact.additionalInfo, sms_issueAffectedArtefact.confirmed, sms_issue.name FROM sms_issueAffectedArtefact LEFT JOIN sms_issue ON sms_issueAffectedArtefact.issue_id = sms_issue.issue_id WHERE sms_issueAffectedArtefact.artefact_id = ?;`

// SMS_SecurityReport
var SELECT_all_reports = `SELECT report_id, report_name, scanner_name, scanner_version, creation_date, upload_date, uploaded_by, scan_scope, vulnerability_count, component_count FROM sms_securityReport;`
var SELECT_report_by_id = `SELECT report_id, report_name, scanner_name, scanner_version, creation_date, upload_date, uploaded_by, scan_scope, vulnerability_count, component_count FROM sms_securityReport WHERE report_id = ?;`
var INSERT_new_report = `INSERT INTO sms_securityReport ( report_name, scanner_name, scanner_version, creation_date, upload_date, uploaded_by, scan_scope, vulnerability_count, component_count ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`
var DELETE_report = ` DELETE FROM sms_securityReport WHERE report_id = ?;`
var UPDATE_report = ` UPDATE sms_securityReport SET report_name = ?, scanner_name = ?, scanner_version = ?, creation_date = ?, upload_date = ?, uploaded_by = ?, scan_scope = ?, vulnerability_count = ?, component_count = ? WHERE report_id = ?;`

// SMS_SecurityReportLink
// Hinzufügen eines neuen Links
var INSERT_new_securityReport = `INSERT INTO sms_securityReportLink (report_id, linked_object_id, linked_object_type) VALUES (?, ?, ?);`
// Alle Links für einen bestimmten Report abrufen
var SELECT_securityReport_by_ID = `SELECT linked_object_id, linked_object_type FROM sms_securityReportLink WHERE report_id = ?;`
// Alle Reports für ein bestimmtes Objekt abrufen
var SELECT_securityReport_by_ObjectID = `SELECT sr.report_id, sr.report_name, sr.scanner_name, sr.scanner_version, sr.creation_date, sr.upload_date, sr.uploaded_by, sr.scan_scope, sr.vulnerability_count, sr.component_count FROM sms_securityReportLink AS srl JOIN sms_securityReport AS sr ON srl.report_id = sr.report_id WHERE srl.linked_object_id = ? AND srl.linked_object_type = ?;`
// Alle Reports für einen bestimmten Typ abrufen
var SELECT_securityReport_by_ObjectType = `SELECT report_id, linked_object_id FROM sms_securityReportLink WHERE linked_object_type = ?;`
// Einen spezifischen Link löschen
var DELETE_securityReport_by_IDs = `DELETE FROM sms_securityReportLink WHERE report_id = ? AND linked_object_id = ? AND linked_object_type = ?;`
// Alle Links für einen bestimmten Report löschen
var DELETE_securityReport_by_reportID = `DELETE FROM sms_securityReportLink WHERE report_id = ?;`
// Alle Links für ein bestimmtes Objekt löschen
var DELETE_securityReport_by_objectID = `DELETE FROM sms_securityReportLink WHERE linked_object_id = ? AND linked_object_type = ?;`

// SMS_projectSetting & SMS_projectSettingLink
var INSERT_new_projectSettings = `INSERT INTO sms_projectSettings (key_name, value_type, default_value) VALUES (?, ?, ?);`
var UPDATE_projectSettings = `UPDATE sms_projectSettings SET key_name = ?, value_type = ?, default_value = ? WHERE setting_id = ?;`
var INSERT_new_projectSettingsLink = `INSERT INTO sms_projectSettingsLink (project_id, setting_id, value) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE value = VALUES(value);`
var SELECT_settings_for_project = `SELECT ps.setting_id, ps.key_name, COALESCE(psl.value, ps.default_value) AS value, ps.value_type FROM sms_projectSettingsLink psl INNER JOIN sms_projectSettings ps ON ps.setting_id = psl.setting_id WHERE psl.project_id = ?;`
var SELECT_projectSetting = `SELECT COALESCE(psl.value, ps.default_value) AS value, ps.value_type FROM sms_projectSettings ps LEFT JOIN sms_projectSettingsLink psl ON ps.setting_id = psl.setting_id AND psl.project_id = ? WHERE ps.key_name = ?;`
var DELETE_projectSettingsLink = `DELETE FROM sms_projectSettingsLink WHERE project_id = ? AND setting_id = ?;`
var DELETE_global_Setting = `DELETE FROM sms_projectSettings WHERE setting_id = ?;`
var SELECT_all_Settings = `SELECT setting_id, key_name, default_value, value_type FROM sms_projectSettings;`
var UPDATE_projectSettingsLink = `UPDATE sms_projectSettingsLink SET value = ? WHERE project_id = ? AND setting_id = ?;`
var SELECT_available_settings_for_project = `SELECT ps.setting_id, ps.key_name, ps.value_type, ps.default_value FROM sms_projectSettings ps WHERE NOT EXISTS (SELECT 1 FROM sms_projectSettingsLink psl WHERE psl.setting_id = ps.setting_id AND psl.project_id = ?);`

// sms_deviceIPDefinition
var INSERT_new_deviceIPDefinition = `INSERT INTO sms_deviceIPDefinition (device_type_id, applicable_versions, ip_address, vlan_id, description, filter_condition) VALUES (?, ?, ?, ?, ?, ?);`;
var UPDATE_deviceIPDefinition = `UPDATE sms_deviceIPDefinition SET device_type_id = ?, applicable_versions = ?, ip_address = ?, vlan_id = ?, description = ?, filter_condition = ? WHERE id = ?;`;
var SELECT_ips_for_deviceType = `SELECT id, device_type_id, applicable_versions, ip_address, vlan_id, description, filter_condition FROM sms_deviceIPDefinition WHERE device_type_id = ?;`;
var SELECT_ips_for_device = `SELECT dip.id, dip.device_type_id, dip.applicable_versions, dip.ip_address, dip.vlan_id, dip.description, dip.filter_condition 
FROM sms_deviceIPDefinition dip 
JOIN sms_device d ON d.devicetype_id = dip.device_type_id 
WHERE d.device_id = ? 
AND (dip.applicable_versions = 'all' OR FIND_IN_SET(d.version, dip.applicable_versions) > 0);`;
var DELETE_deviceIPDefinition = `DELETE FROM sms_deviceIPDefinition WHERE id = ?;`;
var SELECT_ips = `SELECT 
    dip.id, 
    dt.type AS device_type_name, 
    dip.applicable_versions, 
    dip.ip_address, 
    dip.vlan_id, 
    dip.description, 
    dip.filter_condition
FROM sms_deviceIPDefinition dip
JOIN sms_devicetype dt ON dip.device_type_id = dt.devicetype_id;`
var SELECT_ips_for_project = `SELECT 
    dip.ip_address,
    dip.applicable_versions,
    dip.vlan_id,
    dip.description,
    dip.filter_condition,
    dt.type AS device_type,
    COUNT(di.deviceInstance_id) AS instance_count, 
    GROUP_CONCAT(DISTINCT d.version ORDER BY d.version ASC SEPARATOR ', ') AS versions
FROM sms_deviceIPDefinition dip
JOIN sms_devicetype dt ON dip.device_type_id = dt.devicetype_id
JOIN sms_device d ON dt.devicetype_id = d.devicetype_id
JOIN sms_deviceInstance di ON d.device_id = di.device_id
WHERE di.project_id = ?
GROUP BY dip.ip_address, dip.applicable_versions, dip.vlan_id, dip.description, dip.filter_condition, dt.type;`