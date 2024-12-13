package lemin

func  GetPaths(af []Tunnel, Start string, End string, path []string) {
	var paths = &Paths
	path = append(path, Start)
	if Start == End {
		cpy := make([]string, len(path))
		copy(cpy, path)
		*paths = append(*paths, cpy)
		return
	}
	for _, tunnel := range af {
		if Start == tunnel.From && !Contains(path, tunnel.To) {
			GetPaths(af, tunnel.To, End, path,)
		} else if Start == tunnel.To && !Contains(path, tunnel.From) {
			GetPaths(af, tunnel.From, End, path)
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