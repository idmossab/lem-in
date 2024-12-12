package lemin

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (af *AntFarm) ReadFromInput(filename string) error {
	var state string
	var room Room
	var tunnel Tunnel
	nbrstart:=0
	nbrend:=0
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error read file : %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "##") {
			if line == "##start" {
				state = "start"
				nbrstart++
			} else if line == "##end" {
				state = "end"
				nbrend++
			}
			continue
		} else if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Fields(line)
		// Handle the first line (number of ants)
		if len(parts) == 1 && i == 0 {
			af.Ants, err = strconv.Atoi(parts[0])
			if err != nil {
				return fmt.Errorf("invalid number of ants: %v", err)
			}
		} else if len(parts) == 3 {
			if len(parts[0]) > 0 && parts[0][0] == 'L' && parts[0][0] == 'l'{
				return fmt.Errorf("dont enter first charachter L : %v", err)
			}
			err = room.ParseRoom(parts)
			if err != nil {
				return fmt.Errorf("error to parseRoom : %v", err)
			}
			ValidateRoomUniqueness(af.Rooms,room)
			if state == "start"{
				CheckIsDouble(af.Start.Name, state)
				af.Start = room
				state = ""
			} else if state == "end" {
				CheckIsDouble(af.End.Name, state)
				af.End = room
				state = ""
			}
			af.Rooms = append(af.Rooms, room)
		} else if len(parts) == 1 && strings.Contains(parts[0], "-") {
			if parts = strings.Split(line, "-"); len(parts) == 2 {
				tunnel.From = parts[0]
				tunnel.To = parts[1]
				af.Tunnels = append(af.Tunnels, tunnel)
				if tunnel.From == tunnel.To {
					return fmt.Errorf("error to Pqrssing : %v", err)
				}
			} else {
				return fmt.Errorf("error to Parssing : %v", err)
			}
		} else {
			return fmt.Errorf("error to Parssing : %v", err)
		}
	}
	if nbrend !=1 || nbrstart!=1{
		return fmt.Errorf("error start or end: %v", err)
	}
	if (af.Start == Room{}) {
		return fmt.Errorf("missing start room")
	}
	if (af.End == Room{}) {
		return fmt.Errorf("missing end room")
	}

	return nil
}

func (rm *Room) ParseRoom(parts []string) error {
	rm.Name = parts[0]
	x, errX := strconv.Atoi(parts[1])
	y, errY := strconv.Atoi(parts[2])
	if errX != nil {
		return fmt.Errorf("invalid number errX: %v", errX)
	}
	if errY != nil {
		return fmt.Errorf("invalid number errY: %v", errY)
	}
	rm.X = x
	rm.Y = y
	return nil
}

func ValidateRoomUniqueness(rooms []Room, room Room) {
	for i := 0; i < len(rooms); i++ {
		if rooms[i].Name == room.Name {
			fmt.Printf("ERROR: Room %s is defined more than once.\n", rooms[i].Name)
			os.Exit(1)
		} else if rooms[i].X == room.X && rooms[i].Y == room.Y {
			fmt.Printf("ERROR: Duplicate coordinates detected for rooms '%s' and '%s'.\n", rooms[i].Name, room.Name)
			os.Exit(1)
		}
	}
}

func CheckIsDouble(room, state string) {
	if len(room) > 0 {
		fmt.Printf("ERROR : room %s is double\n", state)
		os.Exit(1)
	}
}
