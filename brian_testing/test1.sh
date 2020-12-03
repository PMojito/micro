#!/bin/bash

set -e  # Exit on errors.

# Truncate debugging log.
cat /dev/null > /home/sysop/.micro/logs/debug_log.txt

# Build new Micro binary
cd ../
CGO_ENABLED=0 make build

# Run Micro
echo ""
./micro -debug brian_testing/pytest.py

# Show the log file
cat /home/sysop/.micro/logs/debug_log.txt
