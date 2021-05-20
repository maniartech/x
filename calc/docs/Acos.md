# calc.Acos

Acos calculates the inverse cosine of a number and returns an angle in the radians between 0 and π.
Geometrically it is the ratio of a triangle's adjacent side over its hypotenuse.

## Syntax

```go
Acos(number)
```

## Arguments

### number

Required. The inverse cosine of an angle you want, it must be between -1 to 1.

## Remark

+ If the input value of an argument is outside the range of -1 to 1 then it will throw an error.

## Example

```go
calc.Acos(0.235) // Returns 1.3335777587004611
calc.Acos(-0.546)  // Returns 2.1483786013238344
calc.Acos(-1) // Returns 3.141593
calc.Acos(0) // Returns 1.570796
calc.Acos(1) // Returns 0
calc.Acos(-1.5689) // Returns Error Message
calc.Acos(1.5698)  // Returns Error Message
```
