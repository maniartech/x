# calc.Asin

Asin calculates the inverse sine of a number and returns an angle in the radians between -π/2 and π/2.
Geometrically it is the ratio of a triangle's opposite side over its hypotenuse.

## Syntax

```go
Asin(number)
```

## Arguments

### number

Required. The number is an input to Asin function, it must be between -1 and 1.

## Remark

+ If the input value of an argument is outside the range of -1 to 1 then it will throw an error.

## Example

```go
calc.Asin(0.26) // Returns 0.2630222029084689
calc.Asin(-0.546)  // Returns -0.5775822745289381
calc.Asin(0) // Returns 0.0
calc.Asin(-1) // Returns -1.5708
calc.Asin(1) // Returns 1.570796
```
