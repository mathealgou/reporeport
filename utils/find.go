package utils

// FindInSlice returns item if function returns true
func FindInSlice(slice []any, lambda func(any) bool) any {
	for _, item := range slice {
		if lambda(item) {
			return item
		}
	}
	return nil
}
