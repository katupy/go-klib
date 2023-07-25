package klib

import (
	"net/http"
)

// InsertSliceElem inserts element in the given slice and index.
// If index is negative, it is treated as `index+len(slice)+1`:
//   - -1 adds to the end of the slice, i.e., appends the element.
//   - -(len(slice)+1) is the same as index=0:
//     it adds to the beginning of the slice,
//     i.e., prepends the element.
func InsertSliceElem[T any](slice []T, element T, index int) ([]T, error) {
	sliceLen := len(slice)

	if index < 0 {
		index += sliceLen + 1
	}

	if index < 0 || index > sliceLen {
		return nil, &Error{
			ID:     "c0548791-56ff-4ab2-a996-f04b59caa089",
			Status: http.StatusBadRequest,
			Code:   CodeInvalidValue,
			Detail: "Index is out of bounds.",
		}
	}

	switch {
	case index == 0:
		return append([]T{element}, slice...), nil
	case index == sliceLen:
		return append(slice, element), nil
	}

	s := append(slice[:index+1], slice[index:]...)
	s[index] = element

	return s, nil
}
