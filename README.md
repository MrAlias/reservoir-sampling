# Reservoir Sampling Algorithms

Comparison of reservoir sampling algorithms.

## Results

```sh
go test -bench=.

goos: linux
goarch: amd64
pkg: github.com/MrAlias/reservoir-sampling
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkSampler/R-8 	53220938	        21.60 ns/op
BenchmarkSampler/L-8 	664757622	         1.778 ns/op
PASS
```

## References

- [Reservoir Sampling Wiki](https://en.wikipedia.org/wiki/Reservoir_sampling)
- [Vitter, Jeffrey S. (1 March 1985). "Random sampling with a reservoir"](http://www.cs.umd.edu/~samir/498/vitter.pdf)
- [Java implementation version](https://richardstartin.github.io/posts/reservoir-sampling)
