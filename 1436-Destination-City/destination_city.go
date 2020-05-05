package main

import (
	"fmt"
)

func destCity(paths [][]string) string {
	cityDirectPath := map[string]string{}
	for _, pair := range paths {
		city1 := pair[0]
		city2 := pair[1]
		cityDirectPath[city1] = city2
	}
	for _, pair := range paths {
		city2 := pair[1]
		if _, ok := cityDirectPath[city2]; !ok {
			return city2
		}
	}
	return ""
}

func main() {
	path := [][]string{
		{"London", "New York"},
		{"New York", "Lima"},
		{"Lima", "Sao Paulo"},
	}
	fmt.Println(destCity(path))

	path = [][]string{
		{"B", "C"}, {"D", "B"}, {"C", "A"},
	}
	fmt.Println(destCity(path))

	path = [][]string{{"A", "Z"}}
	fmt.Println(destCity(path))
}
