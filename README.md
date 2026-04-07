## Benchmarks
Using redis-benchmark with default loads(50 parallel clients) the following results were achieved:
1. GET
   Summary:
  throughput summary: 58927.52 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.665     0.024     0.615     1.367     1.967     3.727

2. SET
   Summary:
  throughput summary: 59241.71 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.642     0.016     0.583     1.303     1.911     3.895
   
