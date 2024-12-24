package lemin

import (
	"fmt"
	"strings"
)

func SendAnt(groups [][][]string, ants int) ([][]string, []int) {
	// Select the best group of paths (assuming it's the first group)
	bestGroup := groups[0]
	antInPath := make([]int, len(bestGroup))
	pathLengths := make([]int, len(bestGroup))

	// Calculate the length of each path
	for i, path := range bestGroup {
		pathLengths[i] = len(path)
	}

	// Distribute ants across paths
	for ants > 0 {
		shortestPath := 0

		for i := 1; i < len(antInPath); i++ {
			if antInPath[i]+pathLengths[i] < antInPath[shortestPath]+pathLengths[shortestPath] {
				shortestPath = i
			}
		}

		antInPath[shortestPath]++
		ants--
	}

	return bestGroup, antInPath
}

func PrintAnt(finalPaths [][]string, path []int) {
	type AntPosition struct {
		ant  int
		path int
		step int
	}

	var antPositions []AntPosition
	var finalResult string

	// Initialize ant positions
	antID := 1
	for pathIndex, antCount := range path {
		for i := 0; i < antCount; i++ {
			antPositions = append(antPositions, AntPosition{antID, pathIndex, 0})
			antID++
		}
	}

	// Simulate ant movements
	for len(antPositions) > 0 {
		var moves []string
		var newPositions []AntPosition
		usedLinks := make(map[string]bool)

		for _, pos := range antPositions {
			if pos.step < len(finalPaths[pos.path])-1 {
				currentRoom := finalPaths[pos.path][pos.step]
				nextRoom := finalPaths[pos.path][pos.step+1]
				link := currentRoom + "-" + nextRoom
				if !usedLinks[link] {
					moves = append(moves, fmt.Sprintf("L%d-%s", pos.ant, nextRoom))
					newPositions = append(newPositions, AntPosition{pos.ant, pos.path, pos.step + 1})
					usedLinks[link] = true
				} else {
					newPositions = append(newPositions, pos)
				}
			}
		}

		if len(moves) > 0 {
			finalResult += strings.Join(moves, " ")
			finalResult += "\n"
		}

		antPositions = newPositions
	}

	fmt.Print(finalResult)
}
