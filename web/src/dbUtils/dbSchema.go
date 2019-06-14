/**
 * Author:    Admiral Helmut
 * Created:   12.06.2019
 *
 * (C)
 **/

package dbUtils

var projectSchema = `
CREATE TABLE IF NOT EXISTS projects (
	project_id INT NOT NULL AUTO_INCREMENT,
	name VARCHAR(150) NOT NULL,
	uploads INT DEFAULT NULL,
	date VARCHAR(20) DEFAULT NULL,
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
	PRIMARY KEY (firmware_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
`
