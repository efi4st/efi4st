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


// Results
