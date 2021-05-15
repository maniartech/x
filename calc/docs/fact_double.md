# calc.FactDouble

FactDouble func calculates the double factorial of the number. It is represented using double exclamation mark (!!).

## Syntax

```go
FactDouble(number interface{}) int
```

## Arguments

### number

Required. The number is a non-negative number for which the double factorial is calculated.

## Remark

+ If number is not an integer, it is truncated.
+ If number is even : n!!= n(n-2)(n-4)...(4)(2)
+ If number is odd : n!!= n(n-2)(n-4)...(3)(1)
+ The double factorial for both zero and -1 are defined as 1.
+ For numbers less than -1, a double factorial is not defined and FactDouble function will throw an error.

## Example

```Go
FactDouble(0) // Returns 1
FactDouble(12)  // Returns 46080
```
