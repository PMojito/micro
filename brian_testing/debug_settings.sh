cd ../ && \
CGO_ENABLED=0 make build && \
cd brian_testing && \
cat /dev/null > log.txt && \
cat settings_test.json > ~/.config/micro/settings.json && \
../micro -debug whatever && \
cat log.txt && \
cat ~/.config/micro/settings.json
