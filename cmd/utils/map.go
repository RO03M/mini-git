package utils

func Map[T any, R any](slice []T, callback func(item T, key int) R) []R {
	var newArray []R = make([]R, 0, len(slice))

	for index, item := range slice {
		newArray = append(newArray, callback(item, index))
	}

	return newArray
}
