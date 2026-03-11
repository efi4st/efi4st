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
  hardwaredesign_id INT(11) NOT NULL,
  hardwaredesign_variant_id INT(11) NOT NULL,
  orderNumber VARCHAR(80) DEFAULT NULL,
  additionalInfo VARCHAR(150) DEFAULT NULL,
  assigned_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (projectBOM_id),
  KEY idx_pbom_proj (project_id),
  KEY idx_pbom_sys (system_id),
  KEY idx_pbom_hd (hardwaredesign_id),
  KEY idx_pbom_var (hardwaredesign_variant_id),
  CONSTRAINT fk_pbom_project FOREIGN KEY (project_id) REFERENCES sms_project (project_id) ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT fk_pbom_system FOREIGN KEY (system_id) REFERENCES sms_system (system_id) ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT fk_pbom_sys_hd FOREIGN KEY (system_id, hardwaredesign_id) REFERENCES sms_hardwaredesignPartOfSystem (system_id, hardwaredesign_id) ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT fk_pbom_var_hd FOREIGN KEY (hardwaredesign_variant_id, hardwaredesign_id) REFERENCES sms_hardwaredesign_variant (hardwaredesign_variant_id, hardwaredesign_id) ON UPDATE CASCADE ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
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

var sms_artefactPartOfSystem_schema = `
CREATE TABLE IF NOT EXISTS sms_artefactPartOfSystem (
system_id INT(11) NOT NULL,
artefact_id INT(11) NOT NULL,
additionalInfo VARCHAR(150) DEFAULT NULL,
PRIMARY KEY (system_id, artefact_id),
CONSTRAINT sms_artefactPartOfSystem_ibfk_1 FOREIGN KEY (system_id) REFERENCES sms_system (system_id) ON UPDATE CASCADE ON DELETE NO ACTION,
CONSTRAINT sms_artefactPartOfSystem_ibfk_2 FOREIGN KEY (artefact_id) REFERENCES sms_artefact (artefact_id) ON UPDATE CASCADE ON DELETE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_project_status_log_schema = `
CREATE TABLE IF NOT EXISTS sms_project_status_log (
    status_id INT(11) NOT NULL AUTO_INCREMENT,
    project_id INT(11) NOT NULL,
    status ENUM('ordered', 'in_construction', 'in_commission', 'active') NOT NULL,
    note TEXT DEFAULT NULL,
    access_group VARCHAR(100) DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (status_id),
    CONSTRAINT sms_project_status_log_ibfk_1 FOREIGN KEY (project_id) REFERENCES sms_project (project_id) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_sms_elementSearch_schema = `
CREATE OR REPLACE VIEW sms_elementSearch AS
(
SELECT
'Artefact' AS entity_type,
a.artefact_id AS entity_id,
a.name AS name,
a.version AS version,
at.artefactType AS type,
IFNULL(GROUP_CONCAT(DISTINCT CONCAT(stype.type, ' ', sys.version) SEPARATOR ', '), '–') AS systems
FROM sms_artefact a
JOIN sms_artefacttype at ON a.artefacttype_id = at.artefacttype_id
LEFT JOIN sms_artefactPartOfDevice apd ON a.artefact_id = apd.artefact_id
LEFT JOIN sms_devicePartOfSystem dps ON apd.device_id = dps.device_id
LEFT JOIN sms_system sys ON dps.system_id = sys.system_id
LEFT JOIN sms_systemtype stype ON sys.systemtype_id = stype.systemtype_id
GROUP BY a.artefact_id, a.name, a.version, at.artefactType
)
UNION
(
SELECT
'Software' AS entity_type,
sft.software_id AS entity_id,
st.typeName AS name,               -- ✅ Softwaretyp als Name (z. B. "Firmware")
sft.version AS version,
sft.license AS type,               -- ✅ z. B. Lizenztyp als Kategorie (optional)
IFNULL(GROUP_CONCAT(DISTINCT CONCAT(stype.type, ' ', sys.version) SEPARATOR ', '), '–') AS systems
FROM sms_software sft
JOIN sms_softwaretype st ON sft.softwaretype_id = st.softwaretype_id
LEFT JOIN sms_softwarePartOfDevice spd ON sft.software_id = spd.software_id
LEFT JOIN sms_devicePartOfSystem dps ON spd.device_id = dps.device_id
LEFT JOIN sms_system sys ON dps.system_id = sys.system_id
LEFT JOIN sms_systemtype stype ON sys.systemtype_id = stype.systemtype_id
GROUP BY sft.software_id, sft.version, st.typeName, sft.license
)
UNION
(
SELECT
'Component' AS entity_type,
c.component_id AS entity_id,
c.name AS name,
c.version AS version,
c.componentType AS type,
IFNULL(GROUP_CONCAT(DISTINCT CONCAT(stype.type, ' ', sys.version) SEPARATOR ', '), '–') AS systems
FROM sms_component c
LEFT JOIN sms_componentPartOfSoftware cps ON c.component_id = cps.component_id
LEFT JOIN sms_software sft ON cps.software_id = sft.software_id
LEFT JOIN sms_softwarePartOfDevice spd ON sft.software_id = spd.software_id
LEFT JOIN sms_devicePartOfSystem dps ON spd.device_id = dps.device_id
LEFT JOIN sms_system sys ON dps.system_id = sys.system_id
LEFT JOIN sms_systemtype stype ON sys.systemtype_id = stype.systemtype_id
GROUP BY c.component_id, c.name, c.version, c.componentType
);
`

var sms_hardwaredesign_schema = `
CREATE TABLE IF NOT EXISTS sms_hardwaredesign (
hardwaredesign_id INT(11) NOT NULL AUTO_INCREMENT,
name VARCHAR(150) NOT NULL,
version VARCHAR(80) NOT NULL,
date DATE NOT NULL,
description TEXT DEFAULT NULL,
image BLOB DEFAULT NULL,
author VARCHAR(150) DEFAULT NULL,
isApproved BOOLEAN DEFAULT FALSE,
revision_note TEXT DEFAULT NULL,
document_number VARCHAR(100) DEFAULT NULL,
PRIMARY KEY (hardwaredesign_id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_hardwaredesignPartOfSystem_schema = `
CREATE TABLE IF NOT EXISTS sms_hardwaredesignPartOfSystem (
system_id INT(11) NOT NULL,
hardwaredesign_id INT(11) NOT NULL,
additionalInfo VARCHAR(150) DEFAULT NULL,
is_default BOOLEAN NOT NULL DEFAULT FALSE,
compatibility_status ENUM('recommended','compatible','deprecated') NOT NULL DEFAULT 'compatible',
default_system_id INT(11) GENERATED ALWAYS AS (IF(is_default, system_id, NULL)) STORED,
PRIMARY KEY (system_id, hardwaredesign_id),
UNIQUE KEY uq_one_default_per_system (default_system_id),
KEY idx_hwps_status (system_id, compatibility_status),
KEY idx_hwps_hw (hardwaredesign_id),
CONSTRAINT fk_hwdesign_system FOREIGN KEY (system_id) REFERENCES sms_system (system_id) ON UPDATE CASCADE ON DELETE NO ACTION,
CONSTRAINT fk_hwdesign_design FOREIGN KEY (hardwaredesign_id) REFERENCES sms_hardwaredesign (hardwaredesign_id) ON UPDATE CASCADE ON DELETE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
`


var sms_hardwaredesign_variant_schema = `
CREATE TABLE IF NOT EXISTS sms_hardwaredesign_variant (
hardwaredesign_variant_id INT(11) NOT NULL AUTO_INCREMENT,
hardwaredesign_id INT(11) NOT NULL,
code VARCHAR(50) NOT NULL,
name VARCHAR(150) NOT NULL,
description TEXT DEFAULT NULL,
spec JSON DEFAULT NULL,           -- optional: strukturierte Details (z.B. {"pcs":2,"switches":1})
is_active BOOLEAN NOT NULL DEFAULT TRUE,
created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (hardwaredesign_variant_id),
CONSTRAINT uq_hdvar UNIQUE (hardwaredesign_id, code),
UNIQUE KEY uq_var_pk_hd (hardwaredesign_variant_id, hardwaredesign_id),
KEY idx_hdvar_hd (hardwaredesign_id),
CONSTRAINT fk_hdvar_hd FOREIGN KEY (hardwaredesign_id) REFERENCES sms_hardwaredesign (hardwaredesign_id) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_checklistTemplate_schema = `
CREATE TABLE IF NOT EXISTS sms_checklistTemplate (
    checklistTemplate_id INT(11) NOT NULL AUTO_INCREMENT,
    name VARCHAR(150) NOT NULL,
    description TEXT DEFAULT NULL,
    PRIMARY KEY (checklistTemplate_id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_checklistTemplateItem_schema = `
CREATE TABLE IF NOT EXISTS sms_checklistTemplateItem (
    checklistTemplateItem_id INT(11) NOT NULL AUTO_INCREMENT,
    checklistTemplate_id INT(11) NOT NULL,
    checkDefinition_id INT(11) DEFAULT NULL,
    artefacttype_id INT(11) DEFAULT NULL,
    targetScope ENUM('system', 'device', 'deviceInstance') NOT NULL,
    expected_value VARCHAR(150) DEFAULT NULL,
    optional BOOLEAN DEFAULT FALSE,
    PRIMARY KEY (checklistTemplateItem_id),
    CONSTRAINT fk_checklist_template FOREIGN KEY (checklistTemplate_id)
      REFERENCES sms_checklistTemplate (checklistTemplate_id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_checklist_checkDefinition FOREIGN KEY (checkDefinition_id)
      REFERENCES sms_deviceCheckDefinition (id) ON UPDATE CASCADE ON DELETE SET NULL,
    CONSTRAINT fk_checklist_artefacttype FOREIGN KEY (artefacttype_id)
      REFERENCES sms_artefacttype (artefacttype_id) ON UPDATE CASCADE ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_checklistInstance_schema = `
CREATE TABLE IF NOT EXISTS sms_checklistInstance (
    checklistInstance_id INT AUTO_INCREMENT PRIMARY KEY,
    checklistTemplate_id INT NOT NULL,
    project_id INT DEFAULT NULL,
    device_id INT DEFAULT NULL,
    generated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    generated_by VARCHAR(150) NOT NULL,
    note TEXT DEFAULT NULL,
    status ENUM('open','in_progress','done') DEFAULT 'open',
    FOREIGN KEY (checklistTemplate_id) REFERENCES sms_checklistTemplate(checklistTemplate_id),
    FOREIGN KEY (project_id) REFERENCES sms_project(project_id),
    FOREIGN KEY (device_id) REFERENCES sms_device(device_id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_checklistItemInstance_schema = `
CREATE TABLE IF NOT EXISTS sms_checklistItemInstance (
    checklistItemInstance_id INT(11) NOT NULL AUTO_INCREMENT,
    checklistInstance_id INT(11) NOT NULL,
    checklistTemplateItem_id INT(11) NOT NULL,
    target_object_id INT(11) NOT NULL, -- device_id, deviceInstance_id, system_id (je nach targetScope)
    target_object_type ENUM('system', 'device', 'deviceInstance') NOT NULL,
    is_ok BOOLEAN DEFAULT NULL,
    actual_value VARCHAR(150) DEFAULT NULL,
    comment TEXT DEFAULT NULL,
	expected_value VARCHAR(400) DEFAULT NULL,
    PRIMARY KEY (checklistItemInstance_id),
    CONSTRAINT fk_checklist_inst FOREIGN KEY (checklistInstance_id)
      REFERENCES sms_checklistInstance (checklistInstance_id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_checklist_itemtemplate FOREIGN KEY (checklistTemplateItem_id)
      REFERENCES sms_checklistTemplateItem (checklistTemplateItem_id) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_checklistTemplateDocAsset_schema = `
CREATE TABLE IF NOT EXISTS sms_checklistTemplateDocAsset (
  docAsset_id INT AUTO_INCREMENT PRIMARY KEY,
  checklistTemplate_id INT NOT NULL,
  kind ENUM('cover','footer') NOT NULL,
  mime ENUM('html','md') NOT NULL DEFAULT 'html',
  storage ENUM('file','db') NOT NULL DEFAULT 'file',
  file_path VARCHAR(500) DEFAULT NULL,
  content LONGBLOB DEFAULT NULL,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY uq_template_kind (checklistTemplate_id, kind),
  CONSTRAINT fk_docasset_template
    FOREIGN KEY (checklistTemplate_id)
    REFERENCES sms_checklistTemplate(checklistTemplate_id)
    ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_deviceInstancePartOfProjectBOM_schema = `
CREATE TABLE IF NOT EXISTS sms_deviceInstancePartOfProjectBOM (
projectBOM_id      INT(11) NOT NULL,
deviceInstance_id  INT(11) NOT NULL,
additionalInfo     VARCHAR(150) DEFAULT NULL,
assigned_at        TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (projectBOM_id, deviceInstance_id),
UNIQUE KEY uq_deviceInstance_single_pbom (deviceInstance_id),
KEY idx_dipb_pbom (projectBOM_id),
KEY idx_dipb_di (deviceInstance_id),
CONSTRAINT fk_dipb_pbom FOREIGN KEY (projectBOM_id)
REFERENCES sms_projectBOM (projectBOM_id)
ON UPDATE CASCADE ON DELETE CASCADE,
CONSTRAINT fk_dipb_deviceInstance FOREIGN KEY (deviceInstance_id)
REFERENCES sms_deviceInstance (deviceInstance_id)
ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_update_execution_schema = `
CREATE TABLE IF NOT EXISTS sms_update_execution (
execution_id INT(11) NOT NULL AUTO_INCREMENT,

-- Wer/was hat das Update ausgelöst/ausgeführt
update_center_id INT(11) NOT NULL,

-- Ziel der Ausführung (konkrete Feld-Instanz)
deviceInstance_id INT(11) NOT NULL,

-- Optional: Verweis auf den "Update-Pfad" (Definition)
update_id INT(11) DEFAULT NULL,

-- Optional: Verwendetes Paket (falls eindeutig)
package_id INT(11) DEFAULT NULL,

-- Denormalisiert für Audit/History (damit spätere Änderungen an sms_update nicht verfälschen)
from_system_id INT(11) NOT NULL,
to_system_id   INT(11) NOT NULL,

-- Status + Metadaten
status ENUM('queued','running','success','failed','rolled_back','canceled') NOT NULL DEFAULT 'queued',
triggered_by VARCHAR(150) DEFAULT NULL,          -- User/Service
source ENUM('manual','scheduled','auto') NOT NULL DEFAULT 'manual',

-- Zeitpunkte
queued_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
started_at  DATETIME DEFAULT NULL,
finished_at DATETIME DEFAULT NULL,

-- Ergebnisdetails
exit_code INT(11) DEFAULT NULL,
message VARCHAR(255) DEFAULT NULL,
log TEXT DEFAULT NULL,                            -- oder später in File/Artefact auslagern

-- Idempotency / Korrelation (praktisch für Updater)
correlation_id VARCHAR(100) DEFAULT NULL,

PRIMARY KEY (execution_id),

-- Indizes (wichtig für Projekt-Doku / Timeline)
KEY idx_exec_deviceInstance (deviceInstance_id, queued_at),
KEY idx_exec_center (update_center_id, queued_at),
KEY idx_exec_status (status, queued_at),
KEY idx_exec_from_to (from_system_id, to_system_id),
KEY idx_exec_corr (correlation_id),

CONSTRAINT sms_update_execution_ibfk_1
FOREIGN KEY (update_center_id)
REFERENCES sms_update_center (update_center_id)
ON UPDATE CASCADE ON DELETE CASCADE,

CONSTRAINT sms_update_execution_ibfk_2
FOREIGN KEY (deviceInstance_id)
REFERENCES sms_deviceInstance (deviceInstance_id)
ON UPDATE CASCADE ON DELETE CASCADE,

CONSTRAINT sms_update_execution_ibfk_3
FOREIGN KEY (update_id)
REFERENCES sms_update (update_id)
ON UPDATE CASCADE ON DELETE SET NULL,

CONSTRAINT sms_update_execution_ibfk_4
FOREIGN KEY (package_id)
REFERENCES sms_update_package (package_id)
ON UPDATE CASCADE ON DELETE SET NULL,

CONSTRAINT sms_update_execution_ibfk_5
FOREIGN KEY (from_system_id)
REFERENCES sms_system (system_id)
ON UPDATE CASCADE ON DELETE NO ACTION,

CONSTRAINT sms_update_execution_ibfk_6
FOREIGN KEY (to_system_id)
REFERENCES sms_system (system_id)
ON UPDATE CASCADE ON DELETE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_projectDocEntry_schema = `
CREATE TABLE IF NOT EXISTS sms_projectDocEntry (
entry_id INT(11) NOT NULL AUTO_INCREMENT,

project_id INT(11) NOT NULL,

-- Inhalt
title VARCHAR(150) DEFAULT NULL,
body TEXT DEFAULT NULL,

-- Klassifizierung
entry_type ENUM('note','event','maintenance','incident','update','photo','attachment') NOT NULL DEFAULT 'note',

-- Wer / wann
created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
created_by VARCHAR(150) NOT NULL,

-- Optional: Sichtbarkeit/Organisation (passt zu deinem access_group Konzept)
access_group VARCHAR(100) DEFAULT NULL,

-- Optional: Verlinkung auf “Quelle”, damit du auch fremde Logs einsortieren kannst
-- Beispiel: source_type='deviceInstance' + source_id=deviceInstance_id
source_type ENUM('project','deviceInstance','updateHistory','updateExecution') NOT NULL DEFAULT 'project',
source_id INT(11) DEFAULT NULL,

-- Optional: “Eventzeit” falls abweichend vom Erfassungszeitpunkt (z.B. Foto aufgenommen vorher)
event_time DATETIME DEFAULT NULL,

PRIMARY KEY (entry_id),

KEY idx_doc_proj_time (project_id, created_at),
KEY idx_doc_proj_eventtime (project_id, event_time),
KEY idx_doc_source (source_type, source_id),

CONSTRAINT sms_projectDocEntry_ibfk_1
FOREIGN KEY (project_id)
REFERENCES sms_project (project_id)
ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_projectDocAsset_schema = `
CREATE TABLE IF NOT EXISTS sms_projectDocAsset (
asset_id INT(11) NOT NULL AUTO_INCREMENT,
entry_id INT(11) NOT NULL,

kind ENUM('image','file') NOT NULL DEFAULT 'image',

-- Storage-Strategie
storage ENUM('file','db') NOT NULL DEFAULT 'file',

-- Metadaten
mime VARCHAR(100) NOT NULL,              -- z.B. image/jpeg, image/png
original_filename VARCHAR(255) DEFAULT NULL,
stored_filename VARCHAR(255) DEFAULT NULL,  -- falls du serverseitig umbenennst
file_path VARCHAR(500) DEFAULT NULL,        -- wenn storage='file'
file_size INT(11) DEFAULT NULL,
sha256 CHAR(64) DEFAULT NULL,

-- Bildoptionen (optional)
width INT(11) DEFAULT NULL,
height INT(11) DEFAULT NULL,

-- Content nur falls storage='db'
content LONGBLOB DEFAULT NULL,

created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
created_by VARCHAR(150) DEFAULT NULL,

PRIMARY KEY (asset_id),

KEY idx_asset_entry (entry_id),
KEY idx_asset_sha (sha256),

CONSTRAINT sms_projectDocAsset_ibfk_1
FOREIGN KEY (entry_id)
REFERENCES sms_projectDocEntry (entry_id)
ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_projectTimeline_schema = `
CREATE OR REPLACE VIEW sms_projectTimeline AS
(
    SELECT
        e.project_id                                       AS project_id,
        _latin1'doc' COLLATE latin1_swedish_ci             AS source,
        e.entry_id                                         AS source_id,
        COALESCE(e.event_time, e.created_at)               AS occurred_at,
        e.created_at                                       AS created_at,
        e.created_by                                       AS actor,
        e.entry_type                                       AS entry_type,
        e.title                                            AS title,
        e.body                                             AS body,
        e.access_group                                     AS access_group,

        CASE WHEN e.source_type = 'deviceInstance' THEN e.source_id ELSE NULL END AS deviceInstance_id,
        NULL                                               AS update_center_id,
        NULL                                               AS update_id,
        NULL                                               AS package_id,
        NULL                                               AS from_system_id,
        NULL                                               AS to_system_id,
        NULL                                               AS exec_status,
        NULL                                               AS exec_exit_code
    FROM sms_projectDocEntry e
)
UNION ALL
(
    SELECT
        uc.project_id                                      AS project_id,
        _latin1'updateExecution' COLLATE latin1_swedish_ci  AS source,
        ex.execution_id                                    AS source_id,
        COALESCE(ex.finished_at, ex.started_at, ex.queued_at) AS occurred_at,
        ex.queued_at                                       AS created_at,
        ex.triggered_by                                    AS actor,
        _latin1'update' COLLATE latin1_swedish_ci           AS entry_type,
        NULL                                               AS title,
        ex.message                                         AS body,
        NULL                                               AS access_group,

        ex.deviceInstance_id                               AS deviceInstance_id,
        ex.update_center_id                                AS update_center_id,
        ex.update_id                                       AS update_id,
        ex.package_id                                      AS package_id,
        ex.from_system_id                                  AS from_system_id,
        ex.to_system_id                                    AS to_system_id,
        ex.status                                          AS exec_status,
        ex.exit_code                                       AS exec_exit_code
    FROM sms_update_execution ex
    JOIN sms_update_center uc ON ex.update_center_id = uc.update_center_id
)
UNION ALL
(
    SELECT
        di.project_id                                      AS project_id,
        _latin1'updateHistory' COLLATE latin1_swedish_ci    AS source,
        uh.updateHistory_id                                AS source_id,
        CAST(uh.date AS DATETIME)                          AS occurred_at,
        CAST(uh.date AS DATETIME)                          AS created_at,
        uh.user                                            AS actor,
        _latin1'log' COLLATE latin1_swedish_ci              AS entry_type,
        uh.updateType                                      AS title,
        uh.description                                     AS body,
        NULL                                               AS access_group,

        uh.deviceInstance_id                               AS deviceInstance_id,
        NULL                                               AS update_center_id,
        NULL                                               AS update_id,
        NULL                                               AS package_id,
        NULL                                               AS from_system_id,
        NULL                                               AS to_system_id,
        NULL                                               AS exec_status,
        NULL                                               AS exec_exit_code
    FROM sms_updateHistory uh
    JOIN sms_deviceInstance di ON uh.deviceInstance_id = di.deviceInstance_id
);
`

var sms_projectTimelineWithDeviceInfo_schema = `
CREATE OR REPLACE VIEW sms_projectTimelineWithDeviceInfo AS
SELECT
tl.project_id,
tl.source,
tl.source_id,
tl.occurred_at,
tl.created_at,
tl.actor,
tl.entry_type,
tl.title,
tl.body,
tl.access_group,

tl.deviceInstance_id,
tl.update_center_id,
tl.update_id,
tl.package_id,
tl.from_system_id,
tl.to_system_id,
tl.exec_status,
tl.exec_exit_code,

-- DeviceInstance Kontext
di.serialnumber                         AS deviceInstance_serialnumber,
di.provisioner                          AS deviceInstance_provisioner,
di.configuration                        AS deviceInstance_configuration,
di.device_id                            AS device_id,

-- Device Kontext
dt.type                                 AS device_type,
d.version                               AS device_version,

-- System Kontext (from/to)
stf.type                                AS from_system_type,
sysf.version                            AS from_system_version,

stt.type                                AS to_system_type,
syst.version                            AS to_system_version

FROM sms_projectTimeline tl
LEFT JOIN sms_deviceInstance di
ON tl.deviceInstance_id = di.deviceInstance_id
LEFT JOIN sms_device d
ON di.device_id = d.device_id
LEFT JOIN sms_devicetype dt
ON d.devicetype_id = dt.devicetype_id

LEFT JOIN sms_system sysf
ON tl.from_system_id = sysf.system_id
LEFT JOIN sms_systemtype stf
ON sysf.systemtype_id = stf.systemtype_id

LEFT JOIN sms_system syst
ON tl.to_system_id = syst.system_id
LEFT JOIN sms_systemtype stt
ON syst.systemtype_id = stt.systemtype_id;
`

var sms_projectTimelinePretty_schema = `
CREATE OR REPLACE VIEW sms_projectTimelinePretty AS
SELECT
  tl.project_id,
  tl.source,
  tl.source_id,
  tl.occurred_at,
  tl.created_at,
  tl.actor,
  tl.entry_type,

  CASE
    WHEN tl.source = 'doc' THEN
      COALESCE(
        NULLIF(tl.title,''),
        CONCAT(_latin1'Doku: ' COLLATE latin1_swedish_ci, tl.entry_type)
      )

    WHEN tl.source = 'updateExecution' THEN
      CONCAT(
        _latin1'Update: ' COLLATE latin1_swedish_ci,
        COALESCE(CONCAT(stf.type, _latin1' ' COLLATE latin1_swedish_ci, sysf.version), CONCAT(_latin1'SYSTEM#' COLLATE latin1_swedish_ci, tl.from_system_id)),
        _latin1' -> ' COLLATE latin1_swedish_ci,
        COALESCE(CONCAT(stt.type, _latin1' ' COLLATE latin1_swedish_ci, syst.version), CONCAT(_latin1'SYSTEM#' COLLATE latin1_swedish_ci, tl.to_system_id)),
        _latin1' (' COLLATE latin1_swedish_ci, tl.exec_status, _latin1')' COLLATE latin1_swedish_ci
      )

    WHEN tl.source = 'updateHistory' THEN
      CONCAT(
        _latin1'DILog: ' COLLATE latin1_swedish_ci,
        COALESCE(NULLIF(tl.title,''), _latin1'Eintrag' COLLATE latin1_swedish_ci)
      )

    ELSE
      COALESCE(NULLIF(tl.title,''), _latin1'Timeline Event' COLLATE latin1_swedish_ci)
  END AS title,

  CASE
    WHEN tl.source = 'doc' THEN
      tl.body

    WHEN tl.source = 'updateExecution' THEN
      CONCAT(
        _latin1'DI: ' COLLATE latin1_swedish_ci,
        COALESCE(di.serialnumber, CONCAT(_latin1'#' COLLATE latin1_swedish_ci, tl.deviceInstance_id)),
        IF(dt.type IS NOT NULL OR d.version IS NOT NULL,
           CONCAT(
             _latin1' (' COLLATE latin1_swedish_ci,
             COALESCE(dt.type, _latin1'' COLLATE latin1_swedish_ci),
             IF(d.version IS NOT NULL, CONCAT(_latin1' ' COLLATE latin1_swedish_ci, d.version), _latin1'' COLLATE latin1_swedish_ci),
             _latin1')' COLLATE latin1_swedish_ci
           ),
           _latin1'' COLLATE latin1_swedish_ci
        ),
        IF(tl.body IS NOT NULL AND tl.body <> '', CONCAT(_latin1'\n' COLLATE latin1_swedish_ci, tl.body), _latin1'' COLLATE latin1_swedish_ci)
      )

    WHEN tl.source = 'updateHistory' THEN
      CONCAT(
        _latin1'DILog DI: ' COLLATE latin1_swedish_ci,
        COALESCE(di.serialnumber, CONCAT(_latin1'#' COLLATE latin1_swedish_ci, tl.deviceInstance_id)),
        IF(tl.body IS NOT NULL AND tl.body <> '', CONCAT(_latin1'\n' COLLATE latin1_swedish_ci, tl.body), _latin1'' COLLATE latin1_swedish_ci)
      )

    ELSE
      tl.body
  END AS body,

  tl.access_group,

  tl.deviceInstance_id,
  tl.update_center_id,
  tl.update_id,
  tl.package_id,
  tl.from_system_id,
  tl.to_system_id,
  tl.exec_status,
  tl.exec_exit_code,

  di.serialnumber AS deviceInstance_serialnumber,
  dt.type         AS device_type,
  d.version       AS device_version,

  stf.type        AS from_system_type,
  sysf.version    AS from_system_version,
  stt.type        AS to_system_type,
  syst.version    AS to_system_version

FROM sms_projectTimeline tl
LEFT JOIN sms_deviceInstance di ON tl.deviceInstance_id = di.deviceInstance_id
LEFT JOIN sms_device d          ON di.device_id = d.device_id
LEFT JOIN sms_devicetype dt     ON d.devicetype_id = dt.devicetype_id
LEFT JOIN sms_system sysf       ON tl.from_system_id = sysf.system_id
LEFT JOIN sms_systemtype stf    ON sysf.systemtype_id = stf.systemtype_id
LEFT JOIN sms_system syst       ON tl.to_system_id = syst.system_id
LEFT JOIN sms_systemtype stt    ON syst.systemtype_id = stt.systemtype_id;
`

var sms_liveReport_schema = `
CREATE TABLE IF NOT EXISTS sms_liveReport (
report_id INT AUTO_INCREMENT PRIMARY KEY,

project_id INT(11) NOT NULL,
update_center_id INT(11) DEFAULT NULL,
projectBOM_id INT(11) DEFAULT NULL,
system_id INT(11) DEFAULT NULL,

source ENUM('upload','update_center','api') NOT NULL DEFAULT 'upload',
report_name VARCHAR(255) DEFAULT NULL,

created_at DATETIME NOT NULL,
received_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
received_by VARCHAR(150) DEFAULT NULL,

schema_version VARCHAR(50) NOT NULL DEFAULT '1',
report_format VARCHAR(50) NOT NULL DEFAULT 'device_software_v1',

payload_json JSON NOT NULL,
payload_sha256 CHAR(64) DEFAULT NULL,
payload_size INT DEFAULT NULL,

note VARCHAR(255) DEFAULT NULL,

KEY idx_lr_project (project_id, received_at),
KEY idx_lr_uc (update_center_id, received_at),
KEY idx_lr_system (system_id, received_at),

CONSTRAINT fk_lr_project FOREIGN KEY (project_id)
REFERENCES sms_project(project_id) ON DELETE CASCADE ON UPDATE CASCADE,

CONSTRAINT fk_lr_uc FOREIGN KEY (update_center_id)
REFERENCES sms_update_center(update_center_id) ON DELETE SET NULL ON UPDATE CASCADE,

CONSTRAINT fk_lr_pbom FOREIGN KEY (projectBOM_id)
REFERENCES sms_projectBOM(projectBOM_id) ON DELETE SET NULL ON UPDATE CASCADE,

CONSTRAINT fk_lr_system FOREIGN KEY (system_id)
REFERENCES sms_system(system_id) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
`

var sms_liveReport_item_schema = `
CREATE TABLE IF NOT EXISTS sms_liveReport_item (
item_id INT AUTO_INCREMENT PRIMARY KEY,

report_id INT NOT NULL,
project_id INT NOT NULL,

deviceInstance_id INT DEFAULT NULL,
serialnumber VARCHAR(150) NOT NULL,

live_device_type VARCHAR(150) NOT NULL,
live_device_version VARCHAR(80) NOT NULL,

match_status ENUM(
'unknown_instance',
'unknown_device_version',
'match_device_only',
'match_strict',
'match_partial',
'no_match'
) NOT NULL,

matched_device_id INT DEFAULT NULL,
mismatch_summary VARCHAR(255) DEFAULT NULL,

created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

KEY idx_lri_report (report_id),
KEY idx_lri_project (project_id),
KEY idx_lri_serial (project_id, serialnumber),

CONSTRAINT fk_lri_report FOREIGN KEY (report_id)
REFERENCES sms_liveReport(report_id) ON DELETE CASCADE ON UPDATE CASCADE,

CONSTRAINT fk_lri_project FOREIGN KEY (project_id)
REFERENCES sms_project(project_id) ON DELETE CASCADE ON UPDATE CASCADE,

CONSTRAINT fk_lri_di FOREIGN KEY (deviceInstance_id)
REFERENCES sms_deviceInstance(deviceInstance_id) ON DELETE SET NULL ON UPDATE CASCADE,

CONSTRAINT fk_lri_device FOREIGN KEY (matched_device_id)
REFERENCES sms_device(device_id) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
`