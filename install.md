 
# efi4st
embedded firmware inspection for security testing

Automatic Firmware Analysis of Embedded Linux Devices as Preparation for Security Testing / Fuzzing


Installation:

Requirements: Git + Docker + build-essentials + dbeaver
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.23.1.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

sudo apt install git

sudo apt install docker.io

sudo add-apt-repository ppa:serge-rider/dbeaver-ce
sudo apt-get update
sudo apt-get install dbeaver-ce

Web App:
- Install Golang 1.23
- (there is a go.mod inside project, clone project execute commands in folder of go.mod)
- go get -u github.com/kataras/iris/v12
- go get github.com/jmoiron/sqlx
- go get github.com/go-sql-driver/mysql
- go run main.go

Database:
- https://computingforgeeks.com/how-to-install-mariadb-10-3-on-ubuntu-16-04-lts-xenial/
- https://websiteforstudents.com/install-and-configure-dbeaver-on-ubuntu-16-04-18-04/ 
- (use a mariadb docker container https://hub.docker.com/_/mariadb)
- sudo docker pull mariadb:10.4
- sudo docker run -p 3306:3306 -d --name maria -e MARIADB_ROOT_PASSWORD=mypassword -e MYSQL_DATABASE=efi4st -e MYSQL_USER=efi4db -e MYSQL_PASSWORD=efi4db mariadb:10.4
- Note the credentials, maybe change them here and in webapp (main.go & dbManager.go)

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
- (into firmware-analysis-toolkit folder)
- sudo ./deps.sh
- (If fails: Remove cramfs dependencies and installation function from deps.sh)
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
- sudo apt-get install ninja-build
- ../configure
- make
