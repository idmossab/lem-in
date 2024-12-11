package lemin

import (
	"bufio"
	structs "lemin/structs"
	"os"
	"strings"
)


func (af *structs.AntFarm) ReadFromInput(filename string)error{
	var state string
	var room structs.Room
	var tunnel structs.Tunnel
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
				state=="end"
			}
			continue
		}else if line==""||strings.HasPrefix(line,"#"){
			continue
		}
		parts:=strings.Fields(line)
		if len(parts)==1 && i==0 && !strings.Contains(parts[0],"-")
	}


}

