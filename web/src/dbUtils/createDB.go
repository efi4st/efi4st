/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package dbUtils

import "github.com/jmoiron/sqlx"

// exec the schema or fail; multi-statement Exec behavior varies between
// database drivers;  pq will exec them all, sqlite3 won't, ymmv
func CreateDB(db *sqlx.DB) {
	db.MustExec(projectSchema)
	db.MustExec(firmwareSchema)
	db.MustExec(relevantAppsSchema)
	db.MustExec(testResultSchema)
	db.MustExec(appContentSchema)
	db.MustExec(analysisToolSchema)
	db.MustExec(binaryAnalysisSchema)
	db.MustExec(sms_projecttype_schema) 			// SMS
	db.MustExec(sms_project_schema)     			// SMS
	db.MustExec(sms_systemtype_schema)  			// SMS
	db.MustExec(sms_system_schema)					// SMS
	db.MustExec(sms_devicetype_schema)  			// SMS
	db.MustExec(sms_device_schema)					// SMS
	db.MustExec(sms_deviceInstance_schema)			// SMS
	db.MustExec(sms_updateHistory_schema)   		// SMS
	db.MustExec(sms_issue_schema)					// SMS
	db.MustExec(sms_issueAffectedDevice_schema)		// SMS
	db.MustExec(sms_solution_schema)				// SMS
	db.MustExec(sms_artefacttype_schema)			// SMS
	db.MustExec(sms_artefact_schema)				// SMS
	db.MustExec(sms_releasenote_schema)				// SMS
	db.MustExec(sms_softwaretype_schema)			// SMS
	db.MustExec(sms_software_schema)				// SMS
	db.MustExec(sms_component_schema)				// SMS
	db.MustExec(sms_componentPartOfSoftware_schema) // SMS
	db.MustExec(sms_softwarePartOfDevice_schema)	// SMS
	db.MustExec(sms_devicePartOfSystem_schema)		// SMS
	db.MustExec(sms_projectBOM_schema)				// SMS
	db.MustExec(sms_issueAffectedSoftware_schema)	// SMS
	db.MustExec(sms_artefactPartOfDevice_schema)	// SMS
	db.MustExec(sms_manufacturingOrder_schema)		// SMS
	db.MustExec(sms_certification_schema)			// SMS
	db.MustExec(sms_systemHasCertification_schema)	// SMS
	db.MustExec(sms_issueAffectedComponent_schema)	// SMS
	db.MustExec(sms_issueAffectedArtefact_schema)	// SMS
	db.MustExec(sms_securityReport_schema)			// SMS
	db.MustExec(sms_securityReportLink_schema)		// SMS
	db.MustExec(sms_projectSettings_schema)			// SMS
	db.MustExec(sms_projectSettingsLink_schema)		// SMS
	db.MustExec(sms_deviceIPDefinition_schema)		// SMS
	db.MustExec(sms_deviceCheckDefinition_schema)	// SMS
	db.MustExec(sms_update_schema)					// SMS
	db.MustExec(sms_update_package_schema)			// SMS
	db.MustExec(sms_update_center_schema)			// SMS
	db.MustExec(sms_artefactPartOfDeviceInstance_schema) // SMS
	db.MustExec(sms_artefactPartOfSystem_schema)	// SMS
	db.MustExec(sms_project_status_log_schema)		// SMS
}

