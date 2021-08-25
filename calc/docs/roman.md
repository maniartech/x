# calc.Roman

Roman function accept an Arabic numeral and convert it to the roman numeral as a string.

## Syntax

```go
Roman(number interface{}) string
```

## Arguments

### number

Required. The number is an Arabic numeral you want to convert to a roman string.

## Remark

+ If the number is negative or greater than 3999 the invalid input error message is returned.

## Example

```go
calc.Roman(1) // Returns I
calc.Roman(3999))  // Returns MMMCMXCIX
calc.Roman(4000)   // Returns Invalid Input Error Message
calc.Roman(-22)  // Returns Invalid Input Error Message
calc.Roman(0) // Returns Invalid Input Error Message
```
