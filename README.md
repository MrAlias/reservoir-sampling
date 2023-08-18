# Reservoir Sampling Algorithms

Comparison of reservoir sampling algorithms.

## Results

```sh
go test -bench=.

goos: linux
goarch: amd64
pkg: github.com/MrAlias/reservoir-sampling
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkSampler/BaseRandomInt64-8         	367731243	         3.259 ns/op
BenchmarkSampler/R-8                       	52917668	        21.85 ns/op
BenchmarkSampler/L-8                       	540581655	         2.224 ns/op
BenchmarkSampler/X-8                       	362520919	         3.325 ns/op
BenchmarkSampler/Z-8                       	500422533	         2.374 ns/op
PASS
```

## References

- [Reservoir Sampling Wiki](https://en.wikipedia.org/wiki/Reservoir_sampling)
- [Vitter, Jeffrey S. (1 March 1985). "Random sampling with a reservoir"](http://www.cs.umd.edu/~samir/498/vitter.pdf)
- [Java implementation version](https://richardstartin.github.io/posts/reservoir-sampling)
