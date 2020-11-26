# efi4st
embedded firmware inspection for security testing

Automatic Firmware Analysis of Embedded Linux Devices as Preparation for Security Testing / Fuzzing


Installation:

Web App:
- Install Golang 1.11
- go get -u github.com/kataras/iris
- go get github.com/jmoiron/sqlx
- go get github.com/go-sql-driver/mysql
- go run main.go

Database:
- Linux:
- https://computingforgeeks.com/how-to-install-mariadb-10-3-on-ubuntu-16-04-lts-xenial/
- https://websiteforstudents.com/install-and-configure-dbeaver-on-ubuntu-16-04-18-04/ 
- Windows:
- Download MariaDB 
- Silent windows msi install
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
- https://www.owasp.org/index.php/IoT_Firmware_Analysis
- https://www.pentestpartners.com/security-blog/how-to-do-firmware-analysis-tools-tips-and-tricks/

Installation of firmadyne:
- git clone --recursive https://github.com/attify/firmware-analysis-toolkit.git
- delete binwalk folder and download from new source
- git clone --recursive https://github.com/ReFirmLabs/binwalk.git
- sudo ./deps.sh
- (wird nicht funktionieren: Entferne cramfs deps auch die funktion)
- dann installiere cramfs per hand: apt-get install cramfsprogs, cramsfsswap
- sudo python3 setup.py install

QEMU:
- https://opensourceforu.com/2011/05/quick-quide-to-qemu-setup/
- https://www.youtube.com/watch?v=G0NNBloGIvs
_________________________________________________-
- git clone https://github.com/qemu/qemu.git
- cd qemu
- mkdir build
- cd build
- sudo apt-get install gcc libc6-dev pkg-config bridge-utils uml-utilities zlib1g-dev libglib2.0-dev autoconf automake libtool libsdl1.2-dev
- sudo apt-get install libpixman-1-dev
- sudo apt-get install ninja
- ../configure
- make

