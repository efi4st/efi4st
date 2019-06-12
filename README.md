# efi4st
embedded firmware inspection for security testing

Automatic Firmware Analysis of Embedded Linux Devices as Preparation for Security Testing / Fuzzing


Installation:

Web App:
- Install Golang 1.10
- go get -u github.com/kataras/iris
- go get github.com/jmoiron/sqlx
- go run main.go

Database:
- Download MariaDB 
# Silent windows msi install
- msiexec /i mariadb-10.3.15-winx64.msi SERVICENAME=MySQL PORT=3307 /qn 


Links:
- https://blog.attify.com/getting-started-with-firmware-emulation/
- http://firmware.re/
- http://check.siemens.com/
- https://github.com/avatartwo/avatar2
- https://media.ccc.de/v/34c3-9195-avatar
- http://s3.eurecom.fr/docs/bar18_muench.pdf
- https://github.com/fkie-cad/FACT_core

- https://stackoverflow.com/questions/41257847/how-to-create-singleton-db-class-in-golang