# calc.Round

Round function rounds a number to a specified number of digits.

## Syntax

```go
Round(number, num_digit interface{}) float64
```

## Arguments

### number

Required. The number that you want to round up.

### num_digits

Required. The number of decimal points to which you want to round the number argument.

## Remark

+ If the value of num_digits is greater than 0 (zero), then number is rounded to the specified number of decimal places.
+ If the value of num_digits is 0, the number is rounded to the nearest integer.
+ If the value of num_digits is less than 0, the number is rounded to the left of the decimal point.

## Example

```Go
Round(7.6549, 3) // Returns 7.655

Round(7.2522, 2)  // Returns 7.25
```
