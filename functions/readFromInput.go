package lemin

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)




func (af *AntFarm) ReadFromInput(filename string)error{
	var state string
	var room Room
	var tunnel *Tunnel
	file,err:=os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner:=bufio.NewScanner(file)
	for i:=0;scanner.Scan();i++{
		line:=strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line,"##"){
			if line=="##start"{
				state="start"
			}else if line == "##end"{
				state="end"
			}
			continue
		}else if line==""||strings.HasPrefix(line,"#"){
			continue
		}
		parts:=strings.Fields(line)
		// Handle the first line (number of ants)
		if len(parts) == 1 && i == 0 {
			af.Ants, err = strconv.Atoi(parts[0])
			if err != nil {
				return fmt.Errorf("invalid number of ants: %v", err)
			}
		}else if len(parts)==3{
			if len(parts[0])>0 && parts[0][0]=='L'{
				return err
			}
			err = room.ParseRoom(parts)
			if err!=nil{
				return err
			}
			
		}else if len(parts)==1 && strings.Contains(parts[0], "-"){
			
		}
	}

	return nil
}

func (rm *Room) ParseRoom(parts []string) error {
	if len(parts)!=3{
		return fmt.Errorf("invalide parts") 
	}
	rm.Name=parts[0]
	x,errX:=strconv.Atoi(parts[1])
	y,errY:=strconv.Atoi(parts[2])
	if errX !=nil || errY !=nil{
		return errX
	}
	rm.X=x
	rm.Y=y
	return nil
}
