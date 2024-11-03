package util

func SafeSubslice[T any](slice []T, total int) []T {
	if len(slice) < total {
		// If the slice is shorter than 3, return a copy of the entire slice
		return append([]T{}, slice...)
	}
	// Otherwise, return the first 3 elements
	return slice[:total]
}
