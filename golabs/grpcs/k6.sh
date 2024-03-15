#!/bin/bash

k6 run -u 1 -d 15s --summary-export ./report_small k6-small.js
k6 run -u 1 -d 15s --summary-export ./report_medium k6-medium.js 
k6 run -u 1 -d 15s --summary-export ./report_large k6-large.js 