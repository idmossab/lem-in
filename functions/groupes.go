package lemin

import (
	"slices"
	"sort"
)

// Function to group paths into compatible groups
func GroupPaths(paths [][]string) [][][]string {
	var groups [][][]string
	used := make(map[int]bool)

	for i, path := range paths {
		if used[i] {
			continue // Skip already grouped paths
		}

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

		if !groupExists(groups, currentGroup) {
			groups = append(groups, currentGroup)
		}
		used = map[int]bool{}
	}

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

