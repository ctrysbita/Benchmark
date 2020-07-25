# Benchmark

A set of platform benchmarks based on golang's benchmark framework.

## Usage

Run test directly with go.

```bash
go test -bench=. -benchmem
```

Run pre-built executables.

```bash
./benchmark -test.bench=. -test.benchmem
```