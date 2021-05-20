# calc.Quotient

Quotient function returns only the integer portion of the result for division operation and discards the remainder.

## Syntax

```go
Quotient(number)
```

## Arguments

### numerator

Required. The dividend

### denominator

Required. The divisor

## Example

```go
calc.Quotient(16, 5) //  Returns 3
calc.Quotient(-26, 5)  // Returns -5
calc.Quotient(-122.15631, 2.546) // Returns -47
calc.Quotient(123.15631, 48)  // Returns 2
```
