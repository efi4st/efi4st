#!/bin/bash
echo "# [start]: default search for all executables then pipe"
cd /home/admiral-helmut/MA/efi4st/modules/python
find ../../working/filesystem/ -type f -executable -print | python ../../modules/python/pipeResultsSendToServer.py $1 "ExecutableFinder"
echo "# [finisched]: default search for all executables then pipe"

