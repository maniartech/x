# calc.Acosh

Acosh calculates the inverse hyperbolic cosine of a number and returns the angle in the radian.

## Syntax

```go
Acosh(number)
```

## Arguments

### number

Required. The hyperbolic cosine of an angle you want, it must be greater than or equal to 1.

## Remark

+ If the input value of an argument is less than 1 then it will throw an error.

## Example

```go
calc.Acosh(561.15) // Returns 7.023134636091518
calc.Acosh(1546)  // Returns 8.036573305109862
calc.Acosh(1) // Returns 0.0
calc.Acosh(0) // Returns Invalid Input Error Message
calc.Acosh(-1) // Returns Invalid Input Error Message
```
