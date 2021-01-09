Supose you have a method (M) that always receives the same struct (S) as a parameter to perform a calculation with some struct receiver values (N).

Values of N could change over time but S is always the same.

We don't want to save S in N because we have lots of N representing different states in our program.

This benchmark tries to anwser what's better in terms of allocations and speed, passing struct S as pointer or as a value.


escape analysis
go build -gcflags=-m func_params.go