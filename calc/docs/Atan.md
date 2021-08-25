# calc.Atan

Atan calculates the inverse tangent of a number and returns an angle in the radians between -π/2 and π/2.
Geometrically it is the ratio of the right triangle's opposite side over its adjacent side.

## Syntax

```go
Atan(number)
```

## Arguments

### number

Required. The number is a tangent of an angle you want. It can be any real number.

## Example

```go
calc.Atan(0.235) // Returns 0.23081196040266755
calc.Atan(-0.546)  // Returns -0.49976700744972385
calc.Atan(1) // Returns 0.785398
calc.Atan(-1) // Returns 0.7854
calc.Atan(0) // Returns 0.0
```
