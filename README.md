# protocol-benchmarks
Protocols test laboratory to see the performance in each dimensions


# REST

## 10 vus, duration 15s

| payload(bytes) | tps   |
|----------------|-------|
| 100            | 60179 |
| 1000           | 50274 |
| 10000          | 21582 |

## 100 vus, duration 15s

| payload(bytes) | tps   |
|----------------|-------|
| 100            | 58631 |
| 1000           | 50479 |
| 10000          | 20931 |

## 1000 vus, duration 15s
got error
```bash
WARN[0000] Request Failed                                error="Post \"http://localhost:8080/small\": read tcp 127.0.0.1:57302->127.0.0.1:8080: read: connection reset by peer"
```
