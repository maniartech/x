# calc.Trunc

Trunc function returns a truncated number to the given precision.

## Syntax

```go
Trunc(number)
```

## Arguments

### number

Required. The number you want to truncate.

## Remark

Trunc and Int looks similar functions as they return integers. Trunc removes fractional part of number and Int rounds down number to closest integer according to number's fractional part. But Trunc and Int works differently for negative numbers. Trunc(-1.6) returns -1 and Int(-1.6) returns -2 as the value of -2 is more lower than -1.

## Example

```Go
Trunc(-16) //  Returns -16.0
Trunc(-122.15631)  // Returns -122.0
```
