
# calc.Atanh

Atanh returns the inverse hyperbolic tangent of a number and returns the angle in radian.

## Syntax

```go
Atanh(number)
```

## Arguments

### number

Required. The inverse hyperbolic tangent of the angle you want must be between -1 to 1.

## Remark

+ If the input value of an argument is outside the range of -1 to 1 then it will throw an error.
+ The inverse hyperbolic tangent of 1 and -1 are Infinity and -Infinity.

## Example

```go
calc.Atanh(0.99) // Returns 2.6466524123622457
calc.Atanh(-0.2686)  // Returns -0.275354351742258
calc.Atanh(0.651) // Returns 0.7770322608588505
calc.Atanh(-23) // Returns an error
calc.Atanh(23) // Returns an error
```
