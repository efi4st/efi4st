#!/bin/bash
echo "# [start]: Binary Analyis for User Mode Emulation"
cd /home/admiral-helmut/MA/efi4st/modules/python
readelf -d $2 | grep NEEDED | python ../../modules/python/pipeResultsSendToServer.py $1 "Binary4EmuAnalysisreadelf" $2
ldd $2 | python ../../modules/python/pipeResultsSendToServer.py $1 "Binary4EmuAnalysisldd" $2
strace $2 2>&1 | grep "open" | python ../../modules/python/pipeResultsSendToServer.py $1 "Binary4EmuAnalysisstrace" $2
echo "# [finisched]:  Binary Analyis for User Mode Emulation"
