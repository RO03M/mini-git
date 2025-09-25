package plumbing

func StringSliceMap(slice []string) map[string]bool {
	var stringMap map[string]bool = make(map[string]bool)

	for _, item := range slice {
		stringMap[item] = true
	}

	return stringMap
}
