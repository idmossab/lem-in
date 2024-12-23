package lemin

import "fmt"


func SendAnt(groups [][][]string, ants int) ([][]string,[]int) {
    bestGroup := groups[0]
    antInPath := make([]int, len(bestGroup))
    pathLengths := make([]int, len(bestGroup))

    // Calculate the length of each path
    for i, path := range bestGroup {
        pathLengths[i] = len(path)
    }

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

    // Print the result
    for i, count := range antInPath {
        fmt.Printf("Path %d: %d ants\n", i+1, count)
    }
	return bestGroup,antInPath
}

func PrintAnt(finalPath [][]string, path []int) {
	// Fixed number of ants
	antNum := 1

	// Determine the maximum length of any path
	maxLength := 0
	for i := 0; i < len(finalPath); i++ {
		if len(finalPath[i]) > maxLength {
			maxLength = len(finalPath[i])
		}
	}

	// Iterate through each step
	for step := 0; step < maxLength; step++ {
		// Loop through the paths and print the corresponding ants
		for i := 0; i < len(path); i++ {
			// Only print if the path has remaining steps and the step is within bounds
			if path[i] > 0 && step < len(finalPath[i]) {
				// Print the ant number and the room it's moving to
				fmt.Printf("L%d-%s ", antNum, finalPath[i][step])
				antNum++

				// Reset antNum to 1 after the third ant
				if antNum > 3 {
					antNum = 1
				}
			}
		}
		fmt.Println()
	}
}