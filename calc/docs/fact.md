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

+ If the number is negative, the function will return Invalid Input Error message.
+ If number is not an integer, it is truncated.

## Example

```Go
Fact(0) // Returns 1
Fact(12)  // Returns 479001600
```
