package raycast

import (
	"fmt"
	"strconv"
)

type ErrMismatchingXsize struct {
	Line     int
	Expected int
	Got      int
}

func (e *ErrMismatchingXsize) Error() string {
	return fmt.Sprintf("error: mismatching xsize on line %d, expected %d, got %d", e.Line, e.Expected, e.Got)
}

type ErrBadNumber struct {
	Line int
	H    string
	Err  error
}

func (e *ErrBadNumber) Error() string {
	return e.Err.Error() + " on line " + strconv.Itoa(e.Line) + " with '" + e.H + "'"
}
