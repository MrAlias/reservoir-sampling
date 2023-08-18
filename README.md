# Reservoir Sampling Algorithms

Comparison of reservoir sampling algorithms.

## Results

```sh
go test -run='^$' -bench=. -count=10 > results.txt && benchstat results.txt
```

```
goos: linux
goarch: amd64
pkg: github.com/MrAlias/reservoir-sampling
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
              │ results.txt │
              │   sec/op    │
RNG/Int63-8     3.347n ± 2%
RNG/Float64-8   3.923n ± 5%
Sampler/R-8     22.08n ± 2%
Sampler/L-8     2.312n ± 5%
Sampler/X-8     3.338n ± 7%
Sampler/Z-8     2.372n ± 8%
geomean         4.176n
```

## References

- [Reservoir Sampling Wiki](https://en.wikipedia.org/wiki/Reservoir_sampling)
- [Vitter, Jeffrey S. (1 March 1985). "Random sampling with a reservoir"](http://www.cs.umd.edu/~samir/498/vitter.pdf)
- [Java implementation version](https://richardstartin.github.io/posts/reservoir-sampling)
