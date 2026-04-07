## Benchmarks
Using redis-benchmark with default loads(50 parallel clients) the following results were achieved:
1. GET
   Summary:
  throughput summary: 64432.99 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.461     0.024     0.407     0.919     1.631     3.871

2. SET
   Summary:
  throughput summary: 60350.03 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.627     0.024     0.575     1.239     1.807     3.847
   
These numbers are 2-3x slower than actual redis instance running on same machine which is expected due to level of complexity of the code and usage of locks, mutexes which adds extra overhead compared to single threaded lock-free redis. 
