#!/bin/bash
echo "# [start]: cron job search in default cron infrastructure"
cd /home/admiral-helmut/MA/efi4st/modules/python
python cronJobSearch.py "../../working/filesystem/" | python ../../modules/python/pipeResultsSendToServer.py $1 "CronJobSearch" ""
echo "# [finisched]: cron job search in default cron infrastructure"

