package lemin

import (
	"fmt"
	"sort"
)

func FilterShortestSlices(groupes [][][]string) [][][]string {
	sort.Slice(groupes, func(i, j int) bool {
		return len(groupes[i]) < len(groupes[j])
	})

	shortPath1 := [][][]string{}
	shortPath2 := [][][]string{}
	shortPath3 := [][][]string{}
	shortPath4 := [][][]string{}
	result := [][][]string{}
	//fmt.Println("GRSS:", groupes)
	if len(groupes) == 1 {
		result = groupes
		return result
	}
	for i := 0; i < len(groupes)-1; i++ {
		if len(groupes[i]) == len(groupes[i+1]) {
			if len(groupes[i]) == 1 {
				if !contains(shortPath1, groupes[i]) {
					shortPath1 = append(shortPath1, groupes[i])
				}
			} else if len(groupes[i]) == 2 {
				if !contains(shortPath2, groupes[i]) {
					shortPath2 = append(shortPath2, groupes[i])
				}
			} else if len(groupes[i]) == 3 {
				if !contains(shortPath3, groupes[i]) {
					shortPath3 = append(shortPath3, groupes[i])
				}
			} else {
				if !contains(shortPath4, groupes[i]) {
					shortPath4 = append(shortPath4, groupes[i])
				}
			}
		}
	}
	sort.Slice(shortPath1, func(i, j int) bool {
		return len(shortPath1[i][0]) < len(shortPath1[j][0])
	})
	sort.Slice(shortPath2, func(i, j int) bool {
		return len(shortPath2[i][0]) < len(shortPath2[j][0])
	})
	sort.Slice(shortPath3, func(i, j int) bool {
		return len(shortPath3[i][0]) < len(shortPath3[j][0])
	})
	sort.Slice(shortPath4, func(i, j int) bool {
		return len(shortPath4[i][0]) < len(shortPath4[j][0])
	})
	if len(shortPath1) > 0  {
		result = append(result, shortPath1[0])
	}
	if len(shortPath2) > 0 {
		result = append(result, shortPath2[0])
	}
	if len(shortPath3) > 0 {
		result = append(result, shortPath3[0])
	}
	if len(shortPath4) > 0 {
		result = append(result, shortPath4[0])
	}
	fmt.Println("shortPath1:", shortPath1)
	fmt.Println("shortPath2:", shortPath2)
	fmt.Println("shortPath3:", shortPath3)

	return result
}

func contains(groupList [][][]string, group [][]string) bool {
	for _, item := range groupList {
		if equalSlices(item, group) {
			return true
		}
	}
	return false
}

func equalSlices(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
