package pointer

import "slices"

func To[T comparable](v T, nilifs ...T) *T {
	if slices.Contains(nilifs, v) {
		return nil
	}

	return &v
}

func Deref[T comparable](v *T, fallback T) T {
	if v == nil {
		return fallback
	}

	return *v
}
