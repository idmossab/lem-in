package lemin

import (
	"slices"
	"sort"
)

// Function to group paths into compatible groups
func GroupPaths(paths [][]string) [][][]string {
	var groups [][][]string
	used := make(map[int]bool)
	sort.Slice(paths, func(i, j int) bool {
		var lenI, lenJ int
		for _, path := range paths[i] {
			lenI += len(path) // Calculate total length of all paths in group i
		}
		for _, path := range paths[j] {
			lenJ += len(path) // Calculate total length of all paths in group j
		}
		return lenI < lenJ // Sort by total length
	})
	// fmt.Println("paths : ",paths)
	if len(paths) > 0 {
		groups = append(groups, [][]string{paths[0]})
		used[0] = true // Mark the first path as used
	}
	for i, path := range paths {
		if used[i] {
			continue // Skip already grouped paths
		}
		// fmt.Println("pathss",path)
		currentGroup := [][]string{path}

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

		if !groupExists(groups[1:], currentGroup) {
			groups = append(groups, currentGroup)
		}
		used = map[int]bool{}
	}
	sort.Slice(groups, func(i, j int) bool {
		var lenI, lenJ int
		for _, path := range groups[i] {
			lenI += len(path) // Calculate total length of all paths in group i
		}
		for _, path := range groups[j] {
			lenJ += len(path) // Calculate total length of all paths in group j
		}
		return lenI < lenJ // Sort by total length
	})
	return groups
}

// Helper function to check if a path is compatible with a group
func isCompatible(group [][]string, path []string) bool {
	for _, groupPath := range group {
		if hasCommonRooms(groupPath, path) {
			return false
		}
	}
	return true
}

// Check if two paths share room (excluding "start" and "end")
func hasCommonRooms(path1, path2 []string) bool {
	nodes1 := path1[1 : len(path1)-1]
	nodes2 := path2[1 : len(path2)-1]
	for _, node := range nodes1 {
		if slices.Contains(nodes2, node) {
			return true
		}
	}
	return false
}

// Check if a group already exists
func groupExists(groups [][][]string, newGroup [][]string) bool {
	for _, group := range groups {
		if areGroupsEqual(group, newGroup) {
			return true
		}
	}
	return false
}

// Helper to compare groups for equality
func areGroupsEqual(group1, group2 [][]string) bool {
	if len(group1) != len(group2) {
		return false
	}

	for _, path1 := range group1 {
		found := false
		for _, path2 := range group2 {
			if slices.Equal(path1, path2) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func SortByLength(slices [][]string) [][]string {
	sort.Slice(slices, func(i, j int) bool {
		return len(slices[i]) < len(slices[j])
	})
	return slices
}
