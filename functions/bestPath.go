package lemin

import "sort"


func BestPaths(paths [][]string) ([][]string, [][]string) {
	bestPaths := [][]string{} 

	seen := make(map[string]bool)

	for _, path := range paths {
		unique := true 

		for i := 1; i < len(path)-1; i++ {
			if seen[path[i]] { // If an intermediate room is already used in another path
				unique = false
				break 
			}
		}

		// If the path has unique intermediate room
		if unique {
			// Mark the intermediate nodes of this path as seen
			for i := 1; i < len(path)-1; i++ {
				seen[path[i]] = true
			}

			bestPaths = append(bestPaths, path)
		}
	}

	return bestPaths, paths
}

func UniquePaths(paths [][]string) ([][]string, [][]string) {
	uniquePaths := [][]string{}
	anotherPaths := [][]string{}
	seen := make(map[string]bool)

	for _, path := range paths {
		unique := true

		for i := 1; i < len(path)-1; i++ {
			if seen[path[i]] {
				unique = false
				break
			}
		}

		if unique {
			for i := 1; i < len(path)-1; i++ {
				seen[path[i]] = true
			}
			uniquePaths = append(uniquePaths, path)
		} else {
			anotherPaths = append(anotherPaths, path)
		}
	}

	return uniquePaths, anotherPaths
}

func SortByLength(slices [][]string) [][]string {
	sort.Slice(slices, func(i, j int) bool {
		return len(slices[i]) < len(slices[j]) 
	})
	return slices 
}
