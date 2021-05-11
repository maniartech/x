# calc.Base

Base function returns the string representation of the provided number converted into different base.

## Syntax

```go
Base(number, redix, minL interface{}) string
```

## Arguments

### number

The number that needs to be converted into another representation

### redix

The base of the reprentation that the number should be converted to

### minL

Optional, if provided, guearantees the minimum length of the resulting number by providing padding with leading zeros.

## Example

```Go
Base(15, 16) // Returns F

Base(10, 2)  // Returns 1010
```
