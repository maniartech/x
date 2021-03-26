
# x - the expression evaluator! (work in progress)

The expression evaluator in Go-lang, similar to Excel. This open source project is part of Processious, a modern no-code, low-code process automation platform.

```go
package main

import "github.com/maniartech/x"

func main() {
  exp := "Upper(NumToWord(2000 + 100 + 222))"
  val := x.Eval(exp)

  print(val) // Prints - TWO THOUSAND THREE HUNDRED TWENTY TWO
}
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

#### github.com/maniartech/x/currency

Provides the currency manipulation functions.

```go
func Num2WordInd(input string) (string , error)
func Num2Word(input string) (string , error)
func Digit2Word(input string) (string , error)
func FormatMoney(input string) (string , error)
func FormatMoneyInd(input string) (string , error)
```
