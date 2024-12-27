package lemin

import (
	"slices"
	"sort"
)
func FilterShortestSlices(groupes [][][]string) [][][]string {
	sort.Slice(groupes, func(i, j int) bool {
		return len(groupes[i]) < len(groupes[j])
	})

	result := [][][]string{}
	shortPath := [][]string{}

	for _, groupe := range groupes {
		for i := 0; i < len(groupe)-1; i++ {
			if len(groupe[i]) == len(groupe[i+1]) {
				if !containsSlice(shortPath, groupe[i]) {
					shortPath = append(shortPath, groupe[i])
				}
			}
		}

		if len(shortPath) > 0 {
			result = append(result, shortPath)
		}
		shortPath = [][]string{}
	}

	return result
}

func containsSlice(slice [][]string, item []string) bool {
	for _, v := range slice {
		if slices.Equal(v, item) {
			return true
		}
	}
	return false
}
