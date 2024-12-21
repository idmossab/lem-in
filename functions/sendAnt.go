package lemin

import "fmt"

func SendAnt(groups [][][]string, ants int) {
    bestGroup := groups[0]
    lenbestGroup := make([]int, len(bestGroup))
    pathLengths := make([]int, len(bestGroup))

    // Calculate the length of each path
    for i, path := range bestGroup {
        pathLengths[i] = len(path)
    }

    // Distribute ants, prioritizing shorter paths
    for ants > 0 {
        shortestPath := 0

        // Find the shortest path in terms of the current load
        for i := 1; i < len(lenbestGroup); i++ {
            if lenbestGroup[i]+pathLengths[i] < lenbestGroup[shortestPath]+pathLengths[shortestPath] {
                shortestPath = i
            }
        }

        // Assign an ant to the shortest path
        lenbestGroup[shortestPath]++
        ants--
    }

    // Print the result
    for i, count := range lenbestGroup {
        fmt.Printf("Path %d: %d ants\n", i+1, count)
    }
}
