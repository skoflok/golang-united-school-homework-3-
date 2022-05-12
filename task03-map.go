package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0)
	for i, _ := range input {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	for _, v := range keys {
		result = append(result, input[v])
	}
	return result
}
