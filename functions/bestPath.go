package lemin

func BestPaths(paths [][]string) ([][]string, [][]string) {
	BestPaths := [][]string{}
	AllPaths := paths

	// Map to store the shortest path to each destination
	bestPathsMap := make(map[string][]string)

	for _, path := range paths {
		// Extract the destination (last element in the path)
		destination := path[len(path)-1]

		// Check if the destination is already in the map
		if bestPath, exists := bestPathsMap[destination]; exists {
			// If the current path is shorter, replace it
			if len(path) < len(bestPath) {
				bestPathsMap[destination] = path
			}
		} else {
			// If destination is not in the map, add it
			bestPathsMap[destination] = path
		}
	}

	// Collect the best paths from the map
	for _, path := range bestPathsMap {
		BestPaths = append(BestPaths, path)
	}

	return BestPaths, AllPaths
}



//[[0 5 8 1 2] [0 5 8 2] [0 8 1 2] [0 8 2]]
//[[0 5 8 2] [0 8 2]]