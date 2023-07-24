package klib

import (
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

// BaseFuncMap returns a template.FuncMap with
// sprig's GenericFuncMap and additional utility functions.
func BaseFuncMap() template.FuncMap {
	fm := sprig.GenericFuncMap()

	fm["osAbs"] = filepath.Abs

	return template.FuncMap(fm)
}
