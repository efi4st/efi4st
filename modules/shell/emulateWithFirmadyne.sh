#!/bin/bash
echo "# [start]: firmadyne emulation"
rm ../../../tools/firmware-analysis-toolkit/firmadyne/images/1.tar.gz
rm ../../../tools/firmware-analysis-toolkit/firmadyne/scratch/*
cd /home/admiral-helmut/MA/tools/firmware-analysis-toolkit/firmadyne
echo $2
gnome-terminal -x ./startEmu.sh $2
echo "# [finisched]: firmadyne emulation"