#!/bin/bash
echo "# [start]: List local network interfaces (netstat)"
cd /home/admiral-helmut/MA/efi4st/modules/python
sudo netstat -tulpen | grep -v 127.0.0.1 | grep LISTEN | python ../../modules/python/pipeResultsSendToServer.py $1 "NetstatLocalInterfaces" ""
echo "# [finisched]:  List local network interfaces (netstat)"

