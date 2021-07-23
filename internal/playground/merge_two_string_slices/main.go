package main

import "fmt"

func main() {
	s1 := []string{"a.txt", "b.txt", "c.txt", "d.txt"}
	s2 := []string{"b.txt", "c.txt", "e.txt", "f.txt"}

	all := append(s1, s2...)
	seen := make(map[string]struct{}, len(all))
	uniq := 0

	for _, val := range all {
		// Check if val is in map
		if _, ok := seen[val]; ok {
			continue
		}

		seen[val] = struct{}{}
		all[uniq] = val
		uniq++
	}

	mergedSliced := all[:uniq]
	fmt.Println(mergedSliced)
}
