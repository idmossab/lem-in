package lemin

import "fmt"

func SendAnt(groupes [][][]string, ant int) {
    bestGroup := groupes[0] 
    lenbestGroup := make([]int, len(bestGroup)) 

    for ant > 0 { 
        shortestPath := 0 
        for i := 1; i < len(lenbestGroup); i++ {
            if lenbestGroup[i] < lenbestGroup[shortestPath] {
                shortestPath = i 
            }
        }
        lenbestGroup[shortestPath]++
        ant-- 
    }

    for i, length := range lenbestGroup {
        fmt.Printf("Path %d: %d ants\n", i+1, length)
    }
}
