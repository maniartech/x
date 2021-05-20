# calc.SqrtPi

SqrtPi function returns the value of number multiplied by Pi.

## Syntax

```go
SqrtPi(number)
```

## Arguments

### number

Required. The number to which Pi is multiplied.

## Remark

+ If the number is less than zero then the function returns an Invalid Input Error message.

## Example

```go
calc.SqrtPi(16) //  Returns 7.0898154036220635
calc.SqrtPi(123.15631)  // Returns 19.669950654214343
calc.SqrtPi(-16) // Returns Invalid Input Error
```
