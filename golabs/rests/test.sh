#!/bin/bash

strategy="$1"
if [[ "$strategy" == "v1" ]]; then
    k6 run -u 10 -d 10m --summary-export ./report_small.json k6-small.js
    k6 run -u 10 -d 10m --summary-export ./k6-medium.js
    k6 run -u 10 -d 10m --summary-export ./report_large.json k6-large.js 
elif [[ "$strategy" == "v2" ]]; then
    K6_INFLUXDB_ORGANIZATION=default \
    K6_INFLUXDB_BUCKET=rest-medium\
    K6_INFLUXDB_TOKEN=cat \
    k6v2 run -u 10 -d 10m -o xk6-influxdb=http://localhost:8086 ./k6-medium.js
else
    echo "unknow strategy"
    exit 1
fi



