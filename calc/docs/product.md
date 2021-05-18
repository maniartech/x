# calc.Product

Product returns the product of all input numbers supplied in the argument.
The values can be individual numbers, arrays of numbers or slices.

## Syntax

```go
Product(number1, number2, ...)
```

## Arguments

### number1

Required. First input to Product.

### number2

Optional. Second input to Product.

### '...'

Optional. Subsequent numbers arrys or slices.

## Example

```Go
calc.Product(10) // Returns 10.0
calc.Product(10, 42)) // Returns 420.0
calc.Product(20, 5, []int{2, 10}, -4)  // Returns -8000.0
```
