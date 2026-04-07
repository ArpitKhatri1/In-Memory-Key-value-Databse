## Benchmarks
Using redis-benchmark with default loads(50 parallel clients) the following results were achieved:
1. GET


| Metric        | Value        |
|--------------|-------------|
| Throughput   | 64432.99 req/s |

| Latency (ms) | Value |
|--------------|-------|
| Avg          | 0.461 |
| Min          | 0.024 |
| P50          | 0.407 |
| P95          | 0.919 |
| P99          | 1.631 |
| Max          | 3.871 |

2. SET

| Metric        | Value        |
|--------------|-------------|
| Throughput   | 60350.03 req/s |

| Latency (ms) | Value |
|--------------|-------|
| Avg          | 0.627 |
| Min          | 0.024 |
| P50          | 0.575 |
| P95          | 1.239 |
| P99          | 1.807 |
| Max          | 3.847 |
   
These numbers are 2-3x slower than actual redis instance running on same machine which is expected due to level of complexity of the code and usage of locks, mutexes which adds extra overhead compared to single threaded lock-free redis. 
