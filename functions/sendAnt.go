package lemin

import "fmt"

func SendAnt(groupes [][][]string, ant int) {
	bestGroup := groupes[0]
	lenbestGroup:=make([][]int,len(bestGroup))
	//lenbestGroup[1] = append(lenbestGroup[1], "1")
	for i,_:=range bestGroup{
		pth:=i
		fmt.Println("pth:::",pth)
		for j,_:=range lenbestGroup{
			antlen:=j
			if pth>=antlen && ant >0{
				lenbestGroup = append(lenbestGroup,lenbestGroup...)
				ant--
			}
		}
	}
	
	fmt.Println("Best Group:", lenbestGroup)
	fmt.Println("Best Group:", bestGroup)

	
}
