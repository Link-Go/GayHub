## ghz
gRPC 压测工具 ghz 

* docs: **[ghz](https://ghz.sh/)**
```txt
ghz --insecure --proto ./helloworld.proto --call helloworld.Greeter.SayHello -d '{"name": "lin"}'  0.0.0.0:50051

Summary:
  Count:        200
  Total:        53.73 ms
  Slowest:      17.14 ms
  Fastest:      0.29 ms
  Average:      7.72 ms
  Requests/sec: 3722.30

Response time histogram:
  0.287  [1]  |∎
  1.972  [9]  |∎∎∎∎∎∎∎∎
  3.657  [27] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  5.342  [17] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  7.027  [29] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  8.712  [30] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  10.397 [44] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  12.082 [26] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  13.767 [7]  |∎∎∎∎∎∎
  15.452 [4]  |∎∎∎∎
  17.137 [6]  |∎∎∎∎∎

Latency distribution:
  10 % in 2.59 ms
  25 % in 5.11 ms
  50 % in 8.07 ms
  75 % in 10.08 ms
  90 % in 11.65 ms
  95 % in 13.27 ms
  99 % in 16.83 ms

Status code distribution:
  [OK]   200 responses

```
不足：最终报告仅接口的速率。cup/内存使用情况需要使用其他的工具进行监测

开发：开发人员要自测协程，内存等情况的话，还是得使用 pprof