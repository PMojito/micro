#!/bin/bash

set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $SCRIPT_DIR &&
cd ../ &&
CGO_ENABLED=0 make build && \
cd brian_testing && \
cat /dev/null > log.txt && \
cat settings_test.json > ~/.config/micro/settings.json && \
../micro -debug whatever && \
cat log.txt && \
cat ~/.config/micro/settings.json
