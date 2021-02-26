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
var INSERT_newbinaryAnalysis = `INSERT INTO binaryAnalysis (toolOutput, analysisTool_id, relevantApps_id) VALUES (?,?,?);`
var DELETE_binaryAnalysis = `DELETE FROM binaryAnalysis WHERE binaryAnalysis_id = ?;`
var DELETE_binaryAnalysisByRelevantAppPath = `DELETE FROM binaryAnalysis WHERE binaryAnalysis_id = ?;`
var SELECT_binaryAnalysistByBinary = `SELECT * FROM binaryAnalysis WHERE binaryAnalysis.relevantApps_id = ?;`
var UPDATE_binaryAnalysis = `UPDATE binaryAnalysis SET binaryAnalysis.toolOutput = ? WHERE binaryAnalysis.binaryAnalysis_id = ?;`

// AnalysisTool
var SELECT_analysisTool = `SELECT analysisTool.analysisTool_id, analysisTool.name, analysisTool.call from analysisTool;`
var SELECT_analysisToolInfo = `SELECT analysisTool_id, name, call FROM analysisTool WHERE analysisTool_id = ?;`
var INSERT_newAnalysisTool = `INSERT INTO analysisTool (name, call) VALUES (?,?);`
var DELETE_analysisTool = `DELETE FROM analysisTool WHERE analysisTool_id = ?;`


