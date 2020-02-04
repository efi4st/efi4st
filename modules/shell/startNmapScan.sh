#!/bin/bash
echo "# [start]: Network interface scan (nmap)"
cd /home/admiral-helmut/MA/efi4st/modules/python
sudo nmap -p 1-1000 -sV -sS 192.168.0.100 | python ../../modules/python/pipeResultsSendToServer.py $1 "NMAPNetworkInterfaces" ""
echo "# [finisched]:  Network interface scan (nmap)"

