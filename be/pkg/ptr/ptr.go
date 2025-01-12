package ptr

func From[T any](v *T, fallback ...T) T {
	if v != nil {
		return *v
	}

	if len(fallback) > 0 {
		return fallback[0]
	}

	var zero T
	return zero
}
