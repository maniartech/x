# calc.Roman

Roman function accept an arabic numeral and return the roman numeral as a string.

## Syntax

```go
Roman(number interface{}) string
```

## Arguments

### number

Required. The number is an arabic numeral you want to convert to roman.

## Remark

If the number is negative or greater than 3999 then the invalid input error message is returned.

## Example

```Go
Roman(1) // Returns I

Roman(3999)  // Returns MMMCMXCIX
```
