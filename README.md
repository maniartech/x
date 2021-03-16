
# go-funcs

Excel like reusable functions in go lang

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
