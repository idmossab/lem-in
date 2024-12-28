package lemin

func ChoosePath(bestPaths [][][]string, ant int) [][]string {
	pathFinal := [][]string{}

	for i := 0; i < len(bestPaths); i++ {
		if ant <= len(bestPaths[i]) {
			pathFinal = bestPaths[i]
			return pathFinal 
		}
	}

	if len(bestPaths) > 0 {
		pathFinal = bestPaths[len(bestPaths)-1]
	}

	return pathFinal
}
