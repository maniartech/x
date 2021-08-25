# calc.LCM

LCM function returns the least common multiple of two or more integers.
The LCM is the smallest positive integer that is a multiple of all integers in the argument.

## Syntax

```go
LCM(number1, number2 ... interface{}) int
```

## Arguments

### number1

number1 is required.

### number2, '...'

number2 and subsequent numbers are optional.

## Remark

If the number is not an integer, it is truncated.

## Example

```Go
calc.LCM(15, 12) // Returns 60
calc.LCM(4.5, 2.9)  // Returns 4
calc.LCM(0, 0, 0, 0) // Returns 0
calc.LCM(4, 16, 8, 20) // Returns 80
calc.LCM(2.468, 4.135, 6.84648, 8.45, 10.11531) // Returns 120
```
