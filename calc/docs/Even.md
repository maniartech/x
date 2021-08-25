# calc.Even

Even function returns nearest even integer rounded up to the given input number.

## Syntax

```go
Even(number)
```

## Arguments

### number

Required. The number you wanted to round up.

## Remark

+ Even function rounds positive and negative number up (away from zero) so the positive number becomes larger and negative number becomes smaller.

## Example

```go
calc.Even(-11.15631) // Returns -12
calc.Even(6.2)  // Returns 8
calc.Even(11) // Returns 12
```
