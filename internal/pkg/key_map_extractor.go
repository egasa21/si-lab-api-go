package pkg

func GetKeysFromMap(m map[int]bool) []int {
	keys := make([]int, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}