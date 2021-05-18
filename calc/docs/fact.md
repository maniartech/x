# calc.Fact

Fact function returns the factorial of a number, that is 1*2*3*...*number.

## Syntax

```go
Fact(number interface{}) int
```

## Arguments

### number

Required. The number is a non-negative value for which the factorial is calculated.

## Remark

+ If the number is negative, the function will return an Invalid Input Error message.
+ If the number is not an integer, it is truncated.

## Example

```Go
calc.Fact(0) // Returns 1
calc.Fact(12)  // Returns 479001600
calc.Fact(1.4315) // Returns 1
calc.Fact(4.731) // Returns 24
calc.Fact(-23)  // Returns Invalid Input Error Message
```
