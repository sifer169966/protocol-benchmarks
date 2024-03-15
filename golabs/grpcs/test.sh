#!/bin/bash

size="$1"
request() {
local size="$1"
payload=""
if [[ "$size" == "s" ]]; then
   payload="$(cat ../small.size)"
elif [[ "$size" == "m" ]]; then
    payload="$(cat ../medium.size)"
elif [[ "$size" == "l" ]]; then
    payload="$(cat ../large.size)"
else
    echo "Invalid size"
    exit 1
fi
local con=$2
local duration="$3"

ghz --count-errors --insecure --async -c $con -z $duration --connections=1 --connect-timeout=1m \
    --proto ./service.proto \
    --call main.Test.Test \
    -d '{"name":"'$payload'"}' \
    0.0.0.0:9001


# ghz --insecure --async -c $con -z $duration --connections 1 -O influx-summary \
#     --proto ./service.proto \
#     --call main.Test.Test \
#     -d '{"name":"'$payload'"}' \
#     0.0.0.0:9001 \
#     > /Users/fer.wasin/github.com/protocol-benchmarks/golabs/grpcs/ghz_metrics.influx

# curl -i -XPOST "http://localhost:8086/api/v2/write?org=default&bucket=default" \
#  --header 'Authorization: Token cat' \
#  --data-binary "$(cat ghz_metrics.influx)"

}

request "$1" $2 "$3"




