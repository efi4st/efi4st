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
