# calc.GCD

GCD function Returns the greatest common divisor of two or more non-zero integers.
GCD divides the given numbers without a remainder.

## Syntax

```go
GCD(number1, number2 ... interface{}) int
```

## Arguments

## number1

number1 is required.

### number2, '...'

number2 and subsequent numbers are optional.

## Remark

+ All the values are evenly divided by 1
+ A Prime number can be evenly divided by itself or 1.
+ If the value is not an integer, it is truncated.

## Example

```GO
calc.GCD(5, 1, 45) // Returns 1
calc.GCD(24, 36)  // Returns 12
calc.GCD(1, 2, 3, 0, 0)  // Returns 1
calc.GCD(0, 0, 0, 0)  // Returns 0
calc.GCD(6, 15, 21, 27, 30) // 3
```
