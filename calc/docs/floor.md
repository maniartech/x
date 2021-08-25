# calc.Floor

The floor function rounds down the given number to the nearest multiple of significance.
That is the returned number is shifted towards zero and closer to the multiple of significance.

## Syntax

```go
Floor(number, significance interface{}) float64
```

## Arguments

### number

Required. The number is a value that you want to round off.

### significance

Required. The multiple is used to round down the given number.

## Remark

+ The value is always rounded down, regardless sign of a given number when adjusted away from zero.
+ If the given value is an exact multiple of significance then no rounding up occurs.
+ If both the number and significance are negative, the value is rounded down, towards zero.
+ If the number is negative, and significance is positive, the value is rounded down away from zero.
+ If the input number is passed as a string the string value is converted to the desired format. If successfully converted it returns the expected outcome or throws an error.

## Example

```Go
calc.Floor(-26, -5) // Returns -25.0
calc.Floor(-27.5, 4)  // Returns -28.0
calc.Floor("1.5", "0.1") // Returns 1.5
calc.Floor(123.456, 1.3) // Returns 122.2
calc.Floor(123.456, 0)  // Returns an Error Message
```
