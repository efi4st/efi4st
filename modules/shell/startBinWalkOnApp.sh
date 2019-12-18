#!/bin/bash
echo "# [start]: BinWalk"
cd /home/admiral-helmut/MA/efi4st/modules/python
binwalk $2 | python ../../modules/python/pipeResultsSendToServer.py $1 "BinWalkAnalysis" $2
echo "# [finisched]: BinWalk"

