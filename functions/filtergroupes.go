package lemin

import (
	"sort"
)

// FilterShortestPaths filters and organizes groups of paths by their lengths
func FilterShortestSlices(groups [][][]string, ant int) [][][]string {
    // Handle single group case
    if len(groups) <= 1 {
        return groups
    }

    // Sort groups by length for consistent processing
    sortGroupsByLength(groups)

    // Store unique paths by their length
    pathsByLength := make(map[int][][][]string)
    
    // Group paths by their length
    for _, group := range groups {
        length := len(group)
        // Only add if this group is unique for its length
        if !containsGroup(pathsByLength[length], group) {
            pathsByLength[length] = append(pathsByLength[length], group)
        }
    }

    // Build result array with shortest path from each length category
    return buildResultArray(pathsByLength)
}
// sortGroupsByLength sorts groups based on their length
func sortGroupsByLength(groups [][][]string) {
	sort.Slice(groups, func(i, j int) bool {
		return len(groups[i]) < len(groups[j])
	})
}

// containsGroup checks if a group already exists in the group list
func containsGroup(groupList [][][]string, target [][]string) bool {
	for _, group := range groupList {
		if areGroupsEqual1(group, target) {
			return true
		}
	}
	return false
}

// areGroupsEqual checks if two groups have identical contents
func areGroupsEqual1(group1, group2 [][]string) bool {
	if len(group1) != len(group2) {
		return false
	}

	for i := range group1 {
		if len(group1[i]) != len(group2[i]) {
			return false
		}
		for j := range group1[i] {
			if group1[i][j] != group2[i][j] {
				return false
			}
		}
	}
	return true
}

// buildResultArray constructs the final result array from the path map
func buildResultArray(pathsByLength map[int][][][]string) [][][]string {
	var result [][][]string

	// Get all lengths and sort them
	lengths := getLengths(pathsByLength)
	sort.Ints(lengths)

	// Add shortest path from each length category
	for _, length := range lengths {
		if paths := pathsByLength[length]; len(paths) > 0 {
			result = append(result, paths[0])
		}
	}

	return result
}

// getLengths extracts and returns all lengths from the path map
func getLengths(pathsByLength map[int][][][]string) []int {
	var lengths []int
	for length := range pathsByLength {
		lengths = append(lengths, length)
	}
	return lengths
}
