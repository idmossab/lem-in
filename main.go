package main

import (
	"fmt"
	"os"

	sr "lemin/functions"
)

func main() {
	path := []string{}
	antFarm := sr.AntFarm{}
	err := antFarm.ReadFromInput(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}
	sr.GetPaths(antFarm.Tunnels, antFarm.Start.Name, antFarm.End.Name, path)
	fmt.Println(sr.Paths)
	sr.FilterShortestPathsWithDuplicates(sr.Paths)
	fmt.Println(sr.Paths)
	fmt.Println("AntFarm successfully parsed!")
}
