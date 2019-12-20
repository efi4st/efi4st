#!/bin/bash
echo "# [start]: List local processes (PS)"
cd /home/admiral-helmut/MA/efi4st/modules/python
ps aux | grep " /" | python ../../modules/python/lineShorter.py | python ../../modules/python/pipeResultsSendToServer.py $1 "PSLocalProcesses" ""
echo "# [finisched]:  List local processes (PS)"

