# calc.Sum

Sum returns the sum of the total numbers supplied in the input argument.
The values can be individual numbers, arrays of numbers or slices.

## Syntax

```go
Sum(number1, number2, ...)
```

## Argument

### number1

Required. First input to sum.

### number2

Optional. Second input to sum.

### '...'

Optional. Subsequent numbers arrys or slices.

## Example

```Go
calc.Sum(10, 20, 30, 50) // Returns 110.0
calc.Sum(12.5, 36.2, 88, 50, []int{55, 86})  // Returns 327.7
calc.Sum(13, 4, 456.2, 25.25456, 3543, []int{13, 48}, []float64{-0.1})) // Returns 4102.35456
```
