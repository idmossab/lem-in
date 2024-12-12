package main

import (
	"fmt"
	"log"
	"os"

	sr "lemin/functions"
)

func main() {
	antFarm := sr.AntFarm{}
	err := antFarm.ReadFromInput(os.Args[1])
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	fmt.Println("AntFarm successfully parsed!")
}
