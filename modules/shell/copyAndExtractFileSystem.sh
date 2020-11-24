#!/bin/bash
echo "# [start]: copy and extract file system"
rm ../../working/*
cp ../../../tools/firmware-analysis-toolkit/firmadyne/images/1.tar.gz ../../working/1.tar.gz
# rm ../../../tools/firmware-analysis-toolkit/firmadyne/images/1.tar.gz
mkdir ../../working/filesystem
tar -zxvf ../../working/1.tar.gz -C ../../working/filesystem
rm ../../working/filesystem/etc/init.d/S045lighttpd.sh
cp ../../S045lighttpd.sh ../../working/filesystem/etc/init.d/S045lighttpd.sh
echo "# [finisched]: copy and extract file system"