# calc.Ceiling

Ceiling function returns number rounded up to the supplied number, that is away from zero, to the nearest multiple of significance.

## Syntax

```go
Ceiling(number, significance interface{}) float64
```

## Arguments

### number

Required. The number is a value that you want to round off.

### significance

Required. The multiple to which you want to round up to.

## Remark

+ The value is rouded up and regardless sign of a number when adjusted away from zero.
+ No rounding up occurs for the number which is exact multiple of significance.
+ If both the number and significance are negative, the value is rounded down, away from zero.
+ If number is negative, and significance is positive, the value is rounded up towards zero.

## Example

```Go
Ceiling(-42.5, 16) // Returns -32.0

Ceiling(-32.25, -5)  // Returns -35.0
```