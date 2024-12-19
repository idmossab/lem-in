package lemin

import "slices"

// Function to group paths into compatible groups
func GroupPaths(paths [][]string) [][][]string {
	var groups [][][]string    // Final groups of paths
	used := make(map[int]bool) // Tracks used paths by index

	for i, path := range paths {
		if used[i] {
			continue // Skip already grouped paths
		}

		currentGroup := [][]string{path} // Start a new group with the current path
		used[i] = true

		for j, otherPath := range paths {
			if used[j] {
				continue // Skip already used paths
			}

			if isCompatible(currentGroup, otherPath) {
				currentGroup = append(currentGroup, otherPath)
				used[j] = true
			}
		}

		groups = append(groups, currentGroup)
	}

	return groups
}

// Helper function to check if a path is compatible with a group
func isCompatible(group [][]string, path []string) bool {
	for _, groupPath := range group {
		if hasCommonIntermediateNodes(groupPath, path) {
			return false // Paths are not compatible
		}
	}
	return true
}

// Check if two paths share intermediate nodes (excluding "start" and "end")
func hasCommonIntermediateNodes(path1, path2 []string) bool {
	nodes1 := path1[1 : len(path1)-1] // Exclude "start" and "end"
	nodes2 := path2[1 : len(path2)-1] // Exclude "start" and "end"

	for _, node := range nodes1 {
		if slices.Contains(nodes2, node) {
			return true // Common intermediate node found
		}
	}
	return false
}
