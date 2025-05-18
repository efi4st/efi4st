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
    plant_number VARCHAR(150) DEFAULT NULL,
    project_reference VARCHAR(150) DEFAULT NULL,
    imo_plant_powerplant_factory VARCHAR(150) DEFAULT NULL,
    plant_type ENUM('IMO', 'Plant', 'PowerPlant', 'Factory') DEFAULT NULL,
    note TEXT DEFAULT NULL,
    PRIMARY KEY (project_id),
    CONSTRAINT sms_project_ibfk_1 FOREIGN KEY (projecttype_id) REFERENCES sms_projecttype (projecttype_id) ON UPDATE CASCADE ON DELETE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
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

var sms_releasenote_schema = `
CREATE TABLE IF NOT EXISTS sms_releasenote (
	releasenote_id INT(11) NOT NULL AUTO_INCREMENT,
	device_id INT(11) NOT NULL,
	type VARCHAR(80) NOT NULL,
	date DATE NOT NULL,
	details VARCHAR(300) NOT NULL,
	PRIMARY KEY (releasenote_id),
	CONSTRAINT sms_releasenote_ibfk_1 FOREIGN KEY (device_id) REFERENCES sms_device (device_id) ON UPDATE CASCADE ON DELETE NO ACTION
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

var sms_artefacttype_schema = `
CREATE TABLE IF NOT EXISTS sms_artefacttype (
	artefacttype_id INT(11) NOT NULL AUTO_INCREMENT,
	artefactType VARCHAR(150) NOT NULL,
	PRIMARY KEY (artefacttype_id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_artefact_schema = `
CREATE TABLE IF NOT EXISTS sms_artefact (
	artefact_id INT(11) NOT NULL AUTO_INCREMENT,
	artefacttype_id INT(11) NOT NULL,
	name VARCHAR(150) NOT NULL,
	version VARCHAR(50) NOT NULL,
	PRIMARY KEY (artefact_id),
	CONSTRAINT sms_artefact_ibfk_1 FOREIGN KEY (artefacttype_id) REFERENCES sms_artefacttype (artefacttype_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_softwaretype_schema = `
CREATE TABLE IF NOT EXISTS sms_softwaretype (
	softwaretype_id INT(11) NOT NULL AUTO_INCREMENT,
	typeName VARCHAR(150) NOT NULL,
	PRIMARY KEY (softwaretype_id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_software_schema = `
CREATE TABLE IF NOT EXISTS sms_software (
	software_id INT(11) NOT NULL AUTO_INCREMENT,
	softwaretype_id INT(11) NOT NULL,
	version VARCHAR(80) NOT NULL,
	date DATE NOT NULL,
	license VARCHAR(50) DEFAULT NULL,
	thirdParty BOOLEAN NOT NULL,
	releaseNote VARCHAR(300) DEFAULT NULL,
	PRIMARY KEY (software_id),
	CONSTRAINT sms_software_ibfk_1 FOREIGN KEY (softwaretype_id) REFERENCES sms_softwaretype (softwaretype_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_component_schema = `
CREATE TABLE IF NOT EXISTS sms_component (
	component_id INT(11) NOT NULL AUTO_INCREMENT,
	name VARCHAR(100) NOT NULL,
	componentType VARCHAR(80) NOT NULL,
	version VARCHAR(80) NOT NULL,
	date DATE NOT NULL,
	license VARCHAR(50) DEFAULT NULL,
	thirdParty BOOLEAN NOT NULL,
	releaseNote VARCHAR(300) DEFAULT NULL,
	PRIMARY KEY (component_id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`
var sms_componentPartOfSoftware_schema = `
CREATE TABLE IF NOT EXISTS sms_componentPartOfSoftware (
	software_id INT(11) NOT NULL,
	component_id INT(11) NOT NULL,
	additionalInfo VARCHAR(150) DEFAULT NULL,
	PRIMARY KEY (software_id, component_id),
	CONSTRAINT sms_componentPartOfSoftware_ibfk_1 FOREIGN KEY (software_id) REFERENCES sms_software (software_id) ON UPDATE CASCADE ON DELETE NO ACTION,
	CONSTRAINT sms_componentPartOfSoftware_ibfk_2 FOREIGN KEY (component_id) REFERENCES sms_component (component_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_softwarePartOfDevice_schema = `
CREATE TABLE IF NOT EXISTS sms_softwarePartOfDevice (
	device_id INT(11) NOT NULL,
	software_id INT(11) NOT NULL,
	additionalInfo VARCHAR(150) DEFAULT NULL,
	PRIMARY KEY (device_id, software_id),
	CONSTRAINT sms_softwarePartOfDevice_ibfk_1 FOREIGN KEY (device_id) REFERENCES sms_device (device_id) ON UPDATE CASCADE ON DELETE NO ACTION,
	CONSTRAINT sms_softwarePartOfDevice_ibfk_2 FOREIGN KEY (software_id) REFERENCES sms_software (software_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_devicePartOfSystem_schema = `
CREATE TABLE IF NOT EXISTS sms_devicePartOfSystem (
	system_id INT(11) NOT NULL,
	device_id INT(11) NOT NULL,
	additionalInfo VARCHAR(150) DEFAULT NULL,
	PRIMARY KEY (system_id, device_id),
	CONSTRAINT sms_devicePartOfSystem_ibfk_1 FOREIGN KEY (system_id) REFERENCES sms_system (system_id) ON UPDATE CASCADE ON DELETE NO ACTION,
	CONSTRAINT sms_devicePartOfSystem_ibfk_2 FOREIGN KEY (device_id) REFERENCES sms_device (device_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_projectBOM_schema = `
CREATE TABLE IF NOT EXISTS sms_projectBOM (
	projectBOM_id INT(11) NOT NULL AUTO_INCREMENT,
	project_id INT(11) NOT NULL,
	system_id INT(11) NOT NULL,
	orderNumber VARCHAR(80) DEFAULT NULL,
	additionalInfo VARCHAR(150) DEFAULT NULL,
	PRIMARY KEY (projectBOM_id),
	CONSTRAINT sms_projectBOM_ibfk_1 FOREIGN KEY (project_id) REFERENCES sms_project (project_id) ON UPDATE CASCADE ON DELETE NO ACTION,
	CONSTRAINT sms_projectBOM_ibfk_2 FOREIGN KEY (system_id) REFERENCES sms_system (system_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_issueAffectedSoftware_schema = `
CREATE TABLE IF NOT EXISTS sms_issueAffectedSoftware (
	software_id INT(11) NOT NULL,
	issue_id INT(11) NOT NULL,
	additionalInfo VARCHAR(150) DEFAULT NULL,
	confirmed BOOLEAN NOT NULL,
	PRIMARY KEY (software_id, issue_id),
	CONSTRAINT sms_issueAffectedSoftware_ibfk_1 FOREIGN KEY (software_id) REFERENCES sms_software (software_id) ON UPDATE CASCADE ON DELETE NO ACTION,
	CONSTRAINT sms_issueAffectedSoftware_ibfk_2 FOREIGN KEY (issue_id) REFERENCES sms_issue (issue_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_artefactPartOfDevice_schema = `
CREATE TABLE IF NOT EXISTS sms_artefactPartOfDevice (
	device_id INT(11) NOT NULL,
	artefact_id INT(11) NOT NULL,
	additionalInfo VARCHAR(150) DEFAULT NULL,
	PRIMARY KEY (device_id, artefact_id),
	CONSTRAINT sms_artefactPartOfDevice_ibfk_1 FOREIGN KEY (device_id) REFERENCES sms_device (device_id) ON UPDATE CASCADE ON DELETE NO ACTION,
	CONSTRAINT sms_artefactPartOfDevice_ibfk_2 FOREIGN KEY (artefact_id) REFERENCES sms_artefact (artefact_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_manufacturingOrder_schema = `
CREATE TABLE IF NOT EXISTS sms_manufacturingOrder (
	manufacturingOrder_id INT(11) NOT NULL AUTO_INCREMENT,
	system_id INT(11) NOT NULL,
	packageReference VARCHAR(100) DEFAULT NULL,
	start DATE NOT NULL,
	end DATE DEFAULT NULL,
	description VARCHAR(150) DEFAULT NULL,
	PRIMARY KEY (manufacturingOrder_id),
	CONSTRAINT sms_manufacturingOrder_ibfk_1 FOREIGN KEY (system_id) REFERENCES sms_system (system_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_certification_schema = `
CREATE TABLE IF NOT EXISTS sms_certification (
	certification_id INT(11) NOT NULL AUTO_INCREMENT,
	name VARCHAR(100) NOT NULL,
	date DATE NOT NULL,
	description VARCHAR(200) NOT NULL,
	PRIMARY KEY (certification_id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_systemHasCertification_schema = `
CREATE TABLE IF NOT EXISTS sms_systemHasCertification (
	system_id INT(11) NOT NULL,
	certification_id INT(11) NOT NULL,
	additionalInfo VARCHAR(150) DEFAULT NULL,
	PRIMARY KEY (system_id, certification_id),
	CONSTRAINT sms_systemHasCertification_ibfk_1 FOREIGN KEY (system_id) REFERENCES sms_system (system_id) ON UPDATE CASCADE ON DELETE NO ACTION,
	CONSTRAINT sms_systemHasCertification_ibfk_2 FOREIGN KEY (certification_id) REFERENCES sms_certification (certification_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_issueAffectedComponent_schema = `
CREATE TABLE IF NOT EXISTS sms_issueAffectedComponent (
    component_id INT(11) NOT NULL,
    issue_id INT(11) NOT NULL,
    additionalInfo VARCHAR(150) DEFAULT NULL,
    confirmed BOOLEAN NOT NULL,
    PRIMARY KEY (component_id, issue_id),
    CONSTRAINT sms_issueAffectedComponent_ibfk_1 FOREIGN KEY (component_id) REFERENCES sms_component (component_id) ON UPDATE CASCADE ON DELETE NO ACTION,
    CONSTRAINT sms_issueAffectedComponent_ibfk_2 FOREIGN KEY (issue_id) REFERENCES sms_issue (issue_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_issueAffectedArtefact_schema = `
CREATE TABLE IF NOT EXISTS sms_issueAffectedArtefact (
    artefact_id INT(11) NOT NULL,
    issue_id INT(11) NOT NULL,
    additionalInfo VARCHAR(150) DEFAULT NULL,
    confirmed BOOLEAN NOT NULL,
    PRIMARY KEY (artefact_id, issue_id),
    CONSTRAINT sms_issueAffectedArtefact_ibfk_1 FOREIGN KEY (artefact_id) REFERENCES sms_artefact (artefact_id) ON UPDATE CASCADE ON DELETE NO ACTION,
    CONSTRAINT sms_issueAffectedArtefact_ibfk_2 FOREIGN KEY (issue_id) REFERENCES sms_issue (issue_id) ON UPDATE CASCADE ON DELETE NO ACTION
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_securityReport_schema = `
CREATE TABLE IF NOT EXISTS sms_securityReport (
    report_id INT(11) AUTO_INCREMENT PRIMARY KEY,
    report_name VARCHAR(255) NOT NULL,
    scanner_name VARCHAR(100) NOT NULL,
    scanner_version VARCHAR(50) DEFAULT NULL,
    creation_date DATETIME NOT NULL,
    upload_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    uploaded_by VARCHAR(100) DEFAULT NULL,
    scan_scope TEXT DEFAULT NULL,
    vulnerability_count INT(11) DEFAULT 0,
    component_count INT(11) DEFAULT 0,
    report_filename VARCHAR(255) DEFAULT NULL,  -- Neue Spalte für den Dateinamen
    UNIQUE(report_name, scanner_name, creation_date)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;`

var sms_securityReportLink_schema = `
CREATE TABLE IF NOT EXISTS sms_securityReportLink (
report_id INT NOT NULL,
linked_object_id INT NOT NULL,
linked_object_type ENUM('sms_device', 'sms_software', 'sms_system') NOT NULL,
PRIMARY KEY (report_id, linked_object_id, linked_object_type),
CONSTRAINT sms_securityReportLink_ibfk_1 FOREIGN KEY (report_id) REFERENCES sms_securityReport(report_id) ON UPDATE CASCADE ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_projectSettings_schema = `
CREATE TABLE IF NOT EXISTS sms_projectSettings (
setting_id INT AUTO_INCREMENT PRIMARY KEY,
key_name VARCHAR(255) NOT NULL UNIQUE,
value_type ENUM('string', 'int', 'boolean', 'json') NOT NULL,
default_value VARCHAR(255) DEFAULT NULL
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_projectSettingsLink_schema = `
CREATE TABLE IF NOT EXISTS sms_projectSettingsLink (
project_id INT NOT NULL,
setting_id INT NOT NULL,
value VARCHAR(255) NOT NULL,
PRIMARY KEY (project_id, setting_id),
CONSTRAINT sms_projectSettingsLink_ibfk_1 FOREIGN KEY (project_id) REFERENCES sms_project(project_id) ON DELETE CASCADE,
CONSTRAINT sms_projectSettingsLink_ibfk_2 FOREIGN KEY (setting_id) REFERENCES sms_projectSettings(setting_id) ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_deviceIPDefinition_schema = `
CREATE TABLE IF NOT EXISTS sms_deviceIPDefinition (
    id INT AUTO_INCREMENT PRIMARY KEY,
    device_type_id INT NOT NULL,
    applicable_versions VARCHAR(255) NOT NULL DEFAULT 'all', -- Kommaseparierte Liste oder "all"
    ip_address VARCHAR(45) NOT NULL, -- Einzelne IP-Adresse (IPv4 oder IPv6)
    vlan_id INT DEFAULT NULL, -- VLAN-ID als freie Zahl
    description VARCHAR(255) DEFAULT NULL, -- Beschreibung der IP
    filter_condition VARCHAR(255) DEFAULT NULL, -- Bedingung für die Nutzung der IP (z. B. "IF COUNT=2" oder "IF SETTING appserver")
    CONSTRAINT fk_deviceipdefinition_deviceType FOREIGN KEY (device_type_id) REFERENCES sms_devicetype(devicetype_id) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_deviceCheckDefinition_schema = `
CREATE TABLE IF NOT EXISTS sms_deviceCheckDefinition (
    id INT AUTO_INCREMENT PRIMARY KEY,
    device_type_id INT NOT NULL,
    applicable_versions VARCHAR(255) NOT NULL DEFAULT 'all',
    test_name VARCHAR(255) NOT NULL,
    test_description TEXT NOT NULL,
    explanation TEXT DEFAULT NULL,
    expected_result TEXT NOT NULL,
    filter_condition VARCHAR(255) DEFAULT NULL,
    check_type VARCHAR(255) NOT NULL DEFAULT '',
    CONSTRAINT fk_devicecheckdefinition_deviceType FOREIGN KEY (device_type_id) 
        REFERENCES sms_devicetype(devicetype_id) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;`

var sms_update_schema = `
CREATE TABLE IF NOT EXISTS sms_update (
    update_id INT(11) NOT NULL AUTO_INCREMENT,
    from_system_id INT(11) NOT NULL,
    to_system_id INT(11) NOT NULL,
    mandatory_system_id INT(11) NOT NULL,  -- NEU: Pflichtsystemversion für das Update
    update_type ENUM('security', 'bugfix', 'feature', 'maintenance') NOT NULL DEFAULT 'bugfix',
    additional_info VARCHAR(255) DEFAULT NULL,
    is_approved BOOLEAN DEFAULT FALSE,
    external_issue_link VARCHAR(255) DEFAULT NULL,  -- NEU: Link zu einem externen Ticket-System
    project_name VARCHAR(255) DEFAULT NULL,  -- NEU: Optional, falls das Update projektspezifisch ist
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (update_id),
    CONSTRAINT sms_update_ibfk_1 FOREIGN KEY (from_system_id) REFERENCES sms_system (system_id) ON DELETE NO ACTION ON UPDATE CASCADE,
    CONSTRAINT sms_update_ibfk_2 FOREIGN KEY (to_system_id) REFERENCES sms_system (system_id) ON DELETE NO ACTION ON UPDATE CASCADE,
    CONSTRAINT sms_update_ibfk_3 FOREIGN KEY (mandatory_system_id) REFERENCES sms_system (system_id) ON DELETE NO ACTION ON UPDATE CASCADE  -- NEU: Beziehung zur Pflichtversion
) ENGINE=InnoDB DEFAULT CHARSET=latin1;`

var sms_update_package_schema = `
CREATE TABLE IF NOT EXISTS sms_update_package (
    package_id INT(11) NOT NULL AUTO_INCREMENT,
    update_id INT(11) NOT NULL,
    device_type_id INT(11) NOT NULL,
    package_identifier VARCHAR(100) NOT NULL UNIQUE,
    package_version VARCHAR(50) NOT NULL,
    package_name VARCHAR(255) NOT NULL,
    package_description TEXT DEFAULT NULL,
    update_package_file VARCHAR(255) NOT NULL,
    creator VARCHAR(255) NOT NULL,
    is_tested BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (package_id),
    CONSTRAINT sms_update_package_ibfk_1 FOREIGN KEY (update_id) REFERENCES sms_update (update_id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;`

var sms_update_center_schema = `
CREATE TABLE IF NOT EXISTS sms_update_center (
    update_center_id INT(11) NOT NULL AUTO_INCREMENT,
    project_id INT(11) NOT NULL,
    updater_id INT(11) NOT NULL,
    updater_type VARCHAR(50) NOT NULL, -- z. B. 'salt', 'ansible', 'custom'
    version VARCHAR(50) DEFAULT NULL,
    environment ENUM('staging', 'production') NOT NULL DEFAULT 'staging',
    status VARCHAR(50) DEFAULT 'active',
    description TEXT DEFAULT NULL,
    note VARCHAR(255) DEFAULT NULL,
    owner VARCHAR(255) DEFAULT NULL,
    last_contact TIMESTAMP NULL DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (update_center_id),
    CONSTRAINT sms_update_center_project_fk FOREIGN KEY (project_id) REFERENCES sms_project (project_id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_artefactPartOfDeviceInstance_schema = `
CREATE TABLE IF NOT EXISTS sms_artefactPartOfDeviceInstance (
	deviceInstance_id INT(11) NOT NULL,
	artefact_id INT(11) NOT NULL,
	additionalInfo VARCHAR(150) DEFAULT NULL,
	PRIMARY KEY (deviceInstance_id, artefact_id),
	CONSTRAINT sms_artefactPartOfDeviceInstance_ibfk_1 FOREIGN KEY (deviceInstance_id) REFERENCES sms_deviceInstance (deviceInstance_id) ON UPDATE CASCADE ON DELETE NO ACTION,
	CONSTRAINT sms_artefactPartOfDeviceInstance_ibfk_2 FOREIGN KEY (artefact_id) REFERENCES sms_artefact (artefact_id) ON UPDATE CASCADE ON DELETE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
`