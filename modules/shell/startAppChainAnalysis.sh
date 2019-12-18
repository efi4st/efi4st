#!/bin/bash
echo "# [start]: search for start of other apps in app (chain)"
cd /home/admiral-helmut/MA/efi4st/modules/python
strings $2 | python ../../modules/python/appChainAnalysis.py $2 | python ../../modules/python/pipeResultsSendToServer.py $1 "AppChainAnalysis" $2
echo "# [finisched]: search for start of other apps in app (chain)"
