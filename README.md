
# x - the expression evaluator! (work in progress)

The expression evaluator in Go-lang, similar to Excel. This open source project is part of Processious, a modern no-code, low-code process automation platform.

```go
package main

import "github.com/maniartech/x"

func main() {
  exp := "Upper(NumToWord(2000 + 100 + 222))"
  val, err := x.Eval(exp, nil)
  if err != nil {
    panic(err)
  }

  print(val) // Prints: TWO THOUSAND THREE HUNDRED TWENTY TWO
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

#### github.com/maniartech/x/finance

Provides the currency manipulation functions.

```go
func Num2WordInd(input string) string
func Num2Word(input string) string
func Digit2Word(input string) string
func FormatMoney(input string) string
func FormatMoneyInd(input string) string
```
