#!/bin/bash
echo "# [start]: firmadyne emulation"
cd /home/admiral-helmut/MA/tools/firmware-analysis-toolkit/firmadyne
echo $2
echo xapioMK616 | sudo -S
go run main.go $2
echo "# [finisched]: firmadyne emulation"