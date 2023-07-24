package klib

// StringHandler is an interface for processing strings.
// E.g.: validation, transformation, etc.
type StringHandler interface {
	Handle(input string) (string, error)
}
