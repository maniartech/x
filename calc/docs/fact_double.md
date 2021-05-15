# calc.FactDouble

FactDouble func calculates the double factorial of the number.

## Syntax

```go
FactDouble(number interface{}) int
```

## Arguments

### number

Required. The number is a non-negative number for which the double factorial is calculated. If number is not an integer, it is truncated.

## Remark

Remark

+ If number is even : n!!= n(n-2)(n-4)...(4)(2)
+ If number is odd : n!!= n(n-2)(n-4)...(3)(1)

## Example

```Go
Fact(0) // Returns 1

Base(12)  // Returns 46080
```
