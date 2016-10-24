#!/bin/sh
cd /home/app
./obfuscation-cell -s=$SERVER_ADDRESS -c=$CLIENT_ADDRESS -n=$CELL_NAME -x=$SEED_VALUE -f=$FREQUENCY >> /var/log/obfuscation-cell.log 2>&1
