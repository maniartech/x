
# Expression (WIP)

Expression evaluator in Go-lang, similar to Excel. This open source project is part of Processious, a modern no-code, low-code proces automation platform.

```go
val := expression.Eval(UpperCase(NumToWord(2000 + 100 + 222)))
print(val) // Prints - TWO THOUSAND THREE HUNDRED TWENTY TWO
```

## Important commands

```sh
# Go the the folder and run following command
go test

# Test with coverage
go test --coverprofile=coverage.out

# Open coverage profile analysis
go tool cover -html=coverage.out
```

## Finished Functions

The following packages are functions are ready for use!

### Packages

#### github.com/maniartech/go-funcs/currency

Provides the currency manipulation functions.

```go
func Num2WordInd(input string) (string , error)
func Num2Word(input string) (string , error)
func Digit2Word(input string) (string , error)
func FormatMoney(input string) (string , error)
func FormatMoneyInd(input string) (string , error)
```
