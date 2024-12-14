package lemin

func  GetPaths(tunnels []Tunnel, Start string, End string, path []string) {
	var paths = &Paths
	path = append(path, Start)
	if Start == End {
		// cpy := make([]string, len(path))
		// copy(cpy, path)
		*paths = append(*paths, path)
		return
	}
	for _, tunnel := range tunnels {
		if Start == tunnel.From && !Contains(path, tunnel.To) {
			GetPaths(tunnels, tunnel.To, End, path)
		} else if Start == tunnel.To && !Contains(path, tunnel.From) {
			GetPaths(tunnels, tunnel.From, End, path)
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