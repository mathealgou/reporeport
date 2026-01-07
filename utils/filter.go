package utils

func FilterStringSlice(slice []string, lambda func(any) bool) []string {
	var result []string
	for _, item := range slice {
		if lambda(item) {
			result = append(result, item)
		}
	}
	return result
}
