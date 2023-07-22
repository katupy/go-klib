// package must contains functions that panic on error.
// The intended usage is to simplify testing.
package must

import "path/filepath"

func FilepathAbs(path string) string {
	abs, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return abs
}
