Supose you have a method (M) that always receives the same struct (S) as a parameter to perform a calculation with some struct receiver values (N).

Values of N could change over time but S is always the same.

We don't want to save S in N because we have lots of N representing different states in our program.

This benchmark tries to anwser what's better in terms of allocations and speed, passing struct S as pointer, as a value or keep S in N.


escape analysis
```
go build -gcflags=-m main.go
```

```
BenchmarkMethodParamAsValue-12            	18870171	        60.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkMethodParamsAsPointer-12         	22077403	        54.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkParameterInStructAsValue-12      	20417619	        59.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkParameterInStructAsPointer-12    	18334802	        63.9 ns/op	       0 B/op	       0 allocs/op
```