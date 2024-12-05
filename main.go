package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"lemin/functions"
)

func CheckDuplicates(lines []string) bool {
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			if lines[i] == lines[j] {
				return false
			}
		}
	}

	return true
}

func main() {
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("File error")
		return
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || (line[0] == '#' && line != "##start" && line != "##end") {
			continue
		}
		lines = append(lines, line)
	}

	if !CheckDuplicates(lines) {
		fmt.Println("error data duplicates")
		return
	}

	if len(lines) < 6 {
		fmt.Println("error len lines")
		return
	}

	infos := functions.NewInfos()

	msg := infos.Treatment(lines)
	if msg != "" {
		fmt.Println(msg)
		return
	}

	fmt.Println(infos.NbOfants)
	fmt.Println(infos.Start)
	fmt.Println(infos.Rooms)
}
