# protocol-benchmarks
Protocols test laboratory to see the performance in each dimensions


# REST

## 1 vus, duration 15s

| payload(bytes) | tps   |
|----------------|-------|
| 100            | 17887 |
| 1000           | 13736 |
| 10000          | 4749 |

## 10 vus, duration 15s

| payload(bytes) | tps   |
|----------------|-------|
| 100            | 61983 |
| 1000           | 49241 |
| 10000          | 14847 |


## 100 vus, duration 15s

| payload(bytes) | tps   |
|----------------|-------|
| 100            | 58536 |
| 1000           | 50352 |
| 10000          | 15754 |


# gRPC

## 1 vus, duration 15s

| payload(bytes) | tps   |
|----------------|-------|
| 100            | 29401 |
| 1000           | 27569 |
| 10000          | 25179 |


## 10 vus, duration 15s

| payload(bytes) | tps   |
|----------------|-------|
| 100            | 34643 |
| 1000           | 31121 |
| 10000          | 28631 |

## 100 vus, duration 15s

| payload(bytes) | tps   |
|----------------|-------|
| 100            | 42376 |
| 1000           | 34940 |
| 10000          | 25038 |


