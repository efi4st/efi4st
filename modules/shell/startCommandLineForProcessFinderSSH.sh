#!/bin/bash
echo "# [start]: Search Process Command Line"
cd /home/admiral-helmut/MA/efi4st/modules/python
/proc/319/cmdline | python ../../modules/python/pipeResultsSendToServer.py $1 "PSLocalProcesses" ""
echo "# [finisched]: Search Process Command Line"

