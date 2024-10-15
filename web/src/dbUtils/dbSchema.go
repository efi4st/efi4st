/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package dbUtils

var projectSchema = `
CREATE TABLE IF NOT EXISTS projects (
	project_id INT(11) NOT NULL AUTO_INCREMENT,
	name VARCHAR(150) NOT NULL,
	uploads INT DEFAULT NULL,
	date DATE NOT NULL,
	PRIMARY KEY (project_id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var firmwareSchema = `
CREATE TABLE IF NOT EXISTS firmware (
	firmware_id INT(11) NOT NULL AUTO_INCREMENT,
   	name VARCHAR(150) NOT NULL,
	version VARCHAR(150) DEFAULT NULL,
	binwalkOutput VARCHAR(1000) DEFAULT NULL,
	sizeInBytes INT DEFAULT NULL,
	project_id INT(11) NOT NULL,
	created DATE NOT NULL,
	PRIMARY KEY (firmware_id),
	CONSTRAINT firmware_ibfk_1 FOREIGN KEY (project_id) REFERENCES projects (project_id) ON UPDATE CASCADE ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
`

var relevantAppsSchema = `
CREATE TABLE IF NOT EXISTS relevantApps (
	relevantApps_id INT(11) NOT NULL AUTO_INCREMENT,
   	name VARCHAR(150) NOT NULL,
	path VARCHAR(300) DEFAULT NULL,
	extPort INT DEFAULT NULL,
	extProtocoll VARCHAR(300) DEFAULT NULL,
	intInterface VARCHAR(300) DEFAULT NULL,
	moduleDefault BOOLEAN,
	moduleInitSystem BOOLEAN,
	moduleFileContent BOOLEAN,
	moduleBash BOOLEAN,
	moduleCronJob BOOLEAN,
	moduleProcesses BOOLEAN,
	moduleInterfaces BOOLEAN,
	moduleSystemControls BOOLEAN,
	moduleFileSystem BOOLEAN,
	modulePortscanner BOOLEAN,
	moduleProtocolls BOOLEAN,
	moduleNetInterfaces BOOLEAN,
	moduleFileSystemInterfaces BOOLEAN,
	moduleFileHandles BOOLEAN,
	firmware_id INT(11) NOT NULL,
	PRIMARY KEY (relevantApps_id),
	CONSTRAINT relevantApps_ibfk_1 FOREIGN KEY (firmware_id) REFERENCES firmware (firmware_id) ON UPDATE CASCADE ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
`

var testResultSchema = `
CREATE TABLE IF NOT EXISTS testResult (
	testResult_id INT(11) NOT NULL AUTO_INCREMENT,
   	moduleName VARCHAR(150) NOT NULL,
	result LONGTEXT DEFAULT NULL,
	created DATE NOT NULL,
	firmware_id INT(11) NOT NULL,
	PRIMARY KEY (testResult_id),
	CONSTRAINT testResult_ibfk_1 FOREIGN KEY (firmware_id) REFERENCES firmware (firmware_id) ON UPDATE CASCADE ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
`

var appContentSchema = `
CREATE TABLE IF NOT EXISTS appContent (
	appContent_id INT(11) NOT NULL AUTO_INCREMENT,
	contentPathList LONGTEXT DEFAULT NULL,
	binwalkOutput LONGTEXT DEFAULT NULL,
	readelfOutput LONGTEXT DEFAULT NULL,
	lddOutput LONGTEXT DEFAULT NULL,
	straceOutput LONGTEXT DEFAULT NULL,
	relevantApps_path VARCHAR(150) NOT NULL,
	PRIMARY KEY (appContent_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
`

var binaryAnalysisSchema = `
CREATE TABLE IF NOT EXISTS binaryAnalysis (
	binaryAnalysis_id INT(11) NOT NULL AUTO_INCREMENT,
	toolOutput LONGTEXT DEFAULT NULL,
	analysisTool_id INT(11) NOT NULL,
	relevantApps_id INT(11) NOT NULL,
	PRIMARY KEY (binaryAnalysis_id),
	CONSTRAINT binaryAnalysis_ibfk_1 FOREIGN KEY (analysisTool_id) REFERENCES analysisTool (analysisTool_id) ON UPDATE CASCADE ON DELETE CASCADE,
	CONSTRAINT binaryAnalysis_ibfk_2 FOREIGN KEY (relevantApps_id) REFERENCES relevantApps (relevantApps_id) ON UPDATE CASCADE ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
`

var analysisToolSchema = `
CREATE TABLE IF NOT EXISTS analysisTool (
	analysisTool_id INT(11) NOT NULL AUTO_INCREMENT,
	name VARCHAR(150) NOT NULL,
	executionString VARCHAR(300) NOT NULL,
	PRIMARY KEY (analysisTool_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
`

/**
 * Security Management System
 * Created:   29.09.2024
 *
 * (C)
 **/


var sms_projecttype_schema = `
CREATE TABLE IF NOT EXISTS sms_projecttype (
	projecttype_id INT(11) NOT NULL AUTO_INCREMENT,
	type VARCHAR(150) NOT NULL,
	PRIMARY KEY (projecttype_id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_project_schema = `
CREATE TABLE IF NOT EXISTS sms_project (
	project_id INT(11) NOT NULL AUTO_INCREMENT,
	name VARCHAR(150) NOT NULL,
	customer VARCHAR(150) NOT NULL,
	projecttype_id INT(11) NOT NULL,
	reference VARCHAR(150) DEFAULT NULL,
	date DATE NOT NULL,
	active BOOLEAN,
	PRIMARY KEY (project_id),
	CONSTRAINT sms_project_ibfk_1 FOREIGN KEY (projecttype_id) REFERENCES sms_projecttype (projecttype_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_systemtype_schema = `
CREATE TABLE IF NOT EXISTS sms_systemtype (
	systemtype_id INT(11) NOT NULL AUTO_INCREMENT,
	type VARCHAR(150) NOT NULL,
	PRIMARY KEY (systemtype_id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_system_schema = `
CREATE TABLE IF NOT EXISTS sms_system (
	system_id INT(11) NOT NULL AUTO_INCREMENT,
	systemtype_id INT(11) NOT NULL,
	version VARCHAR(150) NOT NULL,
	date DATE NOT NULL,
	PRIMARY KEY (system_id),
	CONSTRAINT sms_system_ibfk_1 FOREIGN KEY (systemtype_id) REFERENCES sms_systemtype (systemtype_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_projectBOMSystemInventory_schema = `
CREATE TABLE IF NOT EXISTS sms_projectBOMSystemInventory (
	project_id INT(11) NOT NULL,
	system_id INT(11) NOT NULL,
	amount INT(11) NOT NULL,
	PRIMARY KEY (project_id, system_id),
	CONSTRAINT sms_projectBOMSystemInventory_ibfk_1 FOREIGN KEY (project_id) REFERENCES sms_project (project_id) ON UPDATE CASCADE ON DELETE NO ACTION,
	CONSTRAINT sms_projectBOMSystemInventory_ibfk_2 FOREIGN KEY (system_id) REFERENCES sms_system (system_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_devicetype_schema = `
CREATE TABLE IF NOT EXISTS sms_devicetype (
	devicetype_id INT(11) NOT NULL AUTO_INCREMENT,
	type VARCHAR(150) NOT NULL,
	PRIMARY KEY (devicetype_id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_device_schema = `
CREATE TABLE IF NOT EXISTS sms_device (
	device_id INT(11) NOT NULL AUTO_INCREMENT,
	devicetype_id INT(11) NOT NULL,
	version VARCHAR(150) NOT NULL,
	date DATE NOT NULL,
	PRIMARY KEY (device_id),
	CONSTRAINT sms_device_ibfk_1 FOREIGN KEY (devicetype_id) REFERENCES sms_devicetype (devicetype_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_deviceInstance_schema = `
CREATE TABLE IF NOT EXISTS sms_deviceInstance (
	deviceInstance_id INT(11) NOT NULL AUTO_INCREMENT,
	project_id INT(11) NOT NULL,
	device_id INT(11) NOT NULL,
	serialnumber VARCHAR(150) NOT NULL,
	provisioner VARCHAR(150) NOT NULL,
	configuration VARCHAR(150) NOT NULL,
	date DATE NOT NULL,
	PRIMARY KEY (deviceInstance_id),
	CONSTRAINT sms_deviceInstance_ibfk_1 FOREIGN KEY (project_id) REFERENCES sms_project (project_id) ON UPDATE CASCADE ON DELETE NO ACTION,
	CONSTRAINT sms_deviceInstance_ibfk_2 FOREIGN KEY (device_id) REFERENCES sms_device (device_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_updateHistory_schema = `
CREATE TABLE IF NOT EXISTS sms_updateHistory (
	updateHistory_id INT(11) NOT NULL AUTO_INCREMENT,
	deviceInstance_id INT(11) NOT NULL,
	user VARCHAR(150) NOT NULL,
	updateType VARCHAR(150) NOT NULL,
	date DATE NOT NULL,
	description VARCHAR(150) NOT NULL,
	PRIMARY KEY (updateHistory_id),
	CONSTRAINT sms_updateHistory_ibfk_1 FOREIGN KEY (deviceInstance_id) REFERENCES sms_deviceInstance (deviceInstance_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_devicePartOfSystem_schema = `
CREATE TABLE IF NOT EXISTS sms_devicePartOfSystem (
	device_id INT(11) NOT NULL,
	system_id INT(11) NOT NULL,
	amount INT(11) NOT NULL,
	PRIMARY KEY (device_id, system_id),
	CONSTRAINT sms_devicePartOfSystem_ibfk_1 FOREIGN KEY (device_id) REFERENCES sms_device (device_id) ON UPDATE CASCADE ON DELETE NO ACTION,
	CONSTRAINT sms_devicePartOfSystem_ibfk_2 FOREIGN KEY (system_id) REFERENCES sms_system (system_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_releasenote_schema = `
CREATE TABLE IF NOT EXISTS sms_releasenote (
	releasenote_id INT(11) NOT NULL AUTO_INCREMENT,
	device_id INT(11) NOT NULL,
	type VARCHAR(150) NOT NULL,
	date DATE NOT NULL,
	details VARCHAR(150) NOT NULL,
	PRIMARY KEY (releasenote_id),
	CONSTRAINT sms_releasenote_ibfk_1 FOREIGN KEY (device_id) REFERENCES sms_device (device_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_applicationtype_schema = `
CREATE TABLE IF NOT EXISTS sms_applicationtype (
	applicationtype_id INT(11) NOT NULL AUTO_INCREMENT,
	type VARCHAR(150) NOT NULL,
	PRIMARY KEY (applicationtype_id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_application_schema = `
CREATE TABLE IF NOT EXISTS sms_application (
	application_id INT(11) NOT NULL AUTO_INCREMENT,
	device_id INT(11) NOT NULL,
	applicationtype_id INT(11) NOT NULL,
	version VARCHAR(150) NOT NULL,
	date DATE NOT NULL,
	PRIMARY KEY (application_id),
	CONSTRAINT sms_application_ibfk_1 FOREIGN KEY (device_id) REFERENCES sms_device (device_id) ON UPDATE CASCADE ON DELETE NO ACTION,
	CONSTRAINT sms_application_ibfk_2 FOREIGN KEY (applicationtype_id) REFERENCES sms_applicationtype (applicationtype_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_applicationPartOfDevice_schema = `
CREATE TABLE IF NOT EXISTS sms_applicationPartOfDevice (
	application_id INT(11) NOT NULL,
	device_id INT(11) NOT NULL,
	amount INT(11) NOT NULL,
	PRIMARY KEY (application_id, device_id),
	CONSTRAINT sms_applicationPartOfDevice_ibfk_1 FOREIGN KEY (application_id) REFERENCES sms_application (application_id) ON UPDATE CASCADE ON DELETE NO ACTION,
	CONSTRAINT sms_applicationPartOfDevice_ibfk_2 FOREIGN KEY (device_id) REFERENCES sms_device (device_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_component_schema = `
CREATE TABLE IF NOT EXISTS sms_component (
	component_id INT(11) NOT NULL AUTO_INCREMENT,
	version VARCHAR(150) NOT NULL,
	name VARCHAR(150) NOT NULL,
	date DATE NOT NULL,
	PRIMARY KEY (component_id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_componentPartOfApplication_schema = `
CREATE TABLE IF NOT EXISTS sms_componentPartOfApplication (
    component_id INT(11) NOT NULL,
    application_id INT(11) NOT NULL,
    PRIMARY KEY (component_id, application_id),
    CONSTRAINT sms_componentPartOfApplication_ibfk_1 FOREIGN KEY (component_id) REFERENCES sms_component (component_id) ON UPDATE CASCADE ON DELETE NO ACTION,
    CONSTRAINT sms_componentPartOfApplication_ibfk_2 FOREIGN KEY (application_id) REFERENCES sms_application (application_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_issue_schema = `
CREATE TABLE IF NOT EXISTS sms_issue (
    issue_id INT(11) NOT NULL AUTO_INCREMENT,
	name VARCHAR(60) NOT NULL,
	date DATE NOT NULL,
	issueType VARCHAR(50) NOT NULL,
	reference VARCHAR(150) DEFAULT NULL,
	criticality INT(11) DEFAULT NULL,
	cve VARCHAR(50) DEFAULT NULL,
	description VARCHAR(150) DEFAULT NULL,
	PRIMARY KEY (issue_id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_issueAffectedDevice_schema = `
CREATE TABLE IF NOT EXISTS sms_issueAffectedDevice (
	device_id INT(11) NOT NULL,
	issue_id INT(11) NOT NULL,
	additionalInfo VARCHAR(150) DEFAULT NULL,
	confirmed BOOLEAN NOT NULL,
	PRIMARY KEY (device_id, issue_id),
	CONSTRAINT sms_issueAffectedDevice_ibfk_1 FOREIGN KEY (device_id) REFERENCES sms_device (device_id) ON UPDATE CASCADE ON DELETE NO ACTION,
	CONSTRAINT sms_issueAffectedDevice_ibfk_2 FOREIGN KEY (issue_id) REFERENCES sms_issue (issue_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_solution_schema = `
CREATE TABLE IF NOT EXISTS sms_solution (
    solution_id INT(11) NOT NULL AUTO_INCREMENT,
	issue_id INT(11) NOT NULL,
	devicetype_id INT(11) NOT NULL,
	date DATE NOT NULL,
	name VARCHAR(60) NOT NULL,
	description VARCHAR(150) DEFAULT NULL,
	reference VARCHAR(150) DEFAULT NULL,
	PRIMARY KEY (solution_id),
	CONSTRAINT sms_solution_ibfk_1 FOREIGN KEY (issue_id) REFERENCES sms_issue (issue_id) ON UPDATE CASCADE ON DELETE NO ACTION,
	CONSTRAINT sms_solution_ibfk_2 FOREIGN KEY (devicetype_id) REFERENCES sms_devicetype (devicetype_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`