This directory contains POC/tests related to ghz, gRPC testing & benchmarking tool/
link for installation instructions: https://github.com/bojand/ghz

By default, ghz will basically spam a gRPC call 200 times and save metrics such as latency (min, max, average) or distribution of the status codes.

# run ghz
- run `run-ghz.sh`
- wait a few seconds
- results are displayed in the terminal

Example output:

```
Summary:
  Count:        200
  Total:        7.20 s
  Slowest:      2.52 s
  Fastest:      952.72 ms
  Average:      1.70 s
  Requests/sec: 27.79

Response time histogram:
  952.722 [1]   |∎
  1109.793 [9]  |∎∎∎∎∎∎∎∎∎∎∎∎
  1266.864 [26] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  1423.935 [19] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  1581.006 [21] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  1738.077 [29] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  1895.147 [24] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  2052.218 [30] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  2209.289 [27] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  2366.360 [11] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  2523.431 [3]  |∎∎∎∎

Latency distribution:
  10 % in 1.17 s
  25 % in 1.38 s
  50 % in 1.69 s
  75 % in 2.00 s
  90 % in 2.18 s
  95 % in 2.23 s
  99 % in 2.37 s

Status code distribution:
  [OK]   200 responses
```

## using ghz-web
- run `start-ghz-web.sh`
- open a web browser and go to localhost:300
- Create a new project, name does not matter. When you clic the `save` button, the project will appear in the list.
If you clic on it, it should redirect you to `http://localhost:3000/projects/1`. Do not close the page.
- run `post-benchmark.sh` at least one time. The benchmark should take approximately 30 seconds to run.
- refresh the page that you previously opened and did not close.
    - while the page was previously empty, a graph should appear now.
    - run the benchmark script a few more times to get pretty lines and stuff.