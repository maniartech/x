package core

import "errors"

var ErrInvalidInput error = errors.New("invalid-input")
var ErrDivideBy0 error = errors.New("cannot divide by 0")
var ErrInputOutOfRange error = errors.New("array index out of bounds")
