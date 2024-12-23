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
 
}