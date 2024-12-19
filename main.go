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
	//fmt.Println(sr.Paths)
	//best,all:=sr.BestPaths(sr.SortByLength(sr.Paths))
	//fmt.Println(best)
	//fmt.Println(all)
	unique,anothre:=sr.UniquePaths(sr.SortByLength(sr.Paths))
	fmt.Println("Unique paths :",unique)
	fmt.Println("Another paths :",anothre)
	fmt.Println()
	groupe:=sr.GroupPaths(sr.SortByLength(sr.Paths))
	fmt.Println("groupe",groupe)
	fmt.Println("AntFarm successfully parsed!")
}
