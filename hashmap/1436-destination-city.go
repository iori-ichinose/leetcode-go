package hashmap

func destCity(paths [][]string) string {
	count := map[string]int{}
	for i := range paths {
		count[paths[i][0]]++
	}

	for i := range paths {
		if count[paths[i][1]] == 0 {
			return paths[i][1]
		}
	}
	return ""
}
