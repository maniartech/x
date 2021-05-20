# calc.Int

Int function rounds down a number to the nearest integer.

## Syntax

```go
Int(number)
```

## Arguments

## number

Required. The number you want to round down to the nearest integer.

## Example

```go
calc.Int(-122.15631) //  Returns -123
calc.Int(123.15631)  // Returns 123
calc.Int(16) // Returns 16
calc.Int(-16) // Returns -16
```
