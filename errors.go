package klib

import (
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	CodeAborted            = "ABORTED"
	CodeAlreadyExists      = "ALREADY_EXISTS"
	CodeBufferError        = "BUFFER_ERROR"
	CodeCanceled           = "CANCELED"
	CodeConditionFailed    = "CONDITION_FAILED"
	CodeDatabaseError      = "DATABASE_ERROR"
	CodeExecutionError     = "EXECUTION_ERROR"
	CodeExpired            = "EXPIRED"
	CodeFileError          = "FILE_ERROR"
	CodeForwardedError     = "FORWARDED_ERROR"
	CodeInternalError      = "INTERNAL_ERROR"
	CodeInvalidValue       = "INVALID_VALUE"
	CodeMissingValue       = "MISSING_VALUE"
	CodeNetworkError       = "NETWORK_ERROR"
	CodeNotFound           = "NOT_FOUND"
	CodeOpenFileError      = "OPEN_FILE_ERROR"
	CodeParseError         = "PARSE_ERROR"
	CodeSerializationError = "SERIALIZATION_ERROR"
	CodeTypeAssertionError = "TYPE_ASSERTION_ERROR"
	CodeUnauthenticated    = "UNAUTHENTICATED"
	CodeUnauthorized       = "UNAUTHORIZED"
	CodeUnexpectedContent  = "UNEXPECTED_CONTENT"
	CodeUnimplemented      = "UNIMPLEMENTED"
	CodeUnknownError       = "UNKNOWN_ERROR"
	CodeUnsupportedValue   = "UNSUPPORTED_VALUE"
)

type HelpLink struct {
	// Describes what the link offers.
	Description string `cbor:"1,keyasint,omitempty" json:"description,omitempty"`

	// The URL of the link.
	URL string `cbor:"2,keyasint,omitempty" json:"url,omitempty"`
}

// Inspired by https://jsonapi.org/format/#error-objects.
type Error struct {
	// A globally unique identifier.
	ID string `cbor:"1,keyasint,omitempty" json:"id,omitempty"`

	// The HTTP status code applicable to this problem. This SHOULD be provided.
	Status int `cbor:"2,keyasint,omitempty" json:"status,omitempty"`

	// An application-specific error code, expressed as a string value.
	Code string `cbor:"3,keyasint,omitempty" json:"code,omitempty"`

	// A label that can be used as a sub error code, or to better identify the error.
	Label string `cbor:"4,keyasint,omitempty" json:"label,omitempty"`

	// The path to the error, e.g., the input object's field that caused the error.
	Path string `cbor:"5,keyasint,omitempty" json:"path,omitempty"`

	// A short, human-readable summary of the problem that SHOULD NOT change
	// from occurrence to occurrence of the problem, except for purposes of localization.
	Title string `cbor:"6,keyasint,omitempty" json:"title,omitempty"`

	// A human-readable explanation specific to this occurrence of the problem.
	// Like title, this field’s value can be localized.
	Detail string `cbor:"7,keyasint,omitempty" json:"detail,omitempty"`

	// The original error message, from stdlib or external libraries/packages.
	// This is equivalent to jsonapi's "source" field, although we just care
	// about the error message, not additional information.
	Cause string `cbor:"8,keyasint,omitempty" json:"cause,omitempty"`

	// The amount of time to wait before retrying the request.
	RetryDelay time.Duration `cbor:"9,keyasint,omitempty" json:"retryDelay,omitempty"`

	// Available help for the error.
	Help []*HelpLink `cbor:"10,keyasint,omitempty" json:"help,omitempty"`

	// An object containing non-standard meta-information about the error.
	Meta map[string]any `cbor:"11,keyasint,omitempty" json:"meta,omitempty"`
}

// To-do: support custom formatter.
func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	b := new(strings.Builder)

	switch {
	case e.Code != "":
		b.WriteString("code:" + e.Code)
	case e.ID != "":
		b.WriteString("id:" + e.ID)
	case e.Status != 0:
		b.WriteString("status:" + strconv.Itoa(e.Status))
	}

	switch {
	case e.Detail != "":
		b.WriteString(" detail:" + e.Detail)
	case e.Title != "":
		b.WriteString(" title:" + e.Title)
	case e.Cause != "":
		b.WriteString(" cause:" + e.Cause)
	}

	return b.String()
}

type ErrorChain []*Error

func (e ErrorChain) Error() string {
	s := new(strings.Builder)

	for i := len(e) - 1; i >= 0; i-- {
		o := e[i]

		s.WriteString(o.Error())

		if i > 0 {
			s.WriteString(" => ")
		}
	}

	return s.String()
}

func (e ErrorChain) Add(err ...*Error) ErrorChain {
	return append(e, err...)
}

func (e ErrorChain) First() *Error {
	if len(e) == 0 {
		return nil
	}

	return e[0]
}

func (e ErrorChain) Last() *Error {
	if len(e) == 0 {
		return nil
	}

	return e[len(e)-1]
}

// ForwardError wraps and returns the given error as an ErrorChain.
// The id is used to identify where the forwarding happened.
func ForwardError(id string, err error) ErrorChain {
	var chain ErrorChain

	switch v := err.(type) {
	case *Error:
		chain = ErrorChain{v}
	case ErrorChain:
		chain = v
	default:
		chain = ErrorChain{&Error{
			ID:     "1cd23aa9-1844-4672-90d6-2158268bd2ce",
			Status: http.StatusInternalServerError,
			Code:   CodeUnknownError,
			Cause:  err.Error(),
		}}
	}

	return chain.Add(&Error{
		ID:   id,
		Code: CodeForwardedError,
	})
}

// CheckTestError verifies that the given error is the expected error
// and returns true if err is not null and equals want.
func CheckTestError(t *testing.T, err error, want *Error) bool {
	if err == nil {
		if want == nil {
			// No error, move on.
			return false
		}

		t.Fatal("expected error, got none")
	}

	if want == nil {
		t.Fatalf("unexpected error: %s", err)
	}

	have, ok := err.(*Error)
	if !ok {
		t.Fatal("err is not a *klib.Error")
	}

	assert.Equal(t, want.ID, have.ID, "Error.ID mismatch")
	assert.Equal(t, want.Status, have.Status, "Error.Status mismatch")
	assert.Equal(t, want.Code, have.Code, "Error.Code mismatch")
	assert.Equal(t, want.Path, have.Path, "Error.Path mismatch")

	return true
}
