package main

import (
	"maps"
	"slices"
)

func main() {

}
func groupAnagrams(strs []string) [][]string {
	Map := map[string][]string{}
	for _, s := range strs {
		b := []byte(s)
		slices.Sort(b)
		str := string(b)
		Map[str] = append(Map[str], s)
	}
	return slices.Collect(maps.Values(Map))
}
