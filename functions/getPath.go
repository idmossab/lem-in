package lemin

func GetPaths(af []Tunnel, Start string, End string, path []string) {
	var paths = &Paths
	if len(*paths) == 5000 {
		return
	}
	path = append(path, Start)
	for h := 0; h < len(af); h++ {
		if Start == End {
			cpy := make([]string, len(path))
			copy(cpy, path)
			*paths = append(*paths, cpy)
			return
		} else if Start == af[h].From && !Contains(path, af[h].To) {
			GetPaths(af, af[h].To, End, path)
		} else if Start == af[h].To && !Contains(path, af[h].From) {
			GetPaths(af, af[h].From, End, path)
		}
	}
}
func Contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}