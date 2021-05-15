# calc.Ceiling

Ceiling function rounds up the given number to the nearest multiple of significance.
That is the returned number is shifted away from zero and closer to the multiple of significance.

## Syntax

```go
Ceiling(number, significance interface{}) float64
```

## Arguments

### number

Required. The number is a value that you want to round off.

### significance

Required. The multiple which is used to round up the given number.

## Remark

+ The value is always rouded up regardless sign of a given number when adjusted away from zero.
+ The value will be returned in a float.
+ If the given value is an exact multiple of significance then no rounding up occurs.
+ If both the number and significance are negative, the value is rounded down, away from zero.
+ If number is negative, and significance is positive, the value is rounded up towards zero.
+ If the input number is passed as a string the string value is converted to desired format. If sucessfully converted it retuns the expected outcome or throws an error.

## Example

```Go
Ceiling(-42.5, 16) // Returns -32.0
Ceiling(-32.25, -5)  // Returns -35.0
calc.Ceiling("4.42", "0.05")  // Returns 4.45
```
