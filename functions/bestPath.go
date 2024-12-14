package lemin

import "strings"

func FilterShortestPathsWithDuplicates(paths [][]string) [][]string {
	// Map pour stocker la slice la plus courte associée à chaque valeur
	shortestPaths := make(map[string][]string)

	// Parcours de chaque chemin
	for _, path := range paths {
		for _, node := range path { // Parcours des nodes du chemin
			// Si la node est déjà dans la map, vérifie si le chemin actuel est plus court
			if existingPath, exists := shortestPaths[node]; exists {
				if len(path) < len(existingPath) {
					shortestPaths[node] = path // Met à jour avec le chemin le plus court
				}
			} else {
				// Sinon, ajoute le chemin à la map
				shortestPaths[node] = path
			}
		}
	}

	// Résultat final : Collecte des chemins uniques
	uniquePaths := make([][]string, 0, len(shortestPaths))
	seenPaths := make(map[string]bool) // Pour éviter les duplications de slices

	for _, path := range shortestPaths {
		key := strings.Join(path, ",") // Crée une clé unique pour chaque slice
		if !seenPaths[key] {
			uniquePaths = append(uniquePaths, path)
			seenPaths[key] = true
		}
	}

	return uniquePaths
}
