package slices

import "fmt"

type IndexOutOfBoundsError struct {
	slices []any
	index  int
}

type TypeAssertionError struct {
}

func NewIndexOutOfBoundsError(slices []any, index int) IndexOutOfBoundsError {
	return IndexOutOfBoundsError{
		slices: slices,
		index:  index,
	}
}

func NewTypeAssertionError() TypeAssertionError {
	return TypeAssertionError{}
}

func (e IndexOutOfBoundsError) Error() string {
	return fmt.Sprintf("index '%d' is out of bounds of slice: %v", e.index, e.slices)
}

func (e TypeAssertionError) Error() string {
	return ""
}
