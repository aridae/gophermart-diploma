package pointer

import "slices"

func To[T comparable](v T, nilifs ...T) *T {
	if slices.Contains(nilifs, v) {
		return nil
	}

	return &v
}
