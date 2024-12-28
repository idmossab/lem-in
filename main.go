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
	//fmt.Println("path : ",sr.Paths)
	groupe := sr.GroupPaths(sr.Paths)
	fmt.Println("groupe", groupe)
	nbrAnt := antFarm.Ants
	fmt.Println()
	filterGr := sr.FilterShortestSlices(groupe,nbrAnt)
	fmt.Println("filter :", filterGr)
	bestGr:=sr.ChoosePath(filterGr,antFarm.Ants)
	//fmt.Println("bestt :", bestGr)
	//fmt.Println(nbrAnt)
	antDistribution,ant := sr.SendAnt(bestGr, nbrAnt)
	sr.PrintAnt(antDistribution,ant)
	fmt.Println("AntFarm successfully parsed!")
}

