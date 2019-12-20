#!/bin/bash
echo "# [start]: List System Services (systemctl)"
cd /home/admiral-helmut/MA/efi4st/modules/python
systemctl list-unit-files | grep enabled | python ../../modules/python/pipeResultsSendToServer.py $1 "LocalSystemServices" ""
echo "# [finisched]:  List System Services (systemctl)"

