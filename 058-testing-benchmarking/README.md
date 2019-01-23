# Benchmarking
Go comes with capability to benchmark your code. You might want to use this to test out your different
implementations of a solution before you commit to one.

If you create a method starting with `Bench` in your `<file>_test.go` file you tell go that you want to
benchmark the code inside.

Then you can run the command `go test -bench .` to run all benchmarks
or use a regex to run selected ones. Use `go help testflag` to find out more.

**Example benchmark:**
```go
func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Mary")
	}
}
```

**Example output:**

```
$ go test -bench .
goos: darwin
goarch: amd64
pkg: github.com/clementohNZ/learn-go/058-testing-benchmarking/saying
BenchmarkGreet-16    	20000000	        98.1 ns/op
PASS
ok  	github.com/clementohNZ/learn-go/058-testing-benchmarking/saying	2.070s
```

You can see here that it took **98.1 nanoseconds** on average to run each iteration of the
code and it ran our code **20,000,000** times.

This is such a cool feature to have out of the box! It can really help you write performant
code using Big O.

## Tips
Look at the creators of go code to see what efficient code looks like. Try and write your own implementation
and see how it compares.