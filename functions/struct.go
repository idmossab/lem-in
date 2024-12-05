package functions

import (
	"strconv"
	"strings"
)

type Infos struct {
	NbOfants int
	Start    string
	End      string
	Rooms    map[string]string
	Tunneles map[string]string
}

func NewInfos() *Infos {
	return &Infos{
		Rooms:    make(map[string]string),
		Tunneles: make(map[string]string),
	}
}

func (inf *Infos) Treatment(lines []string) string {
	var err error
	inf.NbOfants, err = strconv.Atoi(lines[0])
	if err != nil {
		return "invalid number of ants"
	}

	if inf.NbOfants <= 0 {
		return "Number of ants is negative"
	}

	for i := 1; i < len(lines); i++ {
		if lines[i] == "##start" {
		
			continue
		}
		room := strings.Fields(lines[i])

		if isValidRomm(room) && inf.Start == "" {
			inf.Start = room[0]
			inf.Rooms[room[0]] = room[1] + " " + room[2]
			
		}

		if isValidRomm(room) {
			inf.Rooms[room[0]] = room[1] + " " + room[2]
		}
	}

	return ""
}

func isValidRomm(room []string) bool {
	if len(room) != 3 {
		return false
	}
	if room[0][0] == 'L' || room[0][0] == 'l'{
		return false
	}
	_, err := strconv.Atoi(room[1])
	_, err2 := strconv.Atoi(room[2])
	if err != nil || err2 != nil {
		return false
	}
	return true
}
