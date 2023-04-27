# srp - Single Responsibility Principle

This is an example project demonstrating the Single Responsibility Principle (SRP) in Go. It includes two implementations of a calculator, one following the SRP and another violating it.

## Installation

To use this package, you need to have Go installed on your machine. You can then install the package using the following command:

```shell
go get github.com/mindulle/misc/principles/SOLID/srp
```

## Usage

You can use this package by importing it in your Go code:

```go
import "github.com/mindulle/misc/principles/SOLID/srp"
```

To use the calculator, create a new instance of the `Calculator` struct and call its `Add` and `Subtract` methods:

```go
calc := srp.NewCalculator()
sum := calc.Add(2, 3) // 5
diff := calc.Subtract(5, 2) // 3
```

## Examples

See `example_test.go` for examples of how to use the `Calculator` struct.

## Tests

Run the tests using the following command:

```shell
go test -v ./...
```
