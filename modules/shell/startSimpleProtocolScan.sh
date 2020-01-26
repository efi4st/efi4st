#!/bin/bash
echo "# [start]: Simple HTTP Test"
cd /home/admiral-helmut/MA/efi4st/modules/python
curl -i -X OPTIONS 192.168.0.100:80 | python ../../modules/python/pipeResultsSendToServer.py $1 "SimpleHTTPTest" ""
echo "# [finisched]: Simple HTTP Test"

