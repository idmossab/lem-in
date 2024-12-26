package lemin

import "sort"

func FilterShortestSlices(data [][][]string) [][][]string {
	result := [][][]string{}

	for _, group := range data {
		if len(group) == 0 {
			continue
		}

		// Sort the 2D slice by length of inner slices
		sort.Slice(group, func(i, j int) bool {
			return len(group[i]) < len(group[j])
		})

		// Find the shortest slice (only one)
		shortestSlice := group[0]
		result = append(result, [][]string{shortestSlice})
	}

	return result
}
