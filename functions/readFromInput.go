package lemin

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (af *AntFarm) ReadFromInput(filename string) error {
	var (
		state            string
		room             Room
		tunnel           Tunnel
		nbrStart, nbrEnd int
	)

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines or comments
		if line == "" || strings.HasPrefix(line, "#") {
			if line == "##start" {
				state = "start"
				nbrStart++
			} else if line == "##end" {
				state = "end"
				nbrEnd++
			}
			continue
		}

		parts := strings.Fields(line)

		switch {
		// Handle number of ants
		case len(parts) == 1 && i == 0:
			af.Ants, err = strconv.Atoi(parts[0])
			if err != nil || af.Ants < 1 {
				return fmt.Errorf("invalid number of ants: %v", err)
			}

		// Handle room
		case len(parts) == 3:
			if strings.HasPrefix(parts[0], "L") {
				return fmt.Errorf("room name cannot start with 'L': %s", parts[0])
			}

			if err := room.ParseRoom(parts); err != nil {
				return fmt.Errorf("failed to parse room: %v", err)
			}

			if err := ValidateRoomUniqueness(af.Rooms, room); err != nil {
				return fmt.Errorf("duplicate room: %v", err)
			}

			switch state {
			case "start":
				if af.Start != (Room{}) {
					return fmt.Errorf("duplicate start room defined")
				}
				af.Start = room
			case "end":
				if af.End != (Room{}) {
					return fmt.Errorf("duplicate end room defined")
				}
				af.End = room
			}
			state = ""
			af.Rooms = append(af.Rooms, room)

		// Handle tunnel

		case len(parts) == 1 && strings.Contains(parts[0], "-"):
			parts := strings.Split(parts[0], "-")
			if parts[0] == "" || parts[1] == "" {
				return fmt.Errorf("invalid tunnel format: %s", line)
			}
			if len(parts) != 2 {
				return fmt.Errorf("invalid tunnel format not equal 2: %s", line)
			}
			tunnel = Tunnel{From: parts[0], To: parts[1]}
			if tunnel.From == tunnel.To {
				return fmt.Errorf("tunnel cannot connect a room to itself: %s", line)
			}
			af.Tunnels = append(af.Tunnels, tunnel)

		default:
			return fmt.Errorf("invalid input format: %s", line)
		}
	}

	/*if nbrStart != 1 {
		return fmt.Errorf("there must be exactly one '##start' directive")
	}
	if nbrEnd != 1 {
		return fmt.Errorf("there must be exactly one '##end' directive")
	}*/
	if af.Start == (Room{}) {
		return fmt.Errorf("missing start room definition")
	}
	if af.End == (Room{}) {
		return fmt.Errorf("missing end room definition")
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

func ValidateRoomUniqueness(rooms []Room, room Room) error {
	for i := 0; i < len(rooms); i++ {
		if rooms[i].Name == room.Name {
			return fmt.Errorf("ERROR: Room %s is defined more than once", rooms[i].Name)
		} else if rooms[i].X == room.X && rooms[i].Y == room.Y {
			return fmt.Errorf("ERROR: Duplicate coordinates detected for rooms '%s' and '%s'", rooms[i].Name, room.Name)
		}
	}
	return nil
}
