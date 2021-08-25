# calc.Sqrt

Sqrt function returns a positive square root for an input value.

## Syntax

```go
Sqrt(number)
```

## Arguments

### number

Required. The number for which you want the square root.

## Remark

+ If the number is less than zero then the function returns an Invalid Input Error message.

## Example

```go
calc.Sqrt(16) //  Returns 4.0
calc.Sqrt(123.15631)  // Returns 11.097581268006106
calc.sqrt(0)  // Returns 0
calc.Sqrt(-16) // Returns Invalid Input Error
```
