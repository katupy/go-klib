package klib

import (
	"net/http"
	"strings"
	"time"
)

const (
	// Common.
	ErrInternalError      = "INTERNAL_ERROR"
	ErrUnknownError       = "UNKNOWN_ERROR"
	ErrForwardedError     = "FORWARDED_ERROR"
	ErrUnimplemented      = "UNIMPLEMENTED"
	ErrExecutionError     = "EXECUTION_ERROR"
	ErrTypeAssertionError = "TYPE_ASSERTION_ERROR"
	ErrConditionFailed    = "CONDITION_FAILED"

	// Variable.
	ErrMissingArgument     = "MISSING_ARGUMENT"
	ErrInvalidArgument     = "INVALID_ARGUMENT"
	ErrUndefinedArgument   = "UNDEFINED_ARGUMENT"
	ErrUnknownArgument     = "UNKNOWN_ARGUMENT"
	ErrUnsupportedArgument = "UNSUPPORTED_ARGUMENT"
	ErrContentExpected     = "CONTENT_EXPECTED"
	ErrUnexpectedContent   = "UNEXPECTED_CONTENT"
	ErrParseError          = "PARSE_ERROR"
	ErrBufferError         = "BUFFER_ERROR"

	// Network.
	ErrNetworkError = "NETWORK_ERROR"

	// Object.
	ErrAlreadyExists     = "ALREADY_EXISTS"
	ErrMarshalingError   = "MARSHALING_ERROR"
	ErrUnmarshalingError = "UNMARSHALING_ERROR"

	// File.
	ErrFilesystemError = "FILESYSTEM_ERROR"
	ErrOpenFileError   = "OPEN_FILE_ERROR"
	ErrReadFileError   = "READ_FILE_ERROR"
	ErrWriteFileError  = "WRITE_FILE_ERROR"

	// Database.
	ErrDatabaseError = "DATABASE_ERROR"

	// Auth.
	ErrUnauthenticated = "UNAUTHENTICATED"
	ErrUnauthorized    = "UNAUTHORIZED"

	// Deadline.
	ErrExpired  = "EXPIRED"
	ErrAborted  = "ABORTED"
	ErrCanceled = "CANCELED"
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

func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	b := new(strings.Builder)

	if e.ID != "" {
		b.WriteString(e.ID + " ")
	}

	switch {
	case e.Detail != "":
		b.WriteString(e.Detail)
	case e.Title != "":
		b.WriteString(e.Title)
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
			ID:     "01GRA68Y70YTBNSVX7SKW6HJ81",
			Status: http.StatusInternalServerError,
			Code:   ErrUnknownError,
			Cause:  err.Error(),
		}}
	}

	return chain.Add(&Error{
		ID:   id,
		Code: ErrForwardedError,
	})
}
