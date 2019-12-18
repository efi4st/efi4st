#!/bin/bash
echo "# [start]: search for bash init tasks and scripts"
cd /home/admiral-helmut/MA/efi4st/modules/python
python bashInitAnalysis.py "../../working/filesystem/" | python ../../modules/python/pipeResultsSendToServer.py $1 "BashInitAnalysis" ""
echo "# [finisched]: search for bash init tasks and scripts"

