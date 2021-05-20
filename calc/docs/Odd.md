# calc.Odd

Odd function returns the nearest odd integer after rounding up the input number.

## Syntax

```go
Odd(number)
```

## Arguments

### number

Required. The number you want to round up.

## Returns

+ Even function rounds a positive and negative number up (away from zero) so the positive number becomes larger and the negative number becomes smaller.

## Example

```go
calc.Odd(-122.15631) //  Returns -123
calc.Odd(123.15631)  // Returns 125
calc.Odd(16) // Returns 17
```
