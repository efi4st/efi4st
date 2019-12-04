#!/bin/bash
echo "# [start]: default search for all executables then pipe"
cd /home/admiral-helmut/MA/efi4st/modules/python
python cronJobSearch.py "../../working/filesystem/" | python ../../modules/python/pipeResultsSendToServer.py $1 "CronJobSearch"
echo "# [finisched]: default search for all executables then pipe"

