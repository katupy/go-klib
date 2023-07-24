package klib

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertSliceElem(t *testing.T) {
	testCases := []*struct {
		name    string
		slice   []any
		element any
		index   int
		err     *Error
		want    []any
	}{
		{
			name:  "out-of-bounds",
			index: 1,
			err: &Error{
				ID:     "c0548791-56ff-4ab2-a996-f04b59caa089",
				Status: http.StatusBadRequest,
				Code:   CodeInvalidValue,
			},
		},
		{
			name:  "out-of-bounds-negative",
			index: -2,
			err: &Error{
				ID:     "c0548791-56ff-4ab2-a996-f04b59caa089",
				Status: http.StatusBadRequest,
				Code:   CodeInvalidValue,
			},
		},
		{
			name:    "prepend",
			element: "a",
			slice: []any{
				"b",
				"c",
			},
			want: []any{
				"a",
				"b",
				"c",
			},
		},
		{
			name:    "prepend-negative",
			element: 1,
			slice: []any{
				2,
				3,
			},
			index: -3,
			want: []any{
				1,
				2,
				3,
			},
		},
		{
			name:    "append",
			element: "a",
			slice: []any{
				"b",
				"c",
			},
			index: 2,
			want: []any{
				"b",
				"c",
				"a",
			},
		},
		{
			name:    "append-negative",
			element: 1,
			slice: []any{
				2,
				3,
			},
			index: -1,
			want: []any{
				2,
				3,
				1,
			},
		},
		{
			name:    "insert-positive-index",
			element: "b",
			slice: []any{
				"a",
				"c",
				"d",
			},
			index: 1,
			want: []any{
				"a",
				"b",
				"c",
				"d",
			},
		},
		{
			name:    "insert-negative-index",
			element: 2,
			slice: []any{
				1,
				3,
				4,
			},
			index: -3,
			want: []any{
				1,
				2,
				3,
				4,
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(fmt.Sprintf("%d:%s", i, tc.name), func(st *testing.T) {
			haveSlice, err := InsertSliceElem(tc.slice, tc.element, tc.index)
			if CheckTestError(st, err, tc.err) {
				return
			}

			wantSlice := tc.want

			if assert.Equal(st, len(wantSlice), len(haveSlice), "Slice length mismatch") {
				for j := range wantSlice {
					have := haveSlice[j]
					want := wantSlice[j]

					assert.Equal(st, want, have, "Slice[%d] value mismatch", j)
				}
			}
		})
	}
}
