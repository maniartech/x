# calc.GCD

GCD function Returns the greatest common divisor of two or more non-zero integers without a remainder.

## Syntax

```go
GCD(number1, number2 ... interface{}) int
```

## Arguments

### number1

number1 is required. If the value is not an integer, it is truncated.

### number2 '...'

number2 and subsequent numbers are optional.

## Remark

Remark

+ All the values are evenly divided by 1.
+ A Prime number can be evenly divided by itself or 1.

## Example

```Go
GCD(5, 1, 45) // Returns 1

GCD(24,36)  // Returns 12
```
