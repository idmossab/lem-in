package lemin

import (
	"fmt"
	"strings"
)

func SendAnt(bestGroup [][]string, ants int) ([][]string, []int) {
	fmt.Println("bestGroup",bestGroup)
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
		antID, pathIndex, step int
	}

	var (
		antPositions []AntPosition
		finalResult  string
	)

	// Initialize ant positions based on their counts in each path
	antID := 1
	for pathIndex, antCount := range path {
		for i := 0; i < antCount; i++ {
			antPositions = append(antPositions, AntPosition{antID, pathIndex, 0})
			antID++
		}
	}

	// Simulate ant movements
	for len(antPositions) > 0 {
		var (
			moves       []string
			newPositions []AntPosition
			usedLinks    = make(map[string]bool)
		)

		for _, ant := range antPositions {
			currentPath := finalPaths[ant.pathIndex]

			// Check if the ant can move to the next room
			if ant.step < len(currentPath)-1 {
				currentRoom, nextRoom := currentPath[ant.step], currentPath[ant.step+1]
				link := currentRoom + "-" + nextRoom

				if !usedLinks[link] {
					moves = append(moves, fmt.Sprintf("L%d-%s", ant.antID, nextRoom))
					newPositions = append(newPositions, AntPosition{ant.antID, ant.pathIndex, ant.step + 1})
					usedLinks[link] = true
				} else {
					newPositions = append(newPositions, ant)
				}
			}
		}

		if len(moves) > 0 {
			finalResult += strings.Join(moves, " ") + "\n"
		}

		antPositions = newPositions
	}

	fmt.Print(finalResult)
}
