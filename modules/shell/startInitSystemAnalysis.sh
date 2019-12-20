#!/bin/bash
echo "# [start]: Init System Analysis"
cd /home/admiral-helmut/MA/efi4st/modules/python
python initSystemAnalysis.py "../../working/filesystem/" | python ../../modules/python/pipeResultsSendToServer.py $1 "InitSystemAnalysis" ""
echo "# [finisched]: Init System Analysis"

