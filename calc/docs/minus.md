# calc.Minus

Minus subtracts second number from the first one.

## Syntax

```go
Minus(number1, number2)
```

## Arguments

### number1

Required. The number from which you want to substract.

### number2

Required. The number which is substracted from the first number.

## Example

```Go
calc.Minus(100, 75) // Returns 25.0
calc.Minus(100.456, 75.123456)  // Returns 25.332544
calc.Minus(25, -75)  // Returns 100.0
calc.Minus(25.125, -75.125) // Returns 100.25
```
