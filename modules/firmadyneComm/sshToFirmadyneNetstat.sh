#!/bin/bash
echo "# Listing netstat processes"
netstat -tulpen | grep -v 127.0.0.1
 
