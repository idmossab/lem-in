package main

import (
	"fmt"
	"os"

	sr "lemin/functions"
)

func main() {
	antFarm := sr.AntFarm{}
	err := antFarm.ReadFromInput(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}
	fmt.Println("AntFarm successfully parsed!")
}
