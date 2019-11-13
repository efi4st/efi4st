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
	firmware_id INT(11) NOT NULL,
	PRIMARY KEY (relevantApps_id),
	CONSTRAINT relevantApps_ibfk_1 FOREIGN KEY (firmware_id) REFERENCES firmware (firmware_id) ON UPDATE CASCADE ON DELETE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
`
