# Gossert

[Godoc](http://pkg.go.dev/github.com/fmstephe/gossert)

This package exports three convenience functions for performing runtime assertions. The functions each take a func as an argument.

By default the asserts are not run, but you can enable the asserts by including '-tags gossert' when compiling your Go program. The intention behind this design is to allow the compiler to _completely_ eliminate assert code when asserts are not being run.
