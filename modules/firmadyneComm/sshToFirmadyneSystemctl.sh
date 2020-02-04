#!/bin/bash
echo "# Listing processes"
systemctl list-unit-files | grep running

