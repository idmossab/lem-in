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

	// fmt.Println("All paths :", sr.SortByLength(sr.Paths))
	fmt.Println()
	groupe := sr.GroupPaths(sr.SortByLength(sr.Paths))
	fmt.Println("groupe", groupe)
	nbrAnt := antFarm.Ants
	fmt.Println(nbrAnt)
	antDistribution := sr.DistributeAnts(groupe, nbrAnt)
	finalResult := sr.SimulateAntMovement(groupe, antDistribution)

	fmt.Println(finalResult)
	fmt.Println("AntFarm successfully parsed!")
}
