# ~~AnyToBytes~~

A collection of functions to convert to bytes. Allows to convert anything to bytes.

### Caution!

> Project is abandoned and renamed to [x2bytes](https://github.com/koykov/x2bytes).
> This repository kept for backward compatibility in other packages.

## Usage

```go
var (
    buf []byte
    x float64 = 1234.56789
)
buf = AnyToBytes(buf, x) // []byte("1234.56789")
```

## Benchmark
```
BenchmarkAnyToBytes-8   	 1000000	      1342 ns/op	       0 B/op	       0 allocs/op
```
