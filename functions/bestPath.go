package lemin

import "sort"

// Fonction BestPaths
func BestPaths(paths [][]string) ([][]string, [][]string) {
	bestPaths := [][]string{}

	// Stocker les nœuds intermédiaires déjà utilisés pour éviter les répétitions
	seen := make(map[string]bool)

	for _, path := range paths {
		unique := true

		// Vérifier les nœuds intermédiaires (exclure le premier et le dernier)
		for i := 1; i < len(path)-1; i++ {
			if seen[path[i]] { // Si le nœud est déjà utilisé dans un autre chemin
				unique = false
				break
			}
		}

		if unique {
			// Ajouter les nœuds intermédiaires du chemin à `seen`
			for i := 1; i < len(path)-1; i++ {
				seen[path[i]] = true
			}

			// Ajouter ce chemin à la liste des meilleurs chemins
			bestPaths = append(bestPaths, path)
		}
	}

	return bestPaths, paths
}

func SortByLength(slices [][]string) [][]string {
	sort.Slice(slices, func(i, j int) bool {
		return len(slices[i]) < len(slices[j]) // Sort in ascending order by length
	})
	return slices
}
