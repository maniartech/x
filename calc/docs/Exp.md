# calc.Exp

Exp function returns the result of a constant e raised to the power of a number.
The value of constant e is approximately equal to 2.71828182845904, the base of the natural logarithm.

## Syntax

```go
Exp(number)
```

## Arguments

### number

Required. The number is an exponent applied to the constant base e.

## Remark

+ Exp function is inverse of LN function.

## Example

```go
calc.Exp (0) //  Returns 1.0
calc.Exp(0.002)  // Returns 1.0020020013340003
calc.Exp(2) // Returns 7.38905609893
calc.Exp(1) // Returns 2.71828182846
```
